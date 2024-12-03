struct State {
    dir: usize,
    x: i32,
    y: i32,
    waypoint: Waypoint,
}
struct Waypoint {
    x: i32,
    y: i32,
}

impl State {
    fn new() -> State {
        State {
            dir: 0,
            x: 0,
            y: 0,
            waypoint: Waypoint { x: 10, y: 1 },
        }
    }
    fn manhattan_distance(&self) -> i64 {
        (self.x.abs() + self.y.abs()) as i64
    }

    fn step1(&mut self, command: char, value: usize) {
        match command {
            'E' => self.x += value as i32,
            'N' => self.y += value as i32,
            'W' => self.x -= value as i32,
            'S' => self.y -= value as i32,
            'F' => match self.dir {
                0 => self.x += value as i32,
                90 => self.y += value as i32,
                180 => self.x -= value as i32,
                270 => self.y -= value as i32,
                _ => panic!("Invalid direction"),
            },
            'L' => self.dir = (self.dir + value) % 360,
            'R' => self.dir = (self.dir + 360 - value) % 360,
            _ => panic!("Invalid command"),
        }
    }

    fn step2(&mut self, command: char, value: usize) {
        match command {
            'F' => {
                self.x += value as i32 * self.waypoint.x;
                self.y += value as i32 * self.waypoint.y
            }
            _ => self.waypoint.step(command, value),
        }
    }
}

impl Waypoint {
    fn step(&mut self, command: char, value: usize) {
        match command {
            'E' => self.x += value as i32,
            'N' => self.y += value as i32,
            'W' => self.x -= value as i32,
            'S' => self.y -= value as i32,
            'L' => self.rotation(-(value as i32)),
            'R' => self.rotation(value as i32),
            _ => panic!("Invalid command"),
        }
    }

    fn rotation(&mut self, degrees: i32) {
        let rotation = degrees / 90;
        let direction = rotation.signum();

        for _ in 0..rotation.abs() {
            let temp = self.y;
            self.y = -self.x * direction;
            self.x = temp * direction;
        }
    }
}

pub fn part1(input: String) -> i64 {
    let mut state = State::new();

    input.lines().for_each(|line| {
        let mut iter = line.chars();
        let command = iter.next().unwrap();
        let value = iter.collect::<String>().parse::<usize>().unwrap();
        state.step1(command, value);
    });

    state.manhattan_distance()
}

pub fn part2(input: String) -> i64 {
    let mut state = State::new();

    input.lines().for_each(|line| {
        let mut iter = line.chars();
        let command = iter.next().unwrap();
        let value = iter.collect::<String>().parse::<usize>().unwrap();
        state.step2(command, value);
    });

    state.manhattan_distance()
}
