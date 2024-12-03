use day13::{part1, part2};

fn main() {
    println!("--2021 day 13 solution--");
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
        let expected = 17;
        let result = part1(include_str!("../input_test.txt").to_string());
        assert_eq!(expected, result);
    }

    #[test]
    fn test_input_part1() {
        let expected = 795;
        let result = part1(include_str!("../input.txt").to_string());
        assert_eq!(expected, result);
    }
}
