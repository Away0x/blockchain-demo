package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// 24 指的是算出来的哈希前 24 位必须是 0
// 如果用 16 进制表示，就是前 6 位必须是 0
const targetBits = 24 // 难度值 (比特币中, 这个难度值是动态的, 确保计算始终在 10 分钟左右)

type ProofOfWork struct {
	// 即将生成的区块对象
	block *Block
	// 生成区块的难度，也就是开头有多少个 0
	target *big.Int
}

// data1 := []byte("I like donuts")
// data2 := []byte("I like donutsca07ca")
// targetBits := 24
// target := big.NewInt(1)
// target.Lsh(target, uint(256-targetBits))
// fmt.Printf("%x\n", sha256.Sum256(data1)) // f80867f6efd4484c23b0e7184e53fe4af6ab49b97f5293fcd50d5b2bfa73a4d0
// fmt.Printf("%64x\n", target)             // 0000010000000000000000000000000000000000000000000000000000000000
// fmt.Printf("%x\n", sha256.Sum256(data2)) // 0000002f7c1fe31cb82acdc082cfec47620b7e4ab94f2bf9e096c436fc8cee06
//
// 第一个 hash 比 target 大, 所以不是有效的工作量证明
// 第二个 hash 比 target 小, 所以是有效的证明
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// 左移 256-targetBits 位 (256 是一个 SHA-256 哈希的位数)
	// target 16 进制为: 0x10000000000000000000000000000000000000000000000000000000000
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// pow 算法核心部分
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	// 限制 maxNonce, 是为了防止 nonce 溢出
	for nonce < maxNonce {
		// 1. 准备数据
		data := pow.prepareData(nonce)
		// 2. 对数据进行 hash
		hash = sha256.Sum256(data)
		// 3. hash 转换为大整数
		hashInt.SetBytes(hash[:])
		// 4. 比较
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate 对工作量证明进行验证
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
