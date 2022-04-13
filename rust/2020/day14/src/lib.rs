use std::collections::HashMap;

struct State {
    or_mask: u64,  // X->0, 1->1, 0->0
    and_mask: u64, // X->1, 1->1, 0->0
    mem: HashMap<u64, u64>,
}

impl State {
    fn set_mask(&mut self, mask: &str) {
        self.or_mask = u64::from_str_radix(&mask.replace("X", "0"), 2).unwrap();
        self.and_mask = u64::from_str_radix(&mask.replace("X", "1"), 2).unwrap();
    }

    fn set_mem(&mut self, addr: u64, val: u64) {
        let masked = (val & self.and_mask) | self.or_mask;
        self.mem.insert(addr, masked);
    }

    fn get_sum(&self) -> u64 {
        self.mem.values().sum()
    }
}
pub fn part1(input: String) -> i64 {
    let mut state = State {
        or_mask: 0,
        and_mask: 0,
        mem: HashMap::new(),
    };

    for line in input.lines() {
        match parse_line(line) {
            Action::Mask(mask) => state.set_mask(&mask),
            Action::Mem(addr, val) => state.set_mem(addr, val),
        }
    }

    state.get_sum() as i64
}

pub fn part2(input: String) -> i64 {
    let mut state = State2 {
        mask: String::new(),
        mem: HashMap::new(),
    };

    for line in input.lines() {
        match parse_line(line) {
            Action::Mask(mask) => state.set_mask(&mask),
            Action::Mem(addr, val) => state.set_mem(addr, val),
        }
    }

    state.get_sum() as i64
}

struct State2 {
    mask: String,
    mem: HashMap<String, u64>,
}

impl State2 {
    fn set_mask(&mut self, mask: &str) {
        self.mask = mask.to_string();
    }

    fn get_sum(&self) -> u64 {
        self.mem.values().sum()
    }

    fn set_mem(&mut self, addr: u64, val: u64) {
        // set 1 in mask for each 1 bit of addr
        let xaddr: String = format!("{:036b}", addr)
            .chars()
            .enumerate()
            .map(|(i, c)| {
                if c == '1' && self.mask.chars().nth(i).unwrap() == '0' {
                    '1'
                } else {
                    self.mask.chars().nth(i).unwrap()
                }
            })
            .collect();

        fn recur_set(addr: &String, n: usize, mem: &mut HashMap<String, u64>, val: u64) {
            if n >= addr.len() {
                mem.insert(addr.clone(), val);
            } else {
                if addr.chars().nth(n).unwrap() == 'X' {
                    let mut new_addr = addr.clone();
                    new_addr.replace_range(n..n + 1, "0");
                    recur_set(&new_addr, n + 1, mem, val);
                    new_addr.replace_range(n..n + 1, "1");
                    recur_set(&new_addr, n + 1, mem, val);
                } else {
                    recur_set(addr, n + 1, mem, val);
                }
            }
        }

        recur_set(&xaddr, 0, &mut self.mem, val);
    }
}
pub enum Action {
    Mask(String),
    Mem(u64, u64),
}

fn parse_line(s: &str) -> Action {
    peg::parser! {
      grammar parser() for str {
        rule _() = [' ' | '\t' | '\r' | '\n']*

        rule number() -> u64
          = n:$(['0'..='9']+) { n.parse().unwrap() }

        pub(crate) rule line() -> (Action)
          = "mask" _  "=" _ mask:$(['X'|'1'|'0']+) { Action::Mask(mask.to_string()) }
          / "mem[" adr:number() "]" _ "=" _ value:number() { Action::Mem(adr,value) }
      }
    }

    parser::line(s).unwrap()
}
