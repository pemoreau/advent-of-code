use advent_of_code::day13::part1;

const INPUT: &str = "6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5";

#[test]
fn test_part1() {
    let expected = 17;
    let result = part1(INPUT.to_string());
    assert_eq!(expected, result);
}

// #[test]
// fn test_part2() {
//     let expected: i64 = 195;
//     let result = part2(INPUT.to_string());
//     assert_eq!(expected, result);
// }

#[test]
fn test_input_part1() {
    assert_eq!(795, part1(include_str!("../inputs/day13.txt").to_string()));
}

// #[test]
// fn test_input_part2() {
//     assert_eq!(276, part2(include_str!("../inputs/day11.txt").to_string()));
// }
