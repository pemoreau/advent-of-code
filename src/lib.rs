// Days
pub mod day01;

pub fn noop(_inp: String) -> i32 {
    return 0;
}

pub type DayFn = fn(String) -> i32;

pub fn get_day(day: u32) -> (DayFn, DayFn) {
    match day {
        1 => (day01::part1, day01::part2),
        _ => {
            println!("Unknown day: {}", day);
            (noop, noop)
        }
    }
}
