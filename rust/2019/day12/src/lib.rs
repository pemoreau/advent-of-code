#[macro_use]
use std::collections::HashMap;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
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

pub fn part1(input: String) -> i64 {
    let mut moons = vec![
        build_moon(-1, 0, 2),
        build_moon(2, -10, -7),
        build_moon(4, -8, 8),
        build_moon(3, 5, -1),
    ];
    let mut moons = vec![
        build_moon(-8, -10, 0),
        build_moon(5, 5, 10),
        build_moon(2, -7, 3),
        build_moon(9, -8, -3),
    ];
    /*
        <x=-13, y=14, z=-7>
    <x=-18, y=9, z=0>
    <x=0, y=-3, z=-3>
    <x=-15, y=3, z=-13>
         */
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

pub fn part2(input: String) -> i64 {
    0
}
