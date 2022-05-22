package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp     int64          // 当前时间戳，也就是区块创建的时间
	PrevBlockHash []byte         // 前一个 block 的 hash
	Hash          []byte         // 当前 block 的 hash
	Transactions  []*Transaction // 交易信息
	Nonce         int            // 工作量
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Transactions:  transactions,
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	// 保存计算结果
	block.Hash = hash[:]
	block.Nonce = nonce // 存储工作量证明

	return block
}

// 创世区块
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

// HashTransactions returns a hash of the transactions in the block
// 通过仅仅一个哈希，就可以识别一个块里面的所有交易
// 为此，先获得每笔交易的哈希，然后将它们关联起来，最后获得一个连接后的组合哈希
// - 它将一个块里面包含的所有交易表示为一个 Merkle tree ，然后在工作量证明系统中使用树的根哈希（root hash）
// - 这个方法能够让我们快速检索一个块里面是否包含了某笔交易，即只需 root hash 而无需下载所有交易即可完成判断
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}
