use std::collections::{HashMap, HashSet};

fn dist_to_com(orbit: &str, direct_orbits: &HashMap<&str, &str>) -> i64 {
    if orbit == "COM" {
        0
    } else {
        1 + dist_to_com(direct_orbits.get(orbit).unwrap(), direct_orbits)
    }
}

fn collect_orbits(orbit: &str, direct_orbits: &HashMap<&str, &str>, bag: &mut HashSet<String>) {
    if orbit == "COM" {
        return;
    }
    bag.insert(orbit.to_string());
    collect_orbits(direct_orbits.get(orbit).unwrap(), direct_orbits, bag);
}

pub fn part1(input: String) -> i64 {
    let mut direct_orbits = HashMap::new();
    for line in input.lines() {
        let mut split = line.split(")");
        let center = split.next().unwrap();
        let orbit = split.next().unwrap();
        direct_orbits.insert(orbit, center);
    }

    let indirect = direct_orbits
        .keys()
        .map(|orbit| dist_to_com(orbit, &direct_orbits))
        .sum::<i64>();
    indirect
}

pub fn part2(input: String) -> i64 {
    let mut direct_orbits = HashMap::new();
    for line in input.lines() {
        let mut split = line.split(")");
        let center = split.next().unwrap();
        let orbit = split.next().unwrap();
        direct_orbits.insert(orbit, center);
    }
    let mut bag1 = HashSet::new();
    let mut bag2 = HashSet::new();
    collect_orbits("YOU", &direct_orbits, &mut bag1);
    collect_orbits("SAN", &direct_orbits, &mut bag2);
    let bag3 = bag1.symmetric_difference(&bag2).collect::<Vec<_>>();
    bag3.len() as i64 - 2
}
