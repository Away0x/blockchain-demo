use std::{borrow::Borrow, ops::ShlAssign};

use data_encoding::HEXLOWER;
use num_bigint::{BigInt, Sign};

use crate::{block::Block, utils::sha256_digest};

/// 难度值
const TARGET_BITS: i32 = 24;
const MAX_NONCE: i64 = i64::MAX;

pub struct ProofOfWork {
    block: Block,
    target: BigInt,
}

impl ProofOfWork {
    pub fn new(block: Block) -> Self {
        let mut target = BigInt::from(1);
        // 左移 256-targetBits 位 (256 是一个 SHA-256 哈希的位数)
        // target 16 进制为: 0x10000000000000000000000000000000000000000000000000000000000
        target.shl_assign(256 - TARGET_BITS);
        Self { block, target }
    }

    pub fn run(&self) -> (i64, String) {
        let mut hash: Vec<u8> = vec![];
        let mut nonce = 0;

        println!("Mining the block containing {}", self.block.get_data());
        while nonce < MAX_NONCE {
            let data = self.prepare_data(nonce);
            hash = sha256_digest(data.as_slice());
            let hash_int = BigInt::from_bytes_be(Sign::Plus, hash.as_slice());

            if hash_int.lt(self.target.borrow()) {
                println!("found a valid hash: {}", HEXLOWER.encode(hash.as_slice()));
                break;
            } else {
                nonce += 1;
            }
        }

        println!("");
        (nonce, HEXLOWER.encode(hash.as_slice()))
    }

    pub fn validate(&self) -> bool {
        let data = self.prepare_data(self.block.get_nonce());
        let hash = sha256_digest(data.as_slice());
        let hash_int = BigInt::from_bytes_be(Sign::Plus, hash.as_slice());

        hash_int.lt(&self.target.borrow())
    }

    fn prepare_data(&self, nonce: i64) -> Vec<u8> {
        let mut result: Vec<u8> = vec![];
        result.extend(self.block.get_prev_block_hash().as_bytes());
        result.extend(self.block.get_data().as_bytes());
        result.extend(self.block.get_timestamp().to_be_bytes());
        result.extend(TARGET_BITS.to_be_bytes());
        result.extend(nonce.to_be_bytes());
        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[ignore]
    #[test]
    fn test_pow() {
        let block = Block::new("".to_string(), "".to_string());
        let pow = ProofOfWork::new(block);

        println!("{:?}", pow.run());
    }
}
