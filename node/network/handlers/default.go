package handlers

import (
	"errors"
	"net"

	"github.com/nem0z/dlchat/message"
)

func Default(conn net.Conn) (*message.Message, error) {
	buf := make([]byte, 24)
	n, err := conn.Read(buf)

	if err != nil || n < 24 {
		return nil, errors.New("Header length is incorrect")
	}

	header := &message.Header{}
	header.From(buf)

	if !header.IsValid() {
		return nil, errors.New("Header is invalid")
	}

	buf = make([]byte, header.Length)
	n, err = conn.Read(buf)

	if err != nil || uint32(n) < header.Length {
		return nil, errors.New("Payload length is incorrect")
	}

	msg := &message.Message{Header: header, Payload: buf}
	if !msg.IsValid() {
		return nil, errors.New("Message is invalid")
	}

	return msg, nil
}
