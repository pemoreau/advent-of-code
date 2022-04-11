use std::collections::HashMap;
fn main() {
    let mut map: HashMap<usize, usize> = HashMap::new();
    let n = 10;

    let index = map.get(&n);
    map.insert(n, 0);
    let res = index.unwrap();
}
