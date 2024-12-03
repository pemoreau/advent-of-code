use std::collections::HashMap;
fn main() {
    let mut map: HashMap<usize, String> = HashMap::new();
    let n = 10;

    let value = map.get(&n);
    map.insert(n, "hello".to_string());
    if value.is_some() {
        println!("{}", value.unwrap());
    }
}
