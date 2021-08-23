package service

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/eth/contract"
	"github.com/ethereum/go-ethereum"
	gethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EventHandler func(types.Log)

type EventService struct {
	pastEvents   bool
	fromBlock    uint64
	lastSyncTime int64
}

func NewEventService() *EventService {
	return &EventService{
		pastEvents: true,
		fromBlock:  config.C.Contract.DeCusSystemStartBlock,
	}
}

func (s *EventService) Sync(handler EventHandler) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	curBlock, err := contract.Client.BlockNumber(ctx)
	if err != nil {
		return err
	}
	topics, err := s.getTopics()
	if err != nil {
		return err
	}

	startTime := time.Now()

	for s.fromBlock < curBlock {
		toBlock := s.fromBlock + 20000 - 1
		if toBlock > curBlock {
			toBlock = curBlock
		}

		logs, err := s.filterLogs(topics, s.fromBlock, toBlock)
		if err != nil {
			return err
		}
		for _, l := range logs {
			handler(l)
		}

		s.fromBlock = toBlock + 1
		s.lastSyncTime = time.Now().Unix()

		if time.Since(startTime) > time.Minute*2 {
			return nil
		}
	}

	if s.pastEvents {
		s.pastEvents = false
		log.Print("past events finished")
	}
	return nil
}

func (s *EventService) SyncMinutes() *uint64 {
	if s.lastSyncTime == 0 {
		return nil
	}

	now := time.Now().Unix()
	if now < s.lastSyncTime {
		now = s.lastSyncTime
	}

	syncMinutes := uint64((now - s.lastSyncTime) / 60)
	return &syncMinutes
}

func (s *EventService) filterLogs(topics [][]common.Hash, fromBlock uint64, toBlock uint64) ([]types.Log, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(config.C.Contract.DeCusSystem)},
		Topics:    topics,
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
	}

	return contract.Client.FilterLogs(ctx, query)
}

func (s *EventService) getTopics() ([][]common.Hash, error) {
	if s.pastEvents {
		query := [][]interface{}{{contract.EventID["GroupAdded"], contract.EventID["GroupDeleted"]}}
		return gethabi.MakeTopics(query...)
	} else {
		return nil, nil
	}
}
