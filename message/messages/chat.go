package messages

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math"
	"math/big"

	"github.com/nem0z/dlchat/keys"
	"github.com/nem0z/dlchat/message"
)

type chat struct {
	magic  []byte
	sender []byte
	length byte
	data   []byte
	sign   []byte
}

func (chat *chat) Data() []byte {
	return chat.data
}

func Chat(data []byte) *chat {
	if len(data) > math.MaxUint8 {
		return nil
	}

	return &chat{
		magic:  message.MagicNum,
		length: byte(len(data)),
		data:   data,
	}
}

func DecodeChat(data []byte) *chat {
	c := &chat{}
	length := data[68]

	c.magic = data[:4]
	c.sender = data[4:68]
	c.length = length
	c.data = data[69 : length+69]
	c.sign = data[length+69:]

	return c
}

func (c *chat) ToByte() []byte {
	return bytes.Join([][]byte{
		c.magic,
		c.sender,
		{c.length},
		c.data,
		c.sign,
	}, []byte{})
}

func (c *chat) Sign(keys *keys.Keys) error {
	c.sender = keys.PubAddr()

	hash := sha256.Sum256(c.data)
	r, s, err := ecdsa.Sign(rand.Reader, keys.Priv, hash[:])
	if err != nil {
		return err
	}

	c.sign = append(r.Bytes(), s.Bytes()...)
	return nil
}

func (c *chat) Verify() bool {
	xBytes := c.sender[:32]
	yBytes := c.sender[32:]

	pub := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     big.NewInt(0).SetBytes(xBytes),
		Y:     big.NewInt(0).SetBytes(yBytes),
	}

	rBytes := c.sign[:len(c.sign)/2]
	sBytes := c.sign[len(c.sign)/2:]

	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)

	hash := sha256.Sum256(c.data)
	return ecdsa.Verify(pub, hash[:], r, s)
}
