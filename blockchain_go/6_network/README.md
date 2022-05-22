# Network
区块链网络是去中心化的，这意味着没有服务器，客户端也不需要依赖服务器来获取或处理数据。在区块链网络中，有的是节点，每个节点是网络的一个完全（full-fledged）成员。节点就是一切：它既是一个客户端，也是一个服务器。这一点需要牢记于心，因为这与传统的网页应用非常不同

## 操作流程
1. 中心节点创建一个区块链
2. 一个其他（钱包）节点连接到中心节点并下载区块链
3. 另一个（矿工）节点连接到中心节点并下载区块链
4. 钱包节点创建一笔交易
5. 矿工节点接收交易，并将交易保存到内存池中
6. 当内存池中有足够的交易时，矿工开始挖一个新块
7. 当挖出一个新块后，将其发送到中心节点
8. 钱包节点与中心节点进行同步
9. 钱包节点的用户检查他们的支付是否成功

```bash
# Node1 3000: 中心节点
# Node2 3001: 钱包节点
# Node3 3002: 矿工节点

# Create blockchain
go run . -n 3000 createwallet
# Your new address: 12UaYNx7aPNenFWNZ4Gj9dhQkknoqZzNEi
# 生成一个仅包含创世块的区块链，并在其他节点使用。创世块承担了一条链标识符的角色
go run . -n 3000 createblockchain 12UaYNx7aPNenFWNZ4Gj9dhQkknoqZzNEi
# 000000954fb6c5c2f3b865bce6359a5abcb46729aac1d80ef982bed5e0f97f72 (创世区块 hash)
# 将数据库 copy 到其他节点
cp blockchain_3000.db blockchain_3001.db
cp blockchain_3000.db blockchain_3002.db

# Create wallet
go run . -n 3001 createwallet
# Your new address: 1HYjvuDiBavjSWXJCFJ7hPFgD6XR18nEht (W1)
go run . -n 3001 createwallet
# Your new address: 16HjApaYskVmzvnYgUxY5ANw3VJD6CyfkX (W2)
go run . -n 3001 createwallet
# Your new address: 1CwJzuxwgxvz6f5K6bbtXFgxLbAanASKuz (W3)

# Transaction
# Node1 -> Node2 W1
go run . -n 3000 send 12UaYNx7aPNenFWNZ4Gj9dhQkknoqZzNEi 1HYjvuDiBavjSWXJCFJ7hPFgD6XR18nEht 10 -m
# 00000094a3cc700bf26a21f470741991e74bdac9a687ef61584c0f9a66b47936
# Node1 -> Node2 W2
go run . -n 3000 send 12UaYNx7aPNenFWNZ4Gj9dhQkknoqZzNEi 16HjApaYskVmzvnYgUxY5ANw3VJD6CyfkX 10 -m
# 000000c487db2295684eeefd191ba448373d43f9f2665df3a07d652513c07c2d

# Start node
go run . -n 3000 startnode
go run . -n 3001 startnode # 会从中心节点下载所有区块

# Check balance (需要先暂停 3001 节点的 server)
go run . -n 3001 getbalance 1HYjvuDiBavjSWXJCFJ7hPFgD6XR18nEht
# Balance of '1HYjvuDiBavjSWXJCFJ7hPFgD6XR18nEht': 10
go run . -n 3001 getbalance 16HjApaYskVmzvnYgUxY5ANw3VJD6CyfkX
# Balance of '16HjApaYskVmzvnYgUxY5ANw3VJD6CyfkX': 10
```
```bash
# Node3 矿工节点生成钱包地址
go run . -n 3002 createwallet
# Your new address: 16Be7PjxVxwC4iQ9ZsvorV8FY8y3FkNvmb

# Start node, 将钱包地址，作为矿工钱包
go run . -n 3002 startnode -m=16Be7PjxVxwC4iQ9ZsvorV8FY8y3FkNvmb
# Mining is on. Address to receive rewards:  16Be7PjxVxwC4iQ9ZsvorV8FY8y3FkNvmb

# 在 Node2 的钱包节点发送一些币
# W1 -> W3
go run . -n 3001 send 1HYjvuDiBavjSWXJCFJ7hPFgD6XR18nEht 1CwJzuxwgxvz6f5K6bbtXFgxLbAanASKuz 2
# W2 -> W3
go run . -n 3001 send 16HjApaYskVmzvnYgUxY5ANw3VJD6CyfkX 1CwJzuxwgxvz6f5K6bbtXFgxLbAanASKuz 3

# 可以在 Node3 矿工节点的命令行看到, 挖出了一个区块, 输出如下
# Mining is on. Address to receive rewards:  16Be7PjxVxwC4iQ9ZsvorV8FY8y3FkNvmb
# Received inv command
# Recevied inventory with 1 tx
# Received tx command
# Received inv command
# Recevied inventory with 1 tx
# Received tx command
# 000000c9266d5b77ce857fd994c8c3d3a00249ab5dd2de9c69744ee2dc8a2f42

# New block is mined!
# Received getdata command

# 启动 Node2 节点, 它会下载最近挖出来的 block (同步数据库)
go run . -n 3001 startnode
# 暂停 Node2, Check balance
go run . -n 3001 getbalance 1HYjvuDiBavjSWXJCFJ7hPFgD6XR18nEht
# Balance of '1HYjvuDiBavjSWXJCFJ7hPFgD6XR18nEht': 10 (W1)
go run . -n 3001 getbalance 16HjApaYskVmzvnYgUxY5ANw3VJD6CyfkX
# Balance of '16HjApaYskVmzvnYgUxY5ANw3VJD6CyfkX': 10 (W2)
go run . -n 3001 getbalance 1CwJzuxwgxvz6f5K6bbtXFgxLbAanASKuz
# Balance of '1CwJzuxwgxvz6f5K6bbtXFgxLbAanASKuz': 0 (W3)
go run . -n 3001 getbalance 16Be7PjxVxwC4iQ9ZsvorV8FY8y3FkNvmb
# Balance of '16Be7PjxVxwC4iQ9ZsvorV8FY8y3FkNvmb': 10 (矿工节点)
```

## 还未实现的地方
需要手动同步数据, 没有实现 P2P 网络, 节点不能互相发现彼此