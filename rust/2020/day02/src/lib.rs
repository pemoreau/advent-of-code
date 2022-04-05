#[derive(Debug)]

pub struct Policy {
    min: usize,
    max: usize,
    letter: char,
}

fn _parse_line1(s: &str) -> (Policy, &str) {
    let re = regex::Regex::new(r"^(\d{1,2})-(\d{1,2}) ([a-z]{1}): ([a-z]+)$").unwrap();
    let cap = re.captures(s).unwrap();
    let password: &str = cap.get(4).unwrap().as_str(); // instead of &cap[4]

    (
        Policy {
            min: cap[1].parse().unwrap(),
            max: cap[2].parse().unwrap(),
            letter: cap[3].chars().next().unwrap(),
        },
        password,
    )
}

fn parse_line2(s: &str) -> (Policy, &str) {
    peg::parser! {
      grammar parser() for str {
        rule number() -> usize
          = n:$(['0'..='9']+) { n.parse().unwrap() }

        rule byte() -> char
          = letter:$(['a'..='z']) { letter.chars().next().unwrap() }

        rule password() -> &'input str
          = letters:$([_]*) { letters }

        pub(crate) rule line() -> (Policy, &'input str)
          = min:number() "-" max:number() " " letter:byte() ": " password:password() {
              (Policy { min,max,letter }, password)
          }
      }
    }

    parser::line(s).unwrap()
}

fn check(policy: &Policy, password: &str) -> bool {
    let c = password.chars().filter(|c| *c == policy.letter).count();
    c >= policy.min && c <= policy.max
}

fn valid(policy: &Policy, password: &str) -> bool {
    if policy.min == 0 || policy.max > password.len() {
        return false;
    }

    let c1 = password.chars().nth(policy.min - 1).unwrap();
    let c2 = password.chars().nth(policy.max - 1).unwrap();
    (policy.letter == c1) ^ (policy.letter == c2)
}

type FilterFn = fn(&Policy, &str) -> bool;

pub fn generic_part(input: String, filter_function: FilterFn) -> i64 {
    let values: Vec<_> = input.lines().map(|line| parse_line2(line)).collect();
    let number = values
        .iter()
        .filter(|(policy, password)| filter_function(policy, password))
        .count();
    number.try_into().unwrap()
}

pub fn part1(input: String) -> i64 {
    generic_part(input, check)
}

pub fn part2(input: String) -> i64 {
    generic_part(input, valid)
}
