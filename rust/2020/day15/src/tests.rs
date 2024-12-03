mod tests {
    use crate::{part1, part2};

    #[test]
    fn test_part1() {
        assert_eq!(436, part1("0,3,6".to_string()));
        assert_eq!(1, part1("1,3,2".to_string()));
        assert_eq!(10, part1("2,1,3".to_string()));
        assert_eq!(27, part1("1,2,3".to_string()));
        assert_eq!(78, part1("2,3,1".to_string()));
        assert_eq!(438, part1("3,2,1".to_string()));
        assert_eq!(1836, part1("3,1,2".to_string()));
    }

    // #[test]
    // fn test_part2() {
    //     assert_eq!(175594, part2("0,3,6".to_string()));
    //     assert_eq!(2578, part2("1,3,2".to_string()));
    //     assert_eq!(3544142, part2("2,1,3".to_string()));
    //     assert_eq!(261214, part2("1,2,3".to_string()));
    //     assert_eq!(6895259, part2("2,3,1".to_string()));
    //     assert_eq!(18, part2("3,2,1".to_string()));
    //     assert_eq!(362, part2("3,1,2".to_string()));
    // }

    #[test]
    fn test_input_part1() {
        assert_eq!(376, part1(include_str!("../input.txt").to_string()));
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(323780, part2(include_str!("../input.txt").to_string()));
    }
}
