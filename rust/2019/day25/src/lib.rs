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
    dir: i64,
    pos: Pos,
}

// north (1), south (2), west (3), and east (4)
const NORTH: i64 = 1;
const SOUTH: i64 = 2;
const WEST: i64 = 3;
const EAST: i64 = 4;

const WALL: i32 = 0;
const MOVED: i32 = 1;

impl Droid {
    fn new(code: Vec<i64>) -> Self {
        Self {
            machine: Machine::new(code, vec![]),
            dir: NORTH,
            pos: Pos { x: 0, y: 0 },
        }
    }

    fn movement(&mut self, command: &str) -> i32 {
        for c in command.chars() {
            self.machine.put_input(c as i64);
        }
        self.machine.put_input('\n' as i64);
        self.machine.run();

        let output = self.machine.get_last_output();
        let out_string = output.iter().map(|&c| c as u8 as char).collect::<String>();
        if out_string.contains("You can't go that way.") {
            return WALL;
        }
        MOVED
    }

    fn move_forward(&mut self) -> i64 {
        let dir_string: &str = match self.dir {
            NORTH => "north",
            SOUTH => "south",
            WEST => "west",
            EAST => "east",
            _ => {
                panic!("invalid direction")
            }
        };
        let res = self.movement(dir_string);
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
    droid: Droid,
}

impl Cabinet {
    fn new(code: Vec<i64>) -> Self {
        let mut res = Self {
            grid: HashMap::new(),
            droid: Droid::new(code),
        };
        res.grid.insert(res.droid.pos, 'o');
        res
    }

    // fn move_forward(&mut self) -> i64 {
    //     let front = self.droid.move_forward();
    //     if front == WALL {
    //         self.grid.insert(self.droid.front_pos(), '#');
    //     } else {
    //         self.grid.insert(self.droid.pos, '.');
    //         if front == OXYGEN {
    //             self.oxygen = Some(self.droid.pos);
    //         }
    //     }
    //     front
    // }

    fn neighbors(&self, pos: Pos) -> Vec<Pos> {
        let mut res = vec![];
        let Pos { x, y } = pos;
        res.push(Pos { x: x + 1, y });
        res.push(Pos { x: x - 1, y });
        res.push(Pos { x, y: y + 1 });
        res.push(Pos { x, y: y - 1 });
        res.iter()
            .filter(|p| self.grid.get(p) == Some(&'.'))
            .cloned()
            .collect()
    }

    fn flood(&mut self) -> i64 {
        let directions = vec!["north", "south", "west", "east"];
        let mut cpt = 0;
        let mut visited = HashSet::new();
        let mut todo = directions.clone();
        while todo.len() > 0 {
            let to_visit: Vec<Pos> = todo.iter().flat_map(|p| self.neighbors(*p)).collect();
            todo = to_visit
                .iter()
                .filter(|p| !visited.contains(*p))
                .cloned()
                .collect();
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
                    print!("D");
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

fn display_output(machine: &mut Machine) {
    let output = machine.get_last_output();
    println!(
        "{}",
        output.iter().map(|&c| c as u8 as char).collect::<String>()
    );
}

fn send_command(machine: &mut Machine, command: &str) {
    println!("> {}", command);
    for c in command.chars() {
        machine.put_input(c as i64);
    }
    machine.put_input('\n' as i64);
    machine.run();
    display_output(machine);
}

pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut machine = Machine::new(code, vec![]);

    machine.run();

    display_output(&mut machine);

    send_command(&mut machine, "south");
    send_command(&mut machine, "north");
    // send_command(&mut machine, "take infinite loop");
    send_command(&mut machine, "inv");
    send_command(&mut machine, "south");
    send_command(&mut machine, "east");
    send_command(&mut machine, "east");
    send_command(&mut machine, "take semiconductor");
    send_command(&mut machine, "east");

    0
}

pub fn part2(input: String) -> i64 {
    0
}
