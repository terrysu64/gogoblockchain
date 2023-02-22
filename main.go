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
	transactions []string
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// -----------BLOCK METHODS----------------
func newBlock(nonce int, prevHash [32]byte) *Block {
	b := new(Block)
	b.timestamp = time.Now().Unix()
	b.nonce = nonce
	b.prevHash = prevHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp: %d\n", b.timestamp)
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("prevHash: %x\n", b.prevHash)
	fmt.Printf("transactions: %s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := b.MarshalJSON()
	// fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PrevHash     [32]byte `json:"prevHash"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PrevHash:     b.prevHash,
		Transactions: b.transactions,
	})
}

// -----------BLOCKCHAIN METHODS----------------
func (bc *Blockchain) createBlock(nonce int, prevHash [32]byte) *Block {
	b := newBlock(nonce, prevHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	fmt.Printf("%s", strings.Repeat("*", 50))
	for i, b := range bc.chain {
		fmt.Printf("%s Block: %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		b.Print()
		fmt.Printf("\n")
	}
	fmt.Printf("%s", strings.Repeat("*", 50))
}

func newBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.createBlock(0, (&Block{}).Hash())
	return bc
}

// -----------MAIN----------------
func init() {
	log.SetPrefix("Terry's Blockchain: ")
}

func main() {
	blockChain := newBlockchain()
	blockChain.createBlock(13, blockChain.LastBlock().Hash())
	blockChain.createBlock(14, blockChain.LastBlock().Hash())
	blockChain.Print()
}
