package service

import (
	"log"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/eth/abi"
)

type Group struct {
	Groups map[string]bool
}

func NewGroup() *Group {
	return &Group{
		Groups: map[string]bool{},
	}
}

func (s *Group) OnGroupAdded(event *abi.DeCusSystemGroupAdded) {
	for _, keeper := range event.Keepers {
		if keeper == config.C.Keeper.Id {
			s.Groups[event.BtcAddress] = true
			log.Print("join group: ", event.BtcAddress)
			break
		}
	}
}

func (s *Group) OnGroupDeleted(event *abi.DeCusSystemGroupDeleted) {
	_, ok := s.Groups[event.BtcAddress]
	if ok {
		delete(s.Groups, event.BtcAddress)
		log.Print("leave group: ", event.BtcAddress)
	}
}
