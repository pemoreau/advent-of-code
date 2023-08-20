use std::collections::HashMap;

use intcode::Machine;
use utils::parsing::comma_separated_to_numbers;

#[derive(Clone, Debug, Eq, Hash, PartialEq, Copy)]
struct Pos {
    x: i64,
    y: i64,
}

type Grid = HashMap<Pos, char>;

fn display(grid: &Grid) {
    if grid.is_empty() {
        return;
    }
    let min_x = grid.keys().map(|p| p.x).min().unwrap();
    let max_x = grid.keys().map(|p| p.x).max().unwrap();
    let min_y = grid.keys().map(|p| p.y).min().unwrap();
    let max_y = grid.keys().map(|p| p.y).max().unwrap();
    for y in (min_y..=max_y).rev() {
        for x in min_x..=max_x {
            let tile = grid.get(&Pos { x, y }).unwrap_or(&' ');
            print!("{}", tile);
        }
        println!();
    }
    println!();
}

fn neighbors(pos: Pos) -> Vec<Pos> {
    vec![
        Pos {
            x: pos.x - 1,
            y: pos.y,
        },
        Pos {
            x: pos.x + 1,
            y: pos.y,
        },
        Pos {
            x: pos.x,
            y: pos.y - 1,
        },
        Pos {
            x: pos.x,
            y: pos.y + 1,
        },
    ]
}

pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut machine = Machine::new(code, vec![]);
    machine.run();
    let output = machine.get_last_output();
    println!("{:?}", output);
    let mut grid = Grid::new();
    let mut x = 0;
    let mut y = 0;
    for c in output {
        if c == 10 {
            y += 1;
            x = 0;
            continue;
        }
        grid.insert(Pos { x, y }, c as u8 as char);
        x += 1;
    }
    display(&grid);
    let mut intersections = vec![];
    for (pos, tile) in &grid {
        if *tile == '#' {
            let mut is_intersection = true;
            for n in neighbors(*pos) {
                if grid.get(&n).unwrap_or(&' ') != &'#' {
                    is_intersection = false;
                    break;
                }
            }
            if is_intersection {
                intersections.push(pos);
            }
        }
    }
    println!("{:?}", intersections);
    intersections.iter().map(|p| p.x * p.y).sum()
}

pub fn part2(input: String) -> i64 {
    0
}
