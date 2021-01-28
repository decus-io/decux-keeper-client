package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"log"
)

func main() {
	txBytes, _ := hex.DecodeString("02000000011da769e44941ab0631f4b488f65df32d5ef204c01c13b95b4b4522a1fd60f5fe0000000000ffffffff01a08601000000000017a9141276148e7c31f4e61eb1f31d0380391092e5e4578700000000")

	tx, err := btcutil.NewTxFromBytes(txBytes)
	if err != nil {
		log.Fatalf("decode tx error: %v", err)
	}

	pk1Bytes, _ := hex.DecodeString("0305fc63bed535ea0e7896625e34b5d65557d95e5db72839070e95a24f84be22b3")
	pk1, _ := btcutil.NewAddressPubKey(pk1Bytes, &chaincfg.TestNet3Params)
	pk2Bytes, _ := hex.DecodeString("02b3ff4cd0b6f271fb372e6114f9d91ca53b26a34031f7f181fae2baee1c5ce855")
	pk2, _ := btcutil.NewAddressPubKey(pk2Bytes, &chaincfg.TestNet3Params)
	pk3Bytes, _ := hex.DecodeString("02d11cbb3f610afd29cdb59f50f78748805ffdfc3e520be043d7e23348451b7044")
	pk3, _ := btcutil.NewAddressPubKey(pk3Bytes, &chaincfg.TestNet3Params)
	pk4Bytes, _ := hex.DecodeString("0262d6daf916d5bfbbb0f3fe8814b0285f35cdb0ae0b17034c9b175b3eabbc1fc2")
	pk4, _ := btcutil.NewAddressPubKey(pk4Bytes, &chaincfg.TestNet3Params)
	pk5Bytes, _ := hex.DecodeString("020cca1ee9d326ea27cfbe0463b0bc9b1ff53f777833f520b70a8735d9736e22a5")
	pk5, _ := btcutil.NewAddressPubKey(pk5Bytes, &chaincfg.TestNet3Params)

	redeemScript, err := txscript.MultiSigScript([]*btcutil.AddressPubKey{
		pk1, pk2, pk3, pk4, pk5,
	}, 3)
	if err != nil {
		log.Fatalf("new multisig script error: %v", err)
	}

	sigHashType := txscript.SigHashAll
	hash, err := txscript.CalcSignatureHash(redeemScript, sigHashType, tx.MsgTx(), 0)
	if err != nil {
		log.Fatalf("calc signature hash error: %v", err)
	}

	wif1, err := btcutil.DecodeWIF("cQ97jdD2aCcb2MuQd7XVo4cPyQ6UH1J6Jz8LcAetghA9CXwDJcxa")
	if err != nil {
		log.Fatalf("decode wif error: %v", err)
	}
	sig1, err := wif1.PrivKey.Sign(hash)
	if err != nil {
		log.Fatalf("signature hash error: %v", err)
	}
	sig1WithHashType := append(sig1.Serialize(), byte(sigHashType))

	wif2, err := btcutil.DecodeWIF("cQ97jdD2aCcb2MuQd7XVo4cPyQ6UH1J6Jz8LcAetghA9CXwDJcxa")
	if err != nil {
		log.Fatalf("decode wif error: %v", err)
	}
	sig2, err := wif2.PrivKey.Sign(hash)
	if err != nil {
		log.Fatalf("signature hash error: %v", err)
	}
	sig2WithHashType := append(sig2.Serialize(), byte(sigHashType))

	wif3, err := btcutil.DecodeWIF("cQ97jdD2aCcb2MuQd7XVo4cPyQ6UH1J6Jz8LcAetghA9CXwDJcxa")
	if err != nil {
		log.Fatalf("decode wif error: %v", err)
	}
	sig3, err := wif3.PrivKey.Sign(hash)
	if err != nil {
		log.Fatalf("signature hash error: %v", err)
	}
	sig3WithHashType := append(sig3.Serialize(), byte(sigHashType))

	sigScript, err := txscript.NewScriptBuilder().AddData(sig1WithHashType).AddData(sig2WithHashType).AddData(sig3WithHashType).AddData(redeemScript).Script()
	if err != nil {
		log.Fatalf("new signature script error: %v", err)
	}
	tx.MsgTx().TxIn[0].SignatureScript = sigScript

	tx1 := &wire.MsgTx{
		Version:  wire.TxVersion,
		TxIn:     tx.MsgTx().TxIn,
		TxOut:    tx.MsgTx().TxOut,
		LockTime: 0,
	}

	fmt.Println(hex.EncodeToString(sigScript))
	buf := bytes.NewBuffer(make([]byte, 0, tx1.SerializeSize()))
	if err := tx1.Serialize(buf); err != nil {
		log.Fatalf("serialize tx data error: %v", err)
	}
	fmt.Println(hex.EncodeToString(buf.Bytes()))
}
