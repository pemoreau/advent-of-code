pub fn lines_to_numbers(input: String) -> Vec<i32> {
    input
        .lines()
        .map(|line| line.trim().parse().unwrap())
        .collect()
}
