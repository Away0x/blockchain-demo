package main

import "fmt"

func main() {
	bc := NewBlockchain()

	// 真正的区块链系统中, 添加一个块需要做很多工作
	// 1. 工作量证明 POW: 通过复杂的计算来获取添加 new block 的权利
	// 2. 共识 Consensus: 区块链是分布式的, 加入 new block 需要被网络的其他参与者确认和同意
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
