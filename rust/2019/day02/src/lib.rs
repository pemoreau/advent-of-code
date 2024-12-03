use utils::parsing::comma_separated_to_numbers;

pub fn intcode(input: &mut Vec<i64>) {
    let mut i = 0;
    loop {
        match input[i] {
            1 => {
                let a = input[i + 1] as usize;
                let b = input[i + 2] as usize;
                let c = input[i + 3] as usize;
                input[c] = input[a] + input[b];
                i += 4;
            }
            2 => {
                let a = input[i + 1] as usize;
                let b = input[i + 2] as usize;
                let c = input[i + 3] as usize;
                input[c] = input[a] * input[b];
                i += 4;
            }
            99 => break,
            _ => panic!("Unknown opcode"),
        }
    }
}

pub fn part1(input: String) -> i64 {
    let mut code = comma_separated_to_numbers(input);
    code[1] = 12;
    code[2] = 2;
    intcode(&mut code);
    code[0]
}

pub fn part2(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    for noun in 0..100 {
        for verb in 0..100 {
            let mut memory = code.clone();
            memory[1] = noun;
            memory[2] = verb;
            intcode(&mut memory);
            if memory[0] == 19690720 {
                return 100 * noun + verb;
            }
        }
    }
    panic!("No solution found");
}

