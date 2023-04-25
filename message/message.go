package message

import (
	"bytes"
	"encoding/binary"
)

type Message struct {
	Header  *Header
	Payload []byte
}

func New(command []byte, data []byte) *Message {
	formatedCommand := toFixLength(command, 12)
	checksum := checksum(data)

	header := &Header{
		MagicNum,
		formatedCommand,
		uint32(len(data)),
		checksum,
	}

	return &Message{Header: header, Payload: data}
}

func Decode(data []byte) *Message {
	header := &Header{}
	msg := &Message{}
	header.From(data)

	msg.Header = header
	msg.Payload = data[24 : header.Length+24]

	return msg
}

func (m *Message) IsValid() bool {
	checksum := checksum(m.Payload)
	return m.Header.IsValid() && bytes.Equal(checksum, m.Header.checksum)
}

func (m *Message) ToByte() []byte {
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, m.Header.Length)

	return bytes.Join([][]byte{
		m.Header.magic,
		m.Header.Command,
		length,
		m.Header.checksum,
		m.Payload,
	}, []byte{})
}
