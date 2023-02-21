package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	timestamp    int64
	prevHash     string
	transactions []string
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func newBlock(nonce int, prevHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().Unix()
	b.nonce = nonce
	b.prevHash = prevHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp: %d\n", b.timestamp)
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("prevHash: %s\n", b.prevHash)
	fmt.Printf("transactions: %s\n", b.transactions)
}

// a method associated with the Blockchain struct
func (bc *Blockchain) createBlock(nonce int, prevHash string) *Block {
	b := newBlock(nonce, prevHash)
	bc.chain = append(bc.chain, b)
	return b
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
	bc.createBlock(0, "first hash")
	return bc
}

func init() {
	log.SetPrefix("Terry's Blockchain: ")
}

func main() {
	blockChain := newBlockchain()
	blockChain.createBlock(13, "hash1")
	blockChain.Print()
}
