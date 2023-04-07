package service

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/decux-io/decux-keeper-client/config"
	"github.com/decux-io/decux-keeper-client/eth/contract"
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
	lastProgressTime := startTime

	for s.fromBlock < curBlock {
		toBlock := s.fromBlock + 2000 - 1
		if toBlock > curBlock {
			toBlock = curBlock
		}

		if err := s.syncImpl(topics, s.fromBlock, toBlock, handler); err != nil {
			return err
		}

		s.fromBlock = toBlock + 1

		if s.pastEvents && time.Since(lastProgressTime) > time.Second*20 {
			lastProgressTime = time.Now()
			startBlock := config.C.Contract.DeCusSystemStartBlock
			log.Print("sync progress: ", 100*(toBlock-startBlock+1)/(curBlock-startBlock+1), "%")
		}

		if time.Since(startTime) > time.Minute {
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

func (s *EventService) syncImpl(topics [][]common.Hash, fromBlock uint64, toBlock uint64, handler EventHandler) error {
	tryTimes := 1
	if s.pastEvents {
		tryTimes = 3
	}

	for i := 0; i < tryTimes; i++ {
		logs, err := s.filterLogs(topics, fromBlock, toBlock)
		if err != nil {
			if i+1 == tryTimes {
				return err
			}
		} else {
			for _, l := range logs {
				handler(l)
			}

			s.lastSyncTime = time.Now().Unix()
			break
		}
	}

	return nil
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
