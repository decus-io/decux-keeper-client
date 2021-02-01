package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"log"
)

func main() {
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

	fmt.Println(hex.EncodeToString(redeemScript))

	//redeemHash := btcutil.Hash160(redeemScript)
	//addr, err := btcutil.NewAddressScriptHashFromHash(redeemHash, &chaincfg.TestNet3Params)
	//if err != nil {
	//	log.Fatalf("new address script error: %v", err)
	//}
	//fmt.Println("multisig address: " + addr.EncodeAddress())
	//PkScript, err := txscript.PayToAddrScript(addr)
	//if err != nil {
	//	log.Fatalf("pay to address script error: %v", err)
	//}
	//
	//preHash, _ := chainhash.NewHashFromStr("fef560fda122454b5bb9131cc004f25e2df35df688b4f43106ab4149e469a71d")
	//tx := &wire.MsgTx{
	//	Version: wire.TxVersion,
	//	TxIn: []*wire.TxIn{
	//		{
	//			PreviousOutPoint: wire.OutPoint{
	//				Hash:  *preHash,
	//				Index: 0,
	//			},
	//		},
	//	},
	//	TxOut: []*wire.TxOut{
	//		{
	//			Value:    900000,
	//			PkScript: PkScript,
	//		},
	//	},
	//	LockTime: 0,
	//}

	txBytes, _ := hex.DecodeString("020000000127598b28d5e019f8c6b73a4a7bcbd53cba3c1948a4ed35f2f6b4446faf0054d20000000000ffffffff0100350c000000000017a9141276148e7c31f4e61eb1f31d0380391092e5e4578700000000")
	ptx, _ := btcutil.NewTxFromBytes(txBytes)
	tx := ptx.MsgTx()

	wif1, err := btcutil.DecodeWIF("cQ97jdD2aCcb2MuQd7XVo4cPyQ6UH1J6Jz8LcAetghA9CXwDJcxa")
	if err != nil {
		log.Fatalf("decode wif error: %v", err)
	}
	sig1, err := txscript.RawTxInSignature(tx, 0, redeemScript, txscript.SigHashAll, wif1.PrivKey)
	if err != nil {
		log.Fatalf("sign tx error: %v", err)
	}

	wif2, err := btcutil.DecodeWIF("cS2cvrAc1PJaixn2v6SLXLRJEH8ULErbV12QjXNX6B5UmZN47WwZ")
	if err != nil {
		log.Fatalf("decode wif error: %v", err)
	}
	sig2, err := txscript.RawTxInSignature(tx, 0, redeemScript, txscript.SigHashAll, wif2.PrivKey)
	if err != nil {
		log.Fatalf("sign tx error: %v", err)
	}

	wif3, err := btcutil.DecodeWIF("cScpHakK8uEbeVJsE1G92FM79ESe93bbWZsTFbcLHThUuSd3X8LK")
	if err != nil {
		log.Fatalf("decode wif error: %v", err)
	}
	sig3, err := txscript.RawTxInSignature(tx, 0, redeemScript, txscript.SigHashAll, wif3.PrivKey)
	if err != nil {
		log.Fatalf("sign tx error: %v", err)
	}

	sigScript, err := txscript.NewScriptBuilder().AddOp(txscript.OP_FALSE).AddData(sig1).AddData(sig2).AddData(sig3).AddData(redeemScript).Script()
	if err != nil {
		log.Fatalf("new signature script error: %v", err)
	}
	tx.TxIn[0].SignatureScript = sigScript

	fmt.Println(hex.EncodeToString(sigScript))

	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	if err := tx.Serialize(buf); err != nil {
		log.Fatalf("serialize tx data error: %v", err)
	}
	fmt.Println(hex.EncodeToString(buf.Bytes()))
}
