package config

import (
	"crypto/ecdsa"
	"io/ioutil"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v2"
)

var C Config

type Config struct {
	Keeper struct {
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
		DeCusSystem    string `yaml:"decus_system"`
		KeeperRegistry string `yaml:"keeper_registry"`
	} `yaml:"contract"`
}

func Init(user string) error {
	fb, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	if err := yaml.UnmarshalStrict(fb, &C); err != nil {
		return err
	}

	if C.Keeper.BtcKey, err = loadBtcKey(user); err != nil {
		return err
	}
	if C.Keeper.EthKey, err = loadEthKey(user); err != nil {
		return err
	}

	infuraId, err := loadInfuraId()
	if err != nil {
		return err
	}
	C.EthClient.WssUrl += infuraId
	log.Print("infura id: ", infuraId)

	C.Keeper.Id = crypto.PubkeyToAddress(C.Keeper.EthKey.PublicKey)
	log.Print("keeper id: ", C.Keeper.Id.Hex())

	return nil
}
