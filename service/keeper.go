package service

import (
	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/coordinator"
	"github.com/decus-io/decus-keeper-proto/golang/message"
)

type Keeper struct {
}

func NewKeeper() *Keeper {
	return &Keeper{}
}

func (s *Keeper) Init() error {
	return s.Heartbeat()
}

func (*Keeper) Heartbeat() error {
	op := &message.Operation{
		Operation: &message.Operation_Heartbeat{
			Heartbeat: &message.Heartbeat{
				BtcPubKey: config.C.Keeper.BtcKey.PubKey().SerializeCompressed(),
			},
		},
	}
	return coordinator.SendOperation(op, nil)
}
