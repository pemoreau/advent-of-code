use advent_of_code::day02::part1;
use advent_of_code::day02::part2;

const INPUT: &str = "forward 5
down 5
forward 8
up 3
down 8
forward 2";

#[test]

fn test_part1() {
    let expected = 150;
    let result = part1(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_part2() {
    let expected = 900;
    let result = part2(INPUT.to_string());
    assert_eq!(expected, result);
}
