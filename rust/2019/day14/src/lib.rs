use std::collections::HashSet;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Chemical {
    n: i64,
    name: String,
}

fn rewrite_step(reactions: &HashSet<(Chemical, Vec<Chemical>)>, subject:Vec<Chemical>) -> Vec<Chemical> {
    let mut result = Vec::new();
    for chem in subject {
        if chem.name == "ORE" {
            result.push(chem);
        } else {
            let mut found = false;
            for (lhs, rhs) in reactions {
                if lhs.name == chem.name {
                    found = true;
                    let mut new_chem = lhs.clone();
                    new_chem.n *= chem.n / lhs.n;
                    if chem.n % lhs.n != 0 {
                        new_chem.n += lhs.n;
                    }
                    result.push(new_chem);
                }
            }
            if !found {
                result.push(chem);
            }
        }
    }
    result

}

pub fn part1(input: String) -> i64 {
    let mut reactions= HashSet::new();
    for line in input.lines() {
        let mut parts = line.split("=>");
        let lhs = parts.next().unwrap();
        let rhs = parts.next().unwrap();
        let mut rhs_parts = rhs.split(" ");
        let rhs_n = rhs_parts.next().unwrap().parse::<i64>().unwrap();
        let rhs_name = rhs_parts.next().unwrap().to_string();
        let mut lhs_parts = lhs.split(",");
        let mut lhs_vec = Vec::new();
        for chem in lhs_parts {
            let mut chem_parts = chem.split(" ");
            let chem_n = chem_parts.next().unwrap().parse::<i64>().unwrap();
            let chem_name = chem_parts.next().unwrap().to_string();
            lhs_vec.push(Chemical { n: chem_n, name: chem_name });
        }
        reactions.insert((Chemical { n: rhs_n, name: rhs_name }, lhs_vec));
    }

    0
}





pub fn part2(input: String) -> i64 {
    0
}
