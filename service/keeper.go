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
	op := &message.Operation{
		Operation: &message.Operation_Init{
			Init: &message.Init{
				BtcPubKey: config.C.Keeper.BtcKey.PubKey().SerializeCompressed(),
				Signature: []byte{1},
			},
		},
	}
	return coordinator.SendOperation(op, nil)
}

func (*Keeper) Heartbeat() error {
	op := &message.Operation{
		Operation: &message.Operation_Heartbeat{
			Heartbeat: &message.Heartbeat{
				Payload: []byte{2},
			},
		},
	}
	return coordinator.SendOperation(op, nil)
}
