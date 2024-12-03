use day12::{part1, part2};
use std::time::Instant;

fn main() {
    println!("--2019 day 12 solution--");
    let now = Instant::now();

    println!(
        "Part 1: {} in {} ms",
        part1(),
        now.elapsed().as_micros() as f64 / 1000.0
    );

    let now = Instant::now();
    println!(
        "Part 2: {} in {} ms",
        part2(),
        now.elapsed().as_micros() as f64 / 1000.0
    );
}

#[cfg(test)]
mod tests;
