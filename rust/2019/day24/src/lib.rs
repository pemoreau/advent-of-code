use std::collections::{HashMap, HashSet};

fn is_occupied(tile: i64, x: i8, y: i8) -> bool {
    if x < 0 || x > 4 || y < 0 || y > 4 {
        return false;
    }
    tile & (1 << (y * 5 + x)) != 0
}

fn nb_neighbors(tile: i64, x: i8, y: i8) -> i64 {
    let mut result = 0;
    if is_occupied(tile, x - 1, y) {
        result += 1;
    }
    if is_occupied(tile, x + 1, y) {
        result += 1;
    }
    if is_occupied(tile, x, y - 1) {
        result += 1;
    }
    if is_occupied(tile, x, y + 1) {
        result += 1;
    }
    result
}

fn build_tile(input: String) -> i64 {
    let mut total = 0;
    let mut power = 1;
    for line in input.lines() {
        for c in line.chars() {
            if c == '#' {
                total += power;
            }
            power *= 2;
        }
    }
    total
}

fn step(tile: i64) -> i64 {
    let mut new_tile = 0;
    for y in 0..5 {
        for x in 0..5 {
            let nb_neighbors = nb_neighbors(tile, x, y);
            if is_occupied(tile, x, y) {
                if nb_neighbors == 1 {
                    new_tile |= 1 << (y * 5 + x);
                }
            } else {
                if nb_neighbors == 1 || nb_neighbors == 2 {
                    new_tile |= 1 << (y * 5 + x);
                }
            }
        }
    }
    new_tile
}

pub fn part1(input: String) -> i64 {
    let mut tile = build_tile(input);

    let mut seen = HashSet::new();
    seen.insert(tile);

    loop {
        tile = step(tile);
        if seen.contains(&tile) {
            return tile;
        }
        seen.insert(tile);
    }
}

fn nb_neighbors2(tile: i64, x: i8, y: i8, level: i8, levels: &HashMap<i8, i64>) -> i64 {
    let mut result = 0;
    let next_level = levels.get(&(level + 1)).unwrap_or(&0);
    let previous_level = levels.get(&(level - 1)).unwrap_or(&0);

    for direction in &[(0, -1), (1, 0), (0, 1), (-1, 0)] {
        let (dx, dy) = direction;
        if x + dx == 2 && y + dy == 2 {
            if y == 2 {
                for y2 in 0..5 {
                    if x == 1 && is_occupied(*next_level, 0, y2)
                        || x == 3 && is_occupied(*next_level, 4, y2)
                    {
                        result += 1;
                    }
                }
            } else {
                for x2 in 0..5 {
                    if y == 1 && is_occupied(*next_level, x2, 0)
                        || y == 3 && is_occupied(*next_level, x2, 4)
                    {
                        result += 1;
                    }
                }
            }
        } else if (x + dx < 0 && is_occupied(*previous_level, 1, 2))
            || (x + dx > 4 && is_occupied(*previous_level, 3, 2))
            || (y + dy < 0 && is_occupied(*previous_level, 2, 1))
            || (y + dy > 4 && is_occupied(*previous_level, 2, 3))
        {
            result += 1;
        } else if is_occupied(tile, x + dx, y + dy) {
            result += 1;
        }
    }
    result
}

pub fn part2(input: String) -> i64 {
    let tile = build_tile(input);
    let mut levels: HashMap<i8, i64> = HashMap::new();
    levels.insert(0, tile);
    levels.insert(1, 0);
    levels.insert(-1, 0);
    for _ in 0..200 {
        let mut new_levels: HashMap<i8, i64> = HashMap::new();
        for (&level, &tile) in levels.iter() {
            let mut new_tile = 0;
            for y in 0..5 {
                for x in 0..5 {
                    if x == 2 && y == 2 {
                        continue;
                    }
                    let n = nb_neighbors2(tile, x, y, level, &levels);
                    if is_occupied(tile, x, y) {
                        if n == 1 {
                            new_tile |= 1 << (y * 5 + x);
                        }
                    } else {
                        if n == 1 || n == 2 {
                            new_tile |= 1 << (y * 5 + x);
                        }
                    }
                }
            }
            new_levels.insert(level, new_tile);
        }
        let min = levels.keys().min().unwrap();
        let max = levels.keys().max().unwrap();
        if *new_levels.get(min).unwrap() != 0 {
            new_levels.insert(*min - 1, 0);
        }
        if *new_levels.get(max).unwrap() != 0 {
            new_levels.insert(*max + 1, 0);
        }

        levels = new_levels;
    }
    let total_occupied = levels.values().map(|tile| tile.count_ones()).sum::<u32>();
    total_occupied as i64
}
