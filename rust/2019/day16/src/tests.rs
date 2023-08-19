mod tests {
    use crate::part1;
    use crate::part2;

// #[test]
    // fn test_part1_1() {
    //     assert_eq!(1029498, part1("12345678".to_string()));
    // }

    #[test]
    fn test_part1_2() {
        assert_eq!(24176176, part1("80871224585914546619083218645595".to_string()));
    }

    #[test]
    fn test_part1_3() {
        assert_eq!(73745418, part1("19617804207202209144916044189917".to_string()));
    }

    #[test]
    fn test_part1_4() {
        assert_eq!(52432133, part1("69317163492948606335995924319873".to_string()));
    }

    #[test]
    fn test_input_part1() {
        assert_eq!(67481260, part1(include_str!("../input.txt").to_string()));
    }

    #[test]
    fn test_part2_1() {
        assert_eq!(84462026, part2("03036732577212944063491565474664".to_string()));
    }

    #[test]
    fn test_part2_2() {
        assert_eq!(78725270, part2("02935109699940807407585447034323".to_string()));
    }

    #[test]
    fn test_part2_3() {
        assert_eq!(53553731, part2("03081770884921959731165446850517".to_string()));
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(42178738, part2(include_str!("../input.txt").to_string()));
    }
}
