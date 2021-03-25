package service

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/decus-io/decus-keeper-client/eth/contract"
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

func (s *System) checkReceipt(groupId *big.Int) error {
	receiptId, err := contract.ReceiptController.GetWorkingReceiptId(nil, groupId)
	if err != nil {
		return err
	}
	receiptStatus, err := contract.ReceiptController.GetReceiptStatus(nil, receiptId)
	if err != nil {
		return err
	}

	status := int(receiptStatus.Int64())
	if status == DepositRequested {
		return s.depositService.ProcessDeposit(groupId, receiptId)
	} else if status == WithdrawRequested {
		return s.withdrawService.ProcessWithdraw(groupId, receiptId)
	}

	return nil
}

func (s *System) checkAllReceipt() {
	for _, groupId := range s.groupService.Groups {
		if err := s.checkReceipt(groupId); err != nil {
			log.Print("check receipt err: ", err)
		}
	}
}

func (s *System) Run() {
	heartbeatTicker := time.NewTicker(time.Minute * 2)
	syncGroupTicker := time.NewTicker(time.Minute * 10)
	checkReceiptTicker := time.NewTicker(time.Minute * 10)

	for {
		select {
		case <-heartbeatTicker.C:
			if err := s.keeperService.Heartbeat(); err != nil {
				log.Print("send heart error: ", err) // TODO: use level based logging
			}
		case <-syncGroupTicker.C:
			if err := s.groupService.SyncGroups(); err != nil {
				log.Print("sync groups error: ", err)
			}
		case <-checkReceiptTicker.C:
			s.checkAllReceipt()
		}
	}
}
