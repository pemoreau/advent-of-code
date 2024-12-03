pub fn part1(input: String) -> i64 {
    let values: Vec<_> = input.lines().map(|line| line.trim()).collect();
    let mut gamma = 0;
    let mut epsilon = 0;
    let len = values[0].len();
    for i in 0..len {
        let mut count0 = 0;
        let mut count1 = 0;
        for line in &values {
            if line.chars().nth(i).unwrap() == '0' {
                count0 += 1;
            } else {
                count1 += 1;
            }
        }
        if count1 > count0 {
            gamma = 2 * gamma + 1;
            epsilon = 2 * epsilon;
        } else {
            gamma = 2 * gamma;
            epsilon = 2 * epsilon + 1;
        }
    }

    return gamma * epsilon;
}

pub fn part2(input: String) -> i64 {
    let values: Vec<_> = input.lines().map(|line| line.trim()).collect();
    return search(&values, '1') as i64 * search(&values, '0') as i64;
}
// convert binary string to decimal
fn to_int(s: &str) -> i32 {
    let mut result = 0;
    for c in s.chars() {
        result = result * 2 + c.to_digit(10).unwrap() as i32;
    }
    return result;
}

fn search(values: &Vec<&str>, bit: char) -> i32 {
    let mut candidates = values.clone();
    let word_len = values[0].len();

    let mut i = 0;
    while candidates.len() > 1 && i < word_len {
        // println!("{:?} i={} len={}", candidates, i, candidates.len());
        let count_bit = candidates
            .iter()
            .filter(|line| line.chars().nth(i).unwrap() == bit)
            .count();

        let major_bit = if bit == '1' {
            if count_bit >= candidates.len() - count_bit {
                '1'
            } else {
                '0'
            }
        } else {
            if count_bit <= candidates.len() - count_bit {
                '0'
            } else {
                '1'
            }
        };
        // println!("count={} major={}", count_bit, major_bit);
        candidates.retain(|line| line.chars().nth(i).unwrap() == major_bit);
        i += 1;
    }
    return to_int(candidates[0]);
}
