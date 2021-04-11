package config

import (
	"crypto/ecdsa"
	"errors"
	"io/ioutil"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ququzone/go-common/env"
	"gopkg.in/yaml.v2"
)

var C Config

type Config struct {
	Network string `yaml:"network"`
	Keeper  struct {
		Id     common.Address    `yaml:"-"`
		BtcKey *btcec.PrivateKey `yaml:"-"`
		EthKey *ecdsa.PrivateKey `yaml:"-"`
	} `yaml:"keeper"`
	Coordinator struct {
		Url string `yaml:"url"`
	} `yaml:"coordinator"`
	EthClient struct {
		WssUrl string `yaml:"wss_url"`
	} `yaml:"ethclient"`
	BtcEsplora struct {
		Url string `yaml:"url"`
	} `yaml:"btcesplora"`
	Contract struct {
		DeCusSystem       string `yaml:"decus_system"`
		GroupRegistry     string `yaml:"group_registry"`
		KeeperRegistry    string `yaml:"keeper_registry"`
		ReceiptController string `yaml:"receipt_controller"`
	} `yaml:"contract"`
}

func Init() error {
	fb, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	if err := yaml.UnmarshalStrict(fb, &C); err != nil {
		return err
	}

	wif, err := btcutil.DecodeWIF(env.GetNonEmpty("BTC_PRIVATE_KEY"))
	if err != nil {
		return err
	}
	C.Keeper.BtcKey = wif.PrivKey

	C.Keeper.EthKey, err = crypto.HexToECDSA(env.GetNonEmpty("ETH_PRIVATE_KEY"))
	if err != nil {
		return err
	}

	ethPubkey, ok := C.Keeper.EthKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error getting eth pubkey")
	}
	C.Keeper.Id = crypto.PubkeyToAddress(*ethPubkey)
	log.Print("keeper id: ", C.Keeper.Id.Hex())

	return nil
}
