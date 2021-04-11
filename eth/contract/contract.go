package contract

import (
	"context"
	"log"
	"time"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/eth/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

var (
	GroupRegistry     *abi.GroupRegistryCaller
	KeeperRegistry    *abi.KeeperRegistryCaller
	ReceiptController *abi.ReceiptControllerCaller

	ChainId int64

	GroupAdded        = make(chan *abi.GroupRegistryGroupAdded)
	WithdrawRequested = make(chan *abi.ReceiptControllerWithdrawRequested)
)

func initContracts(client *ethclient.Client) error {
	var err error
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

	return nil
}

func subscribe(createFn func() (event.Subscription, error)) {
	for {
		sub, err := createFn()
		if err != nil {
			log.Print(err)
		} else {
			err = <-sub.Err()
			sub.Unsubscribe()
			log.Print("subscription error: ", err)
		}
		time.Sleep(time.Second * 10)
	}
}

func subscribeEvents(client *ethclient.Client) {
	go subscribe(func() (event.Subscription, error) {
		filter, err := abi.NewReceiptControllerFilterer(common.HexToAddress(config.C.Contract.ReceiptController), client)
		if err != nil {
			return nil, err
		}
		return filter.WatchWithdrawRequested(nil, WithdrawRequested, nil)
	})
	go subscribe(func() (event.Subscription, error) {
		filter, err := abi.NewGroupRegistryFilterer(common.HexToAddress(config.C.Contract.GroupRegistry), client)
		if err != nil {
			return nil, err
		}
		return filter.WatchGroupAdded(nil, GroupAdded, nil)
	})
}

func Init() error {
	client, err := ethclient.Dial(config.C.EthClient.WssUrl)
	if err != nil {
		return err
	}

	id, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}
	ChainId = id.Int64()
	log.Print("chain id: ", ChainId)

	if err := initContracts(client); err != nil {
		return err
	}
	subscribeEvents(client)
	return nil
}
