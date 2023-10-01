use std::collections::{BinaryHeap, HashMap, HashSet};

#[derive(Debug, Eq, PartialEq, Hash, Clone, Copy)]
struct Pos {
    x: i64,
    y: i64,
}

struct Grid {
    grid: HashMap<Pos, char>,
}

impl Grid {
    fn new() -> Self {
        Self {
            grid: HashMap::new(),
        }
    }

    fn build_grid(input: String) -> Self {
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

    fn start_pos(&self) -> Pos {
        self.grid
            .iter()
            .find(|(_, c)| **c == '@')
            .map(|(pos, _)| *pos)
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

    fn neighbours(&self, pos: &Pos) -> Vec<Pos> {
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

    fn is_branch_position(&self, pos: &Pos) -> bool {
        let &c = self.get(&pos).unwrap();
        if c == '#' {
            return false;
        }
        c.is_ascii_lowercase() || c.is_ascii_uppercase()
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

#[derive(Debug, Clone)]
struct StateCost {
    state: State,
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

    fn node_neighbours(&self, node: &Node) -> Result<&Vec<(Node, i32)>, NodeNotInGraph> {
        match self.adjacency_table.get(node) {
            None => Err(NodeNotInGraph),
            Some(i) => Ok(i),
        }
    }

    fn neighbours(&self, state: &State) -> Vec<(State, i32)> {
        let mut res: Vec<(State, i32)> = Vec::new();
        for i in 0..state.current.len() {
            let possible_destinations: Vec<&(Node, i32)> = self
                .node_neighbours(&state.current[i])
                .unwrap()
                .iter()
                .filter(|(node, _)| state.current.iter().find(|n| n.pos == node.pos).is_none())
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

    while !frontier.is_empty() {
        let StateCost {
            state: current,
            cost,
        } = frontier.pop().unwrap();
        if current.keys.len() == number_of_keys {
            println!("Found all keys: {:?}", current);
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
    0
}

pub fn part1(input: String) -> i64 {
    let grid = Grid::build_grid(input);
    let start = Node {
        pos: grid.start_pos(),
        name: '@',
    };
    let mut graph = Graph::new();
    graph.extend_graph(&grid, &start);
    dijkstra(&grid, &graph, vec![start]) as i64
}

pub fn part2(input: String) -> i64 {
    let mut grid = Grid::build_grid(input);
    let center = grid.start_pos();
    let neighbors = grid.neighbours(&center);
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
    for pos in start_positions.iter() {
        grid.insert(*pos, '@');
    }

    let start_nodes = start_positions
        .iter()
        .map(|pos| Node {
            pos: *pos,
            name: '@',
        })
        .collect::<Vec<Node>>();
    let mut graph = Graph::new();
    for start in start_nodes.iter() {
        graph.extend_graph(&grid, start);
    }
    dijkstra(&grid, &graph, start_nodes) as i64
}
