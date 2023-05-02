package handlers

import (
	"github.com/nem0z/dlchat/message"
	"github.com/nem0z/dlchat/node/storage"
)

func Fetch(store *storage.Store) Handler {
	return func(msg *message.Message) []byte {
		if !msg.IsValid() {
			return nil
		}

		data, err := store.Get(msg.Payload)
		if err != nil {
			return nil
		}

		resp := message.New([]byte("send"), data)
		return resp.ToByte()
	}
}
