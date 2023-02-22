package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	timestamp    int64
	prevHash     [32]byte
	transactions []*Transaction
}

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

type Transaction struct {
	senderAddress    string
	recipientAddress string
	amount           float32
}

// -----------BLOCK METHODS----------------
func newBlock(nonce int, prevHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().Unix()
	b.nonce = nonce
	b.prevHash = prevHash
	b.transactions = transactions
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp: %d\n", b.timestamp)
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("prevHash: %x\n", b.prevHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := b.MarshalJSON()
	// fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PrevHash     [32]byte       `json:"prevHash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PrevHash:     b.prevHash,
		Transactions: b.transactions,
	})
}

// -----------BLOCKCHAIN METHODS----------------
func (bc *Blockchain) createBlock(nonce int, prevHash [32]byte) *Block {
	b := newBlock(nonce, prevHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	fmt.Printf("%s\n", strings.Repeat("*", 50))
	for i, b := range bc.chain {
		fmt.Printf("%s Block: %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		b.Print()
		fmt.Printf("\n")
	}
	fmt.Printf("%s", strings.Repeat("*", 50))
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, amount float32) {
	t := newTransaction(sender, recipient, amount)
	bc.transactionPool = append(bc.transactionPool, t)
}

func newBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.createBlock(0, (&Block{}).Hash())
	return bc
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

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderAddress    string  `json:"senderAddress"`
		RecipientAddress string  `json:"recipientAddress"`
		Amount           float32 `json:"amount"`
	}{
		SenderAddress:    t.senderAddress,
		RecipientAddress: t.recipientAddress,
		Amount:           t.amount,
	})
}

// -----------MAIN----------------
func init() {
	log.SetPrefix("Terry's Blockchain: ")
}

func main() {
	blockChain := newBlockchain()
	blockChain.AddTransaction("Terry", "John", 100.0)
	blockChain.AddTransaction("Terrysu", "Johnny", 100.0)
	blockChain.createBlock(13, blockChain.LastBlock().Hash())
	blockChain.AddTransaction("Terry", "John", 10.0)
	blockChain.createBlock(14, blockChain.LastBlock().Hash())
	blockChain.Print()
}
