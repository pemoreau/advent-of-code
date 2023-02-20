mod tests {
    use crate::{part1, part2};
    #[test]
    fn test_part1() {
        assert_eq!(42, part1(include_str!("../input_test.txt").to_string()));
    }

    #[test]
    fn test_part2() {
        assert_eq!(4, part2(include_str!("../input_test2.txt").to_string()));
    }

    #[test]
    fn test_input_part1() {
        assert_eq!(253104, part1(include_str!("../input.txt").to_string()));
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(499, part2(include_str!("../input.txt").to_string()));
    }
}
