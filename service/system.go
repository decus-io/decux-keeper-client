package service

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/decus-io/decus-keeper-client/config"
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

func (s *System) syncGroups() {
	opts, cancel := s.makeCallOpts()
	defer cancel()

	if err := s.groupService.SyncGroups(opts); err != nil {
		log.Print("sync groups error: ", err)
	}
}

func (s *System) checkReceipt(groupId *big.Int) error {
	opts, cancel := s.makeCallOpts()
	defer cancel()

	receiptId, err := contract.ReceiptController.GetWorkingReceiptId(opts, groupId)
	if err != nil {
		return err
	}
	receiptStatus, err := contract.ReceiptController.GetReceiptStatus(opts, receiptId)
	if err != nil {
		return err
	}

	status := int(receiptStatus.Int64())
	if status == DepositRequested {
		return s.depositService.ProcessDeposit(opts, groupId, receiptId)
	} else if status == WithdrawRequested {
		return s.withdrawService.ProcessWithdraw(opts, groupId, receiptId)
	}

	return nil
}

func (s *System) checkAllReceipt() {
	for _, groupId := range s.groupService.Groups {
		if err := s.checkReceipt(groupId); err != nil {
			log.Print("check receipt error: ", err)
		}
	}
}

func (s *System) onGroupAdded(event *abi.GroupRegistryGroupAdded) {
	log.Print("event GroupAdded: ", event.Id)

	for _, keeperId := range event.Keepers {
		if keeperId == config.C.Keeper.Id {
			s.syncGroups()
			return
		}
	}
}

func (s *System) Run() {
	heartbeatTicker := time.NewTicker(time.Minute * 2)
	syncGroupTicker := time.NewTicker(time.Minute * 10)
	checkReceiptTicker := time.NewTicker(time.Minute * 5)

	s.syncGroups()

	for {
		select {
		case <-heartbeatTicker.C:
			if err := s.keeperService.Heartbeat(); err != nil {
				log.Print("send heartbeat error: ", err) // TODO: use level based logging
			}
		case <-syncGroupTicker.C:
			s.syncGroups()
		case <-checkReceiptTicker.C:
			s.checkAllReceipt()
		case event := <-contract.GroupAdded:
			s.onGroupAdded(event)
		// TODO: WithdrawRequested should provide GroupId
		case event := <-contract.WithdrawRequested:
			log.Print("event WithdrawRequested: ", event.ReceiptId)
			s.checkAllReceipt()
		}
	}
}
