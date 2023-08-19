struct Pattern {
    element_number: i32,
    repeat_cpt: i32,
    index: i32,
    base_pattern: Vec<i32>,
}

impl Iterator for Pattern {
    type Item = i32;

    fn next(&mut self) -> Option<Self::Item> {
        if self.repeat_cpt == 0 {
            self.index = (self.index + 1) % self.base_pattern.len() as i32;
            self.repeat_cpt = self.element_number;
        }
        self.repeat_cpt -= 1;
        Some(self.base_pattern[self.index as usize])
    }
}

// Returns a Pattern sequence generator
fn pattern(number: i32) -> Pattern {
    Pattern {
        element_number: number,
        repeat_cpt: number,
        index: 0,
        base_pattern: vec![0, 1, 0, -1],
    }
}

pub fn part1(input: String) -> i64 {
    let mut input: Vec<i32> = input
        .chars()
        .map(|c| c.to_digit(10).unwrap() as i32)
        .collect();

    let mut output = Vec::new();
    for _ in 1..=100 {
        output.clear();
        for i in 0..input.len() {
            let mut sum = 0;
            let mut j = 0;
            for p in pattern(i as i32 + 1).skip(1).take(input.len()) {
                sum += input[j] * p;
                j += 1;
            }
            output.push(sum.abs() % 10);
        }
        input = output.clone();
    }
    output.iter().take(8).fold(0, |acc, x| acc * 10 + x) as i64
}

pub fn part2(input: String) -> i64 {
    let input: Vec<i32> = input
        .repeat(10000)
        .chars()
        .map(|c| c.to_digit(10).unwrap() as i32)
        .collect();
    let offset = input.iter().take(7).fold(0, |acc, x| acc * 10 + x) as usize;

    let mut output = input[offset..].to_vec();
    for _ in 1..=100 {
        let mut sum = 0;
        for i in (0..output.len()).rev() {
            sum += output[i];
            output[i] = sum.abs() % 10;
        }
    }

    output.iter().take(8).fold(0, |acc, x| acc * 10 + x) as i64
}
