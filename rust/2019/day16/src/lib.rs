struct Pattern {
    element_number: i64,
    repeat_cpt: i64,
    index: i64,
    base_pattern: Vec<i64>,
}

impl Iterator for Pattern {
    type Item = i64;

    fn next(&mut self) -> Option<Self::Item> {
        if self.repeat_cpt == 0 {
            self.index = (self.index + 1) % self.base_pattern.len() as i64;
            self.repeat_cpt = self.element_number;
        }
        self.repeat_cpt -= 1;
        Some(self.base_pattern[self.index as usize])
    }
}

// Returns a Pattern sequence generator
fn pattern(number: i64) -> Pattern {
    Pattern {
        element_number: number,
        repeat_cpt: number,
        index: 0,
        base_pattern: vec![0, 1, 0, -1],
    }
}

pub fn part1(input: String) -> i64 {
    let mut input: Vec<i64> = input
        .chars()
        .map(|c| c.to_digit(10).unwrap() as i64)
        .collect();
    for i in pattern(2).skip(1).take(20) {
        println!("> {}", i);
    }

    let mut output = Vec::new();
    for phase in 1..=100 {
        output.clear();
        for i in 0..input.len() {
            let mut sum = 0;
            let mut j = 0;
            for p in pattern(i as i64 + 1).skip(1).take(input.len()) {
                // println!("{} * {} = {}", input[j], p, input[j] * p);
                sum += input[j] * p;
                j += 1;
            }
            output.push(sum.abs() % 10);
        }
        println!("Phase {}: {:?}", phase, output);
        input = output.clone();
    }
    output.iter().take(8).fold(0, |acc, x| acc * 10 + x)
}


pub fn part2(input: String) -> i64 {
    0
}
