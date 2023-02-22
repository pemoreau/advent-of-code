use itertools::Itertools;
use utils::parsing::comma_separated_to_numbers;

enum State {
    Running,
    Halted,
    Suspended,
}

struct Machine {
    pc: usize,
    memory: Vec<i64>,
    input: Vec<i64>,
    output: Vec<i64>,
}

impl Machine {
    pub fn new(memory: Vec<i64>, input: Vec<i64>, output: Vec<i64>) -> Self {
        Self {
            pc: 0,
            memory,
            input,
            output,
        }
    }

    pub fn run(&mut self) -> Vec<i64> {
        loop {
            let state = self.step();
            match state {
                State::Running => {}
                State::Halted => break,
                State::Suspended => break,
            }
        }
        self.output.clone()
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

    fn step(&mut self) -> State {
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
                    return State::Suspended;
                }
                self.memory[a] = self.input.remove(0);
                self.pc += 2;
            }
            4 => {
                // output
                let arg1 = self.decode_arg1();
                self.output.push(arg1);
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
                return State::Halted;
            }
            _ => panic!("Unknown opcode"),
        }
        State::Running
    }
}

/*fn decode(instr: i64) -> (usize, usize, usize, usize) {
    let opcode = instr % 100;
    let mode1 = (instr / 100) % 10;
    let mode2 = (instr / 1000) % 10;
    let mode3 = (instr / 10000) % 10;
    (
        opcode as usize,
        mode1 as usize,
        mode2 as usize,
        mode3 as usize,
    )
}

fn decode_args(program: &Vec<i64>, pc: usize, mode1: usize, mode2: usize) -> (i64, i64) {
    let a = program[pc + 1] as usize;
    let arg1 = if mode1 == 0 { program[a] } else { a as i64 };
    let b = program[pc + 2] as usize;
    let arg2 = if mode2 == 0 { program[b] } else { b as i64 };
    (arg1, arg2)
}

pub fn intcode(program: &mut Vec<i64>, input: Vec<i64>, output: &mut Vec<i64>) {
    let mut pc = 0;
    let mut ic = 0;
    loop {
        let (opcode, mode1, mode2, _) = decode(program[pc]);
        match opcode {
            1 => {
                // add
                let (arg1, arg2) = decode_args(program, pc, mode1, mode2);
                let c = program[pc + 3] as usize;
                program[c] = arg1 + arg2;
                pc += 4;
            }
            2 => {
                // mul
                let (arg1, arg2) = decode_args(program, pc, mode1, mode2);
                let c = program[pc + 3] as usize;
                program[c] = arg1 * arg2;
                pc += 4;
            }
            3 => {
                // input
                let a = program[pc + 1] as usize;
                program[a] = input[ic];
                ic += 1;
                pc += 2;
            }
            4 => {
                // output
                let a = program[pc + 1] as usize;
                let arg1 = if mode1 == 0 { program[a] } else { a as i64 };
                output.push(arg1);
                pc += 2;
            }
            5 => {
                // jump-if-true
                let (arg1, arg2) = decode_args(program, pc, mode1, mode2);
                if arg1 != 0 {
                    pc = arg2 as usize;
                } else {
                    pc += 3;
                }
            }
            6 => {
                // jump-if-false
                let (arg1, arg2) = decode_args(program, pc, mode1, mode2);
                if arg1 == 0 {
                    pc = arg2 as usize;
                } else {
                    pc += 3;
                }
            }
            7 => {
                // less than
                let (arg1, arg2) = decode_args(program, pc, mode1, mode2);
                let c = program[pc + 3] as usize;
                if arg1 < arg2 {
                    program[c] = 1;
                } else {
                    program[c] = 0;
                }
                pc += 4;
            }
            8 => {
                // equals
                let (arg1, arg2) = decode_args(program, pc, mode1, mode2);
                let c = program[pc + 3] as usize;
                if arg1 == arg2 {
                    program[c] = 1;
                } else {
                    program[c] = 0;
                }
                pc += 4;
            }
            99 => break,
            _ => panic!("Unknown opcode"),
        }
    }
}*/

fn run_amplifiers(program: &Vec<i64>, phases: Vec<&i64>) -> i64 {
    let mut last_output = 0;
    for phase in phases {
        let code = program.clone();
        let input = vec![*phase, last_output];
        let output = Vec::new();
        let mut amp = Machine::new(code, input, output);
        last_output = amp.run().last().unwrap().clone();
    }
    last_output
}

pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let phase_setting = vec![0, 1, 2, 3, 4];
    let mut max_signal = 0;
    for phase in phase_setting.iter().permutations(5) {
        let signal = run_amplifiers(&code, phase);
        if signal > max_signal {
            max_signal = signal;
        }
    }
    max_signal
}

// fn run_amplifiers2(program: &Vec<i64>, phases: Vec<&i64>) -> i64 {
//     let mut last_output = 0;
//     // init amps
//     let mut amps = Vec::new();
//     for phase in phases {
//         let code = program.clone();
//         let input = vec![*phase];
//         let output = Vec::new();
//         let mut amp = Machine::new(code, input, output);
//         amps.push(amp);
//     }
//     loop {
//         for amp in amps.iter_mut() {
//             amp.input.push(last_output);
//             let output = amp.run();
//             last_output = output.last().unwrap().clone();
//         }
//         if amps.last().unwrap().halted {
//             break;
//         }
//     }
//     for phase in phases {
//         let code = program.clone();
//         let input = vec![*phase, last_output];

//         let mut amp = Machine::new(code, input);
//         let output = amp.run();
//         last_output = output.last().unwrap().clone();
//     }
//     last_output
// }

pub fn part2(input: String) -> i64 {
    let code =
        "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5";
    let phase = vec![9, 8, 7, 6, 5];

    0
}
