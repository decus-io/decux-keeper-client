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
	"github.com/decus-io/decus-keeper-client/service/helper"
	"github.com/decus-io/decus-keeper-proto/golang/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gorm.io/gorm"
)

type Withdraw struct {
	signed map[int64]bool
}

func NewWithdraw() *Withdraw {
	return &Withdraw{
		signed: map[int64]bool{},
	}
}

func (s *Withdraw) ProcessWithdraw(opts *bind.CallOpts, groupId *big.Int, receiptId *big.Int) error {
	if s.signed[receiptId.Int64()] {
		return nil
	}

	utxo, err := helper.FindUtxo(opts, groupId, receiptId)
	if err != nil {
		return err
	}
	if utxo == nil {
		s.signed[receiptId.Int64()] = true
		log.Print("withdrawing should have been finished for uxto not found: ", receiptId)
		return nil
	}

	amount, err := contract.ReceiptController.GetAmountInSatoshi(opts, receiptId)
	if err != nil {
		return err
	}
	receiptInfo, err := contract.ReceiptController.GetReceiptInfo(opts, receiptId)
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
	if err := coordinator.SendOperation(op, &withdraw); err != nil {
		return err
	}
	if err := s.processWithdrawImpl(&withdraw); err != nil {
		return err
	}

	s.signed[receiptId.Int64()] = true
	log.Print("withdraw signed: ", receiptId)
	return nil
}

func (s *Withdraw) processWithdrawImpl(withdraw *message.Withdraw) error {
	return db.Transaction(context.Background(), func(ctx context.Context, db *gorm.DB) error {
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
	})
}
