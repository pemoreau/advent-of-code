use intcode::Machine;
use utils::parsing::comma_separated_to_numbers;
const PROGRAM1: &str = r"NOT B J
NOT C T
OR T J
AND D J
NOT A T
OR T J
WALK
";
const PROGRAM2: &str = r"NOT B J
NOT C T
OR T J
AND D J
AND H J
NOT A T
OR T J
RUN
";

pub fn solve(input: String, program: &str) -> i64 {
    let code = comma_separated_to_numbers(input);
    let input = program.chars().map(|c| c as i64).collect::<Vec<_>>();
    let mut machine = Machine::new(code.clone(), input);
    machine.run();
    let output = machine.get_output();
    for c in output {
        if c < 256 {
            print!("{}", c as u8 as char);
        } else {
            return c;
        }
    }
    0
}

pub fn part1(input: String) -> i64 {
    solve(input, PROGRAM1)
}

pub fn part2(input: String) -> i64 {
    solve(input, PROGRAM2)
}
