use crate::utils::sha256_digest;
use chrono::Utc;

pub struct Block {
    // 区块头信息 ------
    timestamp: i64,          // 区块创建的时间
    prev_block_hash: String, // 前一个 block 的 hash
    hash: String,            // 当前 block 的 hash
    // 区块头信息 ------
    data: String, // 区块存储的实际有效信息，也就是交易
}

impl Block {
    pub fn new(data: String, prev_block_hash: String) -> Self {
        let mut block = Self {
            timestamp: Utc::now().timestamp(),
            data,
            prev_block_hash,
            hash: String::new(),
        };
        block.set_hash();
        block
    }

    /// 创世区块
    pub fn new_genesis_block() -> Self {
        Self::new(String::from("Genesis Block"), String::from(""))
    }

    pub fn set_hash(&mut self) {
        let headers = format!("{}{}{}", self.timestamp, self.prev_block_hash, self.data);
        self.hash = sha256_digest(headers)
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
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_set_hash() {
        let mut block = Block::new(
            String::from("test"),
            String::from("25605b6ad194e96ca8c5c9007b1a0bf87b74d130ddda685cd7ccc66d51995541"),
        );
        block.timestamp = 1546300800;
        block.set_hash();
        assert_eq!(
            block.get_hash(),
            "82246ad34ac8c46a967f36b58e18165cad43eb59399a250bc8629e0fe89ab520"
        );
    }
}
