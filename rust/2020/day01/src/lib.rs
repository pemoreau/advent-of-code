use itertools::Itertools;

fn values(input: String) -> Vec<u32> {
    input.lines().map(|line| line.parse().unwrap()).collect()
}

pub fn part1(input: String) -> i64 {
    let values = values(input);
    for (i, value1) in values.iter().enumerate() {
        for value2 in values.iter().skip(i) {
            if value1 + value2 == 2020 {
                return (value1 * value2).try_into().unwrap();
            }
        }
    }
    0
}

pub fn part2(input: String) -> i64 {
    let values = values(input);
    let (a, b, c) = values
        .into_iter()
        .tuple_combinations()
        .find(|(a, b, c)| a + b + c == 2020)
        .expect("no tuple of length 3 had a sum of 2020");
    return (a * b * c).try_into().unwrap();
}

pub fn part3(input: String) -> i64 {
    let values = values(input);

    for (i, value1) in values.iter().enumerate() {
        for (j, value2) in values.iter().skip(i).enumerate() {
            for value3 in values.iter().skip(i + j) {
                if value1 + value2 + value3 == 2020 {
                    return (value1 * value2 * value3).try_into().unwrap();
                }
            }
        }
    }
    0
}
