package coordinator

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/eth"
	"github.com/decus-io/decus-keeper-proto/golang/message"
	"google.golang.org/protobuf/proto"
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

	req, err := http.NewRequest(method, config.C.Url.Coordinator+subUrl, body)
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
	opData, err := proto.Marshal(op)
	if err != nil {
		return err
	}
	signature, err := eth.SignMessage(opData, true)
	if err != nil {
		return err
	}

	req := &message.Request{
		Data:      opData,
		Signature: signature,
	}
	reqData, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	return Reqeust("operation", reqData, result)
}
