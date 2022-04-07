pub fn part1(input: String) -> i64 {
    let mut lines = input.lines();
    let start = lines.next().unwrap().parse::<i64>().unwrap();
    let numbers = lines
        .next()
        .unwrap()
        .split(",")
        .filter_map(|x| x.parse::<i64>().ok())
        .collect::<Vec<i64>>();

    let (line, wait) = numbers
        .iter()
        .map(|x| (x, x - (start % x)))
        .min_by_key(|pair| pair.1)
        .unwrap();
    line * wait
}

pub fn part2(input: String) -> i64 {
    let numbers = input
        .lines()
        .nth(1)
        .unwrap()
        .split(",")
        .enumerate()
        .map(|(i, x)| (i, x.parse::<i64>()))
        .filter(|(_, x)| x.is_ok())
        .map(|(i, x)| (i, x.unwrap()))
        .collect::<Vec<(usize, i64)>>();
    let u = numbers
        .iter()
        .map(|&(i, x)| x - i as i64)
        .collect::<Vec<i64>>();
    let m = numbers.iter().map(|&(_, x)| x).collect::<Vec<i64>>();
    let x = chinese_remainder(&u[..], &m[..]).unwrap();
    x
}

// from https://rosettacode.org/wiki/Chinese_remainder_theorem#Rust
fn egcd(a: i64, b: i64) -> (i64, i64, i64) {
    if a == 0 {
        (b, 0, 1)
    } else {
        let (g, x, y) = egcd(b % a, a);
        (g, y - (b / a) * x, x)
    }
}

fn mod_inv(x: i64, n: i64) -> Option<i64> {
    let (g, x, _) = egcd(x, n);
    if g == 1 {
        Some((x % n + n) % n)
    } else {
        None
    }
}

fn chinese_remainder(residues: &[i64], modulii: &[i64]) -> Option<i64> {
    let prod = modulii.iter().product::<i64>();

    let mut sum = 0;

    for (&residue, &modulus) in residues.iter().zip(modulii) {
        let p = prod / modulus;
        sum += residue * mod_inv(p, modulus)? * p
    }

    Some(sum % prod)
}
