package helper

import (
	"sort"

	"github.com/decus-io/decus-keeper-client/btc"
	"github.com/decus-io/decus-keeper-client/eth/contract"
)

func FindUtxo(receipt *contract.Receipt) (*btc.Utxo, error) {
	utxo, err := btc.QueryUtxo(receipt.GroupBtcAddress)
	if err != nil {
		return nil, err
	}
	// reduce the chance that different keepers select different utxo
	// (normally there won't be multiple utxo)
	sort.Slice(utxo, func(i, j int) bool {
		return utxo[i].Status.Block_Height < utxo[j].Status.Block_Height
	})

	for _, v := range utxo {
		if v.Status.Confirmed &&
			v.Value == receipt.AmountInSatoshi.Uint64() &&
			v.Status.Block_Time > receipt.CreateTimestamp.Uint64() {
			return &v, nil
		}
	}

	return nil, nil
}
