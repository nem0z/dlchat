package network

import (
	"github.com/nem0z/dlchat/message"
)

func (network *Network) Broadcast(msg *message.Message) {
	for _, p := range network.peers {
		p.conn.Write(msg.ToByte())
	}
}
