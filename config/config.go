package config

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v2"
)

var (
	Version string = "v0.0.1"
	C       Config
)

type Config struct {
	Keeper struct {
		User   string
		Id     common.Address
		Email  string
		BtcKey *btcec.PrivateKey
		EthKey *ecdsa.PrivateKey
	} `yaml:"-"`
	Btc struct {
		Network       string           `yaml:"network"`
		NetworkParams *chaincfg.Params `yaml:"-"`
		Confirmations uint32
	} `yaml:"btc"`
	Url struct {
		Coordinator string `yaml:"coordinator"`
		NetworkRpc  string `yaml:"-"`
		BtcEsplora  string `yaml:"btcesplora"`
	} `yaml:"url"`
	Contract struct {
		DecuxSystem           string `yaml:"decux_system"`
		KeeperRegistry        string `yaml:"keeper_registry"`
		DecuxSystemStartBlock uint64 `yaml:"decux_system_start_block"`
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

	C.Keeper.User = user

	log.Print("contract DecuxSystem: ", C.Contract.DecuxSystem)
	log.Print("contract KeeperRegistry: ", C.Contract.KeeperRegistry)

	switch C.Btc.Network {
	case "testnet":
		C.Btc.NetworkParams = &chaincfg.TestNet3Params
	case "mainnet":
		C.Btc.NetworkParams = &chaincfg.MainNetParams
	default:
		return fmt.Errorf("unknown btc networkd: %v", C.Btc.Network)
	}

	if C.Keeper.BtcKey, err = loadBtcKey(); err != nil {
		return err
	}
	if C.Keeper.EthKey, err = loadEthKey(); err != nil {
		return err
	}

	if C.Url.NetworkRpc, err = loadRpcUrl(); err != nil {
		return err
	}
	log.Print("network rpc url: ", C.Url.NetworkRpc)

	// optional
	email, err := loadEmail()
	if err != nil {
		email = ""
	}
	C.Keeper.Email = email

	C.Keeper.Id = crypto.PubkeyToAddress(C.Keeper.EthKey.PublicKey)
	log.Print("keeper id: ", C.Keeper.Id.Hex())

	return nil
}
