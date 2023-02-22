mod tests {
    use crate::{part1, part2};

    #[test]
    fn test_part1() {
        assert_eq!(
            43210,
            part1("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0".to_string())
        );
        assert_eq!(
            54321,
            part1(
                "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
                    .to_string()
            )
        );
        assert_eq!(
            65210,
            part1(
                "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
                    .to_string()
            )
        );
    }

    #[test]
    fn test_input_part1() {
        assert_eq!(398674, part1(include_str!("../input.txt").to_string()));
    }

    // #[test]
    // fn test_input_part2() {
    //     assert_eq!(15163975, part2(include_str!("../input.txt").to_string()));
    // }
}
