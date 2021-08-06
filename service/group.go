package service

import (
	"fmt"
	"log"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/eth/abi"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type GroupService struct {
	groups        map[string]bool // groupId -> isMyGroup
	unknownGroups map[string]bool
}

func NewGroupService() *GroupService {
	return &GroupService{
		groups:        map[string]bool{},
		unknownGroups: map[string]bool{},
	}
}

func (s *GroupService) OnGroupAdded(event *abi.DeCusSystemGroupAdded) {
	s.processGroup(event.Required, event.BtcAddress, event.Keepers)
}

func (s *GroupService) OnGroupDeleted(event *abi.DeCusSystemGroupDeleted) {
	s.deleteGroup(event.BtcAddress)
}

func (s *GroupService) CheckGroup(groupId string) {
	if _, ok := s.groups[groupId]; ok {
		return
	}
	if _, ok := s.unknownGroups[groupId]; !ok {
		s.unknownGroups[groupId] = true
		log.Print("unknown group: ", groupId)
	}
}

func (s *GroupService) OnTimer() {
	for groupId := range s.unknownGroups {
		opts, cancel := contract.MakeCallOpts()
		group, err := contract.DeCusSystem.GetGroup(opts, groupId)
		cancel()

		if err != nil {
			log.Print("get group error: ", err)
			return
		}
		s.processGroup(group.Required, groupId, group.Keepers)
	}
}

func (s *GroupService) MyGroups() []string {
	ret := []string{}
	for groupId, isMyGroup := range s.groups {
		if isMyGroup {
			ret = append(ret, groupId)
		}
	}
	return ret
}

func (s *GroupService) IsMyGroup(groupId string) bool {
	isMygroup, ok := s.groups[groupId]
	return ok && isMygroup
}

func (s *GroupService) ReceiptByGroupId(opts *bind.CallOpts, groupId string) (*contract.Receipt, error) {
	group, err := contract.DeCusSystem.GetGroup(opts, groupId)
	if err != nil {
		return nil, err
	}
	s.processGroup(group.Required, groupId, group.Keepers)

	if group.Required == 0 {
		return nil, fmt.Errorf("group not exist: %v", groupId)
	}

	receipt := contract.Receipt{ReceiptId: contract.ReceiptIdToString(group.WorkingReceiptId)}
	if receipt.IDeCusSystemReceipt, err = contract.DeCusSystem.GetReceipt(opts,
		group.WorkingReceiptId); err != nil {
		return nil, err
	}

	return &receipt, nil
}

func (s *GroupService) containsMe(keepers []common.Address) bool {
	for _, keeper := range keepers {
		if keeper == config.C.Keeper.Id {
			return true
		}
	}
	return false
}

func (s *GroupService) processGroup(requiredNum uint32, groupId string, keepers []common.Address) {
	delete(s.unknownGroups, groupId)

	if requiredNum > 0 {
		if _, ok := s.groups[groupId]; !ok {
			isMyGroup := s.containsMe(keepers)
			s.groups[groupId] = isMyGroup
			if isMyGroup {
				log.Print("join group: ", groupId)
			}
		}
	} else {
		s.deleteGroup(groupId)
	}
}

func (s *GroupService) deleteGroup(groupId string) {
	isMyGroup, ok := s.groups[groupId]
	if ok && isMyGroup {
		log.Print("leave group: ", groupId)
	}

	delete(s.groups, groupId)
}
