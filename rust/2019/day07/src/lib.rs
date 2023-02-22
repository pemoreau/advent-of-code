use itertools::Itertools;
use utils::parsing::comma_separated_to_numbers;

#[derive(Debug, PartialEq)]
enum State {
    Running,
    Halted,
    Suspended,
}

#[derive(Debug)]
struct Machine {
    pc: usize,
    memory: Vec<i64>,
    input: Vec<i64>,
    output: Vec<i64>,
    state: State,
    out: bool, // flag set to true immediately after output
}

impl Machine {
    pub fn new(memory: Vec<i64>, input: Vec<i64>, output: Vec<i64>) -> Self {
        Self {
            pc: 0,
            memory,
            input,
            output,
            state: State::Running,
            out: false,
        }
    }

    pub fn run(&mut self) -> Vec<i64> {
        loop {
            self.step();
            match self.state {
                State::Running => {}
                State::Halted => break,
                State::Suspended => break,
            }
        }
        self.output.clone()
    }

    fn get_last_output(&self) -> i64 {
        self.output.last().unwrap().clone()
    }

    fn put_input(&mut self, input: i64) {
        self.input.push(input);
    }

    fn get_opcode(&self) -> usize {
        let instr = self.memory[self.pc];
        (instr % 100) as usize
    }

    fn get_mode(&self) -> (usize, usize, usize) {
        let instr = self.memory[self.pc];
        let mode1 = (instr / 100) % 10;
        let mode2 = (instr / 1000) % 10;
        let mode3 = (instr / 10000) % 10;
        (mode1 as usize, mode2 as usize, mode3 as usize)
    }

    fn decode_arg1(&self) -> i64 {
        let (mode1, _, _) = self.get_mode();
        let a = self.memory[self.pc + 1] as usize;
        let arg1 = if mode1 == 0 { self.memory[a] } else { a as i64 };
        arg1
    }

    fn decode_args(&self) -> (i64, i64) {
        let (mode1, mode2, _) = self.get_mode();
        let a = self.memory[self.pc + 1] as usize;
        let arg1 = if mode1 == 0 { self.memory[a] } else { a as i64 };
        let b = self.memory[self.pc + 2] as usize;
        let arg2 = if mode2 == 0 { self.memory[b] } else { b as i64 };
        (arg1, arg2)
    }

    fn step(&mut self) {
        self.out = false;
        if self.state == State::Halted {
            return;
        }

        let opcode = self.get_opcode();
        match opcode {
            1 => {
                // add
                let (arg1, arg2) = self.decode_args();
                let c = self.memory[self.pc + 3] as usize;
                self.memory[c] = arg1 + arg2;
                self.pc += 4;
            }
            2 => {
                // mul
                let (arg1, arg2) = self.decode_args();
                let c = self.memory[self.pc + 3] as usize;
                self.memory[c] = arg1 * arg2;
                self.pc += 4;
            }
            3 => {
                // input
                let a = self.memory[self.pc + 1] as usize;
                if self.input.is_empty() {
                    self.state = State::Suspended;
                    return;
                }
                self.memory[a] = self.input.remove(0);
                self.pc += 2;
            }
            4 => {
                // output
                let arg1 = self.decode_arg1();
                self.output.push(arg1);
                self.out = true;
                self.pc += 2;
            }
            5 => {
                // jump-if-true
                let (arg1, arg2) = self.decode_args();
                if arg1 != 0 {
                    self.pc = arg2 as usize;
                } else {
                    self.pc += 3;
                }
            }
            6 => {
                // jump-if-false
                let (arg1, arg2) = self.decode_args();
                if arg1 == 0 {
                    self.pc = arg2 as usize;
                } else {
                    self.pc += 3;
                }
            }
            7 => {
                // less than
                let (arg1, arg2) = self.decode_args();
                let c = self.memory[self.pc + 3] as usize;
                if arg1 < arg2 {
                    self.memory[c] = 1;
                } else {
                    self.memory[c] = 0;
                }
                self.pc += 4;
            }
            8 => {
                // equals
                let (arg1, arg2) = self.decode_args();
                let c = self.memory[self.pc + 3] as usize;
                if arg1 == arg2 {
                    self.memory[c] = 1;
                } else {
                    self.memory[c] = 0;
                }
                self.pc += 4;
            }
            99 => {
                // halt
                self.state = State::Halted;
                return;
            }
            _ => panic!("Unknown opcode"),
        }
        self.state = State::Running;
    }
}

fn run_amplifiers(program: &Vec<i64>, phases: Vec<&i64>) -> i64 {
    let mut last_output = 0;
    for phase in phases {
        let code = program.clone();
        let input = vec![*phase, last_output];
        let output = Vec::new();
        let mut amp = Machine::new(code, input, output);
        amp.run();
        last_output = amp.get_last_output()
    }
    last_output
}

fn run_amplifiers2(program: &Vec<i64>, phases: Vec<&i64>) -> i64 {
    let n = phases.len();

    // init amps
    let mut amps = Vec::new();
    for phase in phases {
        let code = program.clone();
        let amp = Machine::new(code, vec![*phase], vec![]);
        amps.push(amp);
    }
    amps[0].put_input(0);

    loop {
        if amps[n - 1].state == State::Halted {
            return amps[n - 1].get_last_output();
        }

        for i in 0..n {
            amps[i].step();
            if amps[i].out {
                let output = amps[i].get_last_output();
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
