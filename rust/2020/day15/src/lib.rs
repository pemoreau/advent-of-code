use std::collections::HashMap;

fn play(turn: usize, n: usize, map: &mut HashMap<usize, usize>) -> usize {
    // let index = map.get(&n);
    // if index.is_none() {
    //     map.insert(n, turn);
    //     return 0;
    // } else {
    // If map.insert(...) is moved here, the borrow checker complains
    //     let spoken = turn - index.unwrap();
    //     map.insert(n, turn);
    //     return spoken;
    // }
    if map.contains_key(&n) {
        let spoken = turn - map.get(&n).unwrap();
        map.insert(n, turn);
        return spoken;
    } else {
        map.insert(n, turn);
        0
    }
}

pub fn solve(input: String, n: usize) -> i64 {
    // let numbers = input
    //     .split(",")
    //     .filter_map(|s| s.parse::<usize>().ok())
    //     .collect::<Vec<usize>>();

    // // println!"numbers {:?}", numbers);
    // let mut map: HashMap<usize, usize> = HashMap::new();

    // for (i, n) in numbers.iter().enumerate() {
    //     map.insert(*n, i + 1);
    // }

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

    return spoken as i64;
}

pub fn part1(input: String) -> i64 {
    solve(input, 2020)
}

pub fn part2(input: String) -> i64 {
    solve(input, 30000000)
}
