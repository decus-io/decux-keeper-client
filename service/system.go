package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/decus-io/decus-keeper-client/eth/abi"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type System struct {
	keeperService   *Keeper
	groupService    *Group
	depositService  *Deposit
	withdrawService *Withdraw
}

func NewSystem() *System {
	return &System{
		keeperService:   NewKeeper(),
		groupService:    NewGroup(),
		depositService:  NewDeposit(),
		withdrawService: NewWithdraw(),
	}
}

func (s *System) Init() error {
	if err := s.keeperService.Init(); err != nil {
		return fmt.Errorf("init keeper error: %v", err)
	}
	return nil
}

func (s *System) makeCallOpts() (*bind.CallOpts, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	return &bind.CallOpts{Context: ctx}, cancel
}

func (s *System) receiptIdToString(receiptId [32]byte) string {
	return "0x" + hex.EncodeToString(receiptId[:])
}

func (s *System) checkReceipt(groupId string) error {
	opts, cancel := s.makeCallOpts()
	defer cancel()

	group, err := contract.DeCusSystem.GetGroup(opts, groupId)
	if err != nil {
		return err
	}
	receiptId, err := contract.DeCusSystem.GetReceiptId(opts, groupId, group.Nonce)
	if err != nil {
		return err
	}
	receipt := contract.Receipt{ReceiptId: s.receiptIdToString(receiptId)}
	if receipt.IDeCusSystemReceipt, err = contract.DeCusSystem.GetReceipt(opts, receiptId); err != nil {
		return err
	}

	if receipt.Status == contract.DepositRequested {
		return s.depositService.ProcessDeposit(&receipt)
	} else if receipt.Status == contract.WithdrawRequested {
		return s.withdrawService.ProcessWithdraw(&receipt)
	}

	return nil
}

func (s *System) checkAllReceipt() {
	for groupId := range s.groupService.Groups {
		if err := s.checkReceipt(groupId); err != nil {
			log.Print("check receipt error: ", err)
		}
	}
}

func (s *System) onBurnRequested(event *abi.DeCusSystemBurnRequested) error {
	opts, cancel := s.makeCallOpts()
	defer cancel()

	receipt, err := contract.DeCusSystem.GetReceipt(opts, event.ReceiptId)
	if err != nil {
		return err
	}
	if s.groupService.Groups[receipt.GroupBtcAddress] {
		log.Print("event BurnRequested: ", s.receiptIdToString(event.ReceiptId))
		return s.checkReceipt(receipt.GroupBtcAddress)
	}
	return nil
}

func (s *System) Run() {
	heartbeatTicker := time.NewTicker(time.Minute * 2)
	checkReceiptTicker := time.NewTicker(time.Minute * 5)

	for {
		select {
		case <-heartbeatTicker.C:
			if err := s.keeperService.Heartbeat(); err != nil {
				log.Print("error sending heartbeat: ", err)
			}
		case <-checkReceiptTicker.C:
			s.checkAllReceipt()
		case event := <-contract.GroupAdded:
			s.groupService.OnGroupAdded(event)
		case event := <-contract.GroupDeleted:
			s.groupService.OnGroupDeleted(event)
		case event := <-contract.BurnRequested:
			if err := s.onBurnRequested(event); err != nil {
				log.Print("error handling BurnRequested: ", err)
			}
		}
	}
}
