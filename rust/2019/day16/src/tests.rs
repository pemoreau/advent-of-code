mod tests {
    use crate::{part1, part2};

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

    // #[test]
    // fn test_input_part2() {
    //     assert_eq!(404, part2(include_str!("../input.txt").to_string()));
    // }
}
