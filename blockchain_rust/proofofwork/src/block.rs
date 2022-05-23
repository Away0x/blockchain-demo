use chrono::Utc;

use crate::proofofwork::ProofOfWork;

#[derive(Debug, Clone)]
pub struct Block {
    timestamp: i64,          // 区块创建的时间
    prev_block_hash: String, // 前一个 block 的 hash
    hash: String,            // 当前 block 的 hash
    data: String,            // 区块存储的实际有效信息，也就是交易
    nonce: i64,              // 工作量
}

impl Block {
    pub fn new(data: String, prev_block_hash: String) -> Self {
        let mut block = Self {
            timestamp: Utc::now().timestamp(),
            data,
            prev_block_hash,
            hash: String::new(),
            nonce: 0,
        };

        let pow = ProofOfWork::new(block.clone());
        let (nonce, hash) = pow.run();
        // 保存计算结果
        block.nonce = nonce;
        block.hash = hash;

        block
    }

    /// 创世区块
    pub fn new_genesis_block() -> Self {
        Self::new(String::from("Genesis Block"), String::from(""))
    }

    pub fn get_timestamp(&self) -> i64 {
        self.timestamp
    }

    pub fn get_prev_block_hash(&self) -> String {
        self.prev_block_hash.clone()
    }

    pub fn get_hash(&self) -> String {
        self.hash.clone()
    }

    pub fn get_data(&self) -> String {
        self.data.clone()
    }

    pub fn get_nonce(&self) -> i64 {
        self.nonce
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[ignore]
    #[test]
    fn test_new_block() {
        let block = Block::new(String::from(""), String::from(""));
        println!("{}", block.get_hash());
    }
}
