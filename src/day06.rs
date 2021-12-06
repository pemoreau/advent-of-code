pub fn part1(input: String) -> i32 {
    return simulate(input, 80) as i32;
}

pub fn part2(input: String) -> i32 {
    println!("part2 {}", simulate(input, 256));
    return 0;
}

pub fn simulate(input: String, n: usize) -> i64 {
    let values = input
        .split(',')
        .map(|s| s.trim().parse::<usize>().unwrap())
        .collect::<Vec<usize>>();

    let mut mult = values.iter().fold([0; 9], |mut acc, &value| {
        acc[value] += 1;
        acc
    });

    for _i in 0..n {
        // mult = mult
        //     .iter()
        //     .enumerate()
        //     .fold([0; 9], |mut acc, (index, &m)| {
        //         if index == 0 {
        //             acc[6] += m;
        //             acc[8] += m;
        //         } else {
        //             acc[index - 1] += m;
        //         }
        //         acc
        //     });
        mult.rotate_left(1);
        mult[6] += mult[8];
    }
    return mult.iter().sum::<usize>() as i64;
}
