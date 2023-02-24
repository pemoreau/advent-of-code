use std::collections::HashSet;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]

struct Pos {
    x: i64,
    y: i64,
}

fn gcd(mut n: i64, mut m: i64) -> i64 {
    if n == 0 {
        return m;
    }
    while m != 0 {
        if m < n {
            std::mem::swap(&mut m, &mut n);
        }
        m = m % n;
    }
    n
}

fn can_be_seen(station: Pos, grid: &mut HashSet<Pos>, maxx: i64, maxy: i64) -> i64 {
    println!("pos: {:?} maxx: {} maxy: {}", station, maxx, maxy);

    let mut concentric = Vec::new();
    for i in 1..maxx.max(maxy) {
        for x in -i..=i {
            concentric.push(Pos {
                x: station.x + x,
                y: station.y + i,
            });
            concentric.push(Pos {
                x: station.x + x,
                y: station.y - i,
            });
        }
        for y in (-i + 1)..=(i - 1) {
            concentric.push(Pos {
                x: station.x + i,
                y: station.y + y,
            });
            concentric.push(Pos {
                x: station.x - i,
                y: station.y + y,
            });
        }
    }
    concentric.retain(|pos| {
        pos.x >= 0 && pos.x <= maxx && pos.y >= 0 && pos.y <= maxy && pos != &station
    });

    println!("concentric: {:?}", concentric);
    println!("len: {}", concentric.len());

    let mut res = 0;
    for neighbor in concentric {
        println!("neighbor: {:?} ", neighbor);
        if grid.contains(&neighbor) {
            res += 1;
            grid.remove(&neighbor);

            let diff_x = neighbor.x - station.x;
            let diff_y = neighbor.y - station.y;
            let gcd = gcd(diff_x.abs(), diff_y.abs());
            for i in 1.. {
                let new_pos = Pos {
                    x: station.x + diff_x / gcd * i,
                    y: station.y + diff_y / gcd * i,
                };
                if new_pos.x < 0 || new_pos.x > maxx || new_pos.y < 0 || new_pos.y > maxy {
                    break;
                }
                // println!("new_pos: {:?}", new_pos);
                if grid.contains(&new_pos) {
                    grid.remove(&new_pos);
                    println!("removed: {:?}", new_pos);
                }
            }
        }
    }
    res
}

pub fn part1(input: String) -> i64 {
    let mut grid = HashSet::new();
    for (y, line) in input.lines().enumerate() {
        for (x, c) in line.chars().enumerate() {
            if c == '#' {
                grid.insert(Pos {
                    x: x as i64,
                    y: y as i64,
                });
            }
        }
    }
    let maxx = (input.lines().next().unwrap().len() - 1) as i64;
    let maxy = (input.lines().count() - 1) as i64;
    let res = grid
        .iter()
        .map(|pos| can_be_seen(pos.clone(), &mut grid.clone(), maxx, maxy))
        .max()
        .unwrap();
    // let res = canBeSeen(Pos { x: 5, y: 8 }, &mut grid.clone(), maxx, maxy);
    res
}

pub fn part2(input: String) -> i64 {
    0
}
