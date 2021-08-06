package service

import (
	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/coordinator"
	"github.com/decus-io/decus-keeper-proto/golang/message"
)

type KeeperService struct {
}

func NewKeeperService() *KeeperService {
	return &KeeperService{}
}

func (s *KeeperService) Init() error {
	return s.Heartbeat(nil)
}

func (*KeeperService) Heartbeat(groupNum *uint64) error {
	op := &message.Operation{
		Operation: &message.Operation_Heartbeat{
			Heartbeat: &message.Heartbeat{
				DecusSystem: config.C.Contract.DeCusSystem,
				BtcPubKey:   config.C.Keeper.BtcKey.PubKey().SerializeCompressed(),
				Email:       config.C.Keeper.Email,
				GroupNum:    groupNum,
			},
		},
	}
	return coordinator.SendOperation(op, nil)
}
