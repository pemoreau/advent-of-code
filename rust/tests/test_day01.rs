use advent_of_code::day01::part1;
use advent_of_code::day01::part2;

const INPUT: &str = "199
200
208
210
200
207
240
269
260
263";

#[test]
fn test_part1() {
    let expected = 7;
    let result = part1(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_part2() {
    let expected = 5;
    let result = part2(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_input_part1() {
    assert_eq!(1722, part1(include_str!("../inputs/day01.txt").to_string()));
}

#[test]
fn test_input_part2() {
    assert_eq!(1748, part2(include_str!("../inputs/day01.txt").to_string()));
}
