use intcode::Machine;
use utils::parsing::comma_separated_to_numbers;

struct Computer {
    address: i64,
    machine: Machine,
    queue: Vec<Message>,
    waiting_for_input: bool,
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
            waiting_for_input: false,
        }
    }

    fn step(&mut self) -> Vec<Message> {
        // println!("step machine {:}", self.address);
        self.machine.run();
        self.waiting_for_input = true;

        let mut res = vec![];
        let output = self.machine.get_last_output();
        for output in output.chunks(3) {
            let address = output[0];
            let x = output[1];
            let y = output[2];
            println!(
                "computer {:?} send message: {:?}",
                self.address,
                Message { address, x, y }
            );
            res.push(Message { address, x, y });
            self.waiting_for_input = false;
        }

        if self.machine.is_suspended() {
            if let Some(message) = self.queue.pop() {
                println!("computer {:?} read message: {:?}", self.address, message);
                self.machine.put_input(message.x);
                self.machine.put_input(message.y);
                self.waiting_for_input = false;
            } else {
                // println!("no message");
                self.machine.put_input(-1);
                // println!("computer {:?} is_idle: {:?}", self.address, self.is_idle());
            }
        }
        res
    }

    fn send(&mut self, message: Message) {
        self.queue.push(message);
    }

    fn is_idle(&self) -> bool {
        self.queue.is_empty() && (self.machine.is_idle() || self.waiting_for_input)
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

// 19641
pub fn part2(input: String) -> i64 {
    let code = comma_separated_to_numbers(input);
    let mut computers: Vec<Computer> = (0..=49)
        .map(|i| Computer::new(&code, i))
        .collect::<Vec<_>>();
    let mut nat_message: Option<Message> = None;
    let mut last_nat_y: Option<i64> = None;

    loop {
        for i in 0..computers.len() {
            let messages = computers[i].step();
            // println!("messages: {:?}", messages);
            for message in messages {
                if message.address == 255 {
                    nat_message = Some(message);
                    println!("SAVE nat_message: {:?}", nat_message);
                } else {
                    println!("SEND {:?} message: {:?}", message.address, message);
                    computers[message.address as usize].send(message);
                }
            }
        }
        let all_idle = computers.iter().all(|c| c.is_idle());
        if all_idle {
            println!("all_idle: {:?} nat_message: {:?}", all_idle, nat_message);
            if let Some(message) = nat_message.take() {
                if let Some(last_nat_y) = last_nat_y {
                    if last_nat_y == message.y {
                        return message.y;
                    }
                }
                last_nat_y = Some(message.y);
                println!("SEND 0 nat_message: {:?}", message);
                computers[0].send(message);
            }
        }
    }
}
