package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &privateKey.PublicKey
	return w
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

//D is a component of the private key
func (w *Wallet) PrivateKeyString() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

//X and Y are mathematical components of the public key
func (w *Wallet) PublicKeyString() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}