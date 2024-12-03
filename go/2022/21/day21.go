package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Expr interface {
	simplify(mem map[string]Expr) Expr
	isValue() bool
}

type Value struct {
	n int
	d int
}
type Var string
type Op struct {
	op  byte
	lhs Expr
	rhs Expr
}

func (e Value) isValue() bool { return true }
func (e Op) isValue() bool    { return false }
func (e Var) isValue() bool   { return false }

func (e Value) String() string {
	return fmt.Sprintf("%d/%d", e.n, e.d)
}

func (e Var) String() string {
	return string(e)
}

func (e Op) String() string {
	return fmt.Sprintf("(%v %c %v)", e.lhs, e.op, e.rhs)
}

func isInt(e Expr) bool {
	if e.isValue() {
		n := e.(Value).n
		d := e.(Value).d
		return n%d == 0
	}
	return false
}

func getInt(e Expr) int {
	if e.isValue() {
		n := e.(Value).n
		d := e.(Value).d
		return n / d
	}
	panic("not int")
}

func (e Value) simplify(mem map[string]Expr) Expr {
	return e
}

func (e Var) simplify(mem map[string]Expr) Expr {
	v, ok := mem[string(e)]
	if ok {
		return v.simplify(mem)
	}
	//fmt.Println("simplify var ", e, " do noting ")
	return e
}

func (e Op) simplify(mem map[string]Expr) Expr {
	lhs := e.lhs.simplify(mem)
	rhs := e.rhs.simplify(mem)
	switch lhs.(type) {
	case Value:
		lv := lhs.(Value)
		switch rhs.(type) {
		case Value:
			rv := rhs.(Value)
			var tmp Value
			switch e.op {
			case '+':
				tmp = Value{lv.n*rv.d + rv.n*lv.d, lv.d * rv.d}
			case '*':
				tmp = Value{lv.n * rv.n, lv.d * rv.d}
			case '/':
				tmp = Value{lv.n * rv.d, lv.d * rv.n}
			case '-':
				tmp = Value{lv.n*rv.d - rv.n*lv.d, lv.d * rv.d}
			default:
				fmt.Printf("unknown op %c\n", e.op)
				panic("unknown op")
			}
			if isInt(tmp) {
				return Value{getInt(tmp), 1}
			}
			gcd := utils.GCD(tmp.n, tmp.d)
			tmp.n /= gcd
			tmp.d /= gcd
			//fmt.Println("simplify ", e, " to ", tmp)
			return tmp

		}
	}
	return Op{op: e.op, lhs: lhs, rhs: rhs}
}

func (e Op) solve(v Var, mem map[string]Expr) Expr {
	lhs := e.lhs.simplify(mem)
	rhs := e.rhs.simplify(mem)
	if lhs.isValue() {
		lhs, rhs = rhs, lhs
	}
	// lhs is an expr, not a value
	if lhs == v {
		return rhs
	}
	if !rhs.isValue() {
		panic("rhs is not a value")
	}

	switch lhs.(type) {
	case Value:
		if lhs.(Value) == rhs.(Value) {
			return e
		}
		panic("lhs is value")
	case Var:
		if lhs.(Var) == v {
			return e
		}
		panic("lhs is not the correct var")
	case Op:
		// lhs is: a op b
		o := lhs.(Op)
		a := o.lhs.simplify(mem)
		b := o.rhs.simplify(mem)
		//fmt.Printf("\ndo %c %v\n\n", inverse[o.op], b.(Value))
		var newEq Op
		if o.op == '+' {
			if !a.isValue() {
				newEq = Op{op: '=', lhs: a, rhs: Op{op: '-', lhs: rhs, rhs: b}}
			}
			if !b.isValue() {
				newEq = Op{op: '=', lhs: b, rhs: Op{op: '-', lhs: rhs, rhs: a}}
			}
		}
		if o.op == '*' {
			if !a.isValue() {
				newEq = Op{op: '=', lhs: a, rhs: Op{op: '/', lhs: rhs, rhs: b}}
			}
			if !b.isValue() {
				newEq = Op{op: '=', lhs: b, rhs: Op{op: '/', lhs: rhs, rhs: a}}
			}
		}
		if o.op == '-' {
			if !a.isValue() {
				newEq = Op{op: '=', lhs: a, rhs: Op{op: '+', lhs: rhs, rhs: b}}
			}
			if !b.isValue() {
				newEq = Op{op: '=', lhs: b, rhs: Op{op: '-', lhs: a, rhs: rhs}}
			}
		}
		if o.op == '/' {
			if !a.isValue() {
				newEq = Op{op: '=', lhs: a, rhs: Op{op: '*', lhs: rhs, rhs: b}}
			}
			if !b.isValue() {
				newEq = Op{op: '=', lhs: b, rhs: Op{op: '/', lhs: a, rhs: rhs}}
			}
		}

		return newEq.solve(v, mem)
	}
	panic("unknown solve")
}

func parse(input string) map[string]Expr {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	mem := make(map[string]Expr)
	for _, line := range lines {
		values := strings.Split(line, " ")
		vname := strings.TrimSuffix(values[0], ":")
		if len(values) == 2 {
			e, _ := strconv.Atoi(values[1])
			mem[vname] = Value{e, 1}
		} else if len(values) == 4 {
			e := Op{
				op:  values[2][0],
				lhs: Var(values[1]),
				rhs: Var(values[3]),
			}
			mem[vname] = e
		}
	}
	return mem
}

func Part1(input string) int {
	mem := parse(input)
	v := mem["root"].simplify(mem).(Value)
	res := v.n / v.d
	return res
}

func Part2(input string) int {
	mem := parse(input)
	delete(mem, "humn")
	root := mem["root"].(Op)
	root.op = '='
	v := root.solve("humn", mem).(Value)
	res := v.n / v.d
	return res
}

func main() {
	fmt.Println("--2022 day 21 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
