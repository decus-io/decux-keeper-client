package config

import (
	"bytes"
	"crypto/ecdsa"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/natefinch/atomic"
	"github.com/ququzone/go-common/env"
)

const (
	userDir = "./user"

	ethKeyFile = "eth-key"
	btcKeyFile = "btc-key"
	rpcUrlFile = "rpc-url"
	emailFile  = "email"
)

func DataPath(file string) string {
	return filepath.Join(userDir, C.Keeper.User, file)
}

func save(file string, r io.Reader) error {
	path := DataPath(file)
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	return atomic.WriteFile(path, r)
}

func password() string {
	return env.GetNonEmpty("PASSWORD")
}

func saveKey(file string, privateKey *ecdsa.PrivateKey) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	ksKey := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}

	data, err := keystore.EncryptKey(ksKey, password(), keystore.StandardScryptN,
		keystore.StandardScryptP)
	if err != nil {
		return err
	}
	return save(file, bytes.NewReader(data))
}

func SaveEthKey(ethKey *ecdsa.PrivateKey) error {
	return saveKey(ethKeyFile, ethKey)
}

func SaveBtcKey(btcKey *btcec.PrivateKey) error {
	return saveKey(btcKeyFile, btcKey.ToECDSA())
}

func SaveRpcUrl(rpcUrl string) error {
	return save(rpcUrlFile, strings.NewReader(rpcUrl))
}

func SaveEmail(email string) error {
	return save(emailFile, strings.NewReader(email))
}

//

func load(file string) ([]byte, error) {
	return ioutil.ReadFile(DataPath(file))
}

func loadKey(file string) (*ecdsa.PrivateKey, error) {
	data, err := load(file)
	if err != nil {
		return nil, err
	}
	ksKey, err := keystore.DecryptKey(data, password())
	if err != nil {
		return nil, err
	}

	return ksKey.PrivateKey, nil
}

func loadEthKey() (*ecdsa.PrivateKey, error) {
	return loadKey(ethKeyFile)
}

func loadBtcKey() (*btcec.PrivateKey, error) {
	privateKey, err := loadKey(btcKeyFile)
	if err != nil {
		return nil, err
	}
	return (*btcec.PrivateKey)(privateKey), nil
}

func loadRpcUrl() (string, error) {
	data, err := load(rpcUrlFile)
	return string(data), err
}

func loadEmail() (string, error) {
	data, err := load(emailFile)
	return string(data), err
}
