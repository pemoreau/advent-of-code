mod tests {
    use crate::part1;

    #[test]
    fn test_input_part1() {
        assert_eq!(2975, part1(include_str!("../input.txt").to_string()));
    }
}
