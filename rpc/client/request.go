package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/nem0z/dlchat/rpc/types"
)

var NotEnoughArgumentsError = errors.New("Not enough arguments passed to the command")

func CreateReq(args []string) (*bytes.Buffer, error) {
	if len(os.Args) < 2 {
		return nil, NotEnoughArgumentsError
	}

	rpcReq := &types.Request{
		Method: os.Args[1],
		Params: os.Args[2:],
	}

	req, err := json.Marshal(rpcReq)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(req), err
}

func Send(endpoint string, req *bytes.Buffer) (*types.Response, error) {
	resp, err := http.Post(endpoint, "application/json", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(body)

	rpcResp := &types.Response{}
	err = json.Unmarshal(body, rpcResp)

	return rpcResp, err
}
