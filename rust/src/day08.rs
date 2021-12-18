use itertools::Itertools;
use std::collections::HashMap;

pub fn part1(input: String) -> i64 {
    input
        .lines()
        .flat_map(|line| {
            let parts = line.split("|").collect::<Vec<_>>();
            parts[1].trim().split(" ").collect::<Vec<_>>()
        })
        .filter(|d| [2, 4, 3, 7].contains(&d.len()))
        .count() as i64
}

fn include_string(s: &str, sub: &str) -> bool {
    for l in sub.chars() {
        if !s.contains(l) {
            return false;
        }
    }
    true
}

pub fn part2(input: String) -> i64 {
    let values = input
        .lines()
        .map(|line| {
            let parts = line.split("|").collect::<Vec<_>>();
            let lhs = parts[0].trim().split(" ").collect::<Vec<_>>();
            let rhs = parts[1].trim().split(" ").collect::<Vec<_>>();
            (lhs, rhs)
        })
        .collect::<Vec<_>>();

    let res = values
        .iter()
        .map(|(lhs, rhs)| {
            let mut table = [""; 10];
            table[1] = lhs.iter().find(|d| d.len() == 2).unwrap();
            table[4] = lhs.iter().find(|d| d.len() == 4).unwrap();
            table[7] = lhs.iter().find(|d| d.len() == 3).unwrap();
            table[8] = lhs.iter().find(|d| d.len() == 7).unwrap();
            table[3] = lhs
                .iter()
                .find(|d| d.len() == 5 && include_string(d, table[7]))
                .unwrap();
            table[6] = lhs
                .iter()
                .find(|d| d.len() == 6 && !include_string(d, table[7]))
                .unwrap();
            table[5] = lhs
                .iter()
                .find(|d| {
                    d.len() == 5 && !include_string(d, table[7]) && include_string(table[6], d)
                })
                .unwrap();
            table[9] = lhs
                .iter()
                .find(|d| {
                    d.len() == 6 && include_string(d, table[7]) && include_string(d, table[5])
                })
                .unwrap();
            table[2] = lhs
                .iter()
                .find(|d| {
                    d.len() == 5 && !include_string(d, table[7]) && !include_string(table[9], d)
                })
                .unwrap();
            table[0] = lhs
                .iter()
                .find(|d| d.len() == 6 && !table.contains(d))
                .unwrap();
            let map: HashMap<String, u8> =
                HashMap::from_iter(table.iter().enumerate().map(|(i, d)| {
                    let key = d.chars().sorted().collect::<String>();
                    (key, i as u8)
                }));
            let number: i64 = rhs.iter().fold(0, |acc, d| {
                let key = d.chars().sorted().collect::<String>();
                acc * 10 + map[&key] as i64
            });
            number
        })
        .sum();

    res
}
