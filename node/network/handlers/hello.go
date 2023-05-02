package handlers

import (
	"github.com/nem0z/dlchat/message"
)

func Hello(msg *message.Message) bool {
	return msg.IsValid()
}
