package service

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/block-well/decux-keeper-client/config"
	"github.com/block-well/decux-keeper-client/coordinator"
	"github.com/block-well/decux-keeper-client/eth/contract"
	"github.com/block-well/decux-keeper-proto/golang/message"
)

type KeeperService struct {
	opFile      *os.File
	opFileIndex int
}

func NewKeeperService() *KeeperService {
	return &KeeperService{}
}

func (s *KeeperService) Init() error {
	return s.Heartbeat(nil, nil)
}

func (s *KeeperService) Heartbeat(groupNum *uint64, syncMinutes *uint64) error {
	clientInfo := strings.Join([]string{
		config.Version,
		strconv.FormatInt(contract.ChainId, 10),
		config.C.Contract.DecuxSystem,
	}, "|")

	op := &message.Operation{
		Operation: &message.Operation_Heartbeat{
			Heartbeat: &message.Heartbeat{
				ClientInfo:  clientInfo,
				BtcPubKey:   config.C.Keeper.BtcKey.PubKey().SerializeCompressed(),
				Email:       config.C.Keeper.Email,
				GroupNum:    groupNum,
				SyncMinutes: syncMinutes,
			},
		},
	}

	var onlineProof message.OnlineProof
	if err := coordinator.SendOperation(op, &onlineProof); err != nil {
		return err
	}
	if err := s.writeOnlineProof(&onlineProof); err != nil {
		return err
	}

	return nil
}

func (s *KeeperService) writeOnlineProof(onlineProof *message.OnlineProof) error {
	file, err := s.getOpFile()
	if err != nil {
		return err
	}

	parts := []string{
		time.Now().Format(time.RFC3339),
		config.C.Keeper.Id.Hex(),
		strconv.FormatUint(onlineProof.Timestamp, 10),
		string(onlineProof.Signature),
	}
	txt := strings.Join(parts, " ")

	_, err = file.WriteString(txt + "\n")
	return err
}

func (s *KeeperService) getOpFile() (*os.File, error) {
	fileIndex := int(time.Now().Month()) % 2
	if s.opFile != nil && s.opFileIndex == fileIndex {
		return s.opFile, nil
	}

	flag := os.O_RDWR | os.O_CREATE
	if s.opFile == nil {
		flag |= os.O_APPEND
	} else {
		s.opFile.Close()
		s.opFile = nil
		flag |= os.O_TRUNC
	}

	path := config.DataPath("online-proof." + strconv.Itoa(fileIndex))
	file, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		return nil, err
	}

	s.opFile = file
	s.opFileIndex = fileIndex
	return file, nil
}
