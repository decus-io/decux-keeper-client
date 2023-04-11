package eth

import (
	"fmt"

	"github.com/decux-io/decux-keeper-client/config"
	"github.com/decux-io/decux-keeper-client/eth/contract"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
)

func SignMessage(message []byte, legacyV bool) (hexutil.Bytes, error) {
	digestHash := crypto.Keccak256(message)
	signature, err := crypto.Sign(digestHash, config.C.Keeper.EthKey)
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
	signature, err := SignMessage(rawData, true)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func SignDeposit(receiptId, txId, height string) (hexutil.Bytes, error) {
	typesStandard := core.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"MintRequest": {
			{Name: "receiptId", Type: "bytes32"},
			{Name: "txId", Type: "bytes32"},
			{Name: "height", Type: "uint256"},
		},
	}
	domainStandard := core.TypedDataDomain{
		Name:              "Decux",
		Version:           "1.0",
		ChainId:           math.NewHexOrDecimal256(contract.ChainId),
		VerifyingContract: config.C.Contract.DecuxSystem,
		Salt:              "",
	}
	messageStandard := map[string]interface{}{
		"receiptId": receiptId,
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
