package service

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/decus-io/decus-keeper-client/btc"
	"github.com/decus-io/decus-keeper-client/coordinator"
	"github.com/decus-io/decus-keeper-client/eth/abi"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/decus-io/decus-keeper-client/service/helper"
	"github.com/decus-io/decus-keeper-proto/golang/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Refund struct {
	processed map[chainhash.Hash]bool
}

func NewRefund() *Refund {
	return &Refund{
		processed: map[chainhash.Hash]bool{},
	}
}

func (s *Refund) ProcessRefund(opts *bind.CallOpts, refundData *abi.IDeCusSystemBtcRefundData) error {
	if s.processed[refundData.TxId] {
		return nil
	}

	utxo, err := helper.UtxoByRefundData(refundData)
	if err != nil {
		return err
	}
	if utxo == nil {
		s.processed[refundData.TxId] = true
		log.Print("refund should have been finished for uxto not found: ",
			hex.EncodeToString(refundData.TxId[:]))
		return nil
	}

	receipt, err := contract.ReceiptByGroupId(opts, refundData.GroupBtcAddress)
	if err != nil {
		return err
	}
	if receipt.TxId == refundData.TxId {
		s.processed[refundData.TxId] = true
		return fmt.Errorf("refund error, utxo is used by receipt: %v", receipt.ReceiptId)
	}

	if receipt.Status == contract.DepositRequested {
		receiptUtxo, err := helper.UtxoByReceipt(receipt)
		if err != nil {
			return err
		}
		if receiptUtxo != nil && strings.EqualFold(utxo.Txid, receiptUtxo.Txid) {
			s.processed[refundData.TxId] = true
			return fmt.Errorf("refund error, utxo is the candidate of receipt: %v", receipt.ReceiptId)
		}
	}

	op := &message.Operation{
		Operation: &message.Operation_RefundRequest{
			RefundRequest: &message.RefundRequest{TxId: utxo.Txid},
		},
	}
	var refundRsp message.Refund
	if err := coordinator.SendOperation(op, &refundRsp); err != nil {
		return err
	}
	if err := s.processRefundImpl(refundData, utxo, refundRsp.Psbt); err != nil {
		return err
	}

	s.processed[refundData.TxId] = true
	log.Print("refund signed: ", refundData.GroupBtcAddress, " - ", utxo.Txid)
	return nil
}

func (s *Refund) processRefundImpl(refundData *abi.IDeCusSystemBtcRefundData,
	utxo *btc.Utxo, psbt string) error {
	tx, err := btc.QueryTx(utxo.Txid)
	if err != nil {
		return err
	}
	if len(tx.Vin) == 0 {
		return fmt.Errorf("inputs not found in tx: %v", utxo.Txid)
	}

	txInfo := &helper.TxInfo{
		PreTxId:    refundData.TxId,
		TargetAddr: tx.Vin[0].Prevout.Scriptpubkey_Address,
		Amount:     int64(utxo.Value),
	}
	signed, err := helper.SignPsbt(txInfo, psbt)
	if err != nil {
		return err
	}

	op := &message.Operation{
		Operation: &message.Operation_RefundSignature{
			RefundSignature: &message.RefundSignature{
				TxId: utxo.Txid,
				Psbt: signed,
			},
		},
	}
	return coordinator.SendOperation(op, nil)
}
