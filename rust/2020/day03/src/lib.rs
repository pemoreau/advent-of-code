pub fn part1(input: String) -> i64 {
    let values = input.lines().map(|line| line.chars().cycle());
    let cpt = values.enumerate().fold(0, |acc, (i, line)| {
        if line.clone().nth(3 * i).unwrap() == '#' {
            acc + 1
        } else {
            acc
        }
    });
    cpt
}

pub fn part2(input: String) -> i64 {
    let slopes = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)];
    let cpt = slopes.iter().fold(1, |prod, (x, y)| {
        let values = input.lines().map(|line| line.chars().cycle());
        let cpt = values.enumerate().fold(0, |acc, (j, line)| {
            if j % y == 0 {
                let i = (x * j) / y;
                let element = line.clone().nth(i).unwrap();
                if element == '#' {
                    acc + 1
                } else {
                    acc
                }
            } else {
                acc
            }
        });
        prod * cpt
    });
    cpt
}
