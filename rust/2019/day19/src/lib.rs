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
    let mut cpt = 0;
    for y in 0..100 {
        for x in 0..100 {
            let mut machine = Machine::new(code.clone(), vec![x, y]);
            machine.run();
            let output = machine.get_output();
            grid.insert(Pos { x, y }, if output[0] == 0 { '.' } else { '#' });
            if output[0] == 1 {
                cpt += 1;
            }
        }
    }
    display(&grid);
    cpt
}
