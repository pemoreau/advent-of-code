use std::collections::{BinaryHeap, HashMap, HashSet};

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

#[derive(Debug)]
struct State {
    current: Node,
    keys: Vec<char>,
    path: Vec<char>,
    cost: i32,
}

#[derive(Debug, Hash, Eq, PartialEq, Clone)]
struct State2 {
    current: Vec<Node>,
    keys: Vec<char>,
}

#[derive(Debug, Clone)]
struct StateCost {
    state: State2,
    cost: i32,
}

impl Eq for StateCost {}

impl PartialEq for StateCost {
    fn eq(&self, other: &Self) -> bool {
        self.cost == other.cost
            && self.state.current == other.state.current
            && self.state.keys == other.state.keys
    }
}

impl PartialOrd for StateCost {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(other.cmp(self))
    }
}

impl Ord for StateCost {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.cost.cmp(&other.cost)
    }
}

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
                // println!("Found all keys: {:?}", state);
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

    fn display(&self) {
        let mut min_x = i64::MAX;
        let mut max_x = i64::MIN;
        let mut min_y = i64::MAX;
        let mut max_y = i64::MIN;
        for (pos, _) in self.grid.iter() {
            if pos.x < min_x {
                min_x = pos.x;
            }
            if pos.x > max_x {
                max_x = pos.x;
            }
            if pos.y < min_y {
                min_y = pos.y;
            }
            if pos.y > max_y {
                max_y = pos.y;
            }
        }
        for y in min_y..=max_y {
            for x in min_x..=max_x {
                let pos = Pos { x, y };
                if self.contains_key(&pos) {
                    print!("{}", self.get(&pos).unwrap());
                } else {
                    print!(" ");
                }
            }
            println!();
        }
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

fn extend_graph(graph: &mut Graph, grid: &Grid, start: &Node) {
    let mut queue = Vec::new();
    let mut visited: HashSet<(Node, Node)> = HashSet::new();
    queue.push((*start, *start, 0));
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
}

fn build_grid(input: String) -> Grid {
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
    grid
}

pub fn part1(input: String) -> i64 {
    let grid = build_grid(input);
    let start = Node {
        pos: grid.start_pos(),
        name: '@',
    };
    let mut graph = Graph::new();
    extend_graph(&mut graph, &grid, &start);
    // println!("graph {:?}", graph);
    // bfs(&graph, start, grid.number_of_keys()) as i64
    // bfs2(&graph, vec![start], grid.number_of_keys()) as i64
    dijkstra(&graph, vec![start], grid.number_of_keys()) as i64
}

fn neighbours2(graph: &Graph, state: &State2) -> Vec<(State2, i32)> {
    let mut neighbours: Vec<(State2, i32)> = Vec::new();
    for i in 0..state.current.len() {
        let possible_destinations: Vec<&(Node, i32)> = graph
            .neighbours(&state.current[i])
            .unwrap()
            .iter()
            .filter(|(node, _)| state.current.iter().find(|n| n.pos == node.pos).is_none())
            .collect();

        for (to, distance) in possible_destinations {
            if to.name.is_ascii_lowercase() {
                let mut keys = state.keys.clone();
                keys.push(to.name);
                keys.sort();
                keys.dedup();
                let mut current = state.current.clone();
                current[i] = to.clone();
                neighbours.push((
                    State2 {
                        current: current,
                        keys: keys,
                        // cost: state.cost + distance,
                    },
                    *distance,
                ));
            } else if state.keys.contains(&to.name.to_ascii_lowercase())
                || !to.name.is_ascii_uppercase()
            {
                let mut current = state.current.clone();
                current[i] = to.clone();
                neighbours.push((
                    State2 {
                        current: current,
                        keys: state.keys.clone(),
                        // cost: state.cost + distance,
                    },
                    *distance,
                ));
            }
        }
    }
    neighbours
}

fn dijkstra(graph: &Graph, start_nodes: Vec<Node>, number_of_keys: usize) -> i32 {
    let mut frontier: BinaryHeap<StateCost> = BinaryHeap::new();
    let mut cost_so_far: HashMap<State2, i32> = HashMap::new();
    let start = State2 {
        current: start_nodes.clone(),
        keys: Vec::new(),
    };
    frontier.push(StateCost {
        state: start.clone(),
        cost: 0,
    });
    cost_so_far.insert(start.clone(), 0);

    while !frontier.is_empty() {
        let StateCost {
            state: current,
            cost,
        } = frontier.pop().unwrap();
        // println!("pop {:?} {:?}", current, cost);
        if current.keys.len() == number_of_keys {
            println!("Found all keys: {:?}", current);
            return cost_so_far[&current];
        }
        for (next, cost) in neighbours2(graph, &current) {
            let next_cost = cost_so_far.get(&next);
            let current_cost = cost_so_far.get(&current).unwrap();
            let new_cost = current_cost + cost;
            // println!("next_cost {:?} new_cost {:?}", next_cost, new_cost);
            if next_cost.is_none() || new_cost < *next_cost.unwrap() {
                cost_so_far.insert(next.clone(), new_cost);
                frontier.push(StateCost {
                    state: next.clone(),
                    cost: new_cost - next.keys.len() as i32,
                    // cost: new_cost,
                });
                let len = frontier.len();
                if len % 10000 == 0 {
                    println!("len {:?}", len);
                }
                // println!("push {:?} {:?}", next, new_cost);
                // println!("frontier {:?}", frontier);
            }
        }
    }
    0
}

// fn bfs2(graph: &Graph, start_nodes: Vec<Node>, number_of_keys: usize) -> i32 {
//     let mut min_cost = i32::MAX;
//     let mut queue = Vec::new();
//     let mut visited: HashMap<(Vec<Node>, Vec<char>), i32> = HashMap::new();
//     queue.push(State2 {
//         current: start_nodes,
//         keys: Vec::new(),
//         cost: 0,
//     });
//
//     while !queue.is_empty() {
//         let state = queue.remove(0);
//         // println!("{:?}", state);
//         if visited.contains_key(&(state.current.clone(), state.keys.clone())) {
//             let cost = visited
//                 .get(&(state.current.clone(), state.keys.clone()))
//                 .unwrap();
//             if state.cost >= *cost {
//                 // println!("skip {:?} {:?}", state, cost);
//                 continue;
//             }
//         }
//         visited.insert((state.current.clone(), state.keys.clone()), state.cost);
//
//         if state.keys.len() == number_of_keys {
//             if state.cost < min_cost {
//                 // println!("Found all keys: {:?}", state);
//                 min_cost = state.cost;
//             }
//             continue;
//         }
//
//         let neighboors = neighbours2(graph, &state);
//         // println!("neighboors {:?}", neighboors);
//         for neighboor in neighboors {
//             // println!("add {:?}", neighboor);
//             queue.push(neighboor);
//         }
//     }
//     min_cost
// }

pub fn part2(input: String) -> i64 {
    let mut grid = build_grid(input);
    let center = grid.start_pos();
    let neighbors = grid.neighboors(&center);
    grid.insert(center, '#');
    for n in neighbors {
        grid.insert(n, '#');
    }
    let start_positions = vec![
        Pos {
            x: center.x - 1,
            y: center.y - 1,
        },
        Pos {
            x: center.x + 1,
            y: center.y - 1,
        },
        Pos {
            x: center.x - 1,
            y: center.y + 1,
        },
        Pos {
            x: center.x + 1,
            y: center.y + 1,
        },
    ];
    for pos in start_positions.clone() {
        grid.insert(pos, '@');
    }
    // println!("graph {:?}", graph);
    // grid.display();

    let start_nodes = start_positions
        .iter()
        .map(|pos| Node {
            pos: *pos,
            name: '@',
        })
        .collect::<Vec<Node>>();
    let mut graph = Graph::new();
    for start in start_nodes.clone() {
        extend_graph(&mut graph, &grid, &start);
    }
    // println!("graph {:?}", graph);
    // bfs2(&graph, start_nodes, grid.number_of_keys()) as i64
    dijkstra(&graph, start_nodes, grid.number_of_keys()) as i64
}
