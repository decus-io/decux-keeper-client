package service

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/decux-io/decux-keeper-client/btc"
	"github.com/decux-io/decux-keeper-client/coordinator"
	"github.com/decux-io/decux-keeper-client/eth/abi"
	"github.com/decux-io/decux-keeper-client/eth/contract"
	"github.com/decux-io/decux-keeper-client/service/helper"
	"github.com/decux-io/decux-keeper-proto/golang/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type RefundService struct {
	taskManager  *helper.TaskManager
	groupService *GroupService
}

func NewRefundService(groupService *GroupService) *RefundService {
	return &RefundService{
		taskManager:  helper.NewTaskManager(),
		groupService: groupService,
	}
}

func (s *RefundService) ProcessRefund(opts *bind.CallOpts, refundData *abi.IDecuxSystemBtcRefundData) error {
	task := hex.EncodeToString(refundData.TxId[:])
	if !s.taskManager.Available(task) {
		return nil
	}

	utxo, err := helper.UtxoByRefundData(refundData)
	if err != nil {
		return err
	}
	if utxo == nil {
		s.taskManager.Finish(task)
		log.Print("refund should have been finished for uxto not found: ", task)
		return nil
	}

	receipt, err := s.groupService.ReceiptByGroupId(opts, refundData.GroupBtcAddress)
	if err != nil {
		return err
	}
	if receipt.TxId == refundData.TxId {
		s.taskManager.Finish(task)
		return fmt.Errorf("refund error, utxo is used by receipt: %v", receipt.ReceiptId)
	}

	if receipt.Status == contract.DepositRequested {
		receiptUtxo, err := helper.UtxoByReceipt(receipt)
		if err != nil {
			return err
		}
		if receiptUtxo != nil && strings.EqualFold(utxo.Txid, receiptUtxo.Txid) {
			s.taskManager.CheckLater(task)
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

	s.taskManager.CheckLater(task)
	log.Print("refund signed: ", refundData.GroupBtcAddress, " - ", utxo.Txid)
	return nil
}

func (s *RefundService) processRefundImpl(refundData *abi.IDecuxSystemBtcRefundData,
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
