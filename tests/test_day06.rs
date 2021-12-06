use advent_of_code::day06::part1;
use advent_of_code::day06::part2;

const INPUT: &str = "3,4,3,1,2";

#[test]

fn test_part1() {
    let expected = 5934;
    let result = part1(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_part2() {
    let expected: i64 = 26984457539;
    let result = part2(INPUT.to_string());
    assert_eq!(expected, result);
}
