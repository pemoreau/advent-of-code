use advent_of_code::day05::part1;
use advent_of_code::day05::part2;

const INPUT: &str = "0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2";

#[test]

fn test_part1() {
    let expected = 5;
    let result = part1(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_part2() {
    let expected = 12;
    let result = part2(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_input_part1() {
    assert_eq!(7297, part1(include_str!("../inputs/day05.txt").to_string()));
}

#[test]
fn test_input_part2() {
    assert_eq!(
        21038,
        part2(include_str!("../inputs/day05.txt").to_string())
    );
}
