use array2d::Array2D;
use std::collections::HashSet;

#[derive(Debug)]
struct Board {
    width: i32,
    height: i32,
    data: Vec<(i32, i32)>,
}

#[derive(Debug)]
struct Instruction {
    command: String,
    direction: char,
    distance: u32,
}

fn build_board(input: &str) -> (Board, Vec<Instruction>) {
    let (coords, instr) = input.split_once("\n\n").unwrap();

    let data: Vec<(i32, i32)> = coords
        .lines()
        .map(|line| {
            (line
                .split_once(",")
                .map(|(x, y)| (x.parse::<i32>().unwrap(), y.parse::<i32>().unwrap())))
            .unwrap()
        })
        .collect();

    let board = Board {
        width: data.iter().map(|(x, _)| x).max().unwrap() + 1,
        height: data.iter().map(|(x, _)| x).max().unwrap() + 1,
        data,
    };

    let instructions: Vec<Instruction> = instr
        .lines()
        .map(|line| {
            let mut iter = line.split_ascii_whitespace();
            let command = iter.next().unwrap();
            let (direction, distance) = iter.skip(1).next().unwrap().split_once("=").unwrap();
            Instruction {
                command: command.to_string(),
                direction: direction.chars().next().unwrap(),
                distance: distance.parse::<u32>().unwrap(),
            }
        })
        .collect();
    (board, instructions)
}

fn step(board: &mut Board, instruction: &Instruction) {
    match (
        instruction.command.as_str(),
        instruction.direction,
        instruction.distance as i32,
    ) {
        ("fold", 'x', d) => {
            for (x, _) in board.data.iter_mut() {
                *x = if *x > d { d - (*x - d) } else { *x };
            }
            board.width = d;
        }
        ("fold", 'y', d) => {
            for (_, y) in board.data.iter_mut() {
                *y = if *y > d { d - (*y - d) } else { *y };
            }
            board.height = d;
        }
        _ => panic!("Unknown command: {}", instruction.command),
    }
}

fn display(board: &mut Board, instructions: Vec<Instruction>) {
    for instruction in instructions {
        step(board, &instruction);
    }
    let screen = board.data.iter().fold(
        Array2D::filled_with(' ', board.height as usize, board.width as usize),
        |mut acc, (x, y)| {
            acc[(*y as usize, *x as usize)] = '#';
            acc
        },
    );
    for row in screen.rows_iter() {
        for c in row {
            print!("{}", c);
        }
        println!();
    }
}

pub fn part1(input: String) -> i64 {
    let (mut board, instructions) = build_board(&input);
    step(&mut board, instructions.iter().next().unwrap());
    let uniq: HashSet<(i32, i32)> = HashSet::from_iter(board.data.iter().cloned());
    uniq.len() as i64
}

pub fn part2(input: String) -> i64 {
    let (mut board, instructions) = build_board(&input);
    display(&mut board, instructions);
    0
}
