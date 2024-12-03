use itertools::Itertools;

#[derive(Debug)]
struct Octopus {
    energy: u32,
    flashed: bool,
}

fn increase_neighbours_energy(board: &mut Vec<Vec<Octopus>>, x: usize, y: usize) {
    for j in (y as i32 - 1)..(y + 1 + 1) as i32 {
        for i in (x as i32 - 1)..(x + 1 + 1) as i32 {
            if j >= 0 && j < board.len() as i32 && i >= 0 && i < board[j as usize].len() as i32 {
                board[j as usize][i as usize].energy += 1;
            }
        }
    }
    board[y][x].energy -= 1;
}

fn clear_flash(board: &mut Vec<Vec<Octopus>>) {
    for y in 0..board.len() {
        for x in 0..board[y].len() {
            if board[y][x].flashed {
                board[y][x].energy = 0;
                board[y][x].flashed = false;
            }
        }
    }
}

fn flash(board: &mut Vec<Vec<Octopus>>) -> usize {
    let mut flashed = 0;
    let mut continue_flashing = true;
    while continue_flashing {
        continue_flashing = false;
        for y in 0..board.len() {
            for x in 0..board[y].len() {
                if board[y][x].energy > 9 && !board[y][x].flashed {
                    board[y][x].flashed = true;
                    flashed += 1;
                    continue_flashing = true;
                    increase_neighbours_energy(board, x, y);
                }
            }
        }
    }
    clear_flash(board);
    return flashed;
}

fn step(board: &mut Vec<Vec<Octopus>>) -> usize {
    for (y, x) in (0..board.len()).cartesian_product(0..board[0].len()) {
        board[x][y].energy += 1;
    }
    flash(board)
}

fn build_board(input: &str) -> Vec<Vec<Octopus>> {
    input
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| Octopus {
                    energy: c.to_digit(10).unwrap(),
                    flashed: false,
                })
                .collect()
        })
        .collect()
}

pub fn part1(input: String) -> i64 {
    let mut board = build_board(&input);
    (0..100).map(|_| step(&mut board)).sum::<usize>() as i64
}

pub fn part2(input: String) -> i64 {
    let mut board = build_board(&input);
    let n = board.len() * board[0].len();
    (1..).find(|_| step(&mut board) == n).unwrap() as i64
}
