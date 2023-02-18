pub fn lines_to_numbers(input: String) -> Vec<i64> {
    input
        .lines()
        .map(|line| line.trim().parse().unwrap())
        .collect()
}

pub fn comma_separated_to_numbers(input: String) -> Vec<i64> {
    input
        .split(',')
        .map(|line| line.trim().parse().unwrap())
        .collect()
}
