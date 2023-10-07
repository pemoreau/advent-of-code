use std::collections::{HashMap, HashSet};

pub struct Card {
    pos: i64,
    value: i64,
    size: i64,
}

impl Card {
    pub fn new(value: i64, size: i64) -> Card {
        Card {
            pos: value,
            value,
            size,
        }
    }
    pub fn new_stack(&mut self) {
        self.pos = self.size - self.pos - 1;
    }

    pub fn cut(&mut self, n: i64) {
        self.pos = (self.pos - n + self.size) % self.size;
    }

    pub fn deal(&mut self, n: i64) {
        self.pos = (self.pos * n) % self.size;
    }

    pub fn get_pos(&self) -> i64 {
        self.pos
    }

    fn interpret(&mut self, line: &str) {
        if line == "deal into new stack" {
            self.new_stack();
        } else if line.starts_with("cut") {
            let n = line.split(" ").nth(1).unwrap().parse::<i64>().unwrap();
            self.cut(n);
        } else if line.starts_with("deal with increment") {
            let n = line.split(" ").nth(3).unwrap().parse::<i64>().unwrap();
            self.deal(n);
        }
    }

    fn run(&mut self, input: &str) {
        for line in input.lines() {
            self.interpret(line);
        }
    }
}

pub fn part1(input: String) -> i64 {
    let mut card = Card::new(2019, 10007);
    card.run(input.as_str());
    card.get_pos() as i64
}

fn mod_pow(mut base: i128, mut exp: i128, modulus: i128) -> i128 {
    if modulus == 1 {
        return 0;
    }
    let mut result = 1;
    base %= modulus;
    while exp > 0 {
        if exp % 2 == 1 {
            result = (result * base) % modulus;
        }
        exp >>= 1;
        base = (base * base) % modulus;
    }
    result
}

fn step(a: i128, b: i128, size: i128, line: &str) -> (i128, i128) {
    if line == "deal into new stack" {
        return (-a + size, -b + size - 1);
    } else if line.starts_with("cut") {
        let n = line.split(" ").nth(1).unwrap().parse::<i128>().unwrap();
        return (a, (b - n + size) % size);
    } else if line.starts_with("deal with increment") {
        let n = line.split(" ").nth(3).unwrap().parse::<i128>().unwrap();
        return (a * n % size, b * n % size);
    }
    (a, b)
}
pub fn part2(input: String) -> i64 {
    let mut a = 1i128;
    let mut b = 0i128;
    let size = 119315717514047i128;
    let n = 101741582076661i128;

    for line in input.lines() {
        (a, b) = step(a, b, size, line);
    }

    println!("a={} b={}", a, b);

    // x * a^n + b * (a^n - 1) / (a-1)
    // let p = mod_pow(a, n, size);
    // let q = (mod_pow(a, n, size) - 1) * mod_pow(a - 1, size - 2, size) % size;
    // let res = (2020 * p + b * q) % size;

    // f^-1(x) = a^-1 * x -b * a^-1
    let c = mod_pow(a, size - 2, size);
    let res = (2020 * c - b * c) % size;

    // 5691574550498 too low
    // 39485958529798 too low
    // 107800053624821 too high
    // 110057814880073
    res as i64
}
