use intcode::Machine;
use std::collections::{HashMap, HashSet};
use utils::parsing::comma_separated_to_numbers;

#[derive(Clone, Debug, Eq, Hash, PartialEq, Copy)]
struct Pos {
    x: i64,
    y: i64,
}

#[derive(Debug)]
struct Droid {
    machine: Machine,
}

impl Droid {
    fn new(code: Vec<i64>) -> Self {
        Self {
            machine: Machine::new(code, vec![]),
        }
    }

    fn explore(&mut self, dir: i64) -> i64 {
        self.machine.put_input(dir);
        self.machine.run();
        let output = self.machine.get_last_output();
        output[0]
    }
}

type Grid = HashMap<Pos, char>;

#[derive(Debug)]
struct Cabinet {
    grid: Grid,
    bot: Pos,
    dir: i64,
    oxygen: Option<Pos>,
    droid: Droid,
}

// north (1), south (2), west (3), and east (4)
const NORTH: i64 = 1;
const SOUTH: i64 = 2;
const WEST: i64 = 3;
const EAST: i64 = 4;


fn neighbor(p: Pos, d: i64) -> Pos {
    let Pos { x, y } = p;
    match d {
        NORTH => { Pos { x, y: y + 1 } }
        SOUTH => { Pos { x, y: y - 1 } }
        WEST => { Pos { x: x - 1, y } }
        EAST => { Pos { x: x + 1, y } }
        _ => { panic!("invalid direction") }
    }
}

fn dir_right(d: i64) -> i64 {
    match d {
        NORTH => { EAST }
        SOUTH => { WEST }
        WEST => { NORTH }
        EAST => { SOUTH }
        _ => { panic!("invalid direction") }
    }
}

// 0: The repair droid hit a wall. Its position has not changed.
// 1: The repair droid has moved one step in the requested direction.
// 2: The repair droid has moved one step in the requested direction; its new position is the location of the oxygen system.
const WALL: i64 = 0;
const MOVED: i64 = 1;
const OXYGEN: i64 = 2;

impl Cabinet {
    fn new(code: Vec<i64>) -> Self {
        Self {
            grid: HashMap::new(),
            bot: Pos { x: 0, y: 0 },
            dir: NORTH,
            oxygen: None,
            droid: Droid::new(code),
        }
    }

    fn turn_right(&mut self) {
        self.dir = dir_right(self.dir);
    }

    fn turn_left(&mut self) {
        self.turn_right();
        self.turn_right();
        self.turn_right();
    }

    fn move_forward(&mut self) -> i64 {
        let front = self.droid.explore(self.dir);
        let next = neighbor(self.bot, self.dir);
        if front == WALL {
            self.grid.insert(next.clone(), '#');
        } else {
            self.grid.insert(next.clone(), '.');
            self.bot = next.clone();
            if front == OXYGEN {
                self.oxygen = Some(next.clone());
            }
        }
        front
    }

    fn check_right(&mut self) -> i64 {
        self.turn_right();
        let right = self.droid.explore(self.dir);
        if right == WALL {
            self.turn_left();
        } else {
            self.turn_right();
            self.turn_right();
            self.droid.explore(self.dir);
            self.turn_right();
            self.turn_right();
        }
        right
    }


    fn explore(&mut self) {
        while self.oxygen.is_none() {
            if self.check_right() != 0 {
                self.turn_right();
                self.move_forward();
            } else {
                let front = self.move_forward();
                if front == WALL {
                    self.turn_left()
                }
            }
            self.display();
        }
    }

    fn display(&self) {
        if self.grid.is_empty() {
            return;
        }
        let min_x = self.grid.keys().map(|p| p.x).min().unwrap();
        let max_x = self.grid.keys().map(|p| p.x).max().unwrap();
        let min_y = self.grid.keys().map(|p| p.y).min().unwrap();
        let max_y = self.grid.keys().map(|p| p.y).max().unwrap();
        for y in (min_y..=max_y).rev() {
            for x in min_x..=max_x {
                if x == self.bot.x && y == self.bot.y {
                    if self.dir == NORTH {
                        print!("^");
                    } else if self.dir == SOUTH {
                        print!("v");
                    } else if self.dir == WEST {
                        print!("<");
                    } else if self.dir == EAST {
                        print!(">");
                    }
                    continue;
                }
                let tile = self.grid.get(&Pos { x, y }).unwrap_or(&' ');
                print!("{}", tile);
            }
            println!();
        }
        println!();
    }
}

pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut game = Cabinet::new(code);
    game.explore();

    0
}

pub fn part2(input: String) -> i64 {
    0
}
