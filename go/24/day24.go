package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	. "github.com/pemoreau/advent-of-code-2021/go/utils"
)

//go:embed input.txt
var input_day string

type Instr interface {
	isInstr()
}

type DecoratedInstr struct {
	Instr
	before map[Reg]Interval
	after  map[Reg]Interval
}

type Assign struct {
	reg Reg
	rhs Expr
}
type Input struct {
	reg Reg
}

func (a Assign) isInstr() {}
func (i Input) isInstr()  {}

func (i Input) String() string {
	return fmt.Sprintf("Input %s\n", i.reg)
}

func (a Assign) String() string {
	//return fmt.Sprintf("before: %v\n%s = %v\nafter: %v\n", a.before, a.reg, a.rhs, a.after)
	return fmt.Sprintf("%s = %v\n", a.reg, a.rhs)
}

type Expr interface {
	isExpr()
}
type Value int
type Reg string

type Add struct {
	reg Reg
	arg Expr
}

type Mul struct {
	reg Reg
	arg Expr
}

type Div struct {
	reg Reg
	arg Expr
}

type Mod struct {
	reg Reg
	arg Expr
}

type Eql struct {
	reg Reg
	arg Expr
}

func (v Value) isExpr() {}
func (r Reg) isExpr()   {}

func (a Add) isExpr() {}
func (m Mul) isExpr() {}
func (d Div) isExpr() {}
func (m Mod) isExpr() {}
func (e Eql) isExpr() {}

func (v Value) String() string {
	return fmt.Sprintf("%d", v)
}
func (r Reg) String() string {
	return string(r)
}

func (a Add) String() string {
	return fmt.Sprintf("(%s + %s)", a.reg, a.arg)
}
func (m Mul) String() string {
	return fmt.Sprintf("(%s * %s)", m.reg, m.arg)
}
func (m Mod) String() string {
	return fmt.Sprintf("(%s %% %s)", m.reg, m.arg)
}
func (d Div) String() string {
	return fmt.Sprintf("(%s / %s)", d.reg, d.arg)
}
func (e Eql) String() string {
	return fmt.Sprintf("(%s == %s)", e.reg, e.arg)
}

func parse(input string, index *int) (r Instr) {
	cmd := strings.Split(input, " ")

	reg := Reg(cmd[1][0])
	if cmd[0] == "inp" {
		var inputName Reg = Reg(fmt.Sprintf("w%d", *index))
		*index += 1
		return Assign{reg: reg, rhs: inputName}
	}

	var arg Expr
	if cmd[2] == "w" || cmd[2] == "x" || cmd[2] == "y" || cmd[2] == "z" {
		arg = Reg(cmd[2])
	} else {
		num, _ := strconv.Atoi(cmd[2])
		arg = Value(num)
	}
	var rhs Expr
	switch cmd[0] {
	case "add":
		rhs = Add{reg, arg}
	case "mul":
		rhs = Mul{reg, arg}
	case "div":
		rhs = Div{reg, arg}
	case "mod":
		rhs = Mod{reg, arg}
	case "eql":
		rhs = Eql{reg, arg}
	}
	return Assign{reg: reg, rhs: rhs}
}

func value(e Expr, env map[Reg]int) int {
	switch exp := e.(type) {
	case Value:
		return int(exp)
	case Reg:
		return env[exp]
	}
	return 0
}

func eqlInterval(a, b Interval) Interval {
	if a.Min == a.Max && a.Min == b.Min && a.Max == b.Max {
		return Interval{1, 1}
	}
	if a.Max < b.Min || b.Max < a.Min {
		return Interval{0, 0}
	}
	return Interval{0, 1}
}

func forwardAbstractInterpretation(e Expr, env map[Reg]Interval) Interval {
	switch exp := e.(type) {
	case Value:
		return Interval{int(exp), int(exp)}
	case Reg:
		return env[exp]
	case Add:
		return env[exp.reg].Add(forwardAbstractInterpretation(exp.arg, env))
	case Mul:
		return env[exp.reg].Mul(forwardAbstractInterpretation(exp.arg, env))
	case Div:
		return env[exp.reg].Div(forwardAbstractInterpretation(exp.arg, env))
	case Mod:
		return env[exp.reg].Mod2(forwardAbstractInterpretation(exp.arg, env))
	case Eql:
		return eqlInterval(env[exp.reg], forwardAbstractInterpretation(exp.arg, env))
	default:
		panic("unknown exp")
	}
}
func forwardAbstractInterpretationInstr(i Instr, env map[Reg]Interval) DecoratedInstr {
	newEnv := map[Reg]Interval{}
	for key, value := range env {
		newEnv[key] = value
	}

	// remove one-level indirection
	switch ins := i.(type) {
	case DecoratedInstr:
		i = ins.Instr
	}

	switch ins := i.(type) {
	case Assign:
		newEnv[ins.reg] = forwardAbstractInterpretation(ins.rhs, env)
		return DecoratedInstr{Instr: ins, before: env, after: newEnv}
	default:
		errorString := fmt.Sprintf("unknown instr: %v", ins)
		panic(errorString)
	}
	return DecoratedInstr{Instr: i}
}
func backwardAbstractInterpretation(e Expr, before, after map[Reg]Interval) Interval {
	//fmt.Printf("BEFORE: %v\n", before)
	//fmt.Printf("AFTER:  %v\n", after)
	switch exp := e.(type) {
	case Value:
		return Interval{int(exp), int(exp)}
	case Reg:
		fmt.Printf("REG:    %s: %v\n", exp, after[exp])
		return after[exp]
	case Add:
		inverse := after[exp.reg].Sub(backwardAbstractInterpretation(exp.arg, before, after))
		return inverse.Inter(before[exp.reg])
	case Mul:
		fmt.Printf("BEFORE: %s: %v\n", exp.reg, before[exp.reg])
		fmt.Printf("AFTER:  %s: %v\n", exp.reg, after[exp.reg])
		inverse := after[exp.reg].Div(backwardAbstractInterpretation(exp.arg, before, after))
		fmt.Printf("INV  :  %v\n", inverse)
		return inverse
	case Div:
		return after[exp.reg].Mul(backwardAbstractInterpretation(exp.arg, before, after))
	case Mod:
		return after[exp.reg].Mod2(backwardAbstractInterpretation(exp.arg, before, after))
	case Eql:
		return eqlInterval(after[exp.reg], backwardAbstractInterpretation(exp.arg, before, after))
	default:
		panic("backward unknown exp")
	}
}

func backwardAbstractInterpretationInstr(i DecoratedInstr) DecoratedInstr {
	newEnv := map[Reg]Interval{}
	for key, value := range i.before {
		newEnv[key] = value
	}

	// remove one-level indirection
	switch ins := i.Instr.(type) {
	case Assign:
		newEnv[ins.reg] = backwardAbstractInterpretation(ins.rhs, i.before, i.after)
		return DecoratedInstr{Instr: ins, before: newEnv, after: i.after}
	default:
		errorString := fmt.Sprintf("unknown instr: %v", ins)
		panic(errorString)
	}
	return DecoratedInstr{Instr: i}
}

func eval(e Instr, env map[Reg]int, index int) {
	switch instr := e.(type) {
	case Input:
		var inputName Reg = Reg(fmt.Sprintf("w%d", index))
		env[instr.reg] = env[inputName]
	case Assign:
		{
			switch exp := instr.rhs.(type) {
			case Add:
				env[exp.reg] += value(exp.arg, env)
			case Mul:
				env[exp.reg] *= value(exp.arg, env)
			case Div:
				if value(exp.arg, env) == 0 {
					panic("divide by zero")
				}
				env[exp.reg] /= value(exp.arg, env)
			case Mod:
				if env[exp.reg] < 0 || value(exp.arg, env) <= 0 {
					panic("modulo by zero")
				}
				env[exp.reg] %= value(exp.arg, env)
			case Eql:
				if value(exp.arg, env) == env[exp.reg] {
					env[exp.reg] = 1
				} else {
					env[exp.reg] = 0
				}
			}
		}
	}
}

func decrement(inp []int, index int) []int {
	if inp[index] == 1 {
		inp[index] = 9
		decrement(inp, index-1)
	} else {
		inp[index]--
	}
	if index < 0 {
		return []int{}
	}
	return inp
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	instructions := make([]Instr, 0, len(lines))
	index := 1
	for _, line := range lines {
		instructions = append(instructions, parse(line, &index))
	}

	env := map[Reg]Interval{"w": {0, 0}, "x": {0, 0}, "y": {0, 0}, "z": {0, 0}}
	// for i := 1; i <= 14; i++ {
	// 	var inputName Reg = Reg(fmt.Sprintf("w%d", i))
	// 	env[inputName] = Interval{1, 9}
	// }
	env["w1"] = Interval{1, 9}
	env["w2"] = Interval{1, 9}
	env["w3"] = Interval{1, 9}
	env["w4"] = Interval{1, 9}
	env["w5"] = Interval{1, 9}
	env["w6"] = Interval{1, 9}
	env["w7"] = Interval{1, 9}
	env["w8"] = Interval{1, 9}
	env["w9"] = Interval{1, 9}
	env["w10"] = Interval{1, 9}
	env["w11"] = Interval{1, 9}
	env["w12"] = Interval{1, 9}
	env["w13"] = Interval{1, 9}
	env["w14"] = Interval{1, 9}

	fmt.Println("env", env)
	dInstructions := make([]DecoratedInstr, 0)
	for _, instr := range instructions {
		newInstr := forwardAbstractInterpretationInstr(instr, env)
		env = newInstr.after
		fmt.Println(newInstr)
		dInstructions = append(dInstructions, newInstr)
	}

	var last DecoratedInstr = dInstructions[len(dInstructions)-1]
	last.after["z"] = Interval{0, 0}
	for i := len(dInstructions) - 1; i >= 0; i-- {
		fmt.Printf("BACKWARD i=%d\n", i)
		newInstr := backwardAbstractInterpretationInstr(dInstructions[i])
		fmt.Println(newInstr)
	}

	//genCode(lines)

	// inp := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	// env := map[string]int{"w1": 9, "w2": 9}
	// for {
	// 	inpCopy := make([]int, len(inp))
	// 	copy(inpCopy, inp)

	// 	// fmt.Printf("try:  %v\n", inp)
	// 	for _, i := range instructions {
	// 		eval(i, env, &inpCopy)
	// 	}
	// 	if len(inpCopy) != 0 {
	// 		panic("input not consumed")
	// 	}
	// 	fmt.Printf("env=%v\tinp=%v\n", env, inp)
	// 	if env[3] == 0 {
	// 		fmt.Printf("inp=%v\n", inp)
	// 		break
	// 	}
	// 	inp = decrement(inp, len(inp)-1)
	// }

	return 0
}

func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	return 0

}

func main() {
	Part1(string(input_day))
	// fmt.Println("--2021 day 24 solution--")
	// start := time.Now()
	// fmt.Println("part1: ", Part1(string(input_day)))
	// fmt.Println(time.Since(start))

	// start = time.Now()
	// fmt.Println("part2: ", Part2(string(input_day)))
	// fmt.Println(time.Since(start))
}
