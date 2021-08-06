package config

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v2"
)

var C Config

type Config struct {
	Keeper struct {
		Id     common.Address    `yaml:"-"`
		Email  string            `yaml:"-"`
		BtcKey *btcec.PrivateKey `yaml:"-"`
		EthKey *ecdsa.PrivateKey `yaml:"-"`
	} `yaml:"keeper"`
	Btc struct {
		Network       string           `yaml:"network"`
		NetworkParams *chaincfg.Params `yaml:"-"`
	} `yaml:"btc"`
	Url struct {
		Coordinator string `yaml:"coordinator"`
		EthClient   string `yaml:"ethclient"`
		BtcEsplora  string `yaml:"btcesplora"`
	} `yaml:"url"`
	Contract struct {
		DeCusSystem           string `yaml:"decus_system"`
		KeeperRegistry        string `yaml:"keeper_registry"`
		DeCusSystemStartBlock uint64 `yaml:"decus_system_start_block"`
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

	log.Print("contract DeCusSystem: ", C.Contract.DeCusSystem)
	log.Print("contract KeeperRegistry: ", C.Contract.KeeperRegistry)

	switch C.Btc.Network {
	case "testnet":
		C.Btc.NetworkParams = &chaincfg.TestNet3Params
	case "mainnet":
		C.Btc.NetworkParams = &chaincfg.MainNetParams
	default:
		return fmt.Errorf("unknown btc networkd: %v", C.Btc.Network)
	}

	if C.Keeper.BtcKey, err = loadBtcKey(user); err != nil {
		return err
	}
	if C.Keeper.EthKey, err = loadEthKey(user); err != nil {
		return err
	}

	if strings.Contains(C.Url.EthClient, "infura") {
		infuraId, err := loadInfuraId(user)
		if err != nil {
			return err
		}
		C.Url.EthClient += infuraId
		log.Print("infura id: ", infuraId)
	}

	// optional
	email, err := loadEmail(user)
	if err != nil {
		email = ""
	}
	C.Keeper.Email = email

	C.Keeper.Id = crypto.PubkeyToAddress(C.Keeper.EthKey.PublicKey)
	log.Print("keeper id: ", C.Keeper.Id.Hex())

	return nil
}
