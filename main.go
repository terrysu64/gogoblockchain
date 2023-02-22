package main

import (
	"log"
)

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
