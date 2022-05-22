package core

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte // 存储的是之前交易的 ID
	Vout      int    // 存储的是该输出在那笔交易中所有输出的索引 (因为一笔交易可能有多个输出, 需要有信息指明是具体的哪一个)
	ScriptSig string // 脚本, 提供了可解锁 Output 结构里面 ScriptPubKey 字段的数据
	// 如果 ScriptSig 提供的数据是正确的, 那么输出就会被解锁, 然后被解锁的值就可以被用于产生新的输出
	// 如果数据不正确, 输出就无法被引用在输入中, 或者说, 无法使用这个输出
	// 这种机制, 保证了用户无法花费属于其他人的币
	// 由于现在没有实现 address, 所以目前 ScriptSig 将仅仅存储一个用户自定义的任意钱包地址
	// 后面会拓展为 public key 和 signature
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}
