pub fn part1(input: String) -> i64 {
    simulate(input, 80)
}

pub fn part2(input: String) -> i64 {
    simulate(input, 256)
}

fn simulate(input: String, n: usize) -> i64 {
    let mut mult = input.split(',').fold([0; 9], |mut acc, value| {
        acc[value.trim().parse::<usize>().unwrap()] += 1;
        acc
    });

    for _ in 0..n {
        mult.rotate_left(1);
        mult[6] += mult[8];
    }
    mult.iter().sum::<usize>() as i64
}
