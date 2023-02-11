package blockchain

import (
	"log"
	"time"
)

type BlockChain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func NewBlockchain(difficulty int) *BlockChain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return &BlockChain{
		genesisBlock: genesisBlock,
		chain:        []Block{genesisBlock},
		difficulty:   difficulty,
	}
}

func (b *BlockChain) AddBlock(data TransactionData) {
	lastHash := b.chain[len(b.chain)-1].hash
	block := CreateBlock(data, lastHash, b.difficulty)
	b.chain = append(b.chain, block)
	log.Println("Successfully added block to blockhain")
	time.Sleep(1 * time.Second)
}

func (b BlockChain) IsValid() bool {
	for i, block := range b.chain[1:] {
		prev := b.chain[i-1]
		calculatedHash, err := block.calculateHash()
		if err != nil {
			return false
		}
		if block.previousHash != prev.hash ||
			block.hash != calculatedHash {
			return false
		}
	}
	return true
}
