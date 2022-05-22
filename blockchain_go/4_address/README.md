# Address
相对于之前的版本添加了一些特性

1. 基于私钥（private key）的真实地址
2. 挖矿可以获得 reward
3. 实现 UTXO 集
    1. 获取余额需要扫描整个区块链，而当区块非常多的时候，这么做就会花费很长时间。并且，如果我们想要验证后续交易，也需要花费很长时间。而 UTXO 集就是为了解决这些问题，加快交易相关的操作
4. 实现内存池 mempool
    1. 在交易被打包到块之前，这些交易被存储在内存池里面。在我们目前的实现中，一个块仅仅包含一笔交易，这是相当低效的

比特币用 base58 算法将公钥转化成人类可读的形式

- 将一个公钥转换成一个 Base58 地址需要以下步骤：
    1. 使用 RIPEMD160(SHA256(PubKey)) 哈希算法，取公钥并对其哈希两次
    2. 给哈希加上地址生成算法版本的前缀
    3. 对于第二步生成的结果，使用 SHA256(SHA256(payload)) 再哈希，计算校验和。校验和是结果哈希的前四个字节。
    4. 将校验和附加到 version+PubKeyHash 的组合中。
    5. 使用 Base58 对 version+PubKeyHash+checksum 组合进行编码。



```bash
go run . createwallet
# Your new address: 17VrzAc1R6Lp29cdaWSv9phDcNJUWDMHUT
go run . createwallet
# Your new address: 15ttGySvUpgEh3pVQ87VSyC3FnxiMaCT4E

go run . createblockchain 17VrzAc1R6Lp29cdaWSv9phDcNJUWDMHUT
go run . getbalance 17VrzAc1R6Lp29cdaWSv9phDcNJUWDMHUT
# Balance of '17VrzAc1R6Lp29cdaWSv9phDcNJUWDMHUT': 10
go run . send 17VrzAc1R6Lp29cdaWSv9phDcNJUWDMHUT 15ttGySvUpgEh3pVQ87VSyC3FnxiMaCT4E 6
go run . getbalance 17VrzAc1R6Lp29cdaWSv9phDcNJUWDMHUT
# Balance of '17VrzAc1R6Lp29cdaWSv9phDcNJUWDMHUT': 4
go run . getbalance 15ttGySvUpgEh3pVQ87VSyC3FnxiMaCT4E
# Balance of '17VrzAc1R6Lp29cdaWSv9phDcNJUWDMHUT': 6
```