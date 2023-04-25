package message

import (
	"bytes"
	"encoding/binary"
)

var MagicNum = []byte{0xDE, 0xAD, 0xBE, 0xEF}

type Header struct {
	magic    []byte
	Command  []byte
	Length   uint32
	checksum []byte
}

func (h *Header) From(buf []byte) {
	h.magic = buf[:4]
	h.Command = buf[4:16]
	h.Length = binary.BigEndian.Uint32(buf[16:20])
	h.checksum = buf[20:24]
}

func (h *Header) IsValid() bool {
	return bytes.Equal(MagicNum, h.magic)
}
