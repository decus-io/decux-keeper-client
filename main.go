package main

import (
	"flag"
	"fmt"
	"log"
	"net/mail"
	"syscall"

	"github.com/btcsuite/btcutil"
	"github.com/decux-io/decux-keeper-client/config"
	"github.com/decux-io/decux-keeper-client/eth/contract"
	"github.com/decux-io/decux-keeper-client/service"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/term"
)

func initSetting() error {
	var data []byte
	var err error

	fmt.Println("Paste ETH private key:")
	if data, err = term.ReadPassword(int(syscall.Stdin)); err != nil {
		return err
	}
	ethKey, err := crypto.HexToECDSA(string(data))
	if err != nil {
		return err
	}
	if err := config.SaveEthKey(ethKey); err != nil {
		return err
	}

	//
	fmt.Println("Paste BTC private key (WIF):")
	if data, err = term.ReadPassword(int(syscall.Stdin)); err != nil {
		return err
	}
	btcKey, err := btcutil.DecodeWIF(string(data))
	if err != nil {
		return err
	}
	if err := config.SaveBtcKey(btcKey.PrivKey); err != nil {
		return err
	}

	//
	fmt.Println("Email to receive notification when the client is offline:")
	str := ""
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

func run() error {
	if err := config.Init(); err != nil {
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

	config.C.Keeper.User = *user

	fmt.Println("Enter password for private keys:")
	pwd, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err.Error())
	}
	config.C.Keeper.Password = string(pwd)

	if !config.UserSettingReady() {
		err = initSetting()
	}
	if err == nil {
		err = run()
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
