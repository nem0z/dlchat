package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/nem0z/dlchat/message"
	"github.com/nem0z/dlchat/node/network"
	"github.com/nem0z/dlchat/node/storage"
	"github.com/nem0z/dlchat/rpc/types"
)

func Send(network *network.Network, store *storage.Store) types.Handler {
	return func(params []string) *types.Response {
		msg := message.New([]byte("send"), []byte(params[0]))
		go network.Broadcast(msg)

		b, err := hex.DecodeString(params[0])
		hash := sha256.Sum256(b)

		err = store.Put(hash[:], b)

		resp := &types.Response{}
		if err != nil {
			resp.Result = "Error when storing the message"
			resp.Err = err.Error()
		} else {
			resp.Result = fmt.Sprintf("%x", hash)
		}

		return resp
	}
}
