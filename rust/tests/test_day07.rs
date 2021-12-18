use advent_of_code::day07::part1;
use advent_of_code::day07::part2;

const INPUT: &str = "16,1,2,0,4,2,7,1,2,14";

#[test]
fn test_part1() {
    let expected = 37;
    let result = part1(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_part2() {
    let expected: i64 = 168;
    let result = part2(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_input_part1() {
    assert_eq!(
        355989,
        part1(include_str!("../inputs/day07.txt").to_string())
    );
}

#[test]
fn test_input_part2() {
    assert_eq!(
        102245489,
        part2(include_str!("../inputs/day07.txt").to_string())
    );
}
