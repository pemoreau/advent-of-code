mod tests {
    use crate::{part1, part2};

    #[test]
    fn test_input_part1() {
        assert_eq!(150, part1(include_str!("../input.txt").to_string()));
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(12201460, part2(include_str!("../input.txt").to_string()));
    }
}
