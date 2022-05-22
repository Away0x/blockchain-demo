package main

import (
	"time"
)

type Block struct {
	// 区块头信息 ------
	Timestamp     int64  // 当前时间戳，也就是区块创建的时间
	PrevBlockHash []byte // 前一个 block 的 hash
	Hash          []byte // 当前 block 的 hash
	// 区块头信息 ------
	Data  []byte // 区块存储的实际有效信息，也就是交易
	Nonce int    // 工作量
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	// 保存计算结果
	block.Hash = hash[:]
	block.Nonce = nonce // 存储工作量证明

	return block
}
