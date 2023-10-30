use std::collections::{BinaryHeap, HashMap, HashSet};
use std::ops::{Deref, DerefMut};

#[derive(Debug, Eq, PartialEq, Hash, Clone, Copy)]
struct Pos(i64, i64);

#[derive(Debug)]
struct Grid {
    grid: HashMap<Pos, char>,
}

impl Deref for Grid {
    type Target = HashMap<Pos, char>;
    fn deref(&self) -> &Self::Target {
        &self.grid
    }
}

impl DerefMut for Grid {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.grid
    }
}

impl From<&str> for Grid {
    fn from(value: &str) -> Self {
        let mut grid = Grid::new();
        for (y, line) in value.lines().enumerate() {
            for (x, c) in line.chars().enumerate() {
                grid.insert(Pos(x as i64, y as i64), c);
            }
        }
        grid
    }
}

impl Grid {
    fn new() -> Self {
        Self {
            grid: HashMap::new(),
        }
    }

    fn start_pos(&self, start_char: char) -> Option<Pos> {
        self.grid
            .iter()
            .find_map(|(&pos, &c)| if c == start_char { Some(pos) } else { None })
    }

    fn is_wall(&self, pos: &Pos) -> bool {
        self.get(pos) == Some(&'#')
    }

    fn is_empty(&self, pos: &Pos) -> bool {
        !self.contains_key(pos) || self.get(pos) == Some(&' ')
    }

    fn neighbours(&self, Pos(x, y): &Pos) -> Vec<Pos> {
        let mut res = Vec::new();
        for (dx, dy) in [(0, 1), (0, -1), (1, 0), (-1, 0)] {
            let new_pos = Pos(x + dx, y + dy);
            if !self.is_wall(&new_pos) && !self.is_empty(&new_pos) {
                res.push(new_pos);
            }
        }
        res
    }

    fn is_branch_position(&self, pos: &Pos) -> bool {
        if let Some(c) = self.get(pos) {
            return c.is_ascii_lowercase() || c.is_ascii_uppercase();
        }
        false
    }
}

#[derive(Debug, Eq, PartialEq, Hash, Clone, Copy)]
struct Node {
    pos: Pos,
    name: char,
}

#[derive(Debug, Hash, Eq, PartialEq, Clone)]
struct State {
    current: Node,
}

#[derive(Debug, Clone, Eq, PartialEq)]
struct StateCost {
    state: State,
    cost: i32,
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

    fn add_edge(&mut self, edge: (&Node, &Node, i32)) {
        self.adjacency_table.entry(*edge.0).or_insert(Vec::new());
        self.adjacency_table.entry(*edge.1).or_insert(Vec::new());

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

    fn node_neighbours(&self, node: &Node) -> Result<&Vec<(Node, i32)>, NodeNotInGraph> {
        println!("node: {:?}", node);
        println!("adjacency_table: {:?}", self.adjacency_table);
        println!("get: {:?}", self.adjacency_table.get(node));
        return self.adjacency_table.get(node).ok_or(NodeNotInGraph);
    }

    fn neighbours(&self, state: &State) -> Vec<(State, i32)> {
        let mut res: Vec<(State, i32)> = Vec::new();
        let possible_destinations: Vec<&(Node, i32)> = self
            .node_neighbours(&state.current)
            .unwrap()
            .iter()
            .collect();

        for (to, distance) in possible_destinations {
            let mut s = state.clone();
            s.current = *to;
            res.push((s, *distance));
        }
        res
    }

    fn extend_graph(&mut self, grid: &Grid, start_nodes: Vec<Node>) {
        let mut queue: Vec<(Node, Node, i32)> = Vec::new();
        let mut visited: HashSet<(Node, Node)> = HashSet::new();
        for start in start_nodes {
            println!("start: {:?}", start);
            queue.push((start, start, 0));
        }
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
            let new_from = if grid.is_branch_position(&pos) && cell != from {
                self.add_edge((&cell, &from, distance));
                cell
            } else {
                from
            };

            grid.neighbours(&cell.pos).iter().for_each(|p| {
                let n = Node {
                    pos: *p,
                    name: *grid.get(p).unwrap(),
                };
                queue.push((n, new_from, if new_from == from { distance + 1 } else { 1 }))
            });
        }
    }

    fn add_tunnels(&mut self) {
        println!("add tunnels {:?}", self.adjacency_table.keys());
        let lower_case_nodes: Vec<Node> = self
            .adjacency_table
            .keys()
            .filter(|n| n.name.is_ascii_lowercase())
            .cloned()
            .collect();
        for lowercase_node in lower_case_nodes {
            let mut uppercase_node = lowercase_node.clone();
            uppercase_node.name = lowercase_node.name.to_ascii_uppercase();
            self.adjacency_table.entry(lowercase_node).and_modify(|e| {
                e.push((uppercase_node, 1));
            });
            self.adjacency_table.entry(uppercase_node).and_modify(|e| {
                e.push((lowercase_node, 1));
            });
        }
    }
}

fn dijkstra(graph: &Graph, start_node: Node) -> i32 {
    let mut frontier: BinaryHeap<StateCost> = BinaryHeap::new();
    let mut cost_so_far: HashMap<State, i32> = HashMap::new();
    let start = State {
        current: start_node.clone(),
    };
    frontier.push(StateCost {
        state: start.clone(),
        cost: 0,
    });
    cost_so_far.insert(start, 0);

    while let Some(StateCost {
        state: current,
        cost: _,
    }) = frontier.pop()
    {
        if current.current.name == 'Z' {
            println!("found Z");
            return cost_so_far[&current];
        }
        for (next, cost) in graph.neighbours(&current) {
            let next_cost = cost_so_far.get(&next);
            let current_cost = cost_so_far.get(&current).unwrap();
            let new_cost = current_cost + cost;
            if next_cost.is_none() || new_cost < *next_cost.unwrap() {
                cost_so_far.insert(next.clone(), new_cost);
                frontier.push(StateCost {
                    state: next,
                    cost: new_cost,
                });
            }
        }
    }
    i32::MAX
}

pub fn part1(input: String) -> i64 {
    let grid = Grid::from(input.as_str());

    let mut graph = Graph::new();
    // for start_char in 'A'..='Z' {
    //     let start = Node {
    //         pos: grid.start_pos(start_char),
    //         name: start_char,
    //     };
    //     graph.extend_graph(&grid, &start);
    // }
    let start_nodes = ('A'..='Z')
        .filter(|&start_char| grid.start_pos(start_char).is_some())
        .map(|start_char| Node {
            pos: grid.start_pos(start_char).unwrap(),
            name: start_char,
        });
    graph.extend_graph(&grid, start_nodes.collect());
    println!("{:?}", graph);

    let start = Node {
        pos: grid.start_pos('A').unwrap(),
        name: 'A',
    };
    dijkstra(&graph, start) as i64
}

pub fn part2(input: String) -> i64 {
    0
}
