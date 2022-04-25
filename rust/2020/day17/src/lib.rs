use std::collections::HashSet;

pub fn part1(input: String) -> i64 {
    let mut state: HashSet<(i32, i32, i32)> = HashSet::new();
    input.lines().enumerate().for_each(|(y, line)| {
        line.chars().enumerate().for_each(|(x, c)| {
            if c == '#' {
                state.insert((x as i32, y as i32, 0));
            }
        });
    });
    (0..6).for_each(|_| state = step(&state));
    state.len() as i64
}

pub fn part2(input: String) -> i64 {
    let mut state: HashSet<(i32, i32, i32, i32)> = HashSet::new();
    input.lines().enumerate().for_each(|(y, line)| {
        line.chars().enumerate().for_each(|(x, c)| {
            if c == '#' {
                state.insert((x as i32, y as i32, 0, 0));
            }
        });
    });
    (0..6).for_each(|_| state = step4(&state));
    state.len() as i64
}

fn neiboors(state: &HashSet<(i32, i32, i32)>, (x, y, z): (i32, i32, i32)) -> usize {
    let mut active = 0;
    for i in -1..2 {
        for j in -1..2 {
            for k in -1..2 {
                if i == 0 && j == 0 && k == 0 {
                    continue;
                }
                if state.contains(&(x + i, y + j, z + k)) {
                    active += 1;
                }
            }
        }
    }
    active
}

fn step(state: &HashSet<(i32, i32, i32)>) -> HashSet<(i32, i32, i32)> {
    let mut new_state: HashSet<(i32, i32, i32)> = HashSet::new();
    state.iter().for_each(|(x, y, z)| {
        for i in -1..2 {
            for j in -1..2 {
                for k in -1..2 {
                    let cell = (x + i, y + j, z + k);
                    let active = neiboors(state, cell);
                    if state.contains(&cell) && (active == 2 || active == 3) {
                        new_state.insert(cell);
                    }
                    if !state.contains(&cell) && active == 3 {
                        new_state.insert(cell);
                    }
                }
            }
        }
    });
    new_state
}

fn neiboors4(state: &HashSet<(i32, i32, i32, i32)>, (x, y, z, w): (i32, i32, i32, i32)) -> usize {
    let mut active = 0;
    for i in -1..2 {
        for j in -1..2 {
            for k in -1..2 {
                for l in -1..2 {
                    if i == 0 && j == 0 && k == 0 && l == 0 {
                        continue;
                    }
                    if state.contains(&(x + i, y + j, z + k, w + l)) {
                        active += 1;
                    }
                }
            }
        }
    }
    active
}

fn step4(state: &HashSet<(i32, i32, i32, i32)>) -> HashSet<(i32, i32, i32, i32)> {
    let mut new_state: HashSet<(i32, i32, i32, i32)> = HashSet::new();
    state.iter().for_each(|(x, y, z, w)| {
        for i in -1..2 {
            for j in -1..2 {
                for k in -1..2 {
                    for l in -1..2 {
                        let cell = (x + i, y + j, z + k, w + l);
                        let active = neiboors4(state, cell);
                        if state.contains(&cell) && (active == 2 || active == 3) {
                            new_state.insert(cell);
                        }
                        if !state.contains(&cell) && active == 3 {
                            new_state.insert(cell);
                        }
                    }
                }
            }
        }
    });
    new_state
}
