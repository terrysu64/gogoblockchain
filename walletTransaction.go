package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"crypto/rand"
	"math/big"
	"fmt"
)

type WalletTransaction struct {
	senderPrivateKey *ecdsa.PrivateKey
	senderPublicKey  *ecdsa.PublicKey
	senderBlockchainAddress string
	recipientBlockchainAddress string
	value float32
}

type Signature struct {
	R *big.Int
	S *big.Int
}

func (w *Wallet) NewWalletTransaction(senderPrivateKey *ecdsa.PrivateKey, senderPublicKey *ecdsa.PublicKey, 
	senderBlockchainAddress string, recipientBlockchainAddress string, value float32) *WalletTransaction {
	wt := WalletTransaction{}
	wt.senderPrivateKey = senderPrivateKey
	wt.senderPublicKey = senderPublicKey
	wt.senderBlockchainAddress = senderBlockchainAddress
	wt.recipientBlockchainAddress = recipientBlockchainAddress
	wt.value = value
	return &wt
}

func (wt *WalletTransaction) GenerateSignature() *Signature{
	m, _ := json.Marshal(wt)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, wt.senderPrivateKey, h[:])
	return &Signature{R: r, S: s}
}

func (wt *WalletTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender string `json:"sender_blockchain_address"`
		Recipient string `json:"recipient_blockchain_address"`
		Value float32 `json:"value"`
	}{
		Sender: wt.senderBlockchainAddress,
		Recipient: wt.recipientBlockchainAddress,
		Value: wt.value,
	})
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}