package main

import (
	"fmt"
	"log"
	"strings"
)

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

// -----------BLOCKCHAIN METHODS----------------
func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
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

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, len(bc.transactionPool))
	for i, t := range bc.transactionPool {
		transactions[i] = newTransaction(t.senderAddress, t.recipientAddress, t.amount)
	}
	return transactions
}

// valid nonce has 3 0s at the beginning
// hash is made from (potential nonce + prevHash + transactions)
// placeholder timestamp as 0
func (bc *Blockchain) ValidNonce(nonce int, prevHash [32]byte, transactions []*Transaction, difficulty int) bool {
	guessBlock := Block{nonce, 0, prevHash, transactions}
	guessHash := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHash[:difficulty] == strings.Repeat("0", difficulty)
}

func (bc *Blockchain) ProofOfWork(difficulty int) int {
	nonce := 0
	prevHash := bc.LastBlock().Hash()
	transactions := bc.CopyTransactionPool()
	for !bc.ValidNonce(nonce, prevHash, transactions, difficulty) {
		nonce++
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork(MINING_DIFFICULTY)
	bc.CreateBlock(nonce, bc.LastBlock().Hash())
	log.Println("Mining successful!")
	return true

}

func newBlockchain(blockchainAddress string) *Blockchain {
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, (&Block{}).Hash())
	return bc
}
