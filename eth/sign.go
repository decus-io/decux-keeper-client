package eth

import (
	"fmt"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
)

func sign(sighash []byte, legacyV bool) (hexutil.Bytes, error) {
	signature, err := crypto.Sign(sighash, config.C.Keeper.EthKey)
	if err != nil {
		return nil, err
	}
	if legacyV {
		signature[64] += 27
	}
	return signature, nil
}

func signTypedData(typedData core.TypedData) (hexutil.Bytes, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	sighash := crypto.Keccak256(rawData)

	signature, err := sign(sighash, true)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func SignDeposit(recipient, receiptId, amount, txId, height string) (hexutil.Bytes, error) {
	typesStandard := core.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"MintRequest": {
			{Name: "recipient", Type: "address"},
			{Name: "nonce", Type: "uint256"},
			{Name: "amount", Type: "uint256"},
			{Name: "txId", Type: "bytes32"},
			{Name: "height", Type: "uint256"},
		},
	}
	domainStandard := core.TypedDataDomain{
		Name:              "DeCus",
		Version:           "1.0",
		ChainId:           math.NewHexOrDecimal256(contract.ChainID),
		VerifyingContract: config.C.Contract.SignatureValidator,
		Salt:              "",
	}
	messageStandard := map[string]interface{}{
		"recipient": recipient,
		"nonce":     receiptId,
		"amount":    amount,
		"txId":      txId,
		"height":    height,
	}
	typedData := core.TypedData{
		Types:       typesStandard,
		PrimaryType: "MintRequest",
		Domain:      domainStandard,
		Message:     messageStandard,
	}

	return signTypedData(typedData)
}
