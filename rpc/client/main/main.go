package main

import (
	"os"

	"github.com/nem0z/dlchat/rpc/client"
)

func main() {
	client.Process(os.Args)
}
