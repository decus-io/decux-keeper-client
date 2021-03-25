package service

import (
	"context"
	"log"
	"math/big"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/db"
	"github.com/decus-io/decus-keeper-client/db/dao"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/decus-io/decus-keeper-proto/golang/message"
	"gorm.io/gorm"
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

func (*Group) Create(groupMessage *message.Group) error {
	return db.Transaction(context.Background(), func(ctx context.Context, db *gorm.DB) error {
		gd, err := dao.NewGroupDao(db)
		if err != nil {
			return err
		}
		group, err := gd.Create(&dao.Group{
			Common:       dao.Common{Id: uint(groupMessage.Id)},
			Address:      groupMessage.Address,
			RedeemScript: groupMessage.RedeemScript,
		})
		if err != nil {
			return nil
		}

		gkd, err := dao.NewGroupKeeperDao(db)
		if err != nil {
			return err
		}
		for _, keeperId := range groupMessage.Keepers {
			_, err := gkd.Create(&dao.GroupKeeper{
				GroupId:  group.Id,
				KeeperId: uint(keeperId),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *Group) SyncGroups() error {
	nGroups, err := contract.GroupRegistry.NGroups(nil)
	if err != nil {
		return err
	}

	updated := false
	for s.nextSyncIndex.Cmp(nGroups) < 0 {
		grp, err := contract.GroupRegistry.GetKeeperGroups(nil, config.C.Keeper.Id, s.nextSyncIndex)
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
