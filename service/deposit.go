package service

import (
	"log"
	"math/big"
	"strconv"

	"github.com/decus-io/decus-keeper-client/btc"
	"github.com/decus-io/decus-keeper-client/coordinator"
	"github.com/decus-io/decus-keeper-client/eth"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/decus-io/decus-keeper-client/service/helper"
	"github.com/decus-io/decus-keeper-proto/golang/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Deposit struct {
	signed map[int64]bool
}

func NewDeposit() *Deposit {
	return &Deposit{
		signed: map[int64]bool{},
	}
}

func (s *Deposit) signDeposit(opts *bind.CallOpts, groupId *big.Int, receiptId *big.Int, utxo *btc.Utxo) error {
	recipient, err := contract.ReceiptController.GetUserAddress(opts, receiptId)
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

func (s *Deposit) ProcessDeposit(opts *bind.CallOpts, groupId *big.Int, receiptId *big.Int) error {
	if s.signed[receiptId.Int64()] {
		return nil
	}

	utxo, err := helper.FindUtxo(opts, groupId, receiptId)
	if err != nil {
		return err
	}

	if err := s.signDeposit(opts, groupId, receiptId, utxo); err != nil {
		return err
	}
	s.signed[receiptId.Int64()] = true
	log.Print("deposit signed: ", receiptId, " txid: ", utxo.Txid)

	return nil
}
