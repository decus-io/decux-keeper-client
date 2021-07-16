package config

import (
	"bytes"
	"crypto/ecdsa"
	"io"
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

const (
	userDir = "./user"

	ethKeyFile   = "eth-key"
	btcKeyFile   = "btc-key"
	infuraIdFile = "infura-id"
	emailFile    = "email"
)

func dataPath(user string, file string) string {
	return filepath.Join(userDir, user, file)
}

func save(user string, file string, r io.Reader) error {
	path := dataPath(user, file)
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	return atomic.WriteFile(path, r)
}

func password() string {
	return env.GetNonEmpty("PASSWORD")
}

func saveKey(user string, file string, privateKey *ecdsa.PrivateKey) error {
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
	return save(user, file, bytes.NewReader(data))
}

func SaveEthKey(user string, ethKey *ecdsa.PrivateKey) error {
	return saveKey(user, ethKeyFile, ethKey)
}

func SaveBtcKey(user string, btcKey *btcec.PrivateKey) error {
	return saveKey(user, btcKeyFile, btcKey.ToECDSA())
}

func SaveInfuraId(user string, infuraId string) error {
	return save(user, infuraIdFile, strings.NewReader(infuraId))
}

func SaveEmail(user string, email string) error {
	return save(user, emailFile, strings.NewReader(email))
}

//

func load(user string, file string) ([]byte, error) {
	return ioutil.ReadFile(dataPath(user, file))
}

func loadKey(user string, file string) (*ecdsa.PrivateKey, error) {
	data, err := load(user, file)
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
	return loadKey(user, ethKeyFile)
}

func loadBtcKey(user string) (*btcec.PrivateKey, error) {
	privateKey, err := loadKey(user, btcKeyFile)
	if err != nil {
		return nil, err
	}
	return (*btcec.PrivateKey)(privateKey), nil
}

func loadInfuraId(user string) (string, error) {
	data, err := load(user, infuraIdFile)
	return string(data), err
}

func loadEmail(user string) (string, error) {
	data, err := load(user, emailFile)
	return string(data), err
}
