use day06::{part1, part2};

fn main() {
    println!("--2021 day 06 solution--");
    println!(
        "Part 1: {}",
        part1(include_str!("../input.txt").to_string())
    );
    println!(
        "Part 2: {}",
        part2(include_str!("../input.txt").to_string())
    );
}

#[cfg(test)]
mod tests;
