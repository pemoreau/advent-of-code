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

    let mut res = 0;
    for neighbor in concentric {
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
                if grid.contains(&new_pos) {
                    grid.remove(&new_pos);
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

    let mut max = 0;
    for pos in grid.clone() {
        let res = can_be_seen(pos.clone(), &mut grid.clone(), maxx, maxy);
        if res > max {
            max = res;
        }
    }

    max
}

fn angles(grid: &HashSet<Pos>, station: &Pos) -> Vec<(Pos, f64)> {
    let mut angles = Vec::new();
    for pos in grid {
        if pos == station {
            continue;
        }
        let diff_x = pos.x - station.x;
        let diff_y = pos.y - station.y;
        let gcd = gcd(diff_x.abs(), diff_y.abs());
        let p = Pos {
            x: diff_x / gcd,
            y: -diff_y / gcd,
        };

        angles.push((p.clone(), to_angle(p)));
    }
    angles.sort_by(|(_, angle1), (_, angle2)| angle1.partial_cmp(angle2).unwrap());
    angles.dedup();

    angles
}

fn to_angle(pos: Pos) -> f64 {
    let alpha = (pos.x as f64).atan2(pos.y as f64);
    let degree = alpha * 180.0 / std::f64::consts::PI;
    let res = degree + 360.0;
    if res >= 360.0 {
        res - 360.0
    } else {
        res
    }
}

fn search(
    grid: &HashSet<Pos>,
    station: &Pos,
    angle: &(Pos, f64),
    maxx: i64,
    maxy: i64,
) -> Option<Pos> {
    for i in 1.. {
        let pos = Pos {
            x: station.x + angle.0.x * i,
            y: station.y - angle.0.y * i,
        };
        if pos.x < 0 || pos.x > maxx || pos.y < 0 || pos.y > maxy {
            return None;
        }
        if grid.contains(&pos) {
            return Some(pos);
        }
    }
    None
}

pub fn part2(input: String) -> i64 {
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
    let station = Pos { x: 20, y: 19 };
    let angles = angles(&grid, &station);
    let mut last_removed = 0;
    for i in 1..200 {
        let angle = &angles[i];
        if let Some(pos) = search(&grid, &station, &angle, maxx, maxy) {
            grid.remove(&pos);
            last_removed = pos.x * 100 + pos.y;
        }
    }
    last_removed
}
