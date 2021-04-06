package service

import (
	"log"
	"math/big"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Group struct {
	Groups        []*big.Int
	nextSyncIndex *big.Int
}

func NewGroup() *Group {
	return &Group{
		nextSyncIndex: big.NewInt(0),
	}
}

func (s *Group) SyncGroups(opts *bind.CallOpts) error {
	nGroups, err := contract.GroupRegistry.NGroups(opts)
	if err != nil {
		return err
	}

	updated := false
	for s.nextSyncIndex.Cmp(nGroups) < 0 {
		grp, err := contract.GroupRegistry.GetKeeperGroups(opts, config.C.Keeper.Id, s.nextSyncIndex)
		if err != nil {
			return err
		}
		for i := 0; i < grp.BitLen(); i++ {
			if grp.Bit(i) != 0 {
				s.Groups = append(s.Groups, big.NewInt(0).Add(s.nextSyncIndex, big.NewInt(int64(i+1))))
				updated = true
			}
		}
		s.nextSyncIndex.Add(s.nextSyncIndex, big.NewInt(256))
	}

	if s.nextSyncIndex.Cmp(nGroups) > 0 {
		s.nextSyncIndex = nGroups
	}
	if updated {
		log.Print("groups updated: ", s.Groups)
	}

	return nil
}
