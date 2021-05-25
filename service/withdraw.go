package service

import (
	"log"

	"github.com/decus-io/decus-keeper-client/coordinator"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/decus-io/decus-keeper-client/service/helper"
	"github.com/decus-io/decus-keeper-proto/golang/message"
)

type Withdraw struct {
	processed map[string]bool
}

func NewWithdraw() *Withdraw {
	return &Withdraw{
		processed: map[string]bool{},
	}
}

func (s *Withdraw) ProcessWithdraw(receipt *contract.Receipt) error {
	if s.processed[receipt.ReceiptId] {
		return nil
	}

	utxo, err := helper.UtxoByReceipt(receipt)
	if err != nil {
		return err
	}
	if utxo == nil {
		s.processed[receipt.ReceiptId] = true
		log.Print("withdrawing should have been finished for uxto not found: ", receipt.ReceiptId)
		return nil
	}

	op := &message.Operation{
		Operation: &message.Operation_WithdrawRequest{
			WithdrawRequest: &message.WithdrawRequest{ReceiptId: receipt.ReceiptId},
		},
	}
	var withdrawRsp message.Withdraw
	if err := coordinator.SendOperation(op, &withdrawRsp); err != nil {
		return err
	}
	if err := s.processWithdrawImpl(receipt, withdrawRsp.Psbt); err != nil {
		return err
	}

	s.processed[receipt.ReceiptId] = true
	log.Print("withdraw signed: ", receipt.ReceiptId)
	return nil
}

func (s *Withdraw) processWithdrawImpl(receipt *contract.Receipt, psbt string) error {
	txInfo := &helper.TxInfo{
		PreTxId:    receipt.TxId,
		TargetAddr: receipt.WithdrawBtcAddress,
		Amount:     receipt.AmountInSatoshi,
	}
	signed, err := helper.SignPsbt(txInfo, psbt)
	if err != nil {
		return err
	}

	op := &message.Operation{
		Operation: &message.Operation_WithdrawSignature{
			WithdrawSignature: &message.WithdrawSignature{
				ReceiptId: receipt.ReceiptId,
				Psbt:      signed,
			},
		},
	}
	return coordinator.SendOperation(op, nil)
}
