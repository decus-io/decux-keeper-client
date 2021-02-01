package config

import (
	"io/ioutil"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/ququzone/go-common/env"
	"gopkg.in/yaml.v2"
)

var C Config

type Config struct {
	Network string `yaml:"network"`
	Keeper  struct {
		Id  uint              `yaml:"id"`
		Key *btcec.PrivateKey `yaml:"-"`
	} `yaml:"keeper"`
	Coordinator struct {
		Url string `yaml:"url"`
	} `yaml:"coordinator"`
}

func Init() error {
	fb, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(fb, &C); err != nil {
		return err
	}

	wif, err := btcutil.DecodeWIF(env.GetNonEmpty("PRIVATE_KEY"))
	if err != nil {
		return err
	}
	C.Keeper.Key = wif.PrivKey

	return nil
}
