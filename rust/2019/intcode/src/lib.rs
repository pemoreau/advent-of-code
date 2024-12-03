#[derive(Debug, PartialEq)]
pub enum State {
    Running,
    Halted,
    Suspended,
}

#[derive(Debug)]
pub struct Machine {
    pc: usize,
    memory: Vec<i64>,
    input: Vec<i64>,
    output: Vec<i64>,
    state: State,
    pub out: bool, // flag set to true immediately after output
    last_output: usize,
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
            last_output: 0,
            state: State::Running,
            out: false,
            relative_base: 0,
        }
    }

    pub fn run(&mut self) {
        loop {
            self.step();
            match self.state {
                State::Running => {}
                State::Halted => break,
                State::Suspended => break,
            }
        }
    }

    pub fn is_halted(&self) -> bool {
        self.state == State::Halted
    }
    pub fn is_suspended(&self) -> bool {
        self.state == State::Suspended
    }

    pub fn is_idle(&self) -> bool {
        self.input.is_empty()
    }

    pub fn run_one_step(&mut self) {
        self.step();
    }

    pub fn get_output(&self) -> Vec<i64> {
        self.output.clone()
    }

    pub fn get_last_output(&mut self) -> Vec<i64> {
        let res = self.output[self.last_output..].to_vec();
        self.last_output = self.output.len();
        res
    }

    pub fn put_input(&mut self, input: i64) {
        self.input.push(input);
    }

    fn get_opcode(&self) -> usize {
        let instr = self.memory[self.pc];
        (instr % 100) as usize
    }

    fn decode_arg(&self, n: usize) -> i64 {
        let instr = self.memory[self.pc];
        let a = self.memory[self.pc + n];

        let mode = match n {
            1 => (instr / 100) % 10,
            2 => (instr / 1000) % 10,
            3 => (instr / 10000) % 10,
            _ => panic!("Invalid argument"),
        };
        let arg = match mode {
            0 => self.memory[a as usize],
            1 => a as i64,
            2 => self.memory[(self.relative_base as i64 + a) as usize],
            _ => panic!("Invalid mode"),
        };
        arg
    }
    fn decode_dest(&self, n: usize) -> i64 {
        let instr = self.memory[self.pc];
        let a = self.memory[self.pc + n];
        let mode = match n {
            1 => (instr / 100) % 10,
            2 => (instr / 1000) % 10,
            3 => (instr / 10000) % 10,
            _ => panic!("Invalid argument"),
        };
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
                let dest = self.decode_dest(3) as usize;
                self.memory[dest] = arg1 + arg2;
                self.pc += 4;
            }
            2 => {
                // mul
                let arg1 = self.decode_arg(1);
                let arg2 = self.decode_arg(2);
                let dest = self.decode_dest(3) as usize;
                self.memory[dest] = arg1 * arg2;
                self.pc += 4;
            }
            3 => {
                // input
                if self.input.is_empty() {
                    self.state = State::Suspended;
                    return;
                }
                let input = self.input.remove(0);
                let dest = self.decode_dest(1) as usize;
                self.memory[dest] = input;
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
                let c = self.decode_dest(3) as usize;
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
                let c = self.decode_dest(3) as usize;
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

#[cfg(test)]
mod tests {
    #[test]
    fn day07_1() {
        use super::*;
        let code = vec![
            3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0,
            0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4,
            20, 1105, 1, 46, 98, 99,
        ];
        let mut machine = Machine::new(code.clone(), vec![7]);
        machine.run();
        assert_eq!(machine.get_last_output().last().unwrap().clone(), 999);

        let mut machine = Machine::new(code.clone(), vec![8]);
        machine.run();
        assert_eq!(machine.get_last_output().last().unwrap().clone(), 1000);

        let mut machine = Machine::new(code, vec![9]);
        machine.run();
        assert_eq!(machine.get_last_output().last().unwrap().clone(), 1001);
    }

    #[test]
    fn day09_1() {
        use super::*;
        let code = vec![
            109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99,
        ];
        let mut machine = Machine::new(code.clone(), vec![]);
        machine.run();
        let output = machine.get_output();
        assert_eq!(output, code);
    }

    #[test]
    fn day09_2() {
        use super::*;
        let code = vec![1102, 34915192, 34915192, 7, 4, 7, 99, 0];
        let mut machine = Machine::new(code, vec![]);
        machine.run();
        let output = machine.get_last_output().last().unwrap().clone();
        assert_eq!(output, 1219070632396864);
    }

    #[test]
    fn day09_3() {
        use super::*;
        let code = vec![104, 1125899906842624, 99];
        let mut machine = Machine::new(code, vec![]);
        machine.run();
        let output = machine.get_last_output().last().unwrap().clone();
        assert_eq!(output, 1125899906842624);
    }
}
