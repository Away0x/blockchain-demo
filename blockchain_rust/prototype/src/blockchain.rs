use crate::block::Block;

pub struct Blockchain {
    blocks: Vec<Block>,
}

impl Blockchain {
    pub fn new() -> Self {
        Self {
            blocks: vec![Block::new_genesis_block()],
        }
    }

    pub fn add_block(&mut self, data: String) {
        let prev_block = self.blocks.last().unwrap();
        let new_block = Block::new(data, prev_block.get_hash());
        self.blocks.push(new_block);
    }

    pub fn get_blocks(&self) -> &Vec<Block> {
        &self.blocks
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add_block() {
        let mut blockchain = Blockchain::new();
        blockchain.add_block(String::from("test"));
        assert_eq!(blockchain.blocks.len(), 2);
    }
}
