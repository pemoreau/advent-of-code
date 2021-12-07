pub fn part1(input: String) -> i64 {
    search(input, |a, b| (a - b).abs())
}

pub fn part2(input: String) -> i64 {
    search(input, |a, b| {
        let n = (a - b).abs();
        n * (n + 1) / 2
    })
}

pub fn search(input: String, cost: fn(a: i32, b: i32) -> i32) -> i64 {
    let values = input
        .split(',')
        .map(|s| s.trim().parse::<i32>().unwrap())
        .collect::<Vec<i32>>();
    let min_value = values.iter().min().unwrap();
    let max_value = values.iter().max().unwrap();
    let min_sum = (*min_value..*max_value)
        .map(|h| values.iter().map(|v| cost(*v, h)).sum::<i32>())
        .min()
        .unwrap();
    min_sum as i64
}
