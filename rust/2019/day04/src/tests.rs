mod tests {
    use crate::{part1, part2};

    #[test]
    fn test_input_part1() {
        assert_eq!(1864, part1("137683-596253".to_string()));
    }

    #[test]
    fn test_input_part2() { assert_eq!(1258, part2("137683-596253".to_string())); }
}
