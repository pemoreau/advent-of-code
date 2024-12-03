fn check1(password: String) -> bool {
    let mut last = 0;
    let mut double = false;
    for c in password.chars() {
        let digit = c.to_digit(10).unwrap();
        if digit < last {
            return false;
        }
        if digit == last {
            double = true;
        }
        last = digit;
    }
    double
}

fn check_contain_double(password: String) -> bool {
    let mut array6 = [1; 6];
    let mut last = 0;
    for (i, c) in password.chars().enumerate() {
        let digit = c.to_digit(10).unwrap();
        if digit == last {
            array6[i] = array6[i - 1] + 1;
            array6[i - 1] = 1;
        }
        last = digit;
    }
    array6.contains(&2)
}

pub fn part1(input: String) -> i64 {
    let v = input.split("-").map(|x| x.parse::<i64>().unwrap()).collect::<Vec<i64>>();
    let mut count = 0;
    for i in v[0]..v[1] {
        if check1(i.to_string()) {
            count += 1;
        }
    }
    count
}

pub fn part2(input: String) -> i64 {
    let v = input.split("-").map(|x| x.parse::<i64>().unwrap()).collect::<Vec<i64>>();
    let mut count = 0;
    for i in v[0]..v[1] {
        if check1(i.to_string()) && check_contain_double(i.to_string()) {
            count += 1;
        }
    }
    count
}
