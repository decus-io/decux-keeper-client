package service

import (
	"encoding/hex"
	"log"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/coordinator"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/decus-io/decus-keeper-client/service/helper"
	"github.com/decus-io/decus-keeper-proto/golang/message"
)

type Withdraw struct {
	signed map[string]bool
}

func NewWithdraw() *Withdraw {
	return &Withdraw{
		signed: map[string]bool{},
	}
}

func (s *Withdraw) ProcessWithdraw(receipt *contract.Receipt) error {
	if s.signed[receipt.ReceiptId] {
		return nil
	}

	utxo, err := helper.FindUtxo(receipt)
	if err != nil {
		return err
	}
	if utxo == nil {
		s.signed[receipt.ReceiptId] = true
		log.Print("withdrawing should have been finished for uxto not found: ", receipt.ReceiptId)
		return nil
	}

	op := &message.Operation{
		Operation: &message.Operation_WithdrawRequest{
			WithdrawRequest: &message.WithdrawRequest{
				ReceiptId: receipt.ReceiptId,
				GroupId:   receipt.GroupBtcAddress,
				Recipient: receipt.WithdrawBtcAddress,
				Amount:    receipt.AmountInSatoshi.Uint64(),
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

	s.signed[receipt.ReceiptId] = true
	log.Print("withdraw signed: ", receipt.ReceiptId)
	return nil
}

func (s *Withdraw) processWithdrawImpl(withdraw *message.Withdraw) error {
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
}
