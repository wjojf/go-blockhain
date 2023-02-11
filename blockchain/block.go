package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	data         TransactionData
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

func (b Block) calculateHash() (string, error) {
	data, err := json.Marshal(b.data)
	if err != nil {
		return "", ErrSerializingData
	}

	blockData := b.previousHash + string(data) + b.timestamp.String()
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%v", blockHash), nil
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		newHash, err := b.calculateHash()
		if err != nil {
			continue
		}
		b.hash = newHash
	}
}

func CreateBlock(data TransactionData, prevHash string, chainDifficulty int) Block {
	block := Block{
		data:         data,
		previousHash: prevHash,
		timestamp:    time.Now(),
	}
	log.Println("Start Mining block")
	block.mine(chainDifficulty)
	return block
}
