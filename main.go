package main

import (
	"flag"
	"fmt"
	"log"
	"net/mail"
	"net/url"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/decux-io/decux-keeper-client/config"
	"github.com/decux-io/decux-keeper-client/eth/contract"
	"github.com/decux-io/decux-keeper-client/service"
	"github.com/ethereum/go-ethereum/crypto"
)

func initSetting() error {
	var str string

	fmt.Println("Paste network RPC URL:")
	if _, err := fmt.Scanln(&str); err != nil {
		return err
	}
	if _, err := url.ParseRequestURI(str); err != nil {
		return err
	}
	if err := config.SaveRpcUrl(str); err != nil {
		return err
	}

	//
	fmt.Println("Paste ETH private key:")
	if _, err := fmt.Scanln(&str); err != nil {
		return err
	}
	ethKey, err := crypto.HexToECDSA(str)
	if err != nil {
		return err
	}
	if err := config.SaveEthKey(ethKey); err != nil {
		return err
	}

	//
	fmt.Println("Paste BTC private key (WIF):")
	if _, err := fmt.Scanln(&str); err != nil {
		return err
	}
	btcKey, err := btcutil.DecodeWIF(str)
	if err != nil {
		return err
	}
	if err := config.SaveBtcKey(btcKey.PrivKey); err != nil {
		return err
	}

	//
	fmt.Println("Email to receive notification when the client is offline:")
	str = ""
	if _, err := fmt.Scanln(&str); err != nil && err.Error() != "unexpected newline" {
		return err
	}
	if len(str) > 0 {
		if _, err := mail.ParseAddress(str); err != nil {
			return err
		}
	}
	if err := config.SaveEmail(str); err != nil {
		return err
	}

	log.Print("init setting ok")
	return nil
}

func run(user string) error {
	if err := config.Init(user); err != nil {
		return fmt.Errorf("init config error: %v", err)
	}
	if err := contract.Init(); err != nil {
		return fmt.Errorf("init contract error: %v", err)
	}

	system := service.NewSystem()
	if err := system.Init(); err != nil {
		return err
	}

	log.Print("init ok")

	system.Run()
	return nil
}

func main() {
	user := flag.String("user", "default", "user name")
	flag.Parse()

	var err error
	if flag.NArg() > 0 {
		if flag.Arg(0) == "init" {
			config.C.Keeper.User = *user
			err = initSetting()
		} else {
			err = fmt.Errorf("unkown arguments: %v", flag.Args())
		}
	} else {
		err = run(*user)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
