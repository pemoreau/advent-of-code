use advent_of_code::day11::part1;
use advent_of_code::day11::part2;

const INPUT: &str = "5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526";

#[test]
fn test_part1() {
    let expected = 1656;
    let result = part1(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_part2() {
    let expected: i64 = 195;
    let result = part2(INPUT.to_string());
    assert_eq!(expected, result);
}

#[test]
fn test_input_part1() {
    assert_eq!(1681, part1(include_str!("../inputs/day11.txt").to_string()));
}

#[test]
fn test_input_part2() {
    assert_eq!(276, part2(include_str!("../inputs/day11.txt").to_string()));
}
