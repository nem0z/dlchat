package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/nem0z/dlchat/keys"
	"github.com/nem0z/dlchat/node"
	rpc "github.com/nem0z/dlchat/rpc/api"
)

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	execFile, err := os.Executable()
	Handle(err)
	execDir := filepath.Dir(execFile)
	dbDir := filepath.Join(execDir, "database")

	keys, err := keys.Import("keys.pem")
	if err != nil {
		keys = nil
	}
	keys.Export("keys.pem")
	node, err := node.Init(9898, dbDir, keys)
	Handle(err)

	rpcServ := rpc.Init(9999)
	rpcServ.Start(node)

	go node.Start()

	select {}
}
