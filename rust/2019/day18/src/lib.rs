use std::collections::{BinaryHeap, HashMap, HashSet};
use std::ops::{Deref, DerefMut};

#[derive(Debug, Eq, PartialEq, Hash, Clone, Copy)]
struct Pos(i64, i64);

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

    fn start_pos(&self) -> Pos {
        self.grid
            .iter()
            .find_map(|(&pos, &c)| if c == '@' { Some(pos) } else { None })
            .unwrap()
    }

    fn is_wall(&self, pos: &Pos) -> bool {
        self.get(pos) == Some(&'#')
    }

    fn neighbours(&self, Pos(x, y): &Pos) -> Vec<Pos> {
        let mut res = Vec::new();
        for (dx, dy) in [(0, 1), (0, -1), (1, 0), (-1, 0)] {
            let new_pos = Pos(x + dx, y + dy);
            if !self.is_wall(&new_pos) && self.contains_key(&new_pos) {
                res.push(new_pos);
            }
        }
        res
    }

    fn number_of_keys(&self) -> usize {
        self.grid
            .iter()
            .filter(|(_, c)| c.is_ascii_lowercase())
            .count()
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
    current: Vec<Node>,
    keys: Vec<char>,
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
        self.adjacency_table.get(node).ok_or(NodeNotInGraph)
    }

    fn neighbours(&self, state: &State) -> Vec<(State, i32)> {
        let mut res: Vec<(State, i32)> = Vec::new();
        for i in 0..state.current.len() {
            let possible_destinations: Vec<&(Node, i32)> = self
                .node_neighbours(&state.current[i])
                .unwrap()
                .iter()
                .collect();

            for (to, distance) in possible_destinations {
                let mut s = state.clone();
                s.current[i] = *to;
                if to.name.is_ascii_uppercase()
                    && !state.keys.contains(&to.name.to_ascii_lowercase())
                {
                    continue;
                }
                if to.name.is_ascii_lowercase() {
                    s.keys.push(to.name);
                    s.keys.sort();
                    s.keys.dedup();
                }
                res.push((s, *distance));
            }
        }
        res
    }

    fn extend_graph(&mut self, grid: &Grid, start: &Node) {
        let mut queue: Vec<(Node, Node, i32)> = Vec::new();
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
}

fn dijkstra(grid: &Grid, graph: &Graph, start_nodes: Vec<Node>) -> i32 {
    let number_of_keys = grid.number_of_keys();
    let mut frontier: BinaryHeap<StateCost> = BinaryHeap::new();
    let mut cost_so_far: HashMap<State, i32> = HashMap::new();
    let start = State {
        current: start_nodes.clone(),
        keys: Vec::new(),
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
        if current.keys.len() == number_of_keys {
            println!("Found all keys: {:?}", current);
            return cost_so_far[&current];
        }
        let neighbours: Vec<_> = graph.neighbours(&current);
        for (next, cost) in neighbours.iter() {
            let next_cost = cost_so_far.get(&next);
            let current_cost = cost_so_far.get(&current).unwrap();
            let new_cost = current_cost + cost;
            if next_cost.is_none() || new_cost < *next_cost.unwrap() {
                cost_so_far.insert(next.clone(), new_cost);
                frontier.push(StateCost {
                    state: next.clone(),
                    cost: new_cost,
                });
            }
        }
    }
    i32::MAX
}

pub fn part1(input: String) -> i64 {
    let grid = Grid::from(input.as_str());
    let start = Node {
        pos: grid.start_pos(),
        name: '@',
    };
    let mut graph = Graph::new();
    graph.extend_graph(&grid, &start);
    dijkstra(&grid, &graph, vec![start]) as i64
}

pub fn part2(input: String) -> i64 {
    let mut grid = Grid::from(input.as_str());
    let center = grid.start_pos();
    let Pos(cx, cy) = center;
    grid.insert(center, '#');
    for n in grid.neighbours(&center) {
        grid.insert(n, '#');
    }
    let start_positions = [(-1, -1), (1, -1), (-1, 1), (1, 1)]
        .into_iter()
        .map(|(dx, dy)| Pos(cx + dx, cy + dy));

    for pos in start_positions.clone() {
        grid.insert(pos, '@');
    }

    let start_nodes = start_positions
        .map(|pos| Node { pos, name: '@' })
        .collect::<Vec<Node>>();
    let mut graph = Graph::new();
    for start in start_nodes.iter() {
        graph.extend_graph(&grid, start);
    }
    dijkstra(&grid, &graph, start_nodes) as i64
}
