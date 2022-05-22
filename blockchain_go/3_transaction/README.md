# 持久化存储
Bitcoin 将每个 block 存储为不同文件, 这样就不需要为了读取一个单一的块而将所有（或者部分）的 block 都加载到内存中, 本例为了简单, 不考虑这一点 

## 存储结构
> 使用 boltdb, 值只能是 []byte, 例子中使用 gob 进行序列化

1. **32 字节的 block-hash**: block 结构
2. **l**: 链中最后一个块的 hash

***

# Transaction
区块链中, 交易一旦被创建, 就没有任何人能够再去修改或是删除它

比特币采用的是 UTXO 模型，并非账户模型，并不直接存在“余额”这个概念，余额需要通过遍历整个交易历史得来

输出, 就是 “币” 存储的地方。每个输出都会带有一个解锁脚本, 这个脚本定义了解锁该输出的逻辑。
每笔新的交易, 必须至少有一个输入和输出。一个输入引用了之前一笔交易的输出, 并提供了解锁数据（也就是 ScriptSig 字段）,
该数据会被用在输出的解锁脚本中解锁输出, 解锁完成后即可使用它的值去产生新的输出

## coinbase 交易
当矿工挖出一个新的块时，它会向新的块中添加一个 coinbase 交易。coinbase 交易是一种特殊的交易，它不需要引用之前一笔交易的输出。它“凭空”产生了币（也就是产生了新币），这是矿工获得挖出新块的奖励，也可以理解为“发行新币”

```bash
go run . createblockchain User1
go run . getbalance User1 # Balance of 'User1': 10

# transaction
go run . send User1 User2 6
go run . getbalance User1 # Balance of 'User1': 4
go run . getbalance User2 # Balance of 'User2': 6
go run . send User1 User2 6 # ERROR: Not enough fund
```