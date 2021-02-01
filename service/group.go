package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/dao"
	"github.com/decus-io/decus-keeper-client/db"
	"github.com/decus-io/decus-keeper-proto/golang/message"
	"github.com/golang/protobuf/proto"
	"gorm.io/gorm"
)

type Group struct {
}

func NewGroup() *Group {
	return &Group{}
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

func (s *Group) SyncGroup() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.C.Coordinator.Url+"groups", nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var groups message.Groups
	if err := proto.Unmarshal(body, &groups); err != nil {
		return err
	}

	for _, group := range groups.Data {
		fmt.Println(group)
		if err := s.Create(group); err != nil {
			return nil
		}
	}

	return nil
}
