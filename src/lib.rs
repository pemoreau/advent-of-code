// Days
pub mod day01;

pub fn noop(_inp: String) {}

pub type DayFn = fn(String);

pub fn get_day(day: u32) -> (DayFn, DayFn) {
    match day {
        1 => (day01::part1, day01::part2),
        _ => {
            println!("Unknown day: {}", day);
            (noop, noop)
        }
    }
}
