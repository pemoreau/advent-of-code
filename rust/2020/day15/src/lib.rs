use std::collections::HashMap;

fn play(turn: usize, n: usize, map: &mut HashMap<usize, usize>) -> usize {
    let previous = map.insert(n, turn);
    if previous.is_some() {
        turn - previous.unwrap()
    } else {
        0
    }
}

pub fn solve(input: String, n: usize) -> i64 {
    let mut map = input
        .trim()
        .split(",")
        .enumerate()
        .map(|(i, n)| (n.parse::<usize>().unwrap(), i + 1))
        .collect::<HashMap<usize, usize>>();

    let mut spoken = 0;
    for turn in 1 + map.len()..n {
        spoken = play(turn, spoken, &mut map);
    }

    spoken as i64
    // (1 + map.len()..n).fold(0, |spoken, turn| play(turn, spoken, &mut map)) as i64
}

pub fn part1(input: String) -> i64 {
    solve(input, 2020)
}

pub fn part2(input: String) -> i64 {
    solve(input, 30000000)
}
