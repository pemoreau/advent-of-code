use num::integer::lcm;
use std::collections::HashMap;

#[derive(Clone, Debug, Eq, Hash, PartialEq, Copy)]
struct Pos {
    x: i64,
    y: i64,
    z: i64,
}

impl Pos {
    fn new() -> Self {
        Self { x: 0, y: 0, z: 0 }
    }

    fn add(&mut self, other: &Pos) {
        self.x += other.x;
        self.y += other.y;
        self.z += other.z;
    }
}

#[derive(Debug, Eq, PartialEq, Clone)]
struct Moon {
    pos: Pos,
    vel: Pos,
}

fn build_moon(x: i64, y: i64, z: i64) -> Moon {
    Moon {
        pos: Pos { x, y, z },
        vel: Pos::new(),
    }
}

fn apply_gravity(moons: &mut Vec<Moon>) {
    for moon2 in moons.clone().iter() {
        for moon1 in moons.iter_mut() {
            if moon1 == moon2 {
                continue;
            }
            if moon1.pos.x < moon2.pos.x {
                moon1.vel.x += 1;
            } else if moon1.pos.x > moon2.pos.x {
                moon1.vel.x -= 1;
            }
            if moon1.pos.y < moon2.pos.y {
                moon1.vel.y += 1;
            } else if moon1.pos.y > moon2.pos.y {
                moon1.vel.y -= 1;
            }
            if moon1.pos.z < moon2.pos.z {
                moon1.vel.z += 1;
            } else if moon1.pos.z > moon2.pos.z {
                moon1.vel.z -= 1;
            }
        }
    }
}

fn apply_velocity(moons: &mut Vec<Moon>) {
    for moon in moons.iter_mut() {
        moon.pos.add(&moon.vel);
    }
}

fn energy(moon: &Moon) -> i64 {
    let pot = moon.pos.x.abs() + moon.pos.y.abs() + moon.pos.z.abs();
    let kin = moon.vel.x.abs() + moon.vel.y.abs() + moon.vel.z.abs();
    pot * kin
}

pub fn part1() -> i64 {
    let mut moons = vec![
        build_moon(-13, 14, -7),
        build_moon(-18, 9, 0),
        build_moon(0, -3, -3),
        build_moon(-15, 3, -13),
    ];
    for _ in 0..1000 {
        apply_gravity(&mut moons);
        apply_velocity(&mut moons);
    }

    moons.iter().map(|m| energy(m)).sum()
}

fn find_cycle(moons: Vec<Moon>, f: fn(Pos) -> i64) -> i64 {
    let mut moons = moons;
    let mut map: HashMap<String, i64> = HashMap::new();
    for i in 0.. {
        let signature = moons
            .iter()
            .map(|m| format!("{},{} ", f(m.pos), f(m.vel)))
            .collect::<Vec<String>>()
            .join(",");
        let n = map.get(&signature);
        if n.is_some() {
            return i - n.unwrap();
        } else {
            map.insert(signature.clone(), i);
        }
        apply_gravity(&mut moons);
        apply_velocity(&mut moons);
    }
    0
}

pub fn part2() -> i64 {
    let moons = vec![
        build_moon(-13, 14, -7),
        build_moon(-18, 9, 0),
        build_moon(0, -3, -3),
        build_moon(-15, 3, -13),
    ];
    let cycle_x = find_cycle(moons.clone(), |p| p.x);
    let cycle_y = find_cycle(moons.clone(), |p| p.y);
    let cycle_z = find_cycle(moons.clone(), |p| p.z);

    lcm(lcm(cycle_x, cycle_y), cycle_z)
}
