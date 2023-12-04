mod tests {
    use crate::part1;
    use crate::part2;

    #[test]
    fn test_part1_1() {
        assert_eq!(86, part1(include_str!("../inputTest1.txt").to_string()));
    }

    #[test]
    fn test_part1_2() {
        assert_eq!(132, part1(include_str!("../input_test2.txt").to_string()));
    }

    #[test]
    fn test_part1_3() {
        assert_eq!(136, part1(include_str!("../inputTest3.txt").to_string()));
    }

    #[test]
    fn test_part1_4() {
        assert_eq!(81, part1(include_str!("../inputTest4.txt").to_string()));
    }

    #[test]
    fn test_part1_5() {
        assert_eq!(13, part1(include_str!("../inputTest5.txt").to_string()));
    }

    #[test]
    fn test_part2_6() {
        assert_eq!(8, part2(include_str!("../inputTest6.txt").to_string()));
    }

    #[test]
    fn test_part2_7() {
        assert_eq!(32, part2(include_str!("../inputTest7.txt").to_string()));
    }

    #[test]
    fn test_part2_8() {
        assert_eq!(72, part2(include_str!("../inputTest8.txt").to_string()));
    }

    #[test]
    fn test_input_part1() {
        assert_eq!(5406, part1(include_str!("../input.txt").to_string()));
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(1938, part2(include_str!("../input.txt").to_string()));
    }
}
