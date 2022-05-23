use prototype::blockchain::Blockchain;

fn main() {
    let mut bc = Blockchain::new();

    bc.add_block("Send 1 BTC to Ivan".to_string());
    bc.add_block("Send 2 more BTC to Ivan".to_string());

    for block in bc.get_blocks() {
        println!("Prev Hash: {}", block.get_prev_block_hash());
        println!("Data: {}", block.get_data());
        println!("Hash: {}", block.get_hash());
        println!("");
    }
}
