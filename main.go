package main

import (
	"fmt"
	"log"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/decus-io/decus-keeper-client/service"
)

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
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}
