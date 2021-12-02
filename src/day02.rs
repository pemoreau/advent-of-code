fn parse_line1(s: &str) -> (&str, usize) {
    let re = regex::Regex::new(r"([a-z]+) ([0-9]+)").unwrap();
    let cap = re.captures(s).unwrap();
    let order: &str = cap.get(1).unwrap().as_str();
    let value: usize = cap[2].parse().unwrap();
    return (order, value);
}

// using for-loop style
pub fn part1(input: String) -> i32 {
    let values: Vec<_> = input.lines().map(|line| parse_line1(line)).collect();
    let mut horizontal = 0;
    let mut depth = 0;
    for (order, value) in values {
        if order == "forward" {
            horizontal += value;
        }
        if order == "up" {
            depth -= value;
        }
        if order == "down" {
            depth += value;
        }
    }

    return (horizontal * depth).try_into().unwrap();
}

// using fold()
pub fn part2(input: String) -> i32 {
    let values: Vec<_> = input.lines().map(|line| parse_line1(line)).collect();
    let (horizontal, depth, _aim) =
        values
            .iter()
            .fold((0, 0, 0), |(horizontal, depth, aim), (order, value)| {
                if *order == "forward" {
                    (horizontal + value, depth + (aim * value), aim)
                } else if *order == "up" {
                    (horizontal, depth, aim - value)
                } else if *order == "down" {
                    (horizontal, depth, aim + value)
                } else {
                    panic!("Unknown order: {}", order);
                }
            });
    return (horizontal * depth).try_into().unwrap();
}
