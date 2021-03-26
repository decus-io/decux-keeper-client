package service

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/coordinator"
	"github.com/decus-io/decus-keeper-client/db"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/decus-io/decus-keeper-proto/golang/message"
	"gorm.io/gorm"
)

type Withdraw struct {
}

func NewWithdraw() *Withdraw {
	return &Withdraw{}
}

func (s *Withdraw) ProcessWithdraw(groupId *big.Int, receiptId *big.Int) error {
	amount, err := contract.ReceiptController.GetAmountInSatoshi(nil, receiptId)
	if err != nil {
		return err
	}
	receiptInfo, err := contract.ReceiptController.GetReceiptInfo(nil, receiptId)
	if err != nil {
		return err
	}

	op := &message.Operation{
		Operation: &message.Operation_WithdrawRequest{
			WithdrawRequest: &message.WithdrawRequest{
				GroupId:   int32(groupId.Int64()),
				ReceiptId: int32(receiptId.Int64()),
				Recipient: receiptInfo.BtcAddress,
				Amount:    amount.Uint64(),
			},
		},
	}
	var withdraw message.Withdraw
	if err = coordinator.SendOperation(op, &withdraw); err != nil {
		return err
	}
	return s.processWithdrawImpl(&withdraw)
}

func (s *Withdraw) processWithdrawImpl(withdraw *message.Withdraw) error {
	return db.Transaction(context.Background(), func(ctx context.Context, db *gorm.DB) error {
		/*
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
			inWithdraw, err := wd.Get(uint(withdraw.ReceiptId))
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
		*/

		// TODO: need verification ?
		paymentRaw, err := hex.DecodeString(withdraw.PaymentRaw)
		if err != nil {
			return err
		}
		ptx, err := btcutil.NewTxFromBytes(paymentRaw)
		if err != nil {
			return err
		}
		tx := ptx.MsgTx()

		redeemScript, err := hex.DecodeString(withdraw.RedeemScript)
		if err != nil {
			return err
		}
		sig, err := txscript.RawTxInSignature(tx, 0, redeemScript,
			txscript.SigHashAll, config.C.Keeper.BtcKey)
		if err != nil {
			log.Fatal("sign tx error: ", err)
		}

		op := &message.Operation{
			Operation: &message.Operation_WithdrawSignature{
				WithdrawSignature: &message.WithdrawSignature{
					ReceiptId: withdraw.ReceiptId,
					Signature: sig,
				},
			},
		}
		return coordinator.SendOperation(op, nil)

		/*
			_, err = wd.Create(&dao.Withdraw{
				Common:    dao.Common{Id: uint(withdraw.ReceiptId)},
				GroupId:   uint(withdraw.GroupId),
				Recipient: withdraw.Recipient,
				Amount:    withdraw.Amount,
			})
			return err
		*/
	})
}
