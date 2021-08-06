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
	gethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Client *ethclient.Client

	DeCusSystem    *abi.DeCusSystemCaller
	KeeperRegistry *abi.KeeperRegistryCaller

	ChainId int64
	EventID = map[string]common.Hash{}

	DecusSystemFilterer *abi.DeCusSystemFilterer
)

func initContracts() error {
	var err error
	DeCusSystem, err = abi.NewDeCusSystemCaller(common.HexToAddress(config.C.Contract.DeCusSystem), Client)
	if err != nil {
		return err
	}
	KeeperRegistry, err = abi.NewKeeperRegistryCaller(common.HexToAddress(config.C.Contract.KeeperRegistry), Client)
	if err != nil {
		return err
	}
	DecusSystemFilterer, err = abi.NewDeCusSystemFilterer(common.HexToAddress(config.C.Contract.DeCusSystem), Client)
	if err != nil {
		return err
	}

	parsed, err := gethabi.JSON(strings.NewReader(abi.DeCusSystemABI))
	if err != nil {
		return err
	}
	for name, event := range parsed.Events {
		EventID[name] = event.ID
	}

	return nil
}

func MakeCallOpts() (*bind.CallOpts, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	return &bind.CallOpts{Context: ctx}, cancel
}

func Init() error {
	var err error
	Client, err = ethclient.Dial(config.C.Url.EthClient)
	if err != nil {
		return err
	}

	id, err := Client.NetworkID(context.Background())
	if err != nil {
		return err
	}
	ChainId = id.Int64()
	log.Print("chain id: ", ChainId)

	if err := initContracts(); err != nil {
		return err
	}

	amount, err := KeeperRegistry.GetCollateralWei(nil, config.C.Keeper.Id)
	if err != nil {
		return err
	}
	if amount.Cmp(&big.Int{}) == 0 {
		return errors.New("keeper not exist: " + config.C.Keeper.Id.Hex())
	}

	return nil
}
