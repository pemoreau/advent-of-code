use utils::parsing::comma_separated_to_numbers;

fn decode(instr: i64) -> (i64, i64, i64, i64) {
    let opcode = instr % 100;
    let mode1 = (instr / 100) % 10;
    let mode2 = (instr / 1000) % 10;
    let mode3 = (instr / 10000) % 10;
    (opcode, mode1, mode2, mode3)
}

pub fn intcode(program: &mut Vec<i64>, input: i64) {
    let mut pc = 0;
    loop {
        let (opcode, mode1, mode2, mode3) = decode(program[pc]);
        // println!("{} : {} {} {} {}", input[pc], opcode, mode1, mode2, mode3);
        match opcode {
            1 => {
                // add
                let a = program[pc + 1] as usize;
                let arg1 = if mode1 == 0 { program[a] } else { a as i64 };
                let b = program[pc + 2] as usize;
                let arg2 = if mode2 == 0 { program[b] } else { b as i64 };
                let c = program[pc + 3] as usize;
                program[c] = arg1 + arg2;
                pc += 4;
            }
            2 => {
                // mul
                let a = program[pc + 1] as usize;
                let arg1 = if mode1 == 0 { program[a] } else { a as i64 };
                let b = program[pc + 2] as usize;
                let arg2 = if mode2 == 0 { program[b] } else { b as i64 };
                let c = program[pc + 3] as usize;
                program[c] = arg1 * arg2;
                pc += 4;
            }
            3 => {
                // input
                let a = program[pc + 1] as usize;
                program[a] = input;
                pc += 2;
            }
            4 => {
                // output
                let a = program[pc + 1] as usize;
                let arg1 = if mode1 == 0 { program[a] } else { a as i64 };
                println!("output: {}", arg1);
                pc += 2;
            }
            5 => {
                // jump-if-true
                let a = program[pc + 1] as usize;
                let arg1 = if mode1 == 0 { program[a] } else { a as i64 };
                let b = program[pc + 2] as usize;
                let arg2 = if mode2 == 0 { program[b] } else { b as i64 };
                if arg1 != 0 {
                    pc = arg2 as usize;
                } else {
                    pc += 3;
                }
            }
            6 => {
                // jump-if-false
                let a = program[pc + 1] as usize;
                let arg1 = if mode1 == 0 { program[a] } else { a as i64 };
                let b = program[pc + 2] as usize;
                let arg2 = if mode2 == 0 { program[b] } else { b as i64 };
                if arg1 == 0 {
                    pc = arg2 as usize;
                } else {
                    pc += 3;
                }
            }
            7 => {
                // less than
                let a = program[pc + 1] as usize;
                let arg1 = if mode1 == 0 { program[a] } else { a as i64 };
                let b = program[pc + 2] as usize;
                let arg2 = if mode2 == 0 { program[b] } else { b as i64 };
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
                let a = program[pc + 1] as usize;
                let arg1 = if mode1 == 0 { program[a] } else { a as i64 };
                let b = program[pc + 2] as usize;
                let arg2 = if mode2 == 0 { program[b] } else { b as i64 };
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
}

pub fn part1(input: String) -> i64 {
    let mut code = comma_separated_to_numbers(input);
    intcode(&mut code, 1);
    0
}

pub fn part2(input: String) -> i64 {
    let mut code = comma_separated_to_numbers(input);
    intcode(&mut code, 5);
    0
}
