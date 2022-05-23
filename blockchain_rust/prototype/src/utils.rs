use data_encoding::HEXLOWER;
use ring::digest::{Context, SHA256};

pub fn sha256_digest(data: String) -> String {
    let mut context = Context::new(&SHA256);
    context.update(data.as_bytes());
    let digest = context.finish();
    return HEXLOWER.encode(digest.as_ref());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_sha256_digest() {
        assert_eq!(
            sha256_digest("123".to_string()),
            "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3"
        );
    }
}
