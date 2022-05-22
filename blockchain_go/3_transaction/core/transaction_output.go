package core

// TXOutput represents a transaction output
type TXOutput struct {
	Value        int    // 一定量的 money
	ScriptPubKey string // 一个锁定脚本, 要花这笔钱, 必须要解锁该脚本
}
