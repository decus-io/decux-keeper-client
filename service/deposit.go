package service

import (
	"log"
	"strconv"

	"github.com/decus-io/decus-keeper-client/btc"
	"github.com/decus-io/decus-keeper-client/coordinator"
	"github.com/decus-io/decus-keeper-client/eth"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/decus-io/decus-keeper-client/service/helper"
	"github.com/decus-io/decus-keeper-proto/golang/message"
)

type Deposit struct {
	processed map[string]bool
}

func NewDeposit() *Deposit {
	return &Deposit{
		processed: map[string]bool{},
	}
}

func (s *Deposit) signDeposit(receipt *contract.Receipt, utxo *btc.Utxo) error {
	signature, err := eth.SignDeposit(receipt.ReceiptId,
		"0x"+utxo.Txid, strconv.FormatInt(int64(utxo.Status.Block_Height), 10))
	if err != nil {
		return err
	}

	op := &message.Operation{
		Operation: &message.Operation_DepositSignature{
			DepositSignature: &message.DepositSignature{
				ReceiptId:   receipt.ReceiptId,
				TxId:        utxo.Txid,
				BlockHeight: uint64(utxo.Status.Block_Height),
				Signature:   signature,
			},
		},
	}
	return coordinator.SendOperation(op, nil)
}

func (s *Deposit) ProcessDeposit(receipt *contract.Receipt) error {
	if s.processed[receipt.ReceiptId] {
		return nil
	}

	utxo, err := helper.UtxoByReceipt(receipt)
	if err != nil {
		return err
	}
	if utxo == nil {
		return nil
	}

	if err := s.signDeposit(receipt, utxo); err != nil {
		return err
	}
	s.processed[receipt.ReceiptId] = true
	log.Print("deposit signed: ", receipt.ReceiptId, " txid: ", utxo.Txid)

	return nil
}
