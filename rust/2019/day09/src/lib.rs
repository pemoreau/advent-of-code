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
    relative_base: usize,
}

impl Machine {
    pub fn new(code: Vec<i64>, input: Vec<i64>) -> Self {
        let mut memory = code.clone();
        memory.resize(32000, 0);
        Self {
            pc: 0,
            memory,
            input,
            output: vec![],
            state: State::Running,
            out: false,
            relative_base: 0,
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

    fn decode_arg(&self, n: usize) -> i64 {
        let instr = self.memory[self.pc];
        // println!("instr: {}", instr);
        let mode = (instr / (10 as i64).pow((n + 1) as u32)) % 10;
        let a = self.memory[self.pc + n];
        // println!("mode: {}, a: {}, base: {}", mode, a, self.relative_base);
        let arg = match mode {
            0 => self.memory[a as usize],
            1 => a as i64,
            2 => self.memory[(self.relative_base as i64 + a) as usize],
            _ => panic!("Invalid mode"),
        };
        arg
    }
    fn decode_litteral_arg(&self, n: usize) -> i64 {
        let instr = self.memory[self.pc];
        let mode = (instr / (10 as i64).pow((n + 1) as u32)) % 10;
        let a = self.memory[self.pc + n];
        // println!("mode: {}, a: {}, base: {}", mode, a, self.relative_base);
        let arg = match mode {
            0 => a as i64,
            1 => a as i64,
            2 => self.relative_base as i64 + a,
            _ => panic!("Invalid mode"),
        };
        arg
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
                let arg1 = self.decode_arg(1);
                let arg2 = self.decode_arg(2);
                // let c = self.memory[self.pc + 3] as usize;
                let c = self.decode_litteral_arg(3) as usize;
                self.memory[c] = arg1 + arg2;
                self.pc += 4;
            }
            2 => {
                // mul
                let arg1 = self.decode_arg(1);
                let arg2 = self.decode_arg(2);
                // let c = self.memory[self.pc + 3] as usize;
                let c = self.decode_litteral_arg(3) as usize;
                self.memory[c] = arg1 * arg2;
                self.pc += 4;
            }
            3 => {
                // input
                if self.input.is_empty() {
                    self.state = State::Suspended;
                    return;
                }
                let input_value = self.input.remove(0);
                // let a = self.memory[self.pc + 1] as usize;
                let a = self.decode_litteral_arg(1);
                self.memory[a as usize] = input_value;
                self.pc += 2;
            }
            4 => {
                // output
                let arg1 = self.decode_arg(1);
                self.output.push(arg1);
                self.out = true;
                self.pc += 2;
            }
            5 => {
                // jump-if-true
                let arg1 = self.decode_arg(1);
                let arg2 = self.decode_arg(2);
                if arg1 != 0 {
                    self.pc = arg2 as usize;
                } else {
                    self.pc += 3;
                }
            }
            6 => {
                // jump-if-false
                let arg1 = self.decode_arg(1);
                let arg2 = self.decode_arg(2);
                if arg1 == 0 {
                    self.pc = arg2 as usize;
                } else {
                    self.pc += 3;
                }
            }
            7 => {
                // less than
                let arg1 = self.decode_arg(1);
                let arg2 = self.decode_arg(2);
                // let c = self.memory[self.pc + 3] as usize;
                let c = self.decode_litteral_arg(3) as usize;
                if arg1 < arg2 {
                    self.memory[c] = 1;
                } else {
                    self.memory[c] = 0;
                }
                self.pc += 4;
            }
            8 => {
                // equals
                let arg1 = self.decode_arg(1);
                let arg2 = self.decode_arg(2);
                // let c = self.memory[self.pc + 3] as usize;
                let c = self.decode_litteral_arg(3) as usize;
                if arg1 == arg2 {
                    self.memory[c] = 1 as i64;
                } else {
                    self.memory[c] = 0;
                }
                self.pc += 4;
            }
            9 => {
                // relative base offset
                let arg1 = self.decode_arg(1);
                self.relative_base = (self.relative_base as i64 + arg1) as usize;
                self.pc += 2;
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

pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    // let code = vec![
    //     // 109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99,
    //     // 1102, 34915192, 34915192, 7, 4, 7, 99, 0,
    //     104,
    //     1125899906842624,
    //     99,
    // ];
    let mut machine = Machine::new(code, vec![1]);
    machine.run();
    machine.get_last_output()
}

pub fn part2(input: String) -> i64 {
    0
}
