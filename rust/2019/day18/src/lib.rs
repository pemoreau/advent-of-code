use std::collections::{HashMap, HashSet};

#[derive(Debug, Eq, PartialEq, Hash, Clone)]
struct Pos {
    x: i64,
    y: i64,
}

#[derive(Debug, Eq, PartialEq, Hash, Clone)]
struct Node {
    pos: Pos,
    name: char,
}

#[derive(Debug, Eq, PartialEq, Hash)]
struct Edge {
    to: Node,
    distance: u32,
}

impl Edge {
    fn new(to: Node, distance: u32) -> Self {
        Self { to, distance }
    }
}

#[derive(Debug)]
struct State {
    current: Node,
    keys: Vec<char>,
    path: Vec<char>,
    cost: u32,
}

type Graph = HashMap<Node, Vec<Edge>>;

fn neighboors(graph: &Graph, state: &State) -> Vec<State> {
    let mut neighboors = Vec::new();
    for Edge { to, distance } in graph.get(&state.current).unwrap() {
        let mut path = state.path.clone();
        path.push(to.name);
        if to.name.is_ascii_lowercase() {
            let mut keys = state.keys.clone();
            if !keys.contains(&to.name) {
                keys.push(to.name);
                keys.sort();
            }
            neighboors.push(State {
                current: to.clone(),
                keys: keys,
                path: path,
                cost: state.cost + distance,
            });
        } else if state.keys.contains(&to.name.to_ascii_lowercase())
            || !to.name.is_ascii_uppercase()
        {
            neighboors.push(State {
                current: to.clone(),
                keys: state.keys.clone(),
                path: path,
                cost: state.cost + distance,
            });
        }
    }
    neighboors
}

fn bfs(graph: &Graph, start: Node, number_of_keys: usize) -> u32 {
    let mut min_cost = u32::MAX;
    let mut queue = Vec::new();
    let mut visited: HashMap<(Node, Vec<char>), u32> = HashMap::new();
    queue.push(State {
        current: start,
        keys: Vec::new(),
        path: Vec::new(),
        cost: 0,
    });
    while !queue.is_empty() {
        let state = queue.remove(0);
        println!("{:?}", state);
        if visited.contains_key(&(state.current.clone(), state.keys.clone())) {
            let cost = visited
                .get(&(state.current.clone(), state.keys.clone()))
                .unwrap();
            if state.cost >= *cost {
                println!("skip {:?} {:?}", state, cost);
                continue;
            }
        }
        visited.insert((state.current.clone(), state.keys.clone()), state.cost);

        if state.keys.len() == number_of_keys {
            if state.cost < min_cost {
                println!("Found all keys: {:?}", state);
                min_cost = state.cost;
            }
            continue;
        }

        let neighboors = neighboors(graph, &state);
        // println!("neighboors {:?}", neighboors);
        for neighboor in neighboors {
            println!("add {:?}", neighboor);
            queue.push(neighboor);
        }
    }
    min_cost
}

// https://kuczma.dev/articles/rust-graphs/

struct Grid {
    grid: HashMap<Pos, char>,
}

impl Grid {
    fn new() -> Self {
        Self {
            grid: HashMap::new(),
        }
    }

    fn start_pos(&self) -> Pos {
        self.grid
            .iter()
            .find(|(_, c)| **c == '@')
            .map(|(pos, _)| pos.clone())
            .unwrap()
    }

    fn insert(&mut self, pos: Pos, c: char) {
        self.grid.insert(pos, c);
    }

    fn get(&self, pos: &Pos) -> Option<&char> {
        self.grid.get(pos)
    }

    fn contains_key(&self, pos: &Pos) -> bool {
        self.grid.contains_key(pos)
    }

    fn is_wall(&self, pos: &Pos) -> bool {
        self.get(pos).unwrap() == &'#'
    }

    fn neighboors(&self, pos: &Pos) -> Vec<Pos> {
        let mut neighboors = Vec::new();
        for (dx, dy) in vec![(0, 1), (0, -1), (1, 0), (-1, 0)] {
            let new_pos = Pos {
                x: pos.x + dx,
                y: pos.y + dy,
            };
            if self.contains_key(&new_pos) && !self.is_wall(&new_pos) {
                neighboors.push(new_pos);
            }
        }
        neighboors
    }

    fn is_branching_position(&self, pos: &Pos) -> bool {
        if self.is_wall(&pos) {
            return false;
        }

        let c = self.get(&pos).unwrap();
        c.is_ascii() || self.neighboors(&pos).len() > 2
    }

    fn number_of_keys(&self) -> usize {
        self.grid
            .iter()
            .filter(|(_, c)| c.is_ascii_lowercase())
            .count()
    }
}

fn is_branch_position(grid: &Grid, pos: &Pos) -> bool {
    let c = grid.get(&pos).unwrap();
    if c == &'#' {
        return false;
    }

    let mut neighboors = 0;
    for (dx, dy) in vec![(0, 1), (0, -1), (1, 0), (-1, 0)] {
        let new_pos = Pos {
            x: pos.x + dx,
            y: pos.y + dy,
        };
        if grid.contains_key(&new_pos) && grid.get(&new_pos).unwrap() != &'#' {
            neighboors += 1;
        }
    }
    // println!("branching {} {:?}", neighboors, pos);
    c.is_ascii_lowercase() || c.is_ascii_uppercase() || neighboors > 2
}

pub fn part1(input: String) -> i64 {
    let mut grid = Grid::new();
    for (y, line) in input.lines().enumerate() {
        for (x, c) in line.chars().enumerate() {
            grid.insert(
                Pos {
                    x: x as i64,
                    y: y as i64,
                },
                c,
            );
        }
    }
    let mut graph = HashMap::new();
    let mut queue = Vec::new();
    let mut visited: HashSet<Node> = HashSet::new();
    let start = Node {
        pos: grid.start_pos(),
        name: '@',
    };
    queue.push((start.clone(), start.clone(), 0));
    while !queue.is_empty() {
        let (cell, from, distance) = queue.remove(0);
        if grid.is_wall(&cell.pos.clone()) {
            continue;
        }
        if visited.contains(&cell.clone()) {
            continue;
        }
        visited.insert(cell.clone());
        let new_from = if is_branch_position(&grid, &cell.pos.clone()) && cell != from {
            graph
                .entry(from.clone())
                .or_insert(Vec::new())
                .push(Edge::new(cell.clone(), distance));
            graph
                .entry(cell.clone())
                .or_insert(Vec::new())
                .push(Edge::new(from.clone(), distance));
            cell.clone()
        } else {
            from.clone()
        };

        grid.neighboors(&cell.pos).iter().for_each(|neighboor| {
            queue.push((
                Node {
                    pos: neighboor.clone(),
                    name: *grid.get(&neighboor).unwrap(),
                },
                new_from.clone(),
                if new_from == from { distance + 1 } else { 1 },
            ))
        });
    }
    println!("graph {:?}", graph);
    bfs(&graph, start, grid.number_of_keys()) as i64
}

// 462 too low
// 5410 too high

pub fn part2(input: String) -> i64 {
    0
}
