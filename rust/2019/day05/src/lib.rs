use utils::parsing::comma_separated_to_numbers;

fn decode(instr: i64) -> (i64, i64, i64, i64) {
    let opcode = instr % 100;
    let mode1 = (instr / 100) % 10;
    let mode2 = (instr / 1000) % 10;
    let mode3 = (instr / 10000) % 10;
    (opcode, mode1, mode2, mode3)
}

pub fn intcode(input: &mut Vec<i64>) {
    let mut pc = 0;
    loop {
        let (opcode, mode1, mode2, mode3) = decode(input[pc]);
        // println!("{} : {} {} {} {}", input[pc], opcode, mode1, mode2, mode3);
        match opcode {
            1 => {
                let a = input[pc + 1] as usize;
                let arg1 = if mode1 == 0 { input[a] } else { a as i64 };
                let b = input[pc + 2] as usize;
                let arg2 = if mode2 == 0 { input[b] } else { b as i64 };
                let c = input[pc + 3] as usize;
                input[c] = arg1 + arg2;
                pc += 4;
            }
            2 => {
                let a = input[pc + 1] as usize;
                let arg1 = if mode1 == 0 { input[a] } else { a as i64 };
                let b = input[pc + 2] as usize;
                let arg2 = if mode2 == 0 { input[b] } else { b as i64 };
                let c = input[pc + 3] as usize;
                input[c] = arg1 * arg2;
                pc += 4;
            }
            3 => {
                let a = input[pc + 1] as usize;
                input[a] = 1; // 1 as input
                pc += 2;
            }
            4 => {
                let a = input[pc + 1] as usize;
                let arg1 = if mode1 == 0 { input[a] } else { a as i64 };
                println!("output: {}", arg1);
                pc += 2;
            }
            99 => break,
            _ => panic!("Unknown opcode"),
        }
    }
}

pub fn part1(input: String) -> i64 {
    let mut code = comma_separated_to_numbers(input);
    intcode(&mut code);
    code[0]
}

pub fn part2(input: String) -> i64 {
    0
}
