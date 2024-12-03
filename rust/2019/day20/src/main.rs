use std::time::Instant;

use day20::{part1, part2};

fn main() {
    println!("--2019 day 20 solution--");
    let now = Instant::now();

    println!(
        "Part 1: {} in {} ms",
        part1(include_str!("../input.txt").to_string()),
        now.elapsed().as_micros() as f64 / 1000.0
    );

    let now = Instant::now();
    println!(
        "Part 2: {} in {} ms",
        part2(include_str!("../input.txt").to_string()),
        now.elapsed().as_micros() as f64 / 1000.0
    );
}

// 2147483647 too high

#[cfg(test)]
mod tests;
