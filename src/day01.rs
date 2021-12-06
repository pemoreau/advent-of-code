fn values(input: String) -> Vec<i32> {
    input
        .lines()
        .map(|line| line.trim().parse().unwrap())
        .collect()
}

pub fn increases(values: &Vec<i32>) -> usize {
    let iter = values.windows(2);
    let diff = iter.map(|pair| pair[1] - pair[0]);
    return diff.filter(|&x| x > 0).count();
}

pub fn part1(input: String) -> i64 {
    let values = values(input);
    return increases(&values).try_into().unwrap();
}

pub fn part2(input: String) -> i64 {
    let values = values(input);
    let sums = values.windows(3).map(|pair| pair[0] + pair[1] + pair[2]);
    let vsums = sums.collect::<Vec<i32>>();
    return increases(&vsums).try_into().unwrap();
}
