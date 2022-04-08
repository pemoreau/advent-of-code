use std::collections::HashMap;

pub fn part1(input: String) -> i64 {
    let mut state = State {
        or_mask: 0,
        and_mask: 0,
        mem: HashMap::new(),
    };

    for line in input.lines() {
        println!("{}", line);
        match parse_line(line) {
            Action::Mask(mask) => state.set_mask(&mask),
            Action::Mem(addr, val) => state.set_mem(addr, val),
        }
    }

    state.get_sum() as i64
}

pub fn part2(input: String) -> i64 {
    0
}

struct State {
    or_mask: u64,  // X->0, 1->1, 0->0
    and_mask: u64, // X->1,, 1->1, 0->0
    mem: HashMap<u64, u64>,
}

impl State {
    fn set_mask(&mut self, mask: &str) {
        self.or_mask = u64::from_str_radix(&mask.replace("X", "0"), 2).unwrap();
        self.and_mask = u64::from_str_radix(&mask.replace("X", "1"), 2).unwrap();
        println!("or={}, and={}", self.or_mask, self.and_mask);
    }

    fn set_mem(&mut self, addr: u64, val: u64) {
        let masked = (val & self.and_mask) | self.or_mask;
        self.mem.insert(addr, masked);
    }

    fn get_sum(&self) -> u64 {
        self.mem.values().sum()
    }
}

pub enum Action {
    Mask(String),
    Mem(u64, u64),
}

fn parse_line(s: &str) -> Action {
    peg::parser! {
      grammar parser() for str {
        rule number() -> u64
          = n:$(['0'..='9']+) { n.parse().unwrap() }

        pub(crate) rule line() -> (Action)
          = "mask = "  mask:$(['X'|'1'|'0']+) { Action::Mask(mask.to_string()) }
          / "mem[" adr:number() "] = " value:number() { Action::Mem(adr,value) }
      }
    }

    parser::line(s).unwrap()
}
