pub fn part1(input: String) -> i64 {
    let width = 25;
    let height = 6;
    let mut layers = Vec::new();
    let mut layer = Vec::new();

    let mut index = 0;
    for c in input.trim().chars() {
        if index == width * height {
            layers.push(layer);
            layer = Vec::new();
            index = 0;
        }
        let digit = c.to_digit(10).unwrap();
        layer.push(digit);
        index += 1;
    }
    layers.push(layer);

    let layer = layers
        .iter()
        .min_by_key(|layer| layer.iter().filter(|&&x| x == 0).count())
        .unwrap();

    let ones = layer.iter().filter(|&&x| x == 1).count();
    let twos = layer.iter().filter(|&&x| x == 2).count();
    (ones * twos) as i64
}

pub fn part2(input: String) -> i64 {
    let width = 25;
    let height = 6;
    let mut layers = Vec::new();
    let mut layer = Vec::new();

    let mut index = 0;
    for c in input.trim().chars() {
        if index == width * height {
            layers.push(layer);
            layer = Vec::new();
            index = 0;
        }
        let digit = c.to_digit(10).unwrap();
        layer.push(digit);
        index += 1;
    }
    layers.push(layer);

    for i in 0..width * height {
        let mut pixel = '.';
        for layer in &layers {
            if layer[i] != 2 {
                pixel = match layer[i] {
                    0 => ' ',
                    1 => '█',
                    _ => panic!("invalid pixel"),
                };
                break;
            }
        }
        print!("{}", pixel);
        if (i + 1) % width == 0 {
            println!();
        }
    }

    0
}
