package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 比特币的 golang 实现 btcd 的 BlockHeader 定义
// https://github.com/btcsuite/btcd/blob/01f26a142be8a55b06db04da906163cd9c31be2b/wire/blockheader.go#L20-L41
// // BlockHeader defines information about a block and is used in the bitcoin
// // block (MsgBlock) and headers (MsgHeaders) messages.
// type BlockHeader struct {
//     // Version of the block.  This is not the same as the protocol version.
//     Version int32

//     // Hash of the previous block in the block chain.
//     PrevBlock chainhash.Hash

//     // Merkle tree reference to hash of all transactions for the block.
//     MerkleRoot chainhash.Hash

//     // Time the block was created.  This is, unfortunately, encoded as a
//     // uint32 on the wire and therefore is limited to 2106.
//     Timestamp time.Time

//     // Difficulty target for the block.
//     Bits uint32

//     // Nonce used to generate the block.
//     Nonce uint32
// }

type Block struct {
	// 区块头信息 ------
	Timestamp     int64  // 当前时间戳，也就是区块创建的时间
	PrevBlockHash []byte // 前一个 block 的 hash
	Hash          []byte // 当前 block 的 hash
	// 区块头信息 ------
	Data []byte // 区块存储的实际有效信息，也就是交易
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// SetHash 计算 block 的 hash
func (b *Block) SetHash() {
	// 字符串转字节数组
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// bytes.Join 可以将多个字节串连接，第二个参数是将字节串连接时的分隔符，这里设置为 []byte{} 即为空
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	// 初始化切片，是数组 hash 的引用
	b.Hash = hash[:]
}
