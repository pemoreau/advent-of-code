use advent_of_code::day03::part1;
use advent_of_code::day03::part2;

const INPUT: &str = "00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010";

#[test]

fn test_part1() {
    let expected = 22 * 9;
    let result = part1(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_part2() {
    let expected = 23 * 10;
    let result = part2(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_input_part1() {
    assert_eq!(
        3242606,
        part1(include_str!("../inputs/day03.txt").to_string())
    );
}

#[test]
fn test_input_part2() {
    assert_eq!(
        4856080,
        part2(include_str!("../inputs/day03.txt").to_string())
    );
}
