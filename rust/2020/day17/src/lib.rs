use std::collections::HashSet;

use itertools::Itertools;

pub fn part1(input: String) -> i64 {
    solve(input, false)
}

pub fn part2(input: String) -> i64 {
    solve(input, true)
}

pub fn solve(input: String, d4: bool) -> i64 {
    let mut state: HashSet<(i32, i32, i32, i32)> = HashSet::new();
    input.lines().enumerate().for_each(|(y, line)| {
        line.chars().enumerate().for_each(|(x, c)| {
            if c == '#' {
                state.insert((x as i32, y as i32, 0, 0));
            }
        });
    });
    (0..6).for_each(|_| state = step(&state, d4));
    state.len() as i64
}

fn neighbors(
    state: &HashSet<(i32, i32, i32, i32)>,
    (x, y, z, w): (i32, i32, i32, i32),
    d4: bool,
) -> usize {
    // let mut active = 0;
    // let domain4 = if d4 { -1..2 } else { 0..1 };
    // for i in -1..2 {
    //     for j in -1..2 {
    //         for k in -1..2 {
    //             for l in domain4.clone() {
    //                 if i == 0 && j == 0 && k == 0 && l == 0 {
    //                     continue;
    //                 }
    //                 if state.contains(&(x + i, y + j, z + k, w + l)) {
    //                     active += 1;
    //                 }
    //             }
    //         }
    //     }
    // }
    // active

    // shorter but slower solution
    (0..4)
        .map(|i| if i == 3 && !d4 { 0..1 } else { -1..2 })
        .multi_cartesian_product()
        .filter(|v| {
            !(v[0] == 0 && v[1] == 0 && v[2] == 0 && v[3] == 0)
                && state.contains(&(x + v[0], y + v[1], z + v[2], w + v[3]))
        })
        .count()
}

fn step(state: &HashSet<(i32, i32, i32, i32)>, d4: bool) -> HashSet<(i32, i32, i32, i32)> {
    let mut new_state: HashSet<(i32, i32, i32, i32)> = HashSet::new();
    // let domain4 = if d4 { -1..2 } else { 0..1 };
    // state.iter().for_each(|(x, y, z, w)| {
    //     for i in -1..2 {
    //         for j in -1..2 {
    //             for k in -1..2 {
    //                 for l in domain4.clone() {
    //                     let cell = (x + i, y + j, z + k, w + l);
    //                     let active = neighbors(state, cell, d4);
    //                     if active == 3 || (state.contains(&cell) && active == 2) {
    //                         new_state.insert(cell);
    //                     }
    //                 }
    //             }
    //         }
    //     }
    // });

    // shorter but slower solution
    state.iter().for_each(|(x, y, z, w)| {
        (0..4)
            .map(|i| if i == 3 && !d4 { 0..1 } else { -1..2 })
            .multi_cartesian_product()
            .for_each(|v| {
                let cell = (x + v[0], y + v[1], z + v[2], w + v[3]);
                let active = neighbors(state, cell, d4);
                if active == 3 || (state.contains(&cell) && active == 2) {
                    new_state.insert(cell);
                }
            })
    });

    new_state
}
