package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	timestamp    int64
	prevHash     string
	transactions []string
}

func newBlock(nonce int, prevHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().Unix()
	b.nonce = nonce
	b.prevHash = prevHash
	return b
}

func PrintBlock(b *Block) {
	fmt.Printf("timestamp: %d\n", b.timestamp)
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("prevHash: %s\n", b.prevHash)
	fmt.Printf("transactions: %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Terry's Blockchain: ")
}

func main() {
	b := newBlock(0, "0")
	PrintBlock(b)
}
