package main

import (
	"log"

	"github.com/nem0z/dlchat/node"
	"github.com/nem0z/dlchat/rpc"
)

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	node, err := node.Init(9898, "./database", nil)
	Handle(err)

	rpcServ := rpc.Init(9999)
	rpcServ.Start(node)

	go node.Start()

	select {}
}
