package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

type Expr interface {
	//eval(mem map[string]Expr, cache map[string]Value) Value
	simplify(mem map[string]Expr) Expr
	isValue() bool
}

type Value struct {
	n int
	d int
}
type Var string
type Op struct {
	Op  byte
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
	return fmt.Sprintf("(%v %c %v)", e.lhs, e.Op, e.rhs)
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

func (e Var) display(mem map[string]Expr) string {
	v, _ := mem[string(e)]
	switch v.(type) {
	case Var:
		return "var " + string(v.(Var))
	case Value:
		v := v.(Value)
		if v.d == 1 {
			return fmt.Sprintf("%d", v.n)
		} else {
			return fmt.Sprintf("%d/%d", v.n, v.d)
		}
	case Op:
		return fmt.Sprintf("(%v)", v.(Op).display(mem))
	}
	return ""
}

func (e Op) display(mem map[string]Expr) string {
	lhs := e.lhs.(Var).display(mem)
	rhs := e.rhs.(Var).display(mem)
	return fmt.Sprintf("%v %c %v", lhs, e.Op, rhs)
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

func GCDRemainderRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return GCDRemainderRecursive(b, a%b)
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
			switch e.Op {
			case '+':
				tmp = Value{lv.n*rv.d + rv.n*lv.d, lv.d * rv.d}
			case '*':
				tmp = Value{lv.n * rv.n, lv.d * rv.d}
			case '/':
				tmp = Value{lv.n * rv.d, lv.d * rv.n}
			case '-':
				tmp = Value{lv.n*rv.d - rv.n*lv.d, lv.d * rv.d}
			default:
				fmt.Printf("unknown op %c\n", e.Op)
				panic("unknown op")
			}
			if isInt(tmp) {
				return Value{getInt(tmp), 1}
			}
			gcd := GCDRemainderRecursive(tmp.n, tmp.d)
			tmp.n /= gcd
			tmp.d /= gcd
			//fmt.Println("simplify ", e, " to ", tmp)
			return tmp

		}
	}
	return Op{Op: e.Op, lhs: lhs, rhs: rhs}
}

func (e Op) solve(v Var, mem map[string]Expr) Expr {
	//fmt.Println("solve e = ", e)
	lhs := e.lhs.simplify(mem)
	rhs := e.rhs.simplify(mem)
	fmt.Println("solve", v, " in ", lhs, " = ", rhs)
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
		// lhs is: a Op b
		o := lhs.(Op)
		a := o.lhs.simplify(mem)
		b := o.rhs.simplify(mem)
		//fmt.Printf("\ndo %c %v\n\n", inverse[o.Op], b.(Value))
		var newEq Op
		if o.Op == '+' {
			if !a.isValue() {
				newEq = Op{Op: '=', lhs: a, rhs: Op{Op: '-', lhs: rhs, rhs: b}}
			}
			if !b.isValue() {
				newEq = Op{Op: '=', lhs: b, rhs: Op{Op: '-', lhs: rhs, rhs: a}}
			}
		}
		if o.Op == '*' {
			if !a.isValue() {
				newEq = Op{Op: '=', lhs: a, rhs: Op{Op: '/', lhs: rhs, rhs: b}}
			}
			if !b.isValue() {
				newEq = Op{Op: '=', lhs: b, rhs: Op{Op: '/', lhs: rhs, rhs: a}}
			}
		}
		if o.Op == '-' {
			if !a.isValue() {
				newEq = Op{Op: '=', lhs: a, rhs: Op{Op: '+', lhs: rhs, rhs: b}}
			}
			if !b.isValue() {
				newEq = Op{Op: '=', lhs: b, rhs: Op{Op: '-', lhs: a, rhs: rhs}}
			}
		}
		if o.Op == '/' {
			if !a.isValue() {
				newEq = Op{Op: '=', lhs: a, rhs: Op{Op: '*', lhs: rhs, rhs: b}}
			}
			if !b.isValue() {
				newEq = Op{Op: '=', lhs: b, rhs: Op{Op: '/', lhs: a, rhs: rhs}}
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
	//cache = make(map[string]Value)
	for _, line := range lines {
		values := strings.Split(line, " ")
		vname := strings.TrimSuffix(values[0], ":")
		//fmt.Println(len(values), values)
		if len(values) == 2 {
			e, _ := strconv.Atoi(values[1])
			mem[vname] = Value{e, 1}
		} else if len(values) == 4 {
			e := Op{
				Op:  values[2][0],
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

	//for e := range mem {
	//	fmt.Println(e, " := ", mem[e])
	//}
	res := 0
	//res := Var("root").eval(mem)

	rootExpr := mem["root"]
	delete(mem, "root")
	e := Op{
		Op:  '=',
		lhs: Var("root"),
		rhs: rootExpr,
	}.solve("root", mem)
	fmt.Println("solve", e)

	return res
}

func Part2(input string) int {
	mem := parse(input)
	delete(mem, "humn")
	root := mem["root"].(Op)
	root.Op = '='

	eq := root.solve("humn", mem)
	fmt.Println("humn", eq)

	root.Op = '-'
	fmt.Println("root", root.simplify(mem))
	switch eq := eq.(type) {
	case Value:
		v := eq.n / eq.d
		fmt.Println("v", v)
		return v
	}
	return 0
}

func main() {
	fmt.Println("--2022 day 21 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
