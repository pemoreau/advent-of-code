pub fn part1(input: String) -> i32 {
    return simulate(input, 80) as i32;
}

pub fn part2(input: String) -> i32 {
    println!("part2 {}", simulate(input, 256));
    return 0;
}

pub fn simulate(input: String, n: usize) -> i64 {
    let mut mult = input.split(',').fold([0; 9], |mut acc, value| {
        acc[value.trim().parse::<usize>().unwrap()] += 1;
        acc
    });

    for _ in 0..n {
        mult.rotate_left(1);
        mult[6] += mult[8];
    }
    return mult.iter().sum::<usize>() as i64;
}
