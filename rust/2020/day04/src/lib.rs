extern crate nom;

use std::collections::{HashMap, HashSet};

use nom::{
    branch::alt,
    bytes::complete::tag,
    character::complete::{alpha1, alphanumeric1, digit1, line_ending, one_of},
    combinator::{map, opt, recognize},
    multi::{count, many1},
    sequence::{pair, terminated},
    IResult,
};

type Entry<'a> = (&'a str, &'a str);
type Passport<'a> = HashMap<&'a str, &'a str>;

fn parse_key(input: &str) -> IResult<&str, &str> {
    terminated(alpha1, tag(":"))(input)
}

fn parse_value(input: &str) -> IResult<&str, &str> {
    recognize(pair(opt(tag("#")), alphanumeric1))(input)
}

fn parse_entry(input: &str) -> IResult<&str, Entry> {
    pair(parse_key, parse_value)(input)
}

fn parse_entries(input: &str) -> IResult<&str, Passport> {
    map(
        many1(terminated(parse_entry, alt((tag(" "), line_ending)))),
        |vec| vec.into_iter().collect(), // transforms list of tuples into hashmap
    )(input)
}

fn parse_passports(input: &str) -> IResult<&str, Vec<Passport>> {
    many1(terminated(parse_entries, line_ending))(input)
}

fn valid_passport(p: &Passport) -> bool {
    let mandatory_keys: HashSet<_> = ["eyr", "iyr", "byr", "ecl", "pid", "hcl", "hgt"]
        .iter()
        .cloned()
        .collect();
    let keys: HashSet<_> = p.keys().cloned().collect();
    mandatory_keys.is_subset(&keys)
}

pub fn part1(input: String) -> i64 {
    assert_eq!(parse_key("abc:"), Ok(("", "abc")));
    assert_eq!(parse_value("cfa07d"), Ok(("", "cfa07d")));
    assert_eq!(parse_value("#cfa07d"), Ok(("", "#cfa07d")));
    assert_eq!(parse_entry("abc:#cfa07d"), Ok(("", ("abc", "#cfa07d"))));

    let (_, passports) = parse_passports(&input).unwrap();
    let res = passports.iter().filter(|p| valid_passport(&p)).count();
    res.try_into().unwrap()
}

fn check_byr(input: &str) -> bool {
    input.len() == 4 && (1920..=2002).contains(&input.parse::<u32>().unwrap())
}

fn check_iyr(input: &str) -> bool {
    input.len() == 4 && (2010..=2020).contains(&input.parse::<u32>().unwrap())
}

fn check_eyr(input: &str) -> bool {
    input.len() == 4 && (2020..=2030).contains(&input.parse::<u32>().unwrap())
}

fn check_hgt(input: &str) -> bool {
    let res: IResult<&str, (&str, &str)> = pair(digit1, alt((tag("cm"), tag("in"))))(input);
    if let Ok((_, (digits, unit))) = res {
        let size: u32 = digits.parse().unwrap_or_default();
        (unit == "cm" && (150..=193).contains(&size)) || (unit == "in" && (59..=76).contains(&size))
    } else {
        false
    }
}

fn check_hcl(input: &str) -> bool {
    let res: IResult<&str, &str> = recognize(pair(
        tag("#"),
        count(
            alt((one_of("0123456789"), one_of("abcdefghijklmnopqrstuvwxyz"))),
            6,
        ),
    ))(input);
    input.len() == 7 && res.is_ok()
}

fn check_pid(input: &str) -> bool {
    let res: IResult<&str, &str> = recognize(count(one_of("0123456789"), 9))(input);
    input.len() == 9 && res.is_ok()
}

fn check_ecl(input: &str) -> bool {
    ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"].contains(&input)
}

fn checked_passport(p: &Passport) -> bool {
    let byr = p.get("byr").unwrap();
    let iyr = p.get("iyr").unwrap();
    let eyr = p.get("eyr").unwrap();
    let hgt = p.get("hgt").unwrap();
    let hcl = p.get("hcl").unwrap();
    let ecl = p.get("ecl").unwrap();
    let pid = p.get("pid").unwrap();

    check_byr(byr)
        && check_iyr(iyr)
        && check_eyr(eyr)
        && check_hgt(hgt)
        && check_hcl(hcl)
        && check_ecl(ecl)
        && check_pid(pid)
}

pub fn part2(input: String) -> i64 {
    let (_, passports) = parse_passports(&input[..]).unwrap();
    let passports = passports
        .iter()
        .filter(|p| valid_passport(&p) && checked_passport(&p));
    passports.count().try_into().unwrap()
}
