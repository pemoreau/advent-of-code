use intcode::Machine;
use std::collections::{HashMap, HashSet};
use utils::parsing::comma_separated_to_numbers;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Pos {
    x: i64,
    y: i64,
}

type Grid = HashMap<Pos, char>;

#[derive(Debug)]
struct Cabinet {
    grid: Grid,
    bot: Pos,
    oxygen: Option<Pos>,
}

// north (1), south (2), west (3), and east (4)
fn neighbors(p: Pos) -> Vec<(i64, Pos)> {
    let Pos { x, y } = p;
    vec![
        (4, Pos { x: x + 1, y }),
        (3, Pos { x: x - 1, y }),
        (2, Pos { x, y: y - 1 }),
        (1, Pos { x, y: y + 1 }),
    ]
}

impl Cabinet {
    fn new() -> Self {
        Self {
            grid: HashMap::new(),
            bot: Pos { x: 0, y: 0 },
            oxygen: None,
        }
    }

    // 0: The repair droid hit a wall. Its position has not changed.
    // 1: The repair droid has moved one step in the requested direction.
    // 2: The repair droid has moved one step in the requested direction; its new position is the location of the oxygen system.

    fn explore(&mut self, code: Vec<i64>) {
        let mut machine = Machine::new(code, vec![]);
        let mut visited = HashSet::new();
        let mut todo = vec![self.bot.clone()];
        while let Some(elt) = todo.pop() {
            for (dir, c) in neighbors(elt) {
                let Pos { x, y } = c;
                if !visited.contains(&c) {
                    visited.insert(c.clone());
                    machine.put_input(dir);
                    machine.run();
                    let output = machine.get_last_output();
                    if output[0] == 0 {
                        self.grid.insert(c, '#');
                    } else if output[0] == 1 {
                        self.grid.insert(c, '.');
                        todo.push(c.clone());
                    } else if output[0] == 2 {
                        self.grid.insert(c, 'O');
                        self.oxygen = Some(c.clone());
                        todo.push(c.clone());
                    }
                }
            }
        }
        self.display();
    }

    fn display(&self) {
        let min_x = self.grid.keys().map(|p| p.x).min().unwrap();
        let max_x = self.grid.keys().map(|p| p.x).max().unwrap();
        let min_y = self.grid.keys().map(|p| p.y).min().unwrap();
        let max_y = self.grid.keys().map(|p| p.y).max().unwrap();
        for y in min_y..=max_y {
            for x in min_x..=max_x {
                if x == self.bot.x && y == self.bot.y {
                    print!("o");
                    continue;
                }
                let tile = self.grid.get(&Pos { x, y }).unwrap_or(&' ');
                print!("{}", tile);
            }
            println!();
        }
    }
}

pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut game = Cabinet::new();
    game.explore(code);

    0
}

pub fn part2(input: String) -> i64 {
    0
}
