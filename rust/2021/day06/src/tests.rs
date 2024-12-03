mod tests {
    use crate::{part1, part2};

    #[test]
    fn test_part1() {
        let expected = 5934;
        let result = part1(include_str!("../input_test.txt").to_string());
        assert_eq!(expected, result);
    }

    #[test]
    fn test_part2() {
        let expected = 26984457539;
        let result = part2(include_str!("../input_test.txt").to_string());
        assert_eq!(expected, result);
    }

    #[test]
    fn test_input_part1() {
        let expected = 351092;
        let result = part1(include_str!("../input.txt").to_string());
        assert_eq!(expected, result);
    }

    #[test]
    fn test_input_part2() {
        let expected = 1595330616005;
        let result = part2(include_str!("../input.txt").to_string());
        assert_eq!(expected, result);
    }
}
