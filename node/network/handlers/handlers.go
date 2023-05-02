package handlers

import "github.com/nem0z/dlchat/message"

type Handler func(msg *message.Message) []byte
