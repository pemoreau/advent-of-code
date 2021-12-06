// Days
pub mod day01;
pub mod day02;
pub mod day03;
pub mod day04;
pub mod day05;
pub mod day06;

pub fn noop(_inp: String) -> i32 {
    return 0;
}

pub type DayFn = fn(String) -> i32;

pub fn get_day(day: u32) -> (DayFn, DayFn) {
    match day {
        1 => (day01::part1, day01::part2),
        2 => (day02::part1, day02::part2),
        3 => (day03::part1, day03::part2),
        4 => (day04::part1, day04::part2),
        5 => (day05::part1, day05::part2),
        6 => (day06::part1, day06::part2),
        _ => {
            println!("Unknown day: {}", day);
            (noop, noop)
        }
    }
}
