package contract

import (
	"encoding/hex"

	"github.com/decus-io/decus-keeper-client/eth/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

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

func ReceiptIdToString(receiptId [32]byte) string {
	return "0x" + hex.EncodeToString(receiptId[:])
}

func ReceiptByGroupId(opts *bind.CallOpts, groupId string) (*Receipt, error) {
	group, err := DeCusSystem.GetGroup(opts, groupId)
	if err != nil {
		return nil, err
	}
	receipt := Receipt{ReceiptId: ReceiptIdToString(group.WorkingReceiptId)}
	if receipt.IDeCusSystemReceipt, err = DeCusSystem.GetReceipt(opts,
		group.WorkingReceiptId); err != nil {
		return nil, err
	}

	return &receipt, nil
}
