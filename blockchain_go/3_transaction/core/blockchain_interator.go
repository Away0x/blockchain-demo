package core

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	// 迭代器的初始状态为链中的 tip
	// 实际上，选择一个 tip 就是意味着给一条链“投票”。一条链可能有多个分支，最长的那条链会被认为是主分支
	// 在获得一个 tip （可以是链中的任意一个块）之后，我们就可以重新构造整条链，找到它的长度和需要构建它的工作。
	// 这同样也意味着，一个 tip 也就是区块链的一种标识符
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// 返回链中的前一个 block, 一直迭代到 genesis block
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}
