pub struct Card {
    pos: i64,
    size: i64,
}

impl Card {
    pub fn new(value: i64, size: i64) -> Card {
        Card { pos: value, size }
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

pub fn mod_euclid(a: i128, b: u128) -> u128 {
    const UPPER: u128 = i128::MAX as u128;
    match b {
        1..=UPPER => a.rem_euclid(b as i128) as u128,
        _ if a >= 0 => a as u128,
        // turn a from two's complement negative into it's
        // equivalent positive value by adding u128::MAX
        // essentialy calculating u128::MAX - |a|
        _ => u128::MAX.wrapping_add_signed(a),
        //_ => a as u128 - (a < 0) as u128,
    }
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

    // println!("a={}\t\tb={}", a, b);

    // f^-1(x) = a^-1 * x -b * a^-1
    let aa = mod_pow(a, size - 2, size);
    // let B = (-b * A) % size;
    let bb = mod_euclid(-b * aa, size as u128) as i128;
    // f^-n(x) = a^-n * x - b * a^-n
    let p = mod_pow(aa, n, size);
    let q = mod_euclid((p - 1) * mod_pow(aa - 1, size - 2, size), size as u128) as i128;
    let res = (2020 * p + bb * q) % size;

    // println!("A={}\tB={}", aa, bb);
    // println!("p={}\tq={}", p, q);

    res as i64
}
