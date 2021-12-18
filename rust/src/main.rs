use advent_of_code::{get_day, noop};
use std::env;
use std::fs;
use std::io;
use std::time::{Duration, Instant};

fn fmt_dur(dur: Duration) -> String {
    format!("{}ms ", dur.as_secs_f64() * 1000.0)
}

fn main() {
    // Get day string
    let args: Vec<String> = env::args().collect();
    let mut day = String::new();

    if args.len() >= 2 {
        day = args[1].clone();
    } else {
        println!("Enter day: ");
        io::stdin()
            .read_line(&mut day)
            .expect("Failed to read line");
    }

    // Parse day as number
    day = day.trim().to_string();
    let day_num: u32 = match day.parse() {
        Ok(num) => num,
        Err(_) => {
            println!("Invalid day number: {}", day);
            return;
        }
    };

    // Read input file
    let cwd = env::current_dir().unwrap();
    let filename = cwd.join("inputs").join(format!("day{:02}.txt", day_num));
    println!("Reading {}", filename.display());
    let input = fs::read_to_string(filename).expect("Error while reading");

    // Get corresponding function
    let to_run = get_day(day_num);

    // Time it
    run(to_run.0, input.clone());
    run(to_run.1, input.clone());
}
pub type DayFn = fn(String) -> i64;

fn run(to_run: DayFn, input: String) {
    if to_run != noop {
        let start = Instant::now();
        let result = to_run(input);
        let duration = start.elapsed();
        println!("Result = {} in {}", result, fmt_dur(duration));
    }
}
