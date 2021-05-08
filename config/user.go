package config

import (
	"bytes"
	"crypto/ecdsa"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/natefinch/atomic"
	"github.com/ququzone/go-common/env"
)

const userDir = "./user"

func ethKeyFile(user string) string {
	return filepath.Join(userDir, user, "eth-key")
}

func btcKeyFile(user string) string {
	return filepath.Join(userDir, user, "btc-key")
}

func infuraIdFile() string {
	return filepath.Join(userDir, "infura-id")
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
	if err := os.MkdirAll(filepath.Dir(file), 0700); err != nil {
		return err
	}
	return atomic.WriteFile(file, bytes.NewReader(data))
}

func SaveEthKey(user string, ethKey *ecdsa.PrivateKey) error {
	return saveKey(ethKeyFile(user), ethKey)
}

func SaveBtcKey(user string, btcKey *btcec.PrivateKey) error {
	return saveKey(btcKeyFile(user), btcKey.ToECDSA())
}

func SaveInfuraId(infuraId string) error {
	file := infuraIdFile()
	if err := os.MkdirAll(filepath.Dir(file), 0700); err != nil {
		return err
	}
	return atomic.WriteFile(file, strings.NewReader(infuraId))
}

func loadKey(file string) (*ecdsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	ksKey, err := keystore.DecryptKey(data, password())
	if err != nil {
		return nil, err
	}

	return ksKey.PrivateKey, nil
}

func loadEthKey(user string) (*ecdsa.PrivateKey, error) {
	return loadKey(ethKeyFile(user))
}

func loadBtcKey(user string) (*btcec.PrivateKey, error) {
	privateKey, err := loadKey(btcKeyFile(user))
	if err != nil {
		return nil, err
	}
	return (*btcec.PrivateKey)(privateKey), nil
}

func loadInfuraId() (string, error) {
	data, err := ioutil.ReadFile(infuraIdFile())
	if err != nil {
		return "", err
	}
	return string(data), nil
}
