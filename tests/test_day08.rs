use advent_of_code::day08::part1;
use advent_of_code::day08::part2;

const INPUT: &str = "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |
fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |
fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |
cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |
efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |
gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |
gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |
cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |
ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |
gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |
fgae cfgab fg bagce";

#[test]
fn test_part1() {
    let expected = 26;
    let result = part1(INPUT.to_string());
    assert_eq!(expected, result);
}

// #[test]
// fn test_part2() {
//     let expected: i64 = 168;
//     let result = part2(INPUT.to_string());
//     assert_eq!(expected, result);
// }

// #[test]
// fn test_input_part1() {
//     assert_eq!(
//         355989,
//         part1(include_str!("../inputs/day08.txt").to_string())
//     );
// }

// #[test]
// fn test_input_part2() {
//     assert_eq!(
//         102245489,
//         part2(include_str!("../inputs/day08.txt").to_string())
//     );
// }
