mod tests {
    use crate::{part1, part2};

    #[test]
    fn test_input_part1() {
        let expected = 911;
        let result = part1(include_str!("../input.txt").to_string());
        assert_eq!(expected, result);
    }

    #[test]
    fn test_input_part2() {
        let expected = 629;
        let result = part2(include_str!("../input.txt").to_string());
        assert_eq!(expected, result);
    }
}
