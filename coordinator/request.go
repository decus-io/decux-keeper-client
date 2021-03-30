package coordinator

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-proto/golang/message"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

var client = &http.Client{
	Timeout: time.Second * 60,
}

type commonReply struct {
	Success bool
	Message string
	Data    []byte
}

func handleResponse(resp *http.Response, result proto.Message) error {
	var reply commonReply
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return err
	}
	if !reply.Success {
		return errors.New("request error: " + reply.Message)
	}
	if result != nil {
		if err := proto.Unmarshal(reply.Data, result); err != nil {
			return err
		}
	}
	return nil
}

func Reqeust(subUrl string, data []byte, result proto.Message) error {
	var method string
	var body io.Reader
	if data == nil {
		method = "GET"
	} else {
		method = "POST"
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, config.C.Coordinator.Url+subUrl, body)
	if err != nil {
		return err
	}
	if data != nil {
		req.Header.Add("Content-Type", "application/x-protobuf")
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return handleResponse(resp, result)
}

func SendOperation(op *message.Operation, result proto.Message) error {
	op.KeeperId = config.C.Keeper.Id.Hex()
	op.Nonce = 1
	op.Signature = []byte{1}
	op.Timestamp = &timestamp.Timestamp{Seconds: time.Now().Unix()}

	data, err := proto.Marshal(op)
	if err != nil {
		return err
	}
	return Reqeust("operation", data, result)
}
