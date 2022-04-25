mod tests {
    use crate::{part1, part2};

    #[test]
    fn test_part1() {
        assert_eq!(71, part1("1 + 2 * 3 + 4 * 5 + 6".to_string()));
        assert_eq!(51, part1("1 + (2 * 3) + (4 * (5 + 6))".to_string()));
        assert_eq!(26, part1("2 * 3 + (4 * 5)".to_string()));
        assert_eq!(437, part1("5 + (8 * 3 + 9 + 3 * 4 * 3)".to_string()));
        assert_eq!(
            12240,
            part1("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))".to_string())
        );
        assert_eq!(
            13632,
            part1("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2".to_string())
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(231, part2("1 + 2 * 3 + 4 * 5 + 6".to_string()));
        assert_eq!(51, part2("1 + (2 * 3) + (4 * (5 + 6))".to_string()));
        assert_eq!(46, part2("2 * 3 + (4 * 5)".to_string()));
        assert_eq!(1445, part2("5 + (8 * 3 + 9 + 3 * 4 * 3)".to_string()));
        assert_eq!(
            669060,
            part2("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))".to_string())
        );
        assert_eq!(
            23340,
            part2("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2".to_string())
        );
    }

    #[test]
    fn test_input_part1() {
        assert_eq!(
            4940631886147,
            part1(include_str!("../input.txt").to_string())
        );
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(
            283582817678281,
            part2(include_str!("../input.txt").to_string())
        );
    }
}
