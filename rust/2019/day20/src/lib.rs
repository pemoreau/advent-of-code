use std::collections::{BinaryHeap, HashMap, HashSet};
use std::hash::Hash;
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
            return c.is_ascii_lowercase() || c.is_ascii_uppercase() || c == &'0' || c == &'1';
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
#[derive(Debug, Hash, Eq, PartialEq, Clone)]
struct State2 {
    current: Node,
    level: i32,
}
#[derive(Debug, Clone, Eq, PartialEq)]
struct StateCost<S> {
    state: S,
    cost: i32,
}

impl<S: Eq> PartialOrd for StateCost<S> {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(other.cmp(self))
    }
}

impl<S: Eq> Ord for StateCost<S> {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.cost.cmp(&other.cost)
    }
}

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

    fn to_string(&self) -> String {
        let mut res = String::new();
        for (node, neighbours) in &self.adjacency_table {
            res.push_str(&format!("{}: ", node.name));
            for (neighbour, distance) in neighbours {
                res.push_str(&format!("{}({}), ", neighbour.name, distance));
            }
            res.push('\n');
        }
        res
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
        return self.adjacency_table.get(node).ok_or(NodeNotInGraph);
    }

    fn neighbours(&self, state: &State) -> Vec<(State, i32)> {
        let possible_destinations = self.node_neighbours(&state.current).unwrap().iter();
        possible_destinations
            .map(|(to, distance)| (State { current: *to }, *distance))
            .collect()
    }

    fn neighbours2(&self, state: &State2) -> Vec<(State2, i32)> {
        let mut res: Vec<(State2, i32)> = Vec::new();
        let possible_destinations = self.node_neighbours(&state.current).unwrap().iter();

        for (to, distance) in possible_destinations {
            let from_name = state.current.name;
            let is_outer = |name: char| name.is_ascii_uppercase() || name == 'É';
            let is_inner = |name: char| name.is_ascii_lowercase() || name == 'é';
            if state.level == 0 && is_outer(from_name) {
                // at level 0 can only exit via 1
                continue;
            }
            let mut level = state.level;
            if is_outer(from_name) && distance == &1 {
                // going from outside to inside port make level decrease
                level -= 1;
            } else if is_inner(from_name) && distance == &1 {
                level += 1;
            }

            res.push((
                State2 {
                    current: *to,
                    level,
                },
                *distance,
            ));
        }

        res
    }

    fn extend_graph(&mut self, grid: &Grid, start_nodes: Vec<Node>) {
        let mut queue: Vec<(Node, Node, i32)> = Vec::new();
        let mut visited: HashSet<(Node, Node)> = HashSet::new();
        for start in start_nodes {
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
        let map_char_node: HashMap<char, Vec<Node>> =
            self.adjacency_table
                .keys()
                .cloned()
                .fold(HashMap::new(), |mut acc, node| {
                    acc.entry(node.name).or_insert(Vec::new()).push(node);
                    acc
                });

        for c in map_char_node.keys() {
            if c == &'é' {
                self.add_edge((&map_char_node[&'é'][0], &map_char_node[&'É'][0], 1));
            } else if c == &'É' {
                self.add_edge((&map_char_node[&'É'][0], &map_char_node[&'é'][0], 1));
            } else if c.is_ascii_lowercase() {
                let uppercase = c.to_ascii_uppercase();
                self.add_edge((&map_char_node[c][0], &map_char_node[&uppercase][0], 1));
            } else if c.is_ascii_uppercase() {
                let lowercase = c.to_ascii_lowercase();
                self.add_edge((&map_char_node[c][0], &map_char_node[&lowercase][0], 1));
            }
        }
    }
}

fn dijkstra<S: Eq + Hash + Clone, Graph>(
    graph: &Graph,
    start: S,
    neighbours: fn(graph: &Graph, state: &S) -> Vec<(S, i32)>,
    goal: fn(state: &S) -> bool,
) -> i32 {
    let mut frontier: BinaryHeap<StateCost<S>> = BinaryHeap::new();
    let mut cost_so_far: HashMap<S, i32> = HashMap::new();
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
        if goal(&current) {
            return cost_so_far[&current];
        }

        for (next, cost) in neighbours(graph, &current) {
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
fn solve<S: Eq + Hash + Clone>(
    grid: &Grid,
    start: S,
    neighbours: fn(graph: &Graph, state: &S) -> Vec<(S, i32)>,
    goal: fn(state: &S) -> bool,
) -> i64 {
    let mut graph = Graph::new();
    let start_nodes: Vec<Node> = ('A'..='Z')
        .chain('a'..='z')
        .chain("01éÉ".chars())
        .filter(|&start_char| grid.start_pos(start_char).is_some())
        .map(|start_char| Node {
            pos: grid.start_pos(start_char).unwrap(),
            name: start_char,
        })
        .collect();
    graph.extend_graph(&grid, start_nodes);
    graph.add_tunnels();
    dijkstra(&graph, start, neighbours, goal) as i64
}

pub fn part1(input: String) -> i64 {
    let grid = Grid::from(input.as_str());

    let start = State {
        current: Node {
            pos: grid.start_pos('0').unwrap(),
            name: '0',
        },
    };

    let neighbours = |graph: &Graph, state: &State| graph.neighbours(state);
    let goal = |state: &State| state.current.name == '1';
    solve(&grid, start, neighbours, goal)
}

pub fn part2(input: String) -> i64 {
    let grid = Grid::from(input.as_str());

    let start = State2 {
        current: Node {
            pos: grid.start_pos('0').unwrap(),
            name: '0',
        },
        level: 0,
    };
    let neighbours = |graph: &Graph, state: &State2| graph.neighbours2(state);
    let goal = |state: &State2| state.current.name == '1' && state.level == 0;
    solve(&grid, start, neighbours, goal)
}
