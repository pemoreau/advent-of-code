use std::collections::HashMap;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Chemical {
    n: i64,
    name: String,
}

fn parse(input: String) -> HashMap<Chemical, Vec<Chemical>> {
    let mut reactions = HashMap::new();
    for line in input.lines() {
        let mut parts = line.split("=>");
        let lhs = parts.next().unwrap();
        let rhs = parts.next().unwrap().trim();
        let mut rhs_parts = rhs.split(" ");
        let rhs_n = rhs_parts.next().unwrap().parse::<i64>().unwrap();
        let rhs_name = rhs_parts.next().unwrap().to_string();
        let lhs_parts = lhs.split(",");
        let mut lhs_vec = Vec::new();
        for chem in lhs_parts {
            let mut chem_parts = chem.trim().split(" ");
            let chem_n = chem_parts.next().unwrap().parse::<i64>().unwrap();
            let chem_name = chem_parts.next().unwrap().to_string();
            lhs_vec.push(Chemical {
                n: chem_n,
                name: chem_name,
            });
        }
        reactions.insert(
            Chemical {
                n: rhs_n,
                name: rhs_name,
            },
            lhs_vec,
        );
    }
    reactions
}

fn update_element(subject: &mut HashMap<String, i64>, n: i64, name: &String) {
    let n = n + subject.get(name).unwrap_or(&0);
    //subject.insert(name.to_string(), n);
    if n == 0 {
        subject.remove(name);
    } else {
        subject.insert(name.to_string(), n);
    }
}

fn rewrite_step(
    reactions: &HashMap<Chemical, Vec<Chemical>>,
    strict: bool,
    subject: &mut HashMap<String, i64>,
    stock: &mut HashMap<String, i64>,
) -> bool {
    for (element, n) in subject.clone() {
        if n > 0 {
            if let Some(stock_n) = stock.get(&element) {
                if *stock_n > 0 {
                    let n = std::cmp::min(*stock_n, n);
                    update_element(stock, -n, &element);
                    update_element(subject, -n, &element);
                    //println!("use stock {:?} on {:?}", stock, subject);
                    return true;
                }
            }
        }
    }
    for (element, n) in subject.clone() {
        if n == 0 {
            continue;
        }
        for (lhs, rhs) in reactions {
            if lhs.name == element.clone() {
                let factor = n / lhs.n;
                let remainder = n % lhs.n;
                if remainder == 0 {
                    // rule can be applied
                    //println!("apply {:?} => {:?} on {:?}", lhs, rhs, subject);
                    // remove lhs
                    subject.remove(&lhs.name);
                    // add rhs
                    for Chemical { n, name } in rhs {
                        update_element(subject, factor * n, name);
                    }
                    return true;
                } else {
                    // rule can be applied
                    //println!("apply+1 {:?} => {:?} on {:?}", lhs, rhs, subject);
                    // remove lhs
                    subject.remove(&lhs.name);
                    update_element(stock, lhs.n - remainder, &lhs.name);
                    // add rhs
                    for Chemical { n, name } in rhs {
                        update_element(subject, (factor + 1) * n, name);
                    }
                    return true;
                }
            }
        }
    }
    false
}

fn fuel1(reactions: &HashMap<Chemical, Vec<Chemical>>, goal: i64) -> i64 {
    let mut subject = HashMap::new();
    let mut stock = HashMap::new();
    subject.insert("FUEL".to_string(), goal);
    let mut modified = true;
    while modified {
        modified = rewrite_step(&reactions, true, &mut subject, &mut stock);
        //println!("subject={:?}, stock={:?}", subject, stock);
    }
    //println!("FINAL subject={:?}, stock={:?}", subject, stock);

    subject.get("ORE").unwrap().clone()
}

pub fn part1(input: String) -> i64 {
    let reactions = parse(input);
    fuel1(&reactions, 1)
}

pub fn part2(input: String) -> i64 {
    let NB_ORE = 1000000000000;
    let reactions = parse(input);
    let cost_fuel1 = fuel1(&reactions, 1);
    let mut a = NB_ORE / cost_fuel1;
    let mut b = 3 * a;
    let mut oa = fuel1(&reactions, a);
    let mut ob = fuel1(&reactions, b);
    while b - a > 1 {
        let c = (a + b) / 2;
        let oc = fuel1(&reactions, c);
        if oc > NB_ORE {
            b = c;
            ob = oc;
        } else {
            a = c;
            oa = oc;
        }
    }
    a
}
