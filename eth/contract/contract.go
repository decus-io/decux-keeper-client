package contract

import (
	"context"
	"log"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/eth/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var GroupRegistry *abi.GroupRegistryCaller

var KeeperRegistry *abi.KeeperRegistryCaller

var ReceiptController *abi.ReceiptControllerCaller

var ChainID int64

func initContracts() error {
	client, err := ethclient.Dial(config.C.EthClient.HttpUrl)
	if err != nil {
		return err
	}

	GroupRegistry, err = abi.NewGroupRegistryCaller(common.HexToAddress(config.C.Contract.GroupRegistry), client)
	if err != nil {
		return err
	}
	KeeperRegistry, err = abi.NewKeeperRegistryCaller(common.HexToAddress(config.C.Contract.KeeperRegistry), client)
	if err != nil {
		return err
	}
	ReceiptController, err = abi.NewReceiptControllerCaller(
		common.HexToAddress(config.C.Contract.ReceiptController), client)
	if err != nil {
		return err
	}

	id, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}
	ChainID = id.Int64()
	log.Print("chain id: ", ChainID)

	return nil
}

func subscribeEvents() error {
	client, err := ethclient.Dial(config.C.EthClient.WssUrl)
	if err != nil {
		return err
	}

	// client.SubscribeFilterLogs()

	defer client.Close()

	return nil
}

func Init() error {
	if err := initContracts(); err != nil {
		return err
	}
	return subscribeEvents()
}
