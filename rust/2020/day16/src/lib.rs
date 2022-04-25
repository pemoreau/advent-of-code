use std::collections::{HashMap, HashSet};

#[derive(Debug)]
enum Constraint {
    Field(String, Interval, Interval),
}

#[derive(Debug)]
enum Interval {
    Range(u32, u32),
}

#[derive(Debug)]
struct Data {
    ticket: Vec<u32>,
    nearby: Vec<Vec<u32>>,
    constraints: Vec<Constraint>,
}

impl Interval {
    fn contains(&self, val: u32) -> bool {
        match self {
            Interval::Range(start, end) => val >= *start && val <= *end,
        }
    }
}

pub fn part1(input: String) -> i64 {
    let e = parse_entries(&input[..]);
    let valid = e
        .nearby
        .iter()
        .map(|f| invalid_nearby_ticket(f, &e.constraints))
        .sum::<u32>();
    valid as i64
}

pub fn part2(input: String) -> i64 {
    let e = parse_entries(&input[..]);

    let valid_nearby = e
        .nearby
        .iter()
        .filter(|&f| invalid_nearby_ticket(f, &e.constraints) == 0)
        .collect::<Vec<_>>();

    // possible index for each field
    let mut map: HashMap<String, Vec<bool>> = HashMap::new();
    for c in e.constraints.iter() {
        let Constraint::Field(name, _, _) = c;
        let mut v = vec![];
        v.resize(e.ticket.len(), true);
        map.insert(name.clone(), v);
    }

    // eliminate index which are not satisfying constraints
    for c in e.constraints.iter() {
        let Constraint::Field(name, i1, i2) = c;
        for valid in valid_nearby.iter() {
            for (i, n) in valid.iter().enumerate() {
                if !i1.contains(*n) && !i2.contains(*n) {
                    let v = map.get_mut(name).unwrap();
                    v[i] = false;
                }
            }
        }
    }

    // find constraint with only one true and remove index from others
    let mut propagated: HashSet<String> = HashSet::new();
    let mut cpt = map.len();
    let mut res: i64 = 1;
    while cpt > 0 {
        // look for solved constraint
        let key = map
            .iter()
            .find(|(key, flags)| single_index(flags).is_some() && !propagated.contains(key.clone()))
            .unwrap()
            .0
            .to_string(); // needed to release a ref on map
        let flags = map.get(&key).unwrap();
        let index = single_index(flags).unwrap();
        propagated.insert(key.clone());
        cpt -= 1;
        if key.starts_with("departure") {
            res *= e.ticket[index] as i64;
        }
        for (name, flags) in map.iter_mut() {
            if *name != *key {
                flags[index] = false;
            }
        }
    }

    res
}

fn single_index(v: &Vec<bool>) -> Option<usize> {
    let mut res = None;
    for (i, &b) in v.iter().enumerate() {
        if b {
            if res.is_none() {
                res = Some(i);
            } else {
                return None;
            }
        }
    }
    res
}

fn invalid_field(f: u32, constraints: &Vec<Constraint>) -> Option<u32> {
    for c in constraints {
        let Constraint::Field(_, i, j) = c;
        if i.contains(f) || j.contains(f) {
            return None;
        }
    }
    Some(f)
}

fn invalid_nearby_ticket(t: &Vec<u32>, constraints: &Vec<Constraint>) -> u32 {
    t.iter()
        .filter(|f| invalid_field(**f, constraints).is_some())
        .sum()
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
