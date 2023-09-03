use std::collections::{HashMap, HashSet};

#[derive(Debug, Eq, PartialEq, Hash, Clone, Copy)]
struct Pos {
    x: i64,
    y: i64,
}

#[derive(Debug, Eq, PartialEq, Hash, Clone, Copy)]
struct Node {
    pos: Pos,
    name: char,
}

// #[derive(Debug, Eq, PartialEq, Hash)]
// struct Edge {
//     to: Node,
//     distance: u32,
// }
//
// impl Edge {
//     fn new(to: Node, distance: u32) -> Self {
//         Self { to, distance }
//     }
// }

#[derive(Debug)]
struct State {
    current: Node,
    keys: Vec<char>,
    path: Vec<char>,
    cost: i32,
}

// type Graph = HashMap<Node, HashSet<Edge>>;

#[derive(Debug)]
struct Graph {
    adjacency_table: HashMap<Node, Vec<(Node, i32)>>,
}

#[derive(Debug, Clone)]
struct NodeNotInGraph;

impl Graph {
    fn new() -> Graph {
        Graph {
            adjacency_table: HashMap::new(),
        }
    }

    fn add_node(&mut self, node: &Node) -> bool {
        match self.adjacency_table.get(node) {
            None => {
                self.adjacency_table.insert(*node, Vec::new());
                true
            }
            _ => false,
        }
    }

    fn add_edge(&mut self, edge: (&Node, &Node, i32)) {
        self.add_node(edge.0);
        self.add_node(edge.1);

        self.adjacency_table.entry(*edge.0).and_modify(|e| {
            if !e.contains(&(*edge.1, edge.2)) {
                e.push((*edge.1, edge.2));
            }
        });
        self.adjacency_table.entry(*edge.1).and_modify(|e| {
            if !e.contains(&(*edge.0, edge.2)) {
                e.push((*edge.0, edge.2));
            }
        });
    }

    // fn nodes(&self) -> HashSet<&Node> {
    //     self.adjacency_table.keys().collect()
    // }
    //
    // fn edges(&self) -> Vec<(&Node, &Node, i32)> {
    //     let mut edges = Vec::new();
    //     for (from_node, from_node_neighbours) in self.adjacency_table {
    //         for (to_node, weight) in from_node_neighbours {
    //             edges.push((&from_node, &to_node, weight));
    //         }
    //     }
    //     edges
    // }

    fn neighbours(&self, node: &Node) -> Result<&Vec<(Node, i32)>, NodeNotInGraph> {
        match self.adjacency_table.get(node) {
            None => Err(NodeNotInGraph),
            Some(i) => Ok(i),
        }
    }
}

fn neighbours(graph: &Graph, state: &State) -> Vec<State> {
    let mut neighbours = Vec::new();
    for (to, distance) in graph.neighbours(&state.current).unwrap() {
        let mut path = state.path.clone();
        path.push(to.name);
        if to.name.is_ascii_lowercase() {
            let mut keys = state.keys.clone();
            if !keys.contains(&to.name) {
                keys.push(to.name);
                keys.sort();
            }
            neighbours.push(State {
                current: to.clone(),
                keys: keys,
                path: path,
                cost: state.cost + distance,
            });
        } else if state.keys.contains(&to.name.to_ascii_lowercase())
            || !to.name.is_ascii_uppercase()
        {
            neighbours.push(State {
                current: to.clone(),
                keys: state.keys.clone(),
                path: path,
                cost: state.cost + distance,
            });
        }
    }
    neighbours
}

fn bfs(graph: &Graph, start: Node, number_of_keys: usize) -> i32 {
    let mut min_cost = i32::MAX;
    let mut queue = Vec::new();
    let mut visited: HashMap<(Node, Vec<char>), i32> = HashMap::new();
    queue.push(State {
        current: start,
        keys: Vec::new(),
        path: Vec::new(),
        cost: 0,
    });
    while !queue.is_empty() {
        let state = queue.remove(0);
        // println!("{:?}", state);
        if visited.contains_key(&(state.current.clone(), state.keys.clone())) {
            let cost = visited
                .get(&(state.current.clone(), state.keys.clone()))
                .unwrap();
            if state.cost >= *cost {
                // println!("skip {:?} {:?}", state, cost);
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

        let neighboors = neighbours(graph, &state);
        // println!("neighboors {:?}", neighboors);
        for neighboor in neighboors {
            // println!("add {:?}", neighboor);
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

    fn number_of_keys(&self) -> usize {
        self.grid
            .iter()
            .filter(|(_, c)| c.is_ascii_lowercase())
            .count()
    }
}

fn is_branch_position(grid: &Grid, pos: &Pos) -> bool {
    let &c = grid.get(&pos).unwrap();
    if c == '#' {
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

fn build_graph(grid: &Grid) -> Graph {
    let mut graph = Graph::new();
    let mut queue = Vec::new();
    let mut visited: HashSet<(Node, Node)> = HashSet::new();
    let start = Node {
        pos: grid.start_pos(),
        name: '@',
    };
    queue.push((start, start, 0));
    while !queue.is_empty() {
        let (cell, from, distance) = queue.remove(0);
        let pos = cell.pos;
        if grid.is_wall(&pos) {
            continue;
        }
        if visited.contains(&(cell, from)) {
            continue;
        }
        visited.insert((cell, from));
        let new_from = if is_branch_position(&grid, &pos) && cell != from {
            graph.add_edge((&cell, &from, distance));
            cell
        } else {
            from
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
    graph
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

    let start = Node {
        pos: grid.start_pos(),
        name: '@',
    };

    let graph = build_graph(&grid);
    // println!("graph {:?}", graph);
    bfs(&graph, start, grid.number_of_keys()) as i64
}

pub fn part2(input: String) -> i64 {
    0
}
