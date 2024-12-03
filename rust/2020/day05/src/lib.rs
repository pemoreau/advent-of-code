fn to_bin(input: &str, lower: char) -> usize {
    input
        .chars()
        .fold(0, |acc, c| if c == lower { 2 * acc } else { 2 * acc + 1 })
}

pub fn part1(input: String) -> i64 {
    // F:0 B:1
    // L:0 R:1
    let res = input
        .lines()
        .map(|line| to_bin(&line[0..7], 'F') * 8 + to_bin(&line[7..10], 'L'))
        .max()
        .unwrap();
    res.try_into().unwrap()
}

pub fn part2(input: String) -> i64 {
    let mut seats = [false; 912];
    input.lines().for_each(|line| {
        let seat = to_bin(&line[0..7], 'F') * 8 + to_bin(&line[7..10], 'L');
        seats[seat] = true;
    });

    let mut previous = false;
    let mut res = 0;
    seats.iter().enumerate().for_each(|(index, element)| {
        if !element && previous {
            res = index.try_into().unwrap();
        }
        previous = *element;
    });
    res
}
