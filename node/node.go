package node

import (
	"github.com/nem0z/dlchat/keys"
	"github.com/nem0z/dlchat/node/network"
	"github.com/nem0z/dlchat/node/network/handlers"
	"github.com/nem0z/dlchat/node/storage"
)

type Node struct {
	Network *network.Network
	Store   *storage.Store
	Keys    *keys.Keys
}

func Init(port int, path string, k *keys.Keys) (*Node, error) {
	network, err := network.Init(port)
	if err != nil {
		return nil, err
	}

	store, err := storage.Init(path)
	if err != nil {
		return nil, err
	}

	if k == nil {
		k, err = keys.Generate()
	}

	network.Register("send", handlers.Send(store))
	network.Register("fetch", handlers.Fetch(store))

	return &Node{
		Network: network,
		Store:   store,
		Keys:    k,
	}, nil
}

func (node *Node) Start() {
	go node.Network.Start()
}
