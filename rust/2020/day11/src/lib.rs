use std::cmp::min;

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
enum Cell {
    Floor,
    Empty,
    Occupied,
}

struct Game {
    width: usize,
    height: usize,
    cells: Vec<Cell>,
}

impl Game {
    fn new(input: String) -> Game {
        let cells = input
            .replace("\n", "")
            .chars()
            .map(|c| match c {
                '#' => Cell::Occupied,
                'L' => Cell::Empty,
                _ => Cell::Floor,
            })
            .collect::<Vec<Cell>>();

        let height = input.lines().count();
        let width = cells.len() / height;

        Game {
            width,
            height,
            cells,
        }
    }

    fn get_index(&self, row: usize, column: usize) -> usize {
        row * self.width + column
    }

    fn step(&mut self, part2: bool) -> bool {
        let mut next = self.cells.clone();
        let mut state_changed = false;

        for row in 0..self.height {
            for col in 0..self.width {
                let idx = self.get_index(row, col);
                let cell = self.cells[idx];
                let occupied_neighbors = if part2 {
                    self.occupied_neighbor_count(row, col)
                } else {
                    self.occupied_immediate_neighbor_count(row, col)
                };

                let next_cell = match (cell, occupied_neighbors) {
                    (Cell::Empty, 0) => Cell::Occupied,
                    (Cell::Occupied, count) if count >= (if part2 { 5 } else { 4 }) => Cell::Empty,
                    (otherwise, _) => otherwise,
                };

                if cell != next_cell {
                    state_changed = true;
                }

                next[idx] = next_cell;
            }
        }

        self.cells = next;
        state_changed
    }

    fn occupied_seats(&self) -> u32 {
        self.cells
            .iter()
            .filter(|&cell| *cell == Cell::Occupied)
            .count() as u32
    }

    fn occupied_immediate_neighbor_count(&self, row: usize, column: usize) -> u32 {
        let mut count = 0;

        for r in row as i32 - 1..=row as i32 + 1 {
            for c in column as i32 - 1..=column as i32 + 1 {
                if r == row as i32 && c == column as i32 {
                    continue;
                }
                if self.is_occupied(r, c) {
                    count += 1;
                }
            }
        }

        count
    }

    fn is_occupied(&self, row: i32, column: i32) -> bool {
        column >= 0
            && column < self.width as i32
            && row >= 0
            && row < self.height as i32
            && self.cells[self.get_index(row as usize, column as usize)] == Cell::Occupied
    }
    fn is_empty(&self, row: i32, column: i32) -> bool {
        column >= 0
            && column < self.width as i32
            && row >= 0
            && row < self.height as i32
            && self.cells[self.get_index(row as usize, column as usize)] == Cell::Empty
    }

    fn occupied_neighbor_count(&self, row: usize, column: usize) -> u32 {
        let mut count = 0;
        let modifiers = [
            (1, 0),
            (-1, 0),
            (0, 1),
            (0, -1),
            (1, 1),
            (-1, -1),
            (1, -1),
            (-1, 1),
        ];
        for (mx, my) in modifiers.iter() {
            for i in 1..min(self.width, self.height) {
                let c = column as i32 + mx * i as i32;
                let r = row as i32 + my * i as i32;
                if self.is_occupied(r, c) {
                    count += 1;
                    break;
                }
                if self.is_empty(r, c) {
                    break;
                }
            }
        }
        count
    }
}

pub fn part1(input: String) -> i64 {
    let mut game = Game::new(input);
    let mut cont = true;
    while cont {
        cont = game.step(false);
    }
    game.occupied_seats() as i64
}

pub fn part2(input: String) -> i64 {
    let mut game = Game::new(input);
    let mut cont = true;
    while cont {
        cont = game.step(true);
    }
    game.occupied_seats() as i64
}
