package main

import (
	"fmt"
	"log"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

func init() {
	log.SetPrefix("Terry's Blockchain: ")
}

func main() {

	w := NewWallet()
	fmt.Println(w.PrivateKeyString())
	fmt.Println(w.PublicKeyString())
	fmt.Println(w.BlockchainAddress())

	// blockChain := newBlockchain("tempblockchainaddress")
	// blockChain.AddTransaction("Terry", "John", 100.0)
	// blockChain.AddTransaction("Terrysu", "Johnny", 100.0)
	// blockChain.Mining()
	// blockChain.AddTransaction("Terry", "John", 10.0)
	// blockChain.Mining()
	// blockChain.Print()
	// fmt.Println("\n", blockChain.TotalBalance("Terry"))
	// fmt.Println("\n", blockChain.TotalBalance("tempblockchainaddress")) //the miner
}
