use std::collections::{HashMap, HashSet};

use intcode::Machine;
use utils::parsing::comma_separated_to_numbers;

#[derive(Clone, Debug, Eq, Hash, PartialEq, Copy)]
struct Pos {
    x: i64,
    y: i64,
}

// north (1), south (2), west (3), and east (4)
const NORTH: i64 = 1;
const SOUTH: i64 = 2;
const WEST: i64 = 3;
const EAST: i64 = 4;

// 0: The repair droid hit a wall. Its position has not changed.
// 1: The repair droid has moved one step in the requested direction.
// 2: The repair droid has moved one step in the requested direction; its new position is the location of the oxygen system.
const WALL: i64 = 0;
const MOVED: i64 = 1;
const OXYGEN: i64 = 2;

#[derive(Debug)]
struct Droid {
    machine: Machine,
    dir: i64,
    pos: Pos,
}

impl Droid {
    fn new(code: Vec<i64>) -> Self {
        Self {
            machine: Machine::new(code, vec![]),
            dir: NORTH,
            pos: Pos { x: 0, y: 0 },
        }
    }

    fn move_forward(&mut self) -> i64 {
        self.machine.put_input(self.dir);
        self.machine.run();
        let output = self.machine.get_last_output();
        if output[0] != WALL {
            self.pos = self.front_pos();
        }
        output[0]
    }

    fn turn_right(&mut self) {
        self.dir = match self.dir {
            NORTH => EAST,
            SOUTH => WEST,
            WEST => NORTH,
            EAST => SOUTH,
            _ => {
                panic!("invalid direction")
            }
        }
    }

    fn turn_left(&mut self) {
        self.turn_right();
        self.turn_right();
        self.turn_right();
    }

    fn front_pos(&mut self) -> Pos {
        match self.dir {
            NORTH => Pos {
                x: self.pos.x,
                y: self.pos.y + 1,
            },
            SOUTH => Pos {
                x: self.pos.x,
                y: self.pos.y - 1,
            },
            WEST => Pos {
                x: self.pos.x - 1,
                y: self.pos.y,
            },
            EAST => Pos {
                x: self.pos.x + 1,
                y: self.pos.y,
            },
            _ => {
                panic!("invalid direction")
            }
        }
    }

    fn right_pos(&mut self) -> Pos {
        self.turn_right();
        let pos = self.front_pos();
        self.turn_left();
        pos
    }


    fn check_right(&mut self) -> i64 {
        self.turn_right();
        let right = self.move_forward();
        if right == WALL {
            self.turn_left();
        } else {
            self.turn_right();
            self.turn_right();
            self.move_forward();
            self.turn_right();
        }
        right
    }
}

type Grid = HashMap<Pos, char>;

#[derive(Debug)]
struct Cabinet {
    grid: Grid,
    oxygen: Option<Pos>,
    droid: Droid,
}

impl Cabinet {
    fn new(code: Vec<i64>) -> Self {
        let mut res = Self {
            grid: HashMap::new(),
            oxygen: None,
            droid: Droid::new(code),
        };
        res.grid.insert(res.droid.pos, 'o');
        res
    }

    fn move_forward(&mut self) -> i64 {
        let front = self.droid.move_forward();
        if front == WALL {
            self.grid.insert(self.droid.front_pos(), '#');
        } else {
            self.grid.insert(self.droid.pos, '.');
            if front == OXYGEN {
                self.oxygen = Some(self.droid.pos);
            }
        }
        front
    }

    fn explore_right_hand(&mut self) -> i64 {
        let mut visited = HashSet::new();
        visited.insert(self.droid.pos);
        let mut cpt = 0;
        while cpt == 0 || self.droid.pos != (Pos { x: 0, y: 0 }) {
            let old_pos = self.droid.pos;
            if self.droid.check_right() != WALL {
                self.droid.turn_right();
                self.move_forward();
            } else {
                self.grid.insert(self.droid.right_pos(), '#');
                let front = self.move_forward();
                if front == WALL {
                    self.droid.turn_left()
                }
            }
            if self.droid.pos != old_pos && self.oxygen.is_none() {
                if visited.contains(&self.droid.pos) {
                    cpt -= 1;
                } else {
                    visited.insert(self.droid.pos);
                    cpt += 1;
                }
            }
        }
        cpt + 1
    }

    fn neighbors(&self, pos: Pos) -> Vec<Pos> {
        let mut res = vec![];
        let Pos { x, y } = pos;
        res.push(Pos { x: x + 1, y });
        res.push(Pos { x: x - 1, y });
        res.push(Pos { x, y: y + 1 });
        res.push(Pos { x, y: y - 1 });
        res.iter().filter(|p| self.grid.get(p) == Some(&'.')).cloned().collect()
    }


    fn flood(&mut self) -> i64 {
        let mut cpt = 0;
        let mut visited = HashSet::new();
        let mut todo = vec![self.oxygen.unwrap()];
        while todo.len() > 0 {
            let to_visit: Vec<Pos> = todo.iter().flat_map(|p| self.neighbors(*p)).collect();
            todo = to_visit.iter().filter(|p| !visited.contains(*p)).cloned().collect();
            for c in &todo {
                visited.insert(c.clone());
                self.grid.insert(*c, 'O');
            }

            if todo.len() > 0 {
                cpt += 1;
                // println!("step: {}", cpt);
                // self.display();
            }
        }
        cpt
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
                if x == self.droid.pos.x && y == self.droid.pos.y {
                    if self.droid.dir == NORTH {
                        print!("^");
                    } else if self.droid.dir == SOUTH {
                        print!("v");
                    } else if self.droid.dir == WEST {
                        print!("<");
                    } else if self.droid.dir == EAST {
                        print!(">");
                    }
                    continue;
                }
                if self.oxygen.is_some()
                    && x == self.oxygen.unwrap().x
                    && y == self.oxygen.unwrap().y
                {
                    print!("O");
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
    let cpt = game.explore_right_hand();
    cpt
}

pub fn part2(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut game = Cabinet::new(code);
    game.explore_right_hand();
    // game.display();
    game.flood()
}
