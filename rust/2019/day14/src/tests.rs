mod tests {
    use crate::{part1, part2};

    #[test]
    fn test_part1_1() {
        assert_eq!(31, part1(include_str!("../input_test.txt").to_string()));
    }
    #[test]
    fn test_part1_2() {
        assert_eq!(165, part1(include_str!("../input_test2.txt").to_string()));
    }
    #[test]
    fn test_part1_3() {
        assert_eq!(13312, part1(include_str!("../input_test3.txt").to_string()));
    }

    #[test]
    fn test_input_part1() {
        assert_eq!(628586, part1(include_str!("../input.txt").to_string()));
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(3209254, part2(include_str!("../input.txt").to_string()));
    }
}
