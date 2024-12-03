fn mark(grid: &mut std::collections::HashMap<(i64, i64), i64>, orders: &str) {
    let mut x = 0;
    let mut y = 0;
    let mut step = 0;

    for order in orders.split(',') {
        let direction = order.chars().next().unwrap();
        let distance = order[1..].parse::<i64>().unwrap();
        for _ in 0..distance {
            match direction {
                'U' => y += 1,
                'D' => y -= 1,
                'L' => x -= 1,
                'R' => x += 1,
                _ => panic!("Unknown direction"),
            }
            step += 1;
            grid.insert((x, y), step);
        }
    }
}

fn get_grids(
    input: String,
) -> (
    std::collections::HashMap<(i64, i64), i64>,
    std::collections::HashMap<(i64, i64), i64>,
) {
    let mut grid1 = std::collections::HashMap::new();
    let mut grid2 = std::collections::HashMap::new();
    let mut orders = input.split('\n');
    let order1 = orders.next().unwrap();
    let order2 = orders.next().unwrap();

    mark(&mut grid1, order1);
    mark(&mut grid2, order2);
    (grid1, grid2)
}

pub fn part1(input: String) -> i64 {
    let (grid1, grid2) = get_grids(input);
    let mut cross = std::collections::HashMap::new();
    for (x, y) in grid1.keys() {
        if grid2.contains_key(&(*x, *y)) {
            cross.insert((x, y), 0);
        }
    }

    cross.keys().map(|(x, y)| x.abs() + y.abs()).min().unwrap()
}

pub fn part2(input: String) -> i64 {
    let (grid1, grid2) = get_grids(input);
    let mut cross = std::collections::HashMap::new();
    for ((x, y), v1) in grid1 {
        if grid2.contains_key(&(x, y)) {
            let v2 = grid2.get(&(x, y)).unwrap();
            if !cross.contains_key(&(x, y)) {
                cross.insert((x, y), v1 + v2);
            }
        }
    }
    cross.values().min().unwrap().clone()
}
