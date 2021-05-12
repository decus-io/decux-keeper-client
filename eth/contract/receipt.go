package contract

import "github.com/decus-io/decus-keeper-client/eth/abi"

const (
	Available = iota
	DepositRequested
	DepositReceived
	WithdrawRequested
)

type Receipt struct {
	ReceiptId string
	abi.IDeCusSystemReceipt
}
