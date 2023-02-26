use intcode::Machine;
use utils::parsing::comma_separated_to_numbers;

pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut machine = Machine::new(code, vec![1]);
    machine.run();
    *machine.get_last_output().last().unwrap()
}

pub fn part2(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut machine = Machine::new(code, vec![2]);
    machine.run();
    *machine.get_last_output().last().unwrap()
}
