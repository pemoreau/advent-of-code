pub fn part1(input: String) -> i64 {
    let mut letters = [false; 27];
    let mut scores: Vec<u32> = Vec::new();

    input.lines().for_each(|line| {
        if line.is_empty() {
            scores.push(letters.iter().filter(|&letter| *letter).count() as u32);
            letters = [false; 27];
        } else {
            line.chars()
                .for_each(|c| letters[1 + c as usize - 'a' as usize] = true);
        }
    });
    scores.push(letters.iter().filter(|&letter| *letter).count() as u32); // for last entry
    scores.iter().sum::<u32>().try_into().unwrap()
}

pub fn part2(input: String) -> i64 {
    let mut letters = [0u32; 27];
    let mut scores: Vec<u32> = Vec::new();
    let mut nb_lines = 0;
    input.lines().for_each(|line| {
        if line.is_empty() {
            scores.push(letters.iter().filter(|&letter| *letter == nb_lines).count() as u32);
            letters = [0; 27];
            nb_lines = 0;
        } else {
            line.chars()
                .for_each(|c| letters[1 + c as usize - 'a' as usize] += 1);
            nb_lines += 1;
        }
    });
    scores.push(letters.iter().filter(|&letter| *letter == nb_lines).count() as u32);
    scores.iter().sum::<u32>().try_into().unwrap()
}
