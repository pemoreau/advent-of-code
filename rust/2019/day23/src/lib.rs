use intcode::Machine;
use utils::parsing::comma_separated_to_numbers;

struct Computer {
    address: i64,
    machine: Machine,
    queue: Vec<Message>,
}

#[derive(Debug)]
struct Message {
    address: i64,
    x: i64,
    y: i64,
}

impl Computer {
    fn new(code: &Vec<i64>, address: i64) -> Self {
        Self {
            address,
            machine: Machine::new(code.clone(), vec![address]),
            queue: vec![],
        }
    }

    fn step(&mut self) -> Vec<Message> {
        println!("step machine {:}", self.address);
        self.machine.run();

        let mut res = vec![];
        let output = self.machine.get_last_output();
        for output in output.chunks(3) {
            let address = output[0];
            let x = output[1];
            let y = output[2];
            res.push(Message { address, x, y });
        }

        if self.machine.is_suspended() {
            if let Some(message) = self.queue.pop() {
                println!("read message: {:?}", message);
                self.machine.put_input(message.x);
                self.machine.put_input(message.y);
            } else {
                // println!("no message");
                self.machine.put_input(-1);
            }
        }
        res
    }

    fn send(&mut self, message: Message) {
        self.machine.put_input(message.x);
        self.machine.put_input(message.y);
    }
}

pub fn part1(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut computers: Vec<Computer> = (0..=49)
        .map(|i| Computer::new(&code, i))
        .collect::<Vec<_>>();
    loop {
        for i in 0..computers.len() {
            let messages = computers[i].step();
            for message in messages {
                if message.address == 255 {
                    return message.y;
                } else {
                    computers[message.address as usize].send(message);
                }
            }
        }
    }
}

pub fn part2(input: String) -> i64 {
    0
}
