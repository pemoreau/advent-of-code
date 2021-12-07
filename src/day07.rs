pub fn part1(input: String) -> i64 {
    return search(input, |a, b| (a - b).abs());
}

pub fn part2(input: String) -> i64 {
    return search(input, |a, b| (a - b).abs() * ((a - b).abs() + 1) / 2);
}

pub fn search(input: String, cost: fn(a: i32, b: i32) -> i32) -> i64 {
    let values: Vec<i32> = input
        .split(',')
        .map(|s| s.trim().parse::<i32>().unwrap())
        .collect();
    let min_value = values.iter().min().unwrap();
    let max_value = values.iter().max().unwrap();
    let h_sum =
        (*min_value..*max_value).map(|h| (h, values.iter().map(|v| cost(*v, h)).sum::<i32>()));
    let (_, sum) = h_sum.min_by(|(_, s1), (_, s2)| s1.cmp(s2)).unwrap();
    return sum as i64;
}
