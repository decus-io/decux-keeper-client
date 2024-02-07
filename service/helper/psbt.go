package helper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/psbt"
	"github.com/block-well/decux-keeper-client/config"
)

type TxInfo struct {
	PreTxId    chainhash.Hash
	TargetAddr string
	Amount     int64
}

func rawTxId(txId chainhash.Hash) chainhash.Hash {
	size := len(txId)
	for i := 0; i < size/2; i++ {
		txId[i], txId[size-1-i] = txId[size-1-i], txId[i]
	}
	return txId
}

func verifyTransaction(txInfo *TxInfo, p *psbt.Packet) error {
	tx := p.UnsignedTx
	if len(tx.TxIn) != 1 || len(tx.TxOut) != 1 || len(p.Inputs) != 1 {
		return errors.New("invalid input/ouput num")
	}
	if tx.TxIn[0].PreviousOutPoint.Hash != rawTxId(txInfo.PreTxId) {
		return errors.New("unmatched txid")
	}
	if p.Inputs[0].WitnessUtxo == nil || p.Inputs[0].WitnessScript == nil {
		return errors.New("missing witness fields")
	}
	if p.Inputs[0].WitnessUtxo.Value != txInfo.Amount {
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
	addrFromTarget, err := btcutil.DecodeAddress(txInfo.TargetAddr, config.C.Btc.NetworkParams)
	if err != nil {
		return err
	}
	if addrFromScript.String() != addrFromTarget.String() {
		return errors.New("unmatched output address")
	}

	fee := p.Inputs[0].WitnessUtxo.Value - tx.TxOut[0].Value
	if fee > 1e6 {
		return fmt.Errorf("fee too high: %v", fee)
	}

	return nil
}

func signImpl(txInfo *TxInfo, p *psbt.Packet) error {
	if err := verifyTransaction(txInfo, p); err != nil {
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

	return nil
}

func SignPsbt(txParams *TxInfo, psbtData string) (string, error) {
	p, err := psbt.NewFromRawBytes(strings.NewReader(psbtData), true)
	if err != nil {
		return "", err
	}
	if err := signImpl(txParams, p); err != nil {
		return "", err
	}

	return p.B64Encode()
}
