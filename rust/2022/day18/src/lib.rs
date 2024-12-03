use std::collections::HashSet;
use utils::parsing::comma_separated_to_numbers;

#[derive(Debug, Eq, Hash, PartialEq, Clone)]
struct Cube {
    x: i64,
    y: i64,
    z: i64,
}

fn neighbors(cube: Cube) -> Vec<Cube> {
    let Cube { x, y, z } = cube;
    vec![
        Cube { x: x + 1, y, z },
        Cube { x: x - 1, y, z },
        Cube { x, y: y - 1, z },
        Cube { x, y: y + 1, z },
        Cube { x, y, z: z - 1 },
        Cube { x, y, z: z + 1 },
    ]
}

fn build_grid(input: String) -> HashSet<Cube> {
    let mut grid: HashSet<Cube> = HashSet::new();
    for line in input.lines() {
        let coords = comma_separated_to_numbers(line.to_string());
        let x = coords[0];
        let y = coords[1];
        let z = coords[2];
        grid.insert(Cube { x, y, z });
    }
    grid
}

pub fn part1(input: String) -> i64 {
    let grid = build_grid(input);

    let mut free = 0;
    for cube in &grid {
        for n in neighbors(cube.clone()) {
            if !grid.contains(&n) {
                free += 1;
            }
        }
    }

    free
}

pub fn part2(input: String) -> i64 {
    let grid = build_grid(input);
    let x_min = grid.iter().map(|Cube { x, y, z }| x).min().unwrap() - 1;
    let x_max = grid.iter().map(|Cube { x, y, z }| x).max().unwrap() + 1;
    let y_min = grid.iter().map(|Cube { x, y, z }| y).min().unwrap() - 1;
    let y_max = grid.iter().map(|Cube { x, y, z }| y).max().unwrap() + 1;
    let z_min = grid.iter().map(|Cube { x, y, z }| z).min().unwrap() - 1;
    let z_max = grid.iter().map(|Cube { x, y, z }| z).max().unwrap() + 1;

    let start = Cube {
        x: x_min,
        y: y_min,
        z: z_min,
    };
    let mut exterior = HashSet::new();
    let mut todo = vec![start];
    while let Some(elt) = todo.pop() {
        for c in neighbors(elt) {
            let Cube { x, y, z } = c;
            if !grid.contains(&c)
                && !exterior.contains(&c)
                && x >= x_min
                && x <= x_max
                && y >= y_min
                && y <= y_max
                && z >= z_min
                && z <= z_max
            {
                exterior.insert(c.clone());
                todo.push(c);
            }
        }
    }

    let mut free = 0;
    for cube in &grid {
        for n in neighbors(cube.clone()) {
            if !grid.contains(&n) && exterior.contains(&n) {
                free += 1;
            }
        }
    }
    free
}
