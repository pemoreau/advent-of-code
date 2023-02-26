use intcode::Machine;
use std::collections::HashMap;
use utils::parsing::comma_separated_to_numbers;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Pos {
    x: i64,
    y: i64,
}

#[derive(Debug)]
struct Robot {
    pos: Pos,
    dir: usize,
}

impl Robot {
    fn new() -> Self {
        Self {
            pos: Pos { x: 0, y: 0 },
            dir: 0,
        }
    }

    fn turn(&mut self, dir: i64) {
        match dir {
            0 => self.dir = (self.dir + 3) % 4,
            1 => self.dir = (self.dir + 1) % 4,
            _ => panic!("Invalid direction"),
        }
    }

    fn move_forward(&mut self) {
        match self.dir {
            0 => self.pos.y += 1,
            1 => self.pos.x += 1,
            2 => self.pos.y -= 1,
            3 => self.pos.x -= 1,
            _ => panic!("Invalid direction"),
        }
    }
}

fn run_robot(code: Vec<i64>, start_color: i64) -> HashMap<Pos, i64> {
    let mut machine = Machine::new(code, vec![]);
    let mut grid = HashMap::new();
    let mut robot = Robot::new();
    grid.insert(robot.pos.clone(), start_color);

    loop {
        let color = grid.entry(robot.pos.clone()).or_insert(0);
        machine.put_input(*color);
        machine.run();
        if machine.is_halted() {
            break;
        }
        let output = machine.get_last_output();
        *color = output[output.len() - 2];
        let dir = output[output.len() - 1];
        robot.turn(dir);
        robot.move_forward();
    }
    grid
}

fn display_grid(grid: &HashMap<Pos, i64>) {
    let min_x = grid.keys().map(|p| p.x).min().unwrap();
    let max_x = grid.keys().map(|p| p.x).max().unwrap();
    let min_y = grid.keys().map(|p| p.y).min().unwrap();
    let max_y = grid.keys().map(|p| p.y).max().unwrap();
    for y in (min_y..=max_y).rev() {
        for x in min_x..=max_x {
            let color = grid.get(&Pos { x, y }).unwrap_or(&0);
            if *color == 0 {
                print!(" ");
            } else {
                print!("█");
            }
        }
        println!();
    }
}

pub fn part1(input: String) -> i64 {
    let grid = run_robot(comma_separated_to_numbers(input), 0);
    grid.len() as i64
}

pub fn part2(input: String) -> i64 {
    let grid = run_robot(comma_separated_to_numbers(input), 1);
    display_grid(&grid);
    0
}
