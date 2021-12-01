// use itertools::Itertools;
fn values(input: String) -> Vec<i32> {
    input.lines().map(|line| line.parse().unwrap()).collect()
}

pub fn increases(values: &Vec<i32>) -> usize {
    let iter = values.windows(2);
    let diff = iter.map(|pair| pair[1] - pair[0]);
    return diff.filter(|&x| x > 0).count();
}

pub fn part1(input: String) {
    let values = values(input);
    println!("increases = {}", increases(&values));
}

pub fn part2(input: String) {
    let values = values(input);
    let sums = values.windows(3).map(|pair| pair[0] + pair[1] + pair[2]);
    let vsums = sums.collect::<Vec<i32>>();
    println!("increases = {}", increases(&vsums));
}
