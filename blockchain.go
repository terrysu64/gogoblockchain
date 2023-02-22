package main

import (
	"fmt"
	"strings"
)

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
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
