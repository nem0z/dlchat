package handlers

import (
	"crypto/sha256"

	"github.com/nem0z/dlchat/message"
	"github.com/nem0z/dlchat/message/messages"
	"github.com/nem0z/dlchat/node/storage"
)

func Send(store *storage.Store) Handler {
	return func(msg *message.Message) []byte {
		if !msg.IsValid() {
			return nil
		}

		chat := messages.DecodeChat(msg.Payload)
		if !chat.Verify() {
			return nil
		}

		hash := sha256.Sum256(msg.Payload)
		store.Put(hash[:], msg.Payload)
		return nil
	}
}
