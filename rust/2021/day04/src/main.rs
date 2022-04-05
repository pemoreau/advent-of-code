use day04::{part1, part2};

fn main() {
    println!("--2021 day 04 solution--");
    println!(
        "Part 1: {}",
        part1(include_str!("../input.txt").to_string())
    );
    println!(
        "Part 2: {}",
        part2(include_str!("../input.txt").to_string())
    );
}

#[cfg(test)]
mod tests {
    use crate::{part1, part2};

    #[test]
    fn test_part1() {
        let expected = 24 * 188;
        let result = part1(include_str!("../input_test.txt").to_string());
        assert_eq!(expected, result);
    }

    #[test]
    fn test_part2() {
        let expected = 148 * 13;
        let result = part2(include_str!("../input_test.txt").to_string());
        assert_eq!(expected, result);
    }

    #[test]
    fn test_input_part1() {
        let expected = 74320;
        let result = part1(include_str!("../input.txt").to_string());
        assert_eq!(expected, result);
    }

    #[test]
    fn test_input_part2() {
        let expected = 17884;
        let result = part2(include_str!("../input.txt").to_string());
        assert_eq!(expected, result);
    }
}
