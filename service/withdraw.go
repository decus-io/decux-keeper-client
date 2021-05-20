package service

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/psbt"
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
			},
		},
	}
	var withdraw message.Withdraw
	if err := coordinator.SendOperation(op, &withdraw); err != nil {
		return err
	}
	if err := s.processWithdrawImpl(receipt, &withdraw); err != nil {
		return err
	}

	s.signed[receipt.ReceiptId] = true
	log.Print("withdraw signed: ", receipt.ReceiptId)
	return nil
}

func (s *Withdraw) rawTxId(receipt *contract.Receipt) [32]byte {
	txid := receipt.TxId
	for i, j := 0, len(txid)-1; i < j; i, j = i+1, j-1 {
		txid[i], txid[j] = txid[j], txid[i]
	}
	return txid
}

func (s *Withdraw) verifyTransaction(receipt *contract.Receipt, p *psbt.Packet) error {
	tx := p.UnsignedTx
	if len(tx.TxIn) != 1 || len(tx.TxOut) != 1 || len(p.Inputs) != 1 {
		return errors.New("invalid input/ouput num")
	}
	if tx.TxIn[0].PreviousOutPoint.Hash != s.rawTxId(receipt) {
		return errors.New("unmatched txid")
	}
	if p.Inputs[0].WitnessUtxo == nil || p.Inputs[0].WitnessScript == nil {
		return errors.New("missing witness fields")
	}
	if p.Inputs[0].WitnessUtxo.Value != receipt.AmountInSatoshi.Int64() {
		return errors.New("unmatched input value")
	}

	pkScript, err := txscript.ParsePkScript(tx.TxOut[0].PkScript)
	if err != nil {
		return err
	}
	addrFromScript, err := pkScript.Address(config.C.Btc.NetworkParams)
	if err != nil {
		return err
	}
	addrFromReceipt, err := btcutil.DecodeAddress(receipt.WithdrawBtcAddress, config.C.Btc.NetworkParams)
	if err != nil {
		return err
	}
	if addrFromScript.String() != addrFromReceipt.String() {
		return errors.New("unmatched output address")
	}

	fee := p.Inputs[0].WitnessUtxo.Value - tx.TxOut[0].Value
	if fee > 1e6 {
		return fmt.Errorf("fee too high: %v", fee)
	}

	return nil
}

func (s *Withdraw) processWithdrawImpl(receipt *contract.Receipt, withdraw *message.Withdraw) error {
	p, err := psbt.NewFromRawBytes(strings.NewReader(withdraw.Psbt), true)
	if err != nil {
		return err
	}
	if err := s.verifyTransaction(receipt, p); err != nil {
		return fmt.Errorf("verify transaction failed: %v", err)
	}

	sigHashes := txscript.NewTxSigHashes(p.UnsignedTx)
	sig, err := txscript.RawTxInWitnessSignature(p.UnsignedTx, sigHashes, 0,
		p.Inputs[0].WitnessUtxo.Value, p.Inputs[0].WitnessScript,
		txscript.SigHashAll, config.C.Keeper.BtcKey)
	if err != nil {
		return err
	}

	updater, err := psbt.NewUpdater(p)
	if err != nil {
		return err
	}
	pubKey := config.C.Keeper.BtcKey.PubKey().SerializeCompressed()
	success, err := updater.Sign(0, sig, pubKey, nil, nil)
	if err != nil {
		return err
	}
	if success != psbt.SignSuccesful {
		return errors.New("sign err")
	}

	signed, err := p.B64Encode()
	if err != nil {
		return err
	}

	op := &message.Operation{
		Operation: &message.Operation_WithdrawSignature{
			WithdrawSignature: &message.WithdrawSignature{
				ReceiptId: withdraw.ReceiptId,
				Psbt:      signed,
			},
		},
	}
	return coordinator.SendOperation(op, nil)
}
