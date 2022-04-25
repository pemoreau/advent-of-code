pub fn part1(input: String) -> i64 {
    input
        .lines()
        .fold(0, |acc, line| acc + parse_line1(line).eval()) as i64
}

pub fn part2(input: String) -> i64 {
    input
        .lines()
        .fold(0, |acc, line| acc + parse_line2(line).eval()) as i64
}

#[derive(Clone, Debug)]
pub enum Exp {
    Val(u64),
    Add(Box<Exp>, Box<Exp>),
    Mul(Box<Exp>, Box<Exp>),
}

impl Exp {
    fn eval(&self) -> u64 {
        match self {
            Exp::Val(val) => *val,
            Exp::Add(lhs, rhs) => lhs.eval() + rhs.eval(),
            Exp::Mul(lhs, rhs) => lhs.eval() * rhs.eval(),
        }
    }
}

fn parse_line1(s: &str) -> Exp {
    peg::parser! {
      grammar parser() for str {
        rule _() = [' ' | '\t' | '\r']*

        rule atom() -> Exp
          = n:$(['0'..='9']+) { Exp::Val(n.parse().unwrap()) }
          / "(" e:expr() ")" { e }

        rule followOpt(left:Exp) -> Exp
          = _ "+" _ e:atom() f:followOpt(Exp::Add(Box::new(left.clone()), Box::new(e))) { f }
          / _ "*" _ e:atom() f:followOpt(Exp::Mul(Box::new(left.clone()), Box::new(e))) { f }
          / { left }

        rule expr() -> Exp
          = a:atom() _ f:followOpt(a) { f }
          / a:atom() { a }

        pub(crate) rule line() -> (Exp)
          = e:expr() { e }
      }
    }

    parser::line(s).unwrap()
}

fn parse_line2(s: &str) -> Exp {
    peg::parser! {
      grammar parser() for str {
        rule _() = [' ' | '\t' | '\r']*

        rule atom() -> Exp
          = n:$(['0'..='9']+) { Exp::Val(n.parse().unwrap()) }
          / "(" e:expr() ")" { e }

        rule mult() -> Exp
          = a:add() _ "*" _ e:mult()  { Exp::Mul(Box::new(a), Box::new(e)) }
          / a:add() { a }

        rule add() -> Exp
          = a:atom() _ "+" _ e:add()  { Exp::Add(Box::new(a), Box::new(e)) }
          / a:atom() { a }

        rule expr() -> Exp
          = a:mult() { a }
          / a:atom() { a }

        pub(crate) rule line() -> (Exp)
          = e:expr() { e }
      }
    }

    parser::line(s).unwrap()
}
