use day11::{part1, part2};
use std::time::Instant;

fn main() {
    println!("--2021 day 11 solution--");
    let now = Instant::now();

    println!(
        "Part 1: {} in {} ms",
        part1(include_str!("../input.txt").to_string(),),
        now.elapsed().as_millis()
    );

    let now = Instant::now();
    println!(
        "Part 2: {} in {} ms",
        part2(include_str!("../input.txt").to_string()),
        now.elapsed().as_millis()
    );
}

#[cfg(test)]
mod tests;
