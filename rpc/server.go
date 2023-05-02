package rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nem0z/dlchat/node"
	"github.com/nem0z/dlchat/rpc/handlers"
	"github.com/nem0z/dlchat/rpc/types"
)

type RpcServer struct {
	port     string
	handlers map[string]types.Handler
}

func Init(port int) *RpcServer {
	return &RpcServer{
		fmt.Sprintf(":%v", port),
		map[string]types.Handler{},
	}
}

func (rpc *RpcServer) Start(node *node.Node) {
	rpc.Register("sign", handlers.Sign(node.Keys))
	rpc.Register("send", handlers.Send(node.Network, node.Store))
	rpc.Register("fetch", handlers.Fetch(node.Network, node.Store))

	http.HandleFunc("/", rpc.handle())
	go http.ListenAndServe(rpc.port, nil)

	log.Println("RPC api started")
}

func (rpc *RpcServer) Register(method string, handler types.Handler) {
	rpc.handlers[method] = handler
}

func (rpc *RpcServer) handle() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := make([]byte, r.ContentLength)
		_, err := r.Body.Read(body)
		if err != nil && err != io.EOF {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		req := &types.Request{}
		err = json.Unmarshal(body, req)
		if err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}

		handler, ok := rpc.handlers[req.Method]
		if !ok {
			http.Error(w, "Method not supported", http.StatusBadRequest)
			return
		}

		resp := handler(req.Params)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
