mod tests {
    use day01::fuel;
    use day01::recursive_fuel;
    use crate::{part1, part2};

    #[test]
    fn test_fuel() {
        assert_eq!(2, fuel(12));
        assert_eq!(2, fuel(14));
    }

    #[test]
    fn test_recursive_fuel() {
        assert_eq!(2, recursive_fuel(12));
        assert_eq!(966, recursive_fuel(1969));
        assert_eq!(50346, recursive_fuel(100756));
    }

    #[test]
    fn test_input_part1() {
        assert_eq!(3297626, part1(include_str!("../input.txt").to_string()));
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(
            4943578,
            part2(include_str!("../input.txt").to_string())
        );
    }
}
