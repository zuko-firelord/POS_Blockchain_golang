package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type wallet struct{
	privateKey *ecdsa.PrivateKey
	publicKey *ecdsa.PublicKey
}

func NewWallet () *wallet  {
	w := new(wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey
	return w
}

func (w *wallet)  PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey	
}

func (w *wallet) PrivateKeyStr() string  {
	return fmt.Sprintf("%x",w.privateKey.D.Bytes())
}

func (w *wallet)  PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *wallet) PublicKeyStr() string  {
	return fmt.Sprintf("%x%x",w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}
