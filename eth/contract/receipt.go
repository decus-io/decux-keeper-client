package contract

import (
	"encoding/hex"

	"github.com/block-well/decux-keeper-client/eth/abi"
)

const (
	Available = iota
	DepositRequested
	DepositReceived
	WithdrawRequested
)

type Receipt struct {
	ReceiptId string
	abi.IDecuxSystemReceipt
}

func ReceiptIdToString(receiptId [32]byte) string {
	return "0x" + hex.EncodeToString(receiptId[:])
}
