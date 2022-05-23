use ring::digest::{Context, SHA256};

pub fn sha256_digest(data: &[u8]) -> Vec<u8> {
    let mut context = Context::new(&SHA256);
    context.update(data);
    let digest = context.finish();
    digest.as_ref().to_vec()
}

#[cfg(test)]
mod tests {
    use super::*;
    use data_encoding::HEXLOWER;

    #[test]
    fn test_sha256_digest() {
        let target = "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3";

        assert_eq!(
            HEXLOWER.encode(sha256_digest("123".as_bytes()).as_slice()),
            target
        );
    }
}
