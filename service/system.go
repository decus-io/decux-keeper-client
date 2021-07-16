package service

import (
	"context"
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
	refundService   *Refund
}

func NewSystem() *System {
	return &System{
		keeperService:   NewKeeper(),
		groupService:    NewGroup(),
		depositService:  NewDeposit(),
		withdrawService: NewWithdraw(),
		refundService:   NewRefund(),
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

func (s *System) checkReceipt(groupId string) error {
	opts, cancel := s.makeCallOpts()
	defer cancel()

	receipt, err := contract.ReceiptByGroupId(opts, groupId)
	if err != nil {
		return err
	}

	if receipt.Status == contract.DepositRequested {
		return s.depositService.ProcessDeposit(receipt)
	} else if receipt.Status == contract.WithdrawRequested {
		return s.withdrawService.ProcessWithdraw(receipt)
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
		log.Print("event BurnRequested: ", contract.ReceiptIdToString(event.ReceiptId))
		return s.checkReceipt(receipt.GroupBtcAddress)
	}
	return nil
}

func (s *System) checkBtcRefundImpl() error {
	opts, cancel := s.makeCallOpts()
	defer cancel()

	refundData, err := contract.DeCusSystem.GetRefundData(opts)
	if err != nil {
		return err
	}
	if s.groupService.Groups[refundData.GroupBtcAddress] {
		if err := s.refundService.ProcessRefund(opts, &refundData); err != nil {
			return err
		}
	}
	return nil
}

func (s *System) checkBtcRefund() {
	if err := s.checkBtcRefundImpl(); err != nil {
		log.Print("check refund error: ", err)
	}
}

func (s *System) Run() {
	heartbeatTicker := time.NewTicker(time.Minute * 2)
	checkContractTicker := time.NewTicker(time.Minute * 5)

	for {
		select {
		case <-heartbeatTicker.C:
			groupNum := uint64(len(s.groupService.Groups))
			if err := s.keeperService.Heartbeat(&groupNum); err != nil {
				log.Print("error sending heartbeat: ", err)
			}
		case <-checkContractTicker.C:
			s.checkAllReceipt()
			s.checkBtcRefund()
		case event := <-contract.GroupAdded:
			s.groupService.OnGroupAdded(event)
		case event := <-contract.GroupDeleted:
			s.groupService.OnGroupDeleted(event)
		case event := <-contract.BurnRequested:
			if err := s.onBurnRequested(event); err != nil {
				log.Print("error handling BurnRequested: ", err)
			}
		case event := <-contract.BtcRefunded:
			log.Print("event BtcRefunded: ", event.GroupBtcAddress)
			s.checkBtcRefund()
		}
	}
}
