package service

import (
	"bytes"
	"context"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/dao"
	"github.com/decus-io/decus-keeper-client/db"
	"github.com/decus-io/decus-keeper-proto/golang/message"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"gorm.io/gorm"
)

type Withdraw struct {
}

func NewWithdraw() *Withdraw {
	return &Withdraw{}
}

func (s *Withdraw) CheckNewWithdraw() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.C.Coordinator.Url+"withdraws/new", nil)
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

	var withdraws message.Withdraws
	if err := proto.Unmarshal(body, &withdraws); err != nil {
		return err
	}

	for _, withdraw := range withdraws.Data {
		if err := s.ProcessWithdraw(withdraw); err != nil {
			return err
		}
	}

	return nil
}

func (s *Withdraw) ProcessWithdraw(withdraw *message.Withdraw) error {
	return db.Transaction(context.Background(), func(ctx context.Context, db *gorm.DB) error {
		gkd, err := dao.NewGroupKeeperDao(db)
		if err != nil {
			return err
		}
		inGroup, err := gkd.CheckInGroup(uint(withdraw.GroupId), config.C.Keeper.Id)
		if err != nil {
			return err
		}
		if !inGroup {
			return nil
		}

		wd, err := dao.NewWithdrawDao(db)
		if err != nil {
			return err
		}
		inWithdraw, err := wd.Get(uint(withdraw.Id))
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return err
			}
		}
		if inWithdraw != nil {
			return nil
		}

		gd, err := dao.NewGroupDao(db)
		if err != nil {
			return err
		}
		group, err := gd.Get(uint(withdraw.GroupId))
		if err != nil {
			return err
		}
		redeemScript, _ := hex.DecodeString(group.RedeemScript)

		ptx, err := btcutil.NewTxFromBytes(withdraw.PaymentRaw)
		if err != nil {
			return err
		}
		tx := ptx.MsgTx()

		sig, err := txscript.RawTxInSignature(tx, 0, redeemScript, txscript.SigHashAll, config.C.Keeper.Key)
		if err != nil {
			log.Fatalf("sign tx error: %v", err)
		}

		op := message.Operation{
			Id:        1,
			Nonce:     1,
			PubKey:    config.C.Keeper.Key.PubKey().SerializeCompressed(),
			Signature: []byte{1},
			Timestamp: &timestamp.Timestamp{
				Seconds: time.Now().Unix(),
			},
			Operation: &message.Operation_WithdrawSignature{
				WithdrawSignature: &message.WithdrawSignature{
					WithdrawId: withdraw.Id,
					Signature:  sig,
				},
			},
		}
		data, err := proto.Marshal(&op)
		if err != nil {
			return err
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", config.C.Coordinator.Url+"/operation", bytes.NewReader(data))
		if err != nil {
			return err
		}
		req.Header.Add("Content-Type", "application/x-protobuf")
		_, err = client.Do(req)
		if err != nil {
			return err
		}

		_, err = wd.Create(&dao.Withdraw{
			Common:    dao.Common{Id: uint(withdraw.Id)},
			GroupId:   uint(withdraw.GroupId),
			Recipient: withdraw.Recipient,
			Amount:    withdraw.Amount,
		})
		return err
	})
}
