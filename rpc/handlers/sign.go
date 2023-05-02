package handlers

import (
	"fmt"

	"github.com/nem0z/dlchat/keys"
	"github.com/nem0z/dlchat/message/messages"
	"github.com/nem0z/dlchat/rpc/types"
)

func Sign(keys *keys.Keys) types.Handler {
	return func(params []string) *types.Response {
		chat := messages.Chat([]byte(params[0]))
		err := chat.Sign(keys)

		if err != nil {
			return &types.Response{Result: "", Err: err.Error()}
		}

		return &types.Response{Result: fmt.Sprintf("%x", chat.ToByte()), Err: ""}
	}
}
