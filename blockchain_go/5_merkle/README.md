# 实现有 Merkle 树的交易
之前的 FindUnspentTransactions 会找到有未花费输出的交易。由于交易被保存在区块中，所以它会对区块链里面的每一个区块进行迭代，检查里面的每一笔交易

解决方案是有一个仅有未花费输出的索引，这就是 UTXO 集要做的事情：这是一个从所有区块链交易中构建（对区块进行迭代，但是只须做一次）而来的缓存，然后用它来计算余额和验证新的交易, 其比整个区块链数据库小得多

Merkle 树的好处就是一个节点可以在不下载整个块的情况下，验证是否包含某笔交易。并且这些只需要一个交易哈希，一个 Merkle 树根哈希和一个 Merkle 路径

## 旧的交易版本使用到的方法
> 所有方法都对数据库中的所有块进行迭代

1. Blockchain.FindUnspentTransactions: 找到有未花费输出交易的主要函数。也是在这个函数里面会对所有区块进行迭代。
2. Blockchain.FindSpendableOutputs: 这个函数用于当一个新的交易创建的时候。如果找到有所需数量的输出。使用 Blockchain.FindUnspentTransactions.
3. Blockchain.FindUTXO: 找到一个公钥哈希的未花费输出，然后用来获取余额。使用 Blockchain.FindUnspentTransactions.
4. Blockchain.FindTransation: 根据 ID 在区块链中找到一笔交易。它会在所有块上进行迭代直到找到它

## 新实现的方法
1. Blockchain.FindUTXO: 通过对区块进行迭代找到所有未花费输出。
2. UTXOSet.Reindex: 使用 UTXO 找到未花费输出，然后在数据库中进行存储。这里就是缓存的地方。
3. UTXOSet.FindSpendableOutputs: 类似 Blockchain.FindSpendableOutputs，但是使用 UTXO 集。
4. UTXOSet.FindUTXO: 类似 Blockchain.FindUTXO，但是使用 UTXO 集。
5. Blockchain.FindTransaction: 跟之前一样


```bash
# create wallet
go run . createwallet
# Your new address: 19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF
go run . createwallet
# Your new address: 1MJUWbrKsvUFXe4jZGpWFuePBuW8zZbdwc
go run . createwallet
# Your new address: 19mM64T2NiT9gVTiy6Q1KRPHGrT6AXhiu2

# create blockchain
go run . createblockchain 19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF
# 000000092ddf4ab1f91c3271b7f1be37718e31137a47b4e172237eca9f1f44ee
go run . getbalance 19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF
# Balance of '19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF': 10 (挖出创世块, 每笔奖励为 10)

# transation
go run . send 19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF 1MJUWbrKsvUFXe4jZGpWFuePBuW8zZbdwc 6
# 0000004a5d9f54d055732d47a81c91dda9480bd0e1bb42d64e523143db212e85
go run . getbalance 19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF
# Balance of '19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF': 14 (10 - 6 + 10)
go run . send 19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF 19mM64T2NiT9gVTiy6Q1KRPHGrT6AXhiu2 4
# 000000d4087fda6dce7889465d9fbf2a7f33692157afa92a6e2eafdd162c750b

# get balance
go run . getbalance 19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF
# Balance of '19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF': 20 (10 - 6 + 10 - 4 + 10)
go run . getbalance 1MJUWbrKsvUFXe4jZGpWFuePBuW8zZbdwc
# Balance of '1MJUWbrKsvUFXe4jZGpWFuePBuW8zZbdwc': 6
go run . getbalance 19mM64T2NiT9gVTiy6Q1KRPHGrT6AXhiu2
# Balance of '19mM64T2NiT9gVTiy6Q1KRPHGrT6AXhiu2': 4

# 19WFRA3QD42ZNctsB4eamXfPe49W5TiZvF 收到了三笔奖励
# 1. 挖出创世块
# 2. 挖出块 0000004a5d9f54d055732d47a81c91dda9480bd0e1bb42d64e523143db212e85
# 3. 挖出块 000000d4087fda6dce7889465d9fbf2a7f33692157afa92a6e2eafdd162c750b
```