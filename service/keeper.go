package service

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-proto/golang/message"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Keeper struct {
}

func NewKeeper() *Keeper {
	return &Keeper{}
}

func (*Keeper) Heartbeat() error {
	for {
		op := message.Operation{
			Id:        1,
			Nonce:     1,
			PubKey:    config.C.Keeper.Key.PubKey().SerializeCompressed(),
			Signature: []byte{1},
			Timestamp: &timestamp.Timestamp{
				Seconds: time.Now().Unix(),
			},
			Operation: &message.Operation_Heartbeat{
				Heartbeat: &message.Heartbeat{
					Payload: []byte{2},
				},
			},
		}

		data, err := proto.Marshal(&op)
		if err != nil {
			log.Printf("marshal message error: %v", err)
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", config.C.Coordinator.Url+"operation", bytes.NewReader(data))
		if err != nil {
			log.Printf("new post request error: %v", err)
		}
		req.Header.Add("Content-Type", "application/x-protobuf")
		_, err = client.Do(req)
		if err != nil {
			log.Printf("request error: %v", err)
		}

		time.Sleep(time.Minute)
	}
}
