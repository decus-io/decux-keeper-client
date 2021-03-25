package service

import (
	"math/big"
	"strconv"

	"github.com/decus-io/decus-keeper-client/btc"
	"github.com/decus-io/decus-keeper-client/coordinator"
	"github.com/decus-io/decus-keeper-client/eth"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/decus-io/decus-keeper-proto/golang/message"
)

type Deposit struct {
}

func NewDeposit() *Deposit {
	return &Deposit{}
}

func (s *Deposit) signDeposit(groupId *big.Int, receiptId *big.Int, utxo *btc.Utxo) error {
	recipient, err := contract.ReceiptController.GetUserAddress(nil, receiptId)
	if err != nil {
		return err
	}

	signature, err := eth.SignDeposit(recipient.Hex(), receiptId.String(), strconv.FormatUint(utxo.Value, 10),
		"0x"+utxo.Txid, strconv.FormatInt(int64(utxo.Status.Block_Height), 10))
	if err != nil {
		return err
	}

	op := &message.Operation{
		Operation: &message.Operation_DepositSignature{
			DepositSignature: &message.DepositSignature{
				ReceiptId:   int32(receiptId.Int64()),
				GroupId:     int32(groupId.Int64()),
				Recipient:   recipient.Hex(),
				Amount:      utxo.Value,
				TxId:        utxo.Txid,
				BlockHeight: utxo.Status.Block_Height,
				Signature:   signature,
			},
		},
	}
	return coordinator.SendOperation(op, nil)
}

func (s *Deposit) ProcessDeposit(groupId *big.Int, receiptId *big.Int) error {
	// TODO: check if we have already sign it

	// TODO: cache the address
	groupInfo, err := contract.GroupRegistry.GetGroupInfo(nil, groupId)
	if err != nil {
		return err
	}
	utxo, err := btc.QueryUtxo(groupInfo.BtcAddress)
	if err != nil {
		return err
	}
	amount, err := contract.ReceiptController.GetAmountInSatoshi(nil, receiptId)
	if err != nil {
		return err
	}
	for _, v := range utxo {
		if v.Status.Confirmed && v.Value == amount.Uint64() {
			return s.signDeposit(groupId, receiptId, &v)
		}
	}

	return nil
}
