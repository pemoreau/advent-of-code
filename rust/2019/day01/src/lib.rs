use utils::parsing::lines_to_numbers;

pub fn fuel(mass: i64) -> i64 {
    (mass / 3) - 2
}

pub fn recursive_fuel(mass: i64) -> i64 {
    let fuel = fuel(mass);
    if fuel <= 0 {
        0
    } else {
        fuel + recursive_fuel(fuel)
    }
}

fn sum_fuel(masses: &Vec<i64>, fuel_func: fn(i64) -> i64) -> i64 {
    masses.iter().map(|&mass| fuel_func(mass.into())).sum()
}

pub fn part1(input: String) -> i64 {
    sum_fuel(&lines_to_numbers(input), fuel)
}

pub fn part2(input: String) -> i64 {
    sum_fuel(&lines_to_numbers(input), recursive_fuel)
}
