package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type Keys struct {
	Priv *ecdsa.PrivateKey
	Pub  *ecdsa.PublicKey
}

func GenerateKeys() (*Keys, error) {
	curve := elliptic.P256()
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	return &Keys{priv, &priv.PublicKey}, nil
}
