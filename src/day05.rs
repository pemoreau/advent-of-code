use array2d::Array2D;
fn parse_line1(s: &str) -> (i32, i32, i32, i32) {
    let re = regex::Regex::new(r"([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)").unwrap();
    let cap = re.captures(s).unwrap();
    let x1: i32 = cap[1].parse().unwrap();
    let y1: i32 = cap[2].parse().unwrap();
    let x2: i32 = cap[3].parse().unwrap();
    let y2: i32 = cap[4].parse().unwrap();
    return (x1, y1, x2, y2);
}

const N: usize = 1000;
pub fn part1(input: String) -> i64 {
    let values: Vec<_> = input.lines().map(|line| parse_line1(line)).collect();
    let mut board = Array2D::filled_with(0, N, N);
    mark_lines(&mut board, &values);
    return count(&board) as i64;
}

pub fn part2(input: String) -> i64 {
    let values: Vec<_> = input.lines().map(|line| parse_line1(line)).collect();
    let mut board = Array2D::filled_with(0, N, N);
    mark_lines(&mut board, &values);
    mark_diags(&mut board, &values);
    return count(&board) as i64;
}

fn mark_lines(board: &mut Array2D<i32>, values: &Vec<(i32, i32, i32, i32)>) {
    for (x1, y1, x2, y2) in values {
        if x1 == x2 || y1 == y2 {
            let (xx1, xx2) = if x1 < x2 { (x1, x2) } else { (x2, x1) };
            let (yy1, yy2) = if y1 < y2 { (y1, y2) } else { (y2, y1) };
            for i in *xx1..*xx2 + 1 {
                for j in *yy1..*yy2 + 1 {
                    board[(i as usize, j as usize)] += 1;
                }
            }
        }
    }
}

fn mark_diags(board: &mut Array2D<i32>, values: &Vec<(i32, i32, i32, i32)>) {
    for (x1, y1, x2, y2) in values {
        if (x1 - x2).abs() == (y1 - y2).abs() {
            let incx = (x2 - x1).signum();
            let incy = if y1 <= y2 { 1 } else { -1 };
            let mut i = *x1;
            let mut j = *y1;
            while i != *x2 && j != *y2 {
                board[(i as usize, j as usize)] += 1;
                i = i + incx;
                j = j + incy;
            }
            board[(*x2 as usize, *y2 as usize)] += 1;
        }
    }
}
fn count(board: &Array2D<i32>) -> i32 {
    return board
        .elements_column_major_iter()
        .filter(|x| **x >= 2)
        .count() as i32;
}
