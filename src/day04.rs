use array2d::Array2D;

fn read_boards(input: String) -> (Vec<i32>, Vec<Array2D<i32>>) {
    let mut parts = input.split("\n\n");

    let values = parts
        .next()
        .unwrap()
        .split(',')
        .map(|s| s.trim().parse::<i32>().unwrap())
        .collect::<Vec<i32>>();

    let boards = parts
        .map(|part| {
            let integers = part
                .lines()
                .flat_map(|line| line.split_whitespace().map(|s| s.parse::<i32>().unwrap()))
                .collect::<Vec<_>>();
            return Array2D::from_row_major(&integers, 5, 5);
        })
        .collect();
    return (values, boards);
}

pub fn part1(input: String) -> i64 {
    let (values, boards) = read_boards(input);
    let success = successful_boards(values, boards);
    return *success.first().unwrap() as i64;
}

pub fn part2(input: String) -> i64 {
    let (values, boards) = read_boards(input);
    let success = successful_boards(values, boards);
    return *success.last().unwrap() as i64;
}

fn successful_boards(values: Vec<i32>, mut boards: Vec<Array2D<i32>>) -> Vec<i32> {
    let mut success: Vec<i32> = Vec::new();
    for value in values {
        for board in boards.iter_mut() {
            let res = play(value, board);
            if res.is_some() {
                success.push(res.unwrap());
            }
        }
    }
    return success;
}

fn play(value: i32, board: &mut Array2D<i32>) -> Option<i32> {
    for i in 0..board.column_len() {
        for j in 0..board.row_len() {
            if board[(i, j)] == value {
                board[(i, j)] = -1;
                if completed(board, i, j) {
                    let sum_positive = board
                        .as_row_major()
                        .iter()
                        .filter(|x| x >= &&0)
                        .sum::<i32>();
                    fill_board(-1, board);
                    return Some(value * sum_positive);
                }
            }
        }
    }
    return None;
}

fn fill_board(value: i32, board: &mut Array2D<i32>) {
    for i in 0..board.column_len() {
        for j in 0..board.row_len() {
            board[(i, j)] = value;
        }
    }
}

fn completed(board: &Array2D<i32>, i: usize, j: usize) -> bool {
    let completed_row = board.row_iter(i).all(|x| x < &0);
    let completed_column = board.column_iter(j).all(|x| x < &0);
    return completed_row || completed_column;
}
