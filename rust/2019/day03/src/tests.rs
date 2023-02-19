mod tests {
    use crate::{part1, part2};

    #[test]
    fn test_part1() {
        assert_eq!(6, part1("R8,U5,L5,D3\nU7,R6,D4,L4".to_string()));
        assert_eq!(
            159,
            part1(
                "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83".to_string()
            )
        );
        assert_eq!(
            135,
            part1(
                "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
                    .to_string()
            )
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(30, part2("R8,U5,L5,D3\nU7,R6,D4,L4".to_string()));
        assert_eq!(
            610,
            part2(
                "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83".to_string()
            )
        );
        assert_eq!(
            410,
            part2(
                "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
                    .to_string()
            )
        );
    }

    #[test]
    fn test_input_part1() {
        assert_eq!(293, part1(include_str!("../input.txt").to_string()));
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(27306, part2(include_str!("../input.txt").to_string()));
    }
}
