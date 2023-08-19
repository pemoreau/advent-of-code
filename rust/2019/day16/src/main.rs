use std::time::Instant;

use day16::{part1, part2};

fn main() {
    println!("--2019 day 16 solution--");
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
        // part2("03036732577212944063491565474664".to_string()),
        // part2("02935109699940807407585447034323".to_string()),
        // part2("03081770884921959731165446850517".to_string()),
        now.elapsed().as_micros() as f64 / 1000.0
    );
}

#[cfg(test)]
mod tests;
