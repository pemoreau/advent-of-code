use advent_of_code::day07::part1;
use advent_of_code::day07::part2;

const INPUT: &str = "16,1,2,0,4,2,7,1,2,14";

#[test]

fn test_part1() {
    let expected = 37; // 355989
    let result = part1(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_part2() {
    let expected: i64 = 168; // 102245489
    let result = part2(INPUT.to_string());
    assert_eq!(expected, result);
}
