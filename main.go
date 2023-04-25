package main

import (
	"log"

	"github.com/nem0z/dlchat/handlers"
	"github.com/nem0z/dlchat/network"
	"github.com/nem0z/dlchat/storage"
)

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	store, err := storage.Init("./database")
	Handle(err)

	network, err := network.Init(9898)
	Handle(err)

	network.Register("send", handlers.Send(store))
	network.Register("fetch", handlers.Fetch(store))

	go network.Start()
	select {}
}
