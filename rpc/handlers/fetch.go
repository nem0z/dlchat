package handlers

import (
	"encoding/hex"

	"github.com/nem0z/dlchat/message"
	"github.com/nem0z/dlchat/message/messages"
	"github.com/nem0z/dlchat/network"
	"github.com/nem0z/dlchat/rpc/types"
	"github.com/nem0z/dlchat/storage"
)

func Fetch(network *network.Network, store *storage.Store) types.Handler {
	return func(params []string) *types.Response {

		hash, err := hex.DecodeString(params[0])
		msg, err := store.Get(hash)

		if err != nil {
			m := message.New([]byte("fetch"), hash)
			go network.Broadcast(m)
			return &types.Response{Result: "", Err: err.Error()}
		}

		chat := messages.DecodeChat(msg)
		return &types.Response{Result: string(chat.Data()), Err: ""}
	}
}
