#[derive(Debug)]
struct Octopus {
    energy: u32,
    flashed: bool,
}

fn increase_energy(board: &mut Vec<Vec<Octopus>>) {
    for y in 0..board.len() {
        for x in 0..board[y].len() {
            board[y][x].energy += 1;
        }
    }
}

fn increase_neighbors_energy(board: &mut Vec<Vec<Octopus>>, x: usize, y: usize) {
    for j in (y as i32 - 1)..(y + 1 + 1) as i32 {
        for i in (x as i32 - 1)..(x + 1 + 1) as i32 {
            if j >= 0 && j < board.len() as i32 && i >= 0 && i < board[j as usize].len() as i32 {
                board[j as usize][i as usize].energy += 1;
            }
        }
    }
    board[y][x].energy -= 1;
}

fn clear_flash(board: &mut Vec<Vec<Octopus>>) -> bool {
    let mut all_falshed = true;
    for y in 0..board.len() {
        for x in 0..board[y].len() {
            if board[y][x].flashed {
                board[y][x].energy = 0;
                board[y][x].flashed = false;
            } else {
                all_falshed = false;
            }
        }
    }
    return all_falshed;
}

fn flash(board: &mut Vec<Vec<Octopus>>) -> u32 {
    let mut continue_flashing = true;
    let mut flashed = 0;
    while continue_flashing {
        continue_flashing = false;
        for y in 0..board.len() {
            for x in 0..board[y].len() {
                if board[y][x].energy > 9 && !board[y][x].flashed {
                    board[y][x].flashed = true;
                    flashed += 1;
                    continue_flashing = true;
                    increase_neighbors_energy(board, x, y);
                }
            }
        }
    }
    return flashed;
}

pub fn part1(input: String) -> i64 {
    let mut board: Vec<Vec<Octopus>> = input
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| Octopus {
                    energy: c.to_digit(10).unwrap(),
                    flashed: false,
                })
                .collect()
        })
        .collect();

    let mut res = 0;
    for _ in 1..101 {
        increase_energy(&mut board);
        let flashed = flash(&mut board);
        res += flashed;
        clear_flash(&mut board);
    }

    return res as i64;
}

pub fn part2(input: String) -> i64 {
    let mut board: Vec<Vec<Octopus>> = input
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| Octopus {
                    energy: c.to_digit(10).unwrap(),
                    flashed: false,
                })
                .collect()
        })
        .collect();

    let mut i = 1;
    loop {
        increase_energy(&mut board);
        flash(&mut board);
        let all_flashed = clear_flash(&mut board);
        if all_flashed {
            return i;
        }
        i += 1
    }
}
