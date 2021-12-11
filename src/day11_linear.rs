#[derive(Debug)]
struct Octopus {
    energy: u32,
    flashed: bool,
}

fn indice(x: usize, y: usize) -> usize {
    let n = 10; // hardcoded for now
    y * n + x
}

fn from_indice(i: usize) -> (usize, usize) {
    let n = 10; // hardcoded for now
    (i % n, i / n)
}

fn increase_neighbours_energy(board: &mut Vec<Octopus>, i: usize) {
    let n = 10; // hardcoded for now
    let (x, y) = from_indice(i);
    for j in (y as i32 - 1)..(y + 1 + 1) as i32 {
        for i in (x as i32 - 1)..(x + 1 + 1) as i32 {
            if j >= 0 && j < n && i >= 0 && i < n {
                board[indice(i as usize, j as usize)].energy += 1;
            }
        }
    }
    board[indice(x, y)].energy -= 1;
}

fn clear_flash(board: &mut Vec<Octopus>) {
    for octopus in board.iter_mut() {
        if octopus.flashed {
            octopus.energy = 0;
            octopus.flashed = false;
        }
    }
}

fn flash(board: &mut Vec<Octopus>) -> usize {
    let mut flashed = 0;
    let mut continue_flashing = true;
    while continue_flashing {
        continue_flashing = false;
        for i in 0..board.len() {
            if board[i].energy > 9 && !board[i].flashed {
                board[i].flashed = true;
                flashed += 1;
                continue_flashing = true;
                increase_neighbours_energy(board, i);
            }
        }
    }
    clear_flash(board);
    return flashed;
}

fn step(board: &mut Vec<Octopus>) -> usize {
    for octopus in board.iter_mut() {
        octopus.energy += 1;
    }
    flash(board)
}

fn build_board(input: &str) -> Vec<Octopus> {
    input
        .lines()
        .flat_map(|line| {
            line.chars().map(|c| Octopus {
                energy: c.to_digit(10).unwrap(),
                flashed: false,
            })
        })
        .collect()
}

pub fn part1(input: String) -> i64 {
    let mut board = build_board(&input);
    (0..100).map(|_| step(&mut board)).sum::<usize>() as i64
}

pub fn part2(input: String) -> i64 {
    let mut board = build_board(&input);
    (1..).find(|_| step(&mut board) == board.len()).unwrap() as i64
}
