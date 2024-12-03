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

pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut cpt = 0;
    for y in 0..50 {
        for x in 0..50 {
            let mut machine = Machine::new(code.clone(), vec![x, y]);
            machine.run();
            let output = machine.get_output();
            if output[0] == 1 {
                cpt += 1;
            }
        }
    }
    cpt
}

pub fn part2(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut grid = Grid::new();
    let size = 2500;
    for y in 0..size {
        for x in 0..size {
            let mut machine = Machine::new(code.clone(), vec![x, y]);
            machine.run();
            let output = machine.get_output();
            if output[0] == 1 {
                grid.insert(Pos { x, y }, '#');
            }
        }
    }
    let square_size = 100;
    for y in 0..size - square_size {
        for x in 0..size - square_size {
            if contains_square(&grid, x, y, square_size) {
                return x * 10000 + y;
            }
        }
    }
    0
}

fn contains_square(grid: &Grid, x: i64, y: i64, size: i64) -> bool {
    return grid.get(&Pos { x, y }) == Some(&'#')
        && grid.get(&Pos { x: x + size - 1, y }) == Some(&'#')
        && grid.get(&Pos { x, y: y + size - 1 }) == Some(&'#')
        && grid.get(&Pos {
            x: x + size - 1,
            y: y + size - 1,
        }) == Some(&'#');
}
