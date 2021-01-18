package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/decus-io/decus-keeper-proto/golang/message"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	pubKey, _ := hex.DecodeString("026477115981fe981a6918a6297d9803c4dc04f328f22041bedff886bbc2962e01")

	op := message.Operation{
		Id:        1,
		Nonce:     1,
		PubKey:    pubKey,
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
		log.Fatalf("marshal message error: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:5000/operation", bytes.NewReader(data))
	if err != nil {
		log.Fatalf("new post request error: %v", err)
	}
	req.Header.Add("Content-Type", "application/x-protobuf")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("request error: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read response error: %v", err)
	}
	fmt.Printf("response code %d with body %s\n", resp.StatusCode, string(body))
}
