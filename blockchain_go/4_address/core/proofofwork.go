package core

import (
	"blockchain/tools"
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

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(), // pow 要把交易考虑进去, 从而保证区块链交易存储的一致性和可靠性
			tools.IntToHex(pow.block.Timestamp),
			tools.IntToHex(int64(targetBits)),
			tools.IntToHex(int64(nonce)),
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

	fmt.Printf("Mining a new block")
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		if math.Remainder(float64(nonce), 100000) == 0 {
			fmt.Printf("\r%x", hash)
		}
		hashInt.SetBytes(hash[:])

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
