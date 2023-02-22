package main

import (
	"log"
)

const MINING_DIFFICULTY = 3

func init() {
	log.SetPrefix("Terry's Blockchain: ")
}

func main() {
	blockChain := newBlockchain()
	blockChain.AddTransaction("Terry", "John", 100.0)
	blockChain.AddTransaction("Terrysu", "Johnny", 100.0)
	blockChain.createBlock(blockChain.ProofOfWork(MINING_DIFFICULTY), blockChain.LastBlock().Hash())
	blockChain.AddTransaction("Terry", "John", 10.0)
	blockChain.createBlock(blockChain.ProofOfWork(MINING_DIFFICULTY), blockChain.LastBlock().Hash())
	blockChain.Print()
}
