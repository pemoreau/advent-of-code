mod tests {
    use crate::part1;
    use crate::part2;

    #[test]
    fn test_input_test1() {
        assert_eq!(
            2129920,
            part1(include_str!("../input_test.txt").to_string())
        );
    }

    // #[test]
    // fn test_input_test2() {
    //     assert_eq!(99, part1(include_str!("../input_test.txt").to_string()));
    // }

    #[test]
    fn test_input_part1() {
        assert_eq!(28717468, part1(include_str!("../input.txt").to_string()));
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(2014, part2(include_str!("../input.txt").to_string()));
    }
}
