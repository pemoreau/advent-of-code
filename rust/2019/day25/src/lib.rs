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
    dir: i8,
    pos: Pos,
}

// north (1), south (2), west (3), and east (4)
const NORTH: i8 = 1;
const SOUTH: i8 = 2;
const WEST: i8 = 3;
const EAST: i8 = 4;

const WALL: i8 = 0;
const MOVED: i8 = 1;

const SENSITIVE_FLOOR: i8 = 2;

fn display_output(machine: &mut Machine) {
    let output = machine.get_last_output();
    println!(
        "{}",
        output.iter().map(|&c| c as u8 as char).collect::<String>()
    );
}

fn send_command(machine: &mut Machine, command: &str) -> String {
    println!("> {}", command);
    for c in command.chars() {
        machine.put_input(c as i64);
    }
    machine.put_input('\n' as i64);
    machine.run();
    let output = machine.get_last_output();
    let res = output.iter().map(|&c| c as u8 as char).collect::<String>();
    println!("{}", res);
    res
}

impl Droid {
    fn new(code: Vec<i64>) -> Self {
        Self {
            machine: Machine::new(code, vec![]),
            dir: NORTH,
            pos: Pos { x: 0, y: 0 },
        }
    }

    fn start(&mut self) {
        self.machine.run();
        display_output(&mut self.machine);
    }

    fn send_command(&mut self, command: &str) -> String {
        send_command(&mut self.machine, command)
    }

    fn movement(&mut self, command: &str) -> i8 {
        println!("> {}", command);
        let out_string = send_command(&mut self.machine, command);
        // for c in command.chars() {
        //     self.machine.put_input(c as i64);
        // }
        // self.machine.put_input('\n' as i64);
        // self.machine.run();
        //
        // let output = self.machine.get_last_output();
        // let out_string = output.iter().map(|&c| c as u8 as char).collect::<String>();
        println!("{}", out_string);

        if out_string.contains("You can't go that way.") {
            return WALL;
        }
        if out_string.contains("Alert! Droids on this ship are heavier than the detected value!") {
            return SENSITIVE_FLOOR;
        }
        if out_string.contains("Items here:") {
            let lines = out_string.split("\n").collect::<Vec<&str>>();
            // filter lines after "Items here:"
            let items = lines
                .iter()
                .skip_while(|&&l| l != "Items here:")
                .skip(1)
                .take_while(|&&l| l != "")
                .map(|&l| l.trim_start_matches("- ").to_string())
                .collect::<Vec<String>>();
            println!("items: {:?}", items);
            for item in items {
                if item == "infinite loop"
                    || item == "photons"
                    || item == "giant electromagnet"
                    || item == "escape pod"
                    || item == "molten lava"
                {
                    continue;
                }
                send_command(&mut self.machine, &format!("take {}", item));
            }
        }
        MOVED
    }

    fn move_forward(&mut self) -> i8 {
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
        if res == MOVED {
            self.pos = self.front_pos();
        }
        res
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

    fn check_right(&mut self) -> i8 {
        self.turn_right();
        let right = self.move_forward();
        if right == MOVED {
            self.turn_right();
            self.turn_right();
            self.move_forward();
            self.turn_right();
        } else {
            self.turn_left();
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

    fn start(&mut self) {
        self.droid.start();
    }

    fn send_command(&mut self, command: &str) -> String {
        self.droid.send_command(command)
    }

    fn explore_right_hand(&mut self) {
        let mut visited: HashSet<(Pos, i8)> = HashSet::new();
        visited.insert((self.droid.pos, self.droid.dir));
        let mut cpt = 0;

        while cpt == 0
            || !(self.droid.pos == Pos { x: 0, y: 0 }
                && self.droid.dir == EAST
                && visited.contains(&(self.droid.pos, self.droid.dir)))
        {
            let right = self.droid.check_right();
            if right == MOVED {
                println!("turn right and move forward");
                self.droid.turn_right();
                self.move_forward();
            } else {
                println!(" wall on right, try moving forward");
                if right == WALL {
                    self.grid.insert(self.droid.right_pos(), '#');
                } else if right == SENSITIVE_FLOOR {
                    self.grid.insert(self.droid.right_pos(), 'S');
                }
                let front = self.move_forward();
                if front != MOVED {
                    println!(" wall in front, turn left");
                    self.droid.turn_left()
                }
            }
            visited.insert((self.droid.pos, self.droid.dir));
            cpt += 1;
            self.display();
        }
    }

    fn move_forward(&mut self) -> i8 {
        let front = self.droid.move_forward();
        if front == WALL {
            println!("  wall in front, insert in grid");
            self.grid.insert(self.droid.front_pos(), '#');
        } else {
            println!("  moved forward, insert {:?} in grid", self.droid.pos);
            self.grid.insert(self.droid.pos, '.');
        }
        front
    }

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

    // fn flood(&mut self) -> i64 {
    //     let directions = vec!["north", "south", "west", "east"];
    //     let mut cpt = 0;
    //     let mut visited = HashSet::new();
    //     let mut todo = directions.clone();
    //     while todo.len() > 0 {
    //         let to_visit: Vec<Pos> = todo.iter().flat_map(|p| self.neighbors(*p)).collect();
    //         todo = to_visit
    //             .iter()
    //             .filter(|p| !visited.contains(*p))
    //             .cloned()
    //             .collect();
    //         for c in &todo {
    //             visited.insert(c.clone());
    //             self.grid.insert(*c, 'O');
    //         }
    //
    //         if todo.len() > 0 {
    //             cpt += 1;
    //             // println!("step: {}", cpt);
    //             // self.display();
    //         }
    //     }
    //     cpt
    // }

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
                    match self.droid.dir {
                        NORTH => print!("^"),
                        SOUTH => print!("v"),
                        WEST => print!("<"),
                        EAST => print!(">"),
                        _ => {
                            panic!("invalid direction")
                        }
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

fn powerset<T>(s: &[T]) -> Vec<Vec<&T>> {
    (0..2usize.pow(s.len() as u32))
        .map(|i| {
            s.iter()
                .enumerate()
                .filter(|&(t, _)| (i >> t) % 2 == 1)
                .map(|(_, element)| element)
                .collect()
        })
        .collect()
}
pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);

    let mut cabinet = Cabinet::new(code);
    cabinet.start();
    cabinet.explore_right_hand();
    cabinet.display();

    // move to Sensor
    for command in vec![
        "east", "east", "east", "south", "south", "east", "east", "inv",
    ] {
        cabinet.send_command(command);
    }

    for tuple in powerset(&vec![
        "food ration",
        "weather machine",
        "antenna",
        "space law space brochure",
        "jam",
        "semiconductor",
        "planetoid",
        "monolith",
    ]) {
        println!("{:?}", tuple);
        for command in tuple.clone() {
            let mut buf: String = "".to_owned();
            buf.push_str("drop ");
            buf.push_str(command);
            cabinet.send_command(&buf);
        }
        cabinet.send_command("east");
        for command in tuple {
            let mut buf: String = "".to_owned();
            buf.push_str("take ");
            buf.push_str(command);
            cabinet.send_command(&buf);
        }
    }

    0
}

pub fn part2(input: String) -> i64 {
    0
}
