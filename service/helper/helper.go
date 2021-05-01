package helper

import (
	"math/big"
	"sort"

	"github.com/decus-io/decus-keeper-client/btc"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func FindUtxo(opts *bind.CallOpts, groupId *big.Int, receiptId *big.Int) (*btc.Utxo, error) {
	// TODO: cache the address
	groupInfo, err := contract.GroupRegistry.GetGroupInfo(opts, groupId)
	if err != nil {
		return nil, err
	}
	utxo, err := btc.QueryUtxo(groupInfo.BtcAddress)
	if err != nil {
		return nil, err
	}
	// reduce the chance that different keepers select different utxo
	// (normally there won't be multiple utxo)
	sort.Slice(utxo, func(i, j int) bool {
		return utxo[i].Status.Block_Height < utxo[j].Status.Block_Height
	})

	amount, err := contract.ReceiptController.GetAmountInSatoshi(opts, receiptId)
	if err != nil {
		return nil, err
	}
	for _, v := range utxo {
		if v.Status.Confirmed && v.Value == amount.Uint64() {
			return &v, nil
		}
	}

	return nil, nil
}
