use itertools::Itertools;

fn corrupted_score(c: char) -> i64 {
    match c {
        ')' => 3,
        ']' => 57,
        '}' => 1197,
        '>' => 25137,
        _ => 0,
    }
}

fn auto_score(c: char) -> i64 {
    match c {
        '(' => 1,
        '[' => 2,
        '{' => 3,
        '<' => 4,
        _ => 0,
    }
}

fn parse_line(line: &str) -> (i64, i64) {
    let mut stack = Vec::new();
    for c in line.chars() {
        if c == '(' || c == '[' || c == '{' || c == '<' {
            stack.push(c);
        } else {
            if stack.is_empty() {
                return (0, 0);
            } else {
                let top = stack.pop().unwrap();
                if (top == '(' && c != ')')
                    || (top == '[' && c != ']')
                    || (top == '{' && c != '}')
                    || (top == '<' && c != '>')
                {
                    return (corrupted_score(c), 0);
                }
            }
        }
    }
    let auto = stack
        .iter()
        .rev()
        .fold(0, |acc, c| 5 * acc + auto_score(*c));
    return (0, auto);
}

pub fn part1(input: String) -> i64 {
    input.lines().map(|line| parse_line(line).0).sum()
}

pub fn part2(input: String) -> i64 {
    let l: Vec<i64> = input
        .lines()
        .map(|line| parse_line(line).1)
        .filter(|v| *v > 0)
        .sorted()
        .collect();
    l[(l.len() - 1) / 2]
}
