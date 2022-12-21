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
	isExpr()
	eval(mem map[string]Expr) int
	simplify(mem map[string]Expr) Expr
	solve(mem map[string]Expr) Expr
}

type Value int
type Var string
type Op struct {
	op  byte
	lhs Expr
	rhs Expr
}

func (e Value) isExpr() {}
func (e Op) isExpr()    {}
func (e Var) isExpr()   {}

func (e Value) String() string {
	return strconv.Itoa(int(e))
}
func (e Var) String() string {
	return string(e)
}

func (e Op) String() string {
	return fmt.Sprintf("(%v %c %v)", e.lhs, e.op, e.rhs)
}

func (e Var) display(mem map[string]Expr) string {
	v, _ := mem[string(e)]
	switch v.(type) {
	case Var:
		return "var " + string(v.(Var))
	case Value:
		return strconv.Itoa(int(v.(Value)))
	case Op:
		return fmt.Sprintf("(%v)", v.(Op).display(mem))
	}
	return ""
}

func (e Op) display(mem map[string]Expr) string {
	lhs := e.lhs.(Var).display(mem)
	rhs := e.rhs.(Var).display(mem)
	return fmt.Sprintf("%v %c %v", lhs, e.op, rhs)
}

func (e Value) simplify(mem map[string]Expr) Expr {
	return e
}

func (e Var) simplify(mem map[string]Expr) Expr {
	v, ok := mem[string(e)]
	if ok {
		return v.simplify(mem)
	}
	fmt.Println("simplify var ", e, " := ", v)
	return e
}

func (e Op) simplify(mem map[string]Expr) Expr {
	lhs := e.lhs.simplify(mem)
	rhs := e.rhs.simplify(mem)
	switch lhs.(type) {
	case Value:
		switch rhs.(type) {
		case Value:
			return Value(e.eval(mem))
		}
	}
	return Op{e.op, lhs, rhs}
}

func (e Value) eval(mem map[string]Expr) int {
	//fmt.Println("eval", e, " --> ", int(e))
	return int(e)
}
func (e Var) eval(mem map[string]Expr) int {
	if v, ok := cache[string(e)]; ok {
		//fmt.Println("cache", e, " --> ", v)
		return v
	}
	v := mem[string(e)].eval(mem)
	//fmt.Println("eval", e, " --> ", v)
	cache[string(e)] = v
	return v
}

func (e Op) eval(mem map[string]Expr) int {
	lhs := e.lhs.eval(mem)
	rhs := e.rhs.eval(mem)
	switch e.op {
	case '+':
		return lhs + rhs
	case '*':
		return lhs * rhs
	case '/':
		return lhs / rhs
	case '-':
		return lhs - rhs
	case '=':
		//fmt.Println("lhs", lhs)
		//fmt.Println("rhs", rhs)
		return lhs - rhs
	}
	panic("unknown op")
}

var cache map[string]int

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	mem := make(map[string]Expr)
	cache = make(map[string]int)
	for _, line := range lines {
		values := strings.Split(line, " ")
		vname := strings.TrimSuffix(values[0], ":")
		//fmt.Println(len(values), values)
		if len(values) == 2 {
			e, _ := strconv.Atoi(values[1])
			mem[vname] = Value(e)
		} else if len(values) == 4 {
			e := Op{
				op:  values[2][0],
				lhs: Var(values[1]),
				rhs: Var(values[3]),
			}
			mem[vname] = e
		}
	}

	//for e := range mem {
	//	fmt.Println(e, " := ", mem[e])
	//}
	res := Var("root").eval(mem)
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	mem := make(map[string]Expr)
	for _, line := range lines {
		values := strings.Split(line, " ")
		vname := strings.TrimSuffix(values[0], ":")
		//fmt.Println(len(values), values)
		if len(values) == 2 {
			e, _ := strconv.Atoi(values[1])
			mem[vname] = Value(e)
		} else if len(values) == 4 {
			e := Op{
				op:  values[2][0],
				lhs: Var(values[1]),
				rhs: Var(values[3]),
			}
			mem[vname] = e
		}
	}
	mem["root"] = Op{
		op:  '=',
		lhs: mem["root"].(Op).lhs,
		rhs: mem["root"].(Op).rhs,
	}
	//for e := range mem {
	//	fmt.Println(e, " := ", mem[e])
	//}

	//mem["humn"] = Value(1234)
	//fmt.Println("humn", Var("humn").display(mem))
	cache = make(map[string]int)

	delete(mem, "humn")
	name1 := "mcnw"
	name2 := "wqdw"
	e1 := Var(name1).simplify(mem)
	e2 := Var(name2).simplify(mem)
	fmt.Println("mcnw", e1.simplify(mem))
	fmt.Println("wqdw", e2.simplify(mem))
	fmt.Println(name1, e1)
	fmt.Println(name2, e2)

	//switch expr.(type) {
	//case Op:
	//	fmt.Println(name, expr.(Op).display(mem))
	//}

	//a := 1000000000000
	//for i := a + 0; i < a+30; i = i + 3 {
	//	//if i%10000 == 0 {
	//	//	fmt.Println(i)
	//	//}
	//	cache = make(map[string]int)
	//	mem["humn"] = Value(i)
	//	res := Var("root").eval(mem)
	//	fmt.Println("humn", i, " --> ", res, res-31343426392931)
	//	if res == 0 {
	//		return i
	//	}
	//}
	return 0
}

func main() {
	fmt.Println("--2022 day 21 solution--")
	start := time.Now()
	//fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
