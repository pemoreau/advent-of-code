use std::collections::HashMap;

#[derive(Debug, Eq, PartialEq, Hash)]
struct Edge {
    to: char,
    distance: u32,
}

impl Edge {
    fn new(to: char, distance: u32) -> Self {
        Self { to, distance }
    }
}

#[derive(Debug)]
struct State {
    node: char,
    keys: Vec<char>,
    cost: u32,
}

type Graph = HashMap<char, Vec<Edge>>;

fn neighboors(graph: &Graph, state: &State) -> Vec<State> {
    let mut neighboors = Vec::new();
    for Edge { to, distance } in graph.get(&state.node).unwrap() {
        if to.is_ascii_lowercase() {
            let mut keys = state.keys.clone();
            if !keys.contains(to) {
                keys.push(*to);
                keys.sort();
            }
            neighboors.push(State {
                node: *to,
                keys: keys,
                cost: state.cost + distance,
            });
        } else if state.keys.contains(&to.to_ascii_lowercase()) || !to.is_ascii_uppercase() {
            neighboors.push(State {
                node: *to,
                keys: state.keys.clone(),
                cost: state.cost + distance,
            });
        }
    }
    neighboors
}

fn bfs(graph: &Graph, start: char) -> u32 {
    let mut queue = Vec::new();
    let mut visited = HashMap::new();
    queue.push(State {
        node: start,
        keys: Vec::new(),
        cost: 0,
    });
    while !queue.is_empty() {
        let state = queue.remove(0);
        println!("{:?}", state);
        if visited.contains_key(&(state.node, state.keys.clone())) {
            let cost = visited.get(&(state.node, state.keys.clone())).unwrap();
            if state.cost >= *cost {
                println!("skip {:?} {:?}", state, cost);
                continue;
            }
        }
        if state.keys.len() == 7 {
            println!("Found all keys in {} steps", state.cost);
            continue;
        }

        visited.insert((state.node, state.keys.clone()), state.cost);
        let neighboors = neighboors(graph, &state);
        println!("neighboors {:?}", neighboors);
        for neighboor in neighboors {
            println!("add {:?}", neighboor);
            queue.push(neighboor);
        }
    }
    *visited.values().min().unwrap()
}

// https://kuczma.dev/articles/rust-graphs/

pub fn part1(input: String) -> i64 {
    // create nodes: f-D-C-b-@-a-B-c-d-A-e-F-g
    let mut graph: Graph = HashMap::new();
    graph.insert('f', vec![Edge::new('D', 2)]);
    graph.insert('D', vec![Edge::new('C', 2), Edge::new('f', 2)]);
    graph.insert('C', vec![Edge::new('b', 2), Edge::new('D', 2)]);
    graph.insert('b', vec![Edge::new('@', 22), Edge::new('C', 2)]);
    graph.insert('@', vec![Edge::new('a', 2), Edge::new('b', 22)]);
    graph.insert('a', vec![Edge::new('B', 2), Edge::new('@', 2)]);
    graph.insert('B', vec![Edge::new('c', 2), Edge::new('a', 2)]);
    graph.insert('c', vec![Edge::new('d', 2), Edge::new('B', 2)]);
    graph.insert('d', vec![Edge::new('A', 2), Edge::new('c', 2)]);
    graph.insert('A', vec![Edge::new('e', 2), Edge::new('d', 2)]);
    graph.insert('e', vec![Edge::new('F', 2), Edge::new('A', 2)]);
    graph.insert('F', vec![Edge::new('g', 2), Edge::new('e', 2)]);
    graph.insert('g', vec![Edge::new('F', 2)]);

    bfs(&graph, '@');
    0
}

pub fn part2(input: String) -> i64 {
    0
}
