package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
)

type Keys struct {
	Priv *ecdsa.PrivateKey
	Pub  *ecdsa.PublicKey
}

func Generate() (*Keys, error) {
	curve := elliptic.P256()
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	return &Keys{priv, &priv.PublicKey}, nil
}

func Import(path string) (*Keys, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	pemData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	pemBlock, _ := pem.Decode(pemData)
	privKey, err := x509.ParseECPrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return &Keys{privKey, &privKey.PublicKey}, nil

}

func (k *Keys) PubAddr() []byte {
	xBytes := k.Pub.X.Bytes()
	yBytes := k.Pub.Y.Bytes()
	return append(xBytes, yBytes...)
}

func (k *Keys) Export(path string) error {
	der, err := x509.MarshalECPrivateKey(k.Priv)
	if err != nil {
		return err
	}

	pemBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: der,
	}
	pemData := pem.EncodeToMemory(pemBlock)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(pemData)
	return err
}
