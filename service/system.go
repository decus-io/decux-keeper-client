package service

import (
	"fmt"
	"log"
	"time"

	"github.com/decux-io/decux-keeper-client/eth/contract"
	"github.com/ethereum/go-ethereum/core/types"
)

type System struct {
	heartbeatOK bool

	eventService    *EventService
	keeperService   *KeeperService
	groupService    *GroupService
	depositService  *DepositService
	withdrawService *WithdrawService
	refundService   *RefundService
}

func NewSystem() *System {
	groupService := NewGroupService()

	return &System{
		heartbeatOK:     true,
		eventService:    NewEventService(),
		keeperService:   NewKeeperService(),
		groupService:    groupService,
		depositService:  NewDepositService(),
		withdrawService: NewWithdrawService(),
		refundService:   NewRefundService(groupService),
	}
}

func (s *System) Init() error {
	if err := s.keeperService.Init(); err != nil {
		return fmt.Errorf("init keeper error: %v", err)
	}
	return nil
}

func (s *System) handleEvent(l types.Log) {
	if err := s.handleEventImpl(l); err != nil {
		log.Print("handle event error: ", err)
	}
}

func (s *System) handleEventImpl(l types.Log) error {
	eventId := contract.EventID
	filterer := contract.DecuxSystemFilterer

	switch l.Topics[0] {
	case eventId["GroupAdded"]:
		event, err := filterer.ParseGroupAdded(l)
		if err != nil {
			return err
		}
		s.groupService.OnGroupAdded(event)
	case eventId["GroupDeleted"]:
		event, err := filterer.ParseGroupDeleted(l)
		if err != nil {
			return err
		}
		s.groupService.OnGroupDeleted(event)
	case eventId["MintRequested"]:
		event, err := filterer.ParseMintRequested(l)
		if err != nil {
			return err
		}
		s.groupService.CheckGroup(event.GroupBtcAddress)
		if s.groupService.IsMyGroup(event.GroupBtcAddress) {
			log.Print("event MintRequested: ", contract.ReceiptIdToString(event.ReceiptId))
		}
	case eventId["BurnRequested"]:
		event, err := filterer.ParseBurnRequested(l)
		if err != nil {
			return err
		}
		s.groupService.CheckGroup(event.GroupBtcAddress)
		if s.groupService.IsMyGroup(event.GroupBtcAddress) {
			log.Print("event BurnRequested: ", contract.ReceiptIdToString(event.ReceiptId))
			// call checkReceipt immediately to get the blockNumber
			s.checkReceipt(event.GroupBtcAddress)
		}
	}

	return nil
}

func (s *System) syncEvents() {
	if err := s.eventService.Sync(s.handleEvent); err != nil {
		log.Print("sync events warning: ", err)
	}
}

func (s *System) checkReceipt(groupId string) {
	if err := s.checkReceiptImpl(groupId); err != nil {
		log.Print("check receipt error: ", err)
	}
}

func (s *System) checkReceiptImpl(groupId string) error {
	opts, cancel := contract.MakeCallOpts()
	defer cancel()

	receipt, err := s.groupService.ReceiptByGroupId(opts, groupId)
	if err != nil {
		return err
	}

	if receipt.Status == contract.DepositRequested {
		return s.depositService.ProcessDeposit(receipt)
	} else if receipt.Status == contract.WithdrawRequested {
		return s.withdrawService.ProcessWithdraw(opts, receipt)
	}

	return nil
}

func (s *System) checkAllReceipt() {
	for _, groupId := range s.groupService.MyGroups() {
		s.checkReceipt(groupId)
	}
}

func (s *System) checkBtcRefundImpl() error {
	opts, cancel := contract.MakeCallOpts()
	defer cancel()

	refundData, err := contract.DecuxSystem.GetRefundData(opts)
	if err != nil {
		return err
	}
	if s.groupService.IsMyGroup(refundData.GroupBtcAddress) {
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
	minuteTicker := time.NewTicker(time.Minute)
	checkContractTicker := time.NewTicker(time.Minute * 3)

	s.syncEvents()

	for {
		select {
		case <-heartbeatTicker.C:
			groupNum := uint64(len(s.groupService.MyGroups()))
			syncMinutes := s.eventService.SyncMinutes()
			err := s.keeperService.Heartbeat(&groupNum, syncMinutes)
			s.heartbeatOK = (err == nil)
			if err != nil {
				log.Print("error sending heartbeat: ", err)
			}
		case <-minuteTicker.C:
			s.syncEvents()
			s.groupService.OnTimer()
		case <-checkContractTicker.C:
			if s.heartbeatOK {
				s.checkAllReceipt()
				s.checkBtcRefund()
			}
		}
	}
}
