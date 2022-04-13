use std::collections::HashMap;

pub fn part1(input: String) -> i64 {
    let e = parse_entries(&input[..]);

    // println!("{:?}", e);
    let valid = e
        .nearby
        .iter()
        .map(|f| invalid_nearby_ticket(f, &e.constraints))
        .sum::<u32>();
    println!("{:?}", valid);

    valid as i64
}

pub fn part2(input: String) -> i64 {
    0
}

#[derive(Debug)]
enum Constraint {
    Field(String, Interval, Interval),
}

#[derive(Debug)]
enum Interval {
    Range(u32, u32),
}

impl Interval {
    fn contains(&self, val: u32) -> bool {
        match self {
            Interval::Range(start, end) => val >= *start && val <= *end,
        }
    }
}

#[derive(Debug)]
struct Data {
    ticket: Vec<u32>,
    nearby: Vec<Vec<u32>>,
    constraints: Vec<Constraint>,
}

fn invalid_field(f: u32, constraints: &Vec<Constraint>) -> Option<u32> {
    let mut res = None;
    for c in constraints {
        match c {
            Constraint::Field(_, i, j) => {
                if !i.contains(f) && !j.contains(f) {
                    println!("{} {:?} {:?}", f, i, j);
                    res = Some(f);
                } else {
                    println!("ok {} {:?} {:?}", f, i, j);
                    return None;
                }
            }
        }
    }
    res
}

fn invalid_nearby_ticket(t: &Vec<u32>, constraints: &Vec<Constraint>) -> u32 {
    let res = t
        .iter()
        .filter(|f| invalid_field(**f, constraints).is_some())
        .sum();
    if res > 0 {
        println!("{:?} {}", t, res);
    }
    res

    // for n in t.iter() {
    //     if !valid_field(*n, &constraints) {
    //         return false;
    //     }
    // }
    // true
}

fn parse_entries(s: &str) -> Data {
    peg::parser! {
      grammar parser() for str {
        rule _() = [' ' | '\t' | '\r']*
        rule __() = ['\n']*

        rule identifier() -> String
          = id:$(['a'..='z']+) { id.parse().unwrap() }

        rule number() -> u32
          = n:$(['0'..='9']+) { n.parse().unwrap() }

        rule list() -> Vec<u32>
          = l:(number() ++ ",") __ { l }

        rule interval() -> Interval
          = a:number() "-" b:number() { Interval::Range(a, b) }

        rule constraint() -> Constraint
          = names:(identifier() ++ " ") ":"
            _ a:interval() _ "or" _ b:interval() __ {
              Constraint::Field(names.join(" "), a, b)
            }

        rule constraints() -> (Vec<Constraint>)
          = cc:(constraint())* __ { cc }

        pub(crate) rule entries() -> (Data)
          = cc:constraints() __
            "your ticket:" __ l:list()
            "nearby tickets:" __ ll:(list())* __ {
                Data {constraints:cc, ticket:l, nearby:ll}
            }

      }
    }

    parser::entries(s).unwrap()
}
