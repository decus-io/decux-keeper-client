package btc

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/decus-io/decus-keeper-client/config"
)

type Utxo struct {
	Txid   string
	Value  uint64
	Status struct {
		Confirmed    bool
		Block_Height uint32
		Block_Time   uint32
	}
}

type Tx struct {
	Vin []struct {
		Prevout struct {
			Scriptpubkey_Address string
		}
	}
}

var client = &http.Client{
	Timeout: time.Second * 30,
}

func request(subUrl string, result interface{}) error {
	req, err := http.NewRequest("GET", config.C.Url.BtcEsplora+subUrl, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(result)
}

func QueryUtxo(address string) ([]Utxo, error) {
	var utxo []Utxo
	err := request("address/"+address+"/utxo", &utxo)
	if err != nil {
		return nil, err
	}
	return utxo, nil
}

func QueryTx(txId string) (*Tx, error) {
	var tx Tx
	err := request("tx/"+txId, &tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}
