use std::cmp::Ordering;
use std::collections::HashSet;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Pos {
    x: i64,
    y: i64,
}

impl Ord for Pos {
    fn cmp(&self, other: &Self) -> Ordering {
        if self == other {
            return Ordering::Equal;
        }

        let angle1 = (self.y as f64).atan2(self.x as f64);
        let angle2 = (other.y as f64).atan2(other.x as f64);
        angle1.partial_cmp(&angle2).unwrap()
    }
}

impl PartialOrd for Pos {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
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
    // println!("pos: {:?} maxx: {} maxy: {}", station, maxx, maxy);

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

    // println!("concentric: {:?}", concentric);
    // println!("len: {}", concentric.len());

    let mut res = 0;
    for neighbor in concentric {
        // println!("neighbor: {:?} ", neighbor);
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
                    // println!("removed: {:?}", new_pos);
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
    // grid.iter()
    //     .map(|pos| can_be_seen(pos.clone(), &mut grid.clone(), maxx, maxy))
    //     .max()
    //     .unwrap()

    let mut max = 0;
    for pos in grid.clone() {
        let res = can_be_seen(pos.clone(), &mut grid.clone(), maxx, maxy);
        if res > max {
            max = res;
            println!("pos: {:?} res: {}", pos, res);
        }
    }

    max
}

fn angles(grid: &HashSet<Pos>, station: &Pos) -> Vec<(Pos)> {
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
        // println!("pos: {:?}, p: {:?}, angle: {}", pos, p, to_angle(p.clone()));

        angles.push(p);
    }
    angles.sort_by_key(|pos| to_angle(pos.clone()) as i64);
    angles.dedup();

    angles
}

fn to_angle(pos: Pos) -> f64 {
    let alpha = (pos.x as f64).atan2(pos.y as f64);
    let degree = alpha * 180.0 / std::f64::consts::PI;
    (degree + 360.0) % 360.0
}

fn search(grid: &HashSet<Pos>, station: &Pos, angle: &Pos, maxx: i64, maxy: i64) -> Option<Pos> {
    for i in 1.. {
        let pos = Pos {
            x: station.x + angle.x * i,
            y: station.y - angle.y * i,
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
    println!("angles: {:?}", angles);
    let mut last_removed = 0;
    for i in 1..200 {
        let angle = &angles[i];
        if let Some(pos) = search(&grid, &station, &angle, maxx, maxy) {
            grid.remove(&pos);
            // println!("{}: removed {:?}", i + 1, pos);
            last_removed = pos.x * 100 + pos.y;
        }
    }
    last_removed
}
