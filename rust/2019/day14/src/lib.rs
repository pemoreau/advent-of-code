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

fn update_element(subject: &mut HashMap<String, i64>, n: i64, name: &str) {
    let n = n + subject.get(name).unwrap_or(&0);
    subject.insert(name.to_string(), n);
}

fn rewrite_step(
    reactions: &HashMap<Chemical, Vec<Chemical>>,
    subject: &mut HashMap<String, i64>,
) -> (Map<String, i64>, bool) {
    let mut result = subject.clone();
    for (element, n) in subject.clone() {
        for (lhs, rhs) in reactions {
            if lhs.name == element.clone() && n > 0 && n >= lhs.n {
                // rule can be applied
                let factor = n / lhs.n;
                let remainder = n % lhs.n;
                println!(
                    "apply {}*{:?}=>{:?} on {:?}",
                    factor,
                    lhs,
                    rhs,
                    subject.clone()
                );
                if remainder != 0 {
                    update_element(&mut result, remainder, element.as_str())
                }
                // remove lhs
                update_element(&mut result, -lhs.n * factor, lhs.name.as_str());
                for Chemical { n, name } in rhs {
                    for _ in 0..factor {
                        update_element(&mut result, *n, name);
                    }
                }
                return (result, true);
            } else {
                // rule cannot be applied
            }
        }
    }
    (result, false)
}

pub fn part1(input: String) -> i64 {
    let reactions = parse(input);
    let mut subject = HashMap::new();
    subject.insert("FUEL".to_string(), 1);
    let mut modified = true;
    while modified {
        let res = rewrite_step(&reactions, subject.clone());
        modified = res.1;
        subject = res.0;
        println!("{:?}", subject);
    }

    0
}

pub fn part2(input: String) -> i64 {
    0
}
