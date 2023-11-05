use std::collections::HashSet;

struct Tile {
    biodiversity: i64,
}

impl Tile {
    fn new() -> Tile {
        Tile { biodiversity: 0 }
    }

    fn neighbors(&self, index: i8) -> Vec<i8> {
        let mut result = Vec::new();
        if index % 5 != 0 {
            result.push(index - 1);
        }
        if index % 5 != 4 {
            result.push(index + 1);
        }
        if index > 4 {
            result.push(index - 5);
        }
        if index < 20 {
            result.push(index + 5);
        }
        result
    }
    fn level_neighbors(&self, index: i8, level: i8) -> Vec<i8> {
        let mut result = Vec::new();

        result
    }

    fn nb_adjacent_bugs(&self, index: i8) -> i64 {
        let mut result = 0;
        for neighbor in self.neighbors(index) {
            if self.biodiversity & (1 << neighbor) != 0 {
                result += 1;
            }
        }
        result
    }

    fn step(&mut self) {
        let mut new_biodiversity = 0;
        for i in 0..25 {
            let nb_adjacent_bugs = self.nb_adjacent_bugs(i);
            if self.biodiversity & (1 << i) != 0 {
                if nb_adjacent_bugs == 1 {
                    new_biodiversity |= 1 << i;
                }
            } else {
                if nb_adjacent_bugs == 1 || nb_adjacent_bugs == 2 {
                    new_biodiversity |= 1 << i;
                }
            }
        }
        self.biodiversity = new_biodiversity;
    }
}

pub fn part1(input: String) -> i64 {
    let mut total = 0;
    let mut power = 1;
    for line in input.lines() {
        for c in line.chars() {
            if c == '#' {
                total += power;
            }
            power *= 2;
        }
    }
    let mut tile = Tile {
        biodiversity: total,
    };

    let mut seen = HashSet::new();
    seen.insert(tile.biodiversity);
    loop {
        tile.step();
        if seen.contains(&tile.biodiversity) {
            return tile.biodiversity;
        }
        seen.insert(tile.biodiversity);
    }
}

pub fn part2(input: String) -> i64 {
    let mut total = 0;
    let mut power = 1;
    for line in input.lines() {
        for c in line.chars() {
            if c == '#' {
                total += power;
            }
            power *= 2;
        }
    }
    let mut tile = Tile {
        biodiversity: total,
    };

    0
}
