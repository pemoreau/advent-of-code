use intcode::Machine;
use itertools::Itertools;
use utils::parsing::comma_separated_to_numbers;

fn run_amplifiers(program: &Vec<i64>, phases: Vec<&i64>) -> i64 {
    let mut last_output = 0;
    for phase in phases {
        let code = program.clone();
        let input = vec![*phase, last_output];
        let mut amp = Machine::new(code, input);
        amp.run();
        last_output = *amp.get_last_output().last().unwrap();
    }
    last_output
}

fn run_amplifiers2(program: &Vec<i64>, phases: Vec<&i64>) -> i64 {
    let n = phases.len();

    // init amps
    let mut amps = Vec::new();
    for phase in phases {
        let code = program.clone();
        let amp = Machine::new(code, vec![*phase]);
        amps.push(amp);
    }
    amps[0].put_input(0);

    loop {
        if amps[n - 1].is_halted() {
            return *amps[n - 1].get_output().last().unwrap();
        }

        for i in 0..n {
            amps[i].run_one_step();
            if amps[i].out {
                let output = *amps[i].get_last_output().last().unwrap();
                amps[(i + 1) % n].put_input(output);
            }
        }
    }
}

fn search_max_signal(
    input: String,
    phase_setting: Vec<i64>,
    run_func: &dyn Fn(&Vec<i64>, Vec<&i64>) -> i64,
) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut max_signal = 0;
    for phase in phase_setting.iter().permutations(5) {
        let signal = run_func(&code, phase);
        if signal > max_signal {
            max_signal = signal;
        }
    }
    max_signal
}

pub fn part1(input: String) -> i64 {
    search_max_signal(input, vec![0, 1, 2, 3, 4], &run_amplifiers)
}

pub fn part2(input: String) -> i64 {
    search_max_signal(input, vec![5, 6, 7, 8, 9], &run_amplifiers2)
}
