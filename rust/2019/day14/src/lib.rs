use std::collections::HashMap;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Chemical {
    n: i64,
    name: String,
}

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Reaction {
    name: String,
    n: i64,
    rhs: Vec<Chemical>,
}

type Reactions = HashMap<String, Reaction>;
type Elements = HashMap<String, i64>;

fn parse(input: String) -> Reactions {
    let mut reactions = HashMap::new();
    for line in input.lines() {
        let mut parts = line.split("=>");
        let lhs = parts.next().unwrap();
        let rhs = parts.next().unwrap().trim();
        let mut rhs_parts = rhs.split(' ');
        let rhs_n = rhs_parts.next().unwrap().parse::<i64>().unwrap();
        let rhs_name = rhs_parts.next().unwrap().to_string();
        let lhs_parts = lhs.split(',');
        let mut lhs_vec = Vec::new();
        for chem in lhs_parts {
            let mut chem_parts = chem.trim().split(' ');
            let chem_n = chem_parts.next().unwrap().parse::<i64>().unwrap();
            let chem_name = chem_parts.next().unwrap().to_string();
            lhs_vec.push(Chemical {
                n: chem_n,
                name: chem_name,
            });
        }
        let reaction = Reaction {
            name: rhs_name.clone(),
            n: rhs_n,
            rhs: lhs_vec,
        };

        reactions.insert(rhs_name, reaction);
    }
    reactions
}

fn update_element(subject: &mut Elements, n: i64, name: &String) {
    let n = n + subject.get(name).unwrap_or(&0);
    //subject.insert(name.to_string(), n);
    if n == 0 {
        subject.remove(name);
    } else {
        subject.insert(name.to_string(), n);
    }
}

fn rewrite_step(reactions: &Reactions, subject: &mut Elements, stock: &mut Elements) -> bool {
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
        if reactions.contains_key(&element) {
            let reaction = reactions.get(&element).unwrap();
            if reaction.name == element.clone() {
                let remainder = n % reaction.n;
                let factor = if remainder == 0 {
                    n / reaction.n
                } else {
                    n / reaction.n + 1
                };
                // remove lhs
                subject.remove(&reaction.name);
                // add rhs
                for Chemical { n, name } in &reaction.rhs {
                    update_element(subject, factor * n, name);
                }
                if remainder != 0 {
                    update_element(stock, reaction.n - remainder, &reaction.name);
                }
                return true;
            }
        }
    }
    false
}

fn fuel1(reactions: &Reactions, goal: i64) -> i64 {
    let mut subject = HashMap::new();
    let mut stock = HashMap::new();
    subject.insert("FUEL".to_string(), goal);
    let mut modified = true;
    while modified {
        modified = rewrite_step(reactions, &mut subject, &mut stock);
        //println!("subject={:?}, stock={:?}", subject, stock);
    }
    //println!("FINAL subject={:?}, stock={:?}", subject, stock);

    *subject.get("ORE").unwrap()
}

pub fn part1(input: String) -> i64 {
    let reactions = parse(input);
    fuel1(&reactions, 1)
}

pub fn part2(input: String) -> i64 {
    let nb_ore = 1000000000000;
    let reactions = parse(input);
    let cost_fuel1 = fuel1(&reactions, 1);
    let mut a = nb_ore / cost_fuel1;
    let mut b = 3 * a;
    while b - a > 1 {
        let c = (a + b) / 2;
        let oc = fuel1(&reactions, c);
        if oc > nb_ore {
            b = c;
        } else {
            a = c;
        }
    }
    a
}
