use intcode::Machine;
use std::collections::HashMap;
use utils::parsing::comma_separated_to_numbers;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Pos {
    x: i64,
    y: i64,
}

#[derive(Debug)]
struct Cabinet {
    grid: HashMap<Pos, i64>,
    ball: Pos,
    dir: Pos,
}

impl Cabinet {
    fn new() -> Self {
        Self {
            grid: HashMap::new(),
            ball: Pos { x: 0, y: 0 },
            dir: Pos { x: 0, y: 0 },
        }
    }

    fn init(&mut self, code: Vec<i64>) {
        let mut machine = Machine::new(code, vec![]);
        machine.run();
        let output = machine.get_last_output();
        for i in 0..output.len() / 3 {
            let x = output[i * 3];
            let y = output[i * 3 + 1];
            let tile = output[i * 3 + 2];
            if tile == 4 {
                self.ball = Pos { x, y };
                self.dir = Pos { x: -1, y: -1 };
            } else {
                self.grid.insert(Pos { x, y }, tile);
            }
        }
    }

    fn display(&self) {
        let min_x = self.grid.keys().map(|p| p.x).min().unwrap();
        let max_x = self.grid.keys().map(|p| p.x).max().unwrap();
        let min_y = self.grid.keys().map(|p| p.y).min().unwrap();
        let max_y = self.grid.keys().map(|p| p.y).max().unwrap();
        for y in (min_y..=max_y) {
            for x in min_x..=max_x {
                if x == self.ball.x && y == self.ball.y {
                    print!("o");
                    continue;
                }
                let tile = self.grid.get(&Pos { x, y }).unwrap_or(&0);
                let tile = match tile {
                    0 => " ",
                    1 => "█",
                    2 => "X",
                    3 => "▁",
                    4 => "o",
                    _ => " ",
                };
                print!("{}", tile);
            }
            println!();
        }
    }

    fn game_over(&self) -> bool {
        let min_x = self.grid.keys().map(|p| p.x).min().unwrap();
        let max_x = self.grid.keys().map(|p| p.x).max().unwrap();
        let min_y = self.grid.keys().map(|p| p.y).min().unwrap();
        let max_y = self.grid.keys().map(|p| p.y).max().unwrap();

        !(min_x <= self.ball.x
            && self.ball.x <= max_x
            && min_y <= self.ball.y
            && self.ball.y <= max_y)
    }

    fn step(&mut self) {
        // move ball
        let next = Pos {
            x: self.ball.x + self.dir.x,
            y: self.ball.y + self.dir.y,
        };
        let ball = self.ball.clone();
        let tile = self.grid.get(&next).unwrap_or(&0);
        if *tile == 0 || *tile == 2 {
            self.grid.insert(ball, 0);
            self.ball = next;
        } else if *tile == 1 || *tile == 3 {
            self.dir.x *= -1;
            self.dir.y *= -1;
            self.ball = Pos {
                x: self.ball.x + self.dir.x,
                y: self.ball.y + self.dir.y,
            };
        }
    }
}

pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut game = Cabinet::new();
    game.init(code);
    game.display();
    // while !game.game_over() {
    //     game.step();
    //     game.display();
    // }
    game.grid.values().filter(|&&t| t == 2).count() as i64
}

pub fn part2(input: String) -> i64 {
    0
}
