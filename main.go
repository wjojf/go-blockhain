package main

import (
	"fmt"
	"go-blockchain/blockchain"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	b := blockchain.NewBlockchain(1)

	var i int
	for i <= 3 {
		transaction := blockchain.NewTransaction(
			fmt.Sprintf("address_%v", i),
			fmt.Sprintf("address_%v", i),
			int64(i*100),
		)
		wg.Add(1)
		go func() {
			b.AddBlock(transaction)
			wg.Done()
		}()
		i++
	}
	fmt.Println("Blocks mining...")
	wg.Wait()
	fmt.Println("All blocks mined")
}
