package main

import (
	"fmt"
	"strings"
)

type Transaction struct {
	senderAddress    string
	recipientAddress string
	amount           float32
}

// -----------TRANSACTION METHODS----------------
func newTransaction(sender string, recipient string, amount float32) *Transaction {
	t := new(Transaction)
	t.senderAddress = sender
	t.recipientAddress = recipient
	t.amount = amount
	return t
}

func (t *Transaction) Print() {
	fmt.Printf("%s Block Transactions %s\n", strings.Repeat("-", 25), strings.Repeat("-", 25))
	fmt.Printf("senderAddress: %s\n", t.senderAddress)
	fmt.Printf("recipientAddress: %s\n", t.recipientAddress)
	fmt.Printf("amount: %f\n", t.amount)
}

//	func (t *Transaction) MarshalJSON() ([]byte, error) {
//		return json.Marshal(struct {
//			SenderAddress    string  `json:"senderAddress"`
//			RecipientAddress string  `json:"recipientAddress"`
//			Amount           float32 `json:"amount"`
//		}{
//			SenderAddress:    t.senderAddress,
//			RecipientAddress: t.recipientAddress,
//			Amount:           t.amount,
//		})
//	}
//
