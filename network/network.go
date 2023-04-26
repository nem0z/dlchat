package network

import (
	"fmt"
	"log"
	"net"

	"github.com/nem0z/dlchat/handlers"
)

type peer struct {
	conn net.Conn
	ch   chan bool
}
type Network struct {
	ln     net.Listener
	peers  []*peer
	router map[string]handlers.Handler
}

func Init(port int) (*Network, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		return nil, err
	}

	return &Network{
		listener,
		[]*peer{},
		map[string]handlers.Handler{},
	}, nil
}

func (n *Network) Start() {
	defer n.ln.Close()

	for {
		conn, err := n.ln.Accept()
		if err != nil {
			log.Println("Listener :", err)
			continue
		}

		chStop := n.NewConn(conn)
		go handler(conn, n.router, chStop)
	}
}

func (n *Network) Register(name string, handler handlers.Handler) {
	n.router[name] = handler
}

func (n *Network) NewConn(conn net.Conn) chan bool {
	if len(n.peers) < 5 {
		ch := make(chan bool, 1)
		p := &peer{conn, ch}

		n.peers = append(n.peers, p)
		return ch
	}

	var x *peer
	x, n.peers = n.peers[0], n.peers[:1]
	x.ch <- true

	ch := make(chan bool, 1)
	p := &peer{conn, ch}
	n.peers = append(n.peers, p)
	return ch
}

func handler(conn net.Conn, router map[string]handlers.Handler, stop chan bool) {
	defer conn.Close()
	defer log.Println("Connection closed")

	msg, err := handlers.Default(conn)
	if err != nil || !handlers.Hello(msg) {
		return
	}

	for {
		select {
		case <-stop:
			return
		default:
			msg, err := handlers.Default(conn)
			if err != nil {
				return
			}

			fmt.Println("Message : ", string(msg.Header.Command))

			handler, ok := router[parseFixedSize(msg.Header.Command)]
			if !ok {
				return
			}

			resp := handler(msg)

			if resp != nil {
				conn.Write(resp)
			}
		}
	}
}
