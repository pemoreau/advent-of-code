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

pub fn part2(input: String) -> i64 {
    let mut cache: HashMap<i64, i64> = HashMap::new();
    let mut card = Card::new(2020, 119315717514047);
    let mut deltas: Vec<i64> = Vec::new();
    deltas.push(card.get_pos());
    let mut cpt = 3;
    for i in 0i64..101741582076661 {
        // for i in 0i64..100000 {
        card.run(input.as_str());
        let delta = ((card.get_pos() - deltas[i as usize]) + 119315717514047) % 119315717514047;
        deltas.push(delta);
        // println!("{}", delta);
        if let Some(index) = cache.get(&delta) {
            println!(
                "FOUND index={}\ti={}\tdiff={}\tdelta={}",
                index,
                i,
                i - index,
                delta
            );
            // cpt -= 1;
            // if cpt == 0 {
            //     break;
            // }
        } else {
            cache.insert(delta, i);
        }
    }

    card.get_pos() as i64
}
