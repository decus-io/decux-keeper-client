package contract

import (
	"context"
	"errors"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/eth/abi"
	"github.com/ethereum/go-ethereum"
	gethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

var (
	DeCusSystem    *abi.DeCusSystemCaller
	KeeperRegistry *abi.KeeperRegistryCaller

	ChainId int64

	GroupAdded    = make(chan *abi.DeCusSystemGroupAdded)
	GroupDeleted  = make(chan *abi.DeCusSystemGroupDeleted)
	BurnRequested = make(chan *abi.DeCusSystemBurnRequested)
	BtcRefunded   = make(chan *abi.DeCusSystemBtcRefunded)

	decusSystemFilterer *abi.DeCusSystemFilterer
	groupAddedId        common.Hash
	groupDeletedId      common.Hash
)

func initContracts(client *ethclient.Client) error {
	var err error
	DeCusSystem, err = abi.NewDeCusSystemCaller(common.HexToAddress(config.C.Contract.DeCusSystem), client)
	if err != nil {
		return err
	}
	KeeperRegistry, err = abi.NewKeeperRegistryCaller(common.HexToAddress(config.C.Contract.KeeperRegistry), client)
	if err != nil {
		return err
	}
	decusSystemFilterer, err = abi.NewDeCusSystemFilterer(common.HexToAddress(config.C.Contract.DeCusSystem), client)
	if err != nil {
		return err
	}

	return nil
}

func subscribeLoop(createFn func() (event.Subscription, error)) {
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
	go subscribeLoop(func() (event.Subscription, error) {
		return decusSystemFilterer.WatchBurnRequested(nil, BurnRequested, nil)
	})
	go subscribeLoop(func() (event.Subscription, error) {
		return decusSystemFilterer.WatchBtcRefunded(nil, BtcRefunded)
	})
}

func parseLogs(logs []types.Log) error {
	for _, l := range logs {
		switch l.Topics[0] {
		case groupAddedId:
			event, err := decusSystemFilterer.ParseGroupAdded(l)
			if err != nil {
				return err
			}
			GroupAdded <- event
		case groupDeletedId:
			event, err := decusSystemFilterer.ParseGroupDeleted(l)
			if err != nil {
				return err
			}
			GroupDeleted <- event
		}
	}

	return nil
}

func filterLoop(client *ethclient.Client, topics [][]common.Hash) {
	var startBlock uint64

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		curBlock, err := client.BlockNumber(ctx)

		if err != nil {
			log.Print("BlockNumber err: ", err)
		} else if curBlock >= startBlock {
			query := ethereum.FilterQuery{
				Addresses: []common.Address{common.HexToAddress(config.C.Contract.DeCusSystem)},
				Topics:    topics,
				FromBlock: new(big.Int).SetUint64(startBlock),
				ToBlock:   new(big.Int).SetUint64(curBlock),
			}

			logs, err := client.FilterLogs(ctx, query)
			if err == nil {
				err = parseLogs(logs)
			}
			if err != nil {
				log.Print("FilterLogs err: ", err)
			} else {
				startBlock = curBlock + 1
			}
		}

		cancel()
		time.Sleep(time.Minute * 5)
	}
}

func filterEvents(client *ethclient.Client) error {
	parsed, err := gethabi.JSON(strings.NewReader(abi.DeCusSystemABI))
	if err != nil {
		return err
	}

	groupAddedId = parsed.Events["GroupAdded"].ID
	groupDeletedId = parsed.Events["GroupDeleted"].ID
	query := [][]interface{}{{groupAddedId, groupDeletedId}}
	topics, err := gethabi.MakeTopics(query...)
	if err != nil {
		return err
	}

	go filterLoop(client, topics)
	return nil
}

func Init() error {
	client, err := ethclient.Dial(config.C.Url.EthClient)
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

	amount, err := KeeperRegistry.GetCollateralValue(nil, config.C.Keeper.Id)
	if err != nil {
		return err
	}
	if amount.Cmp(&big.Int{}) == 0 {
		return errors.New("keeper not exist: " + config.C.Keeper.Id.Hex())
	}

	if err := filterEvents(client); err != nil {
		return err
	}
	subscribeEvents(client)

	return nil
}
