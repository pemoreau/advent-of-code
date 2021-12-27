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
		return Assign{reg, inputName}
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
	return Assign{reg, rhs}
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

func abstractInterpretation(e Expr, env map[Reg]Interval) Interval {
	switch exp := e.(type) {
	case Value:
		return Interval{int(exp), int(exp)}
	case Reg:
		return env[exp]
	case Add:
		return env[exp.reg].Add(abstractInterpretation(exp.arg, env))
	case Mul:
		return env[exp.reg].Mul(abstractInterpretation(exp.arg, env))
	case Div:
		return env[exp.reg].Div(abstractInterpretation(exp.arg, env))
	case Mod:
		return env[exp.reg].Mod2(abstractInterpretation(exp.arg, env))
	case Eql:
		return eqlInterval(env[exp.reg], abstractInterpretation(exp.arg, env))
	default:
		panic("unknown exp")
	}
}

func abstractInterpretationInstr(i Instr, env map[Reg]Interval) map[Reg]Interval {
	newEnv := map[Reg]Interval{}
	for key, value := range env {
		newEnv[key] = value
	}
	switch ins := i.(type) {
	case Assign:
		newEnv[ins.reg] = abstractInterpretation(ins.rhs, env)
		// switch exp := ins.rhs.(type) {
		// case Eql:
		// 	inter := asbtractInterpretation(ins.rhs, env)
		// 	if inter.Max < inter.Min {
		// 		// empty Interval ==> not equal
		// 		newEnv[ins.reg] = Interval{0, 0}
		// 	} else {
		// 		newEnv[ins.reg] = Interval{1, 1}
		// 		newEnv[exp.reg] = inter

		// 	}
		// default:
		// }
	default:
		panic("unknown instr")
	}
	return newEnv
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
	if index < 0 {
		return []int{}
	}
	if inp[index] == 1 {
		inp[index] = 9
		decrement(inp, index-1)
	} else {
		inp[index]--
	}
	return inp
}

func compile(input string, index int) (r string, i int) {
	cmd := strings.Split(input, " ")
	i = index
	switch cmd[0] {
	case "inp":
		r = fmt.Sprintf("%v = inp[%d]", cmd[1], index)
		i = index + 1
	case "add":
		r = fmt.Sprintf("%v = %v + %v", cmd[1], cmd[1], cmd[2])
	case "mul":
		if cmd[2] == "0" {
			r = fmt.Sprintf("%v = %v", cmd[1], cmd[2])
		} else {
			r = fmt.Sprintf("%v = %v * %v", cmd[1], cmd[1], cmd[2])
		}
	case "div":
		if cmd[2] != "1" {
			r = fmt.Sprintf("%v = %v / %v", cmd[1], cmd[1], cmd[2])
		}
	case "mod":
		r = fmt.Sprintf("%v = %v %% %v", cmd[1], cmd[1], cmd[2])
	case "eql":
		r = fmt.Sprintf("if %v == %v {\n %v=1\n} else {\n %v=0\n}", cmd[1], cmd[2], cmd[1], cmd[1])
	}
	return
}

func genCode(lines []string) {
	fmt.Println(`package main

	import (
		"fmt"
	)`)
	fmt.Println(`func Run(inp []int) (w, x, y, z int) {
		w = 0
		x = 0
		y = 0
		z = 0
		`)
	i := 0
	for _, line := range lines {
		var s string
		s, i = compile(line, i)
		fmt.Println(s)
	}
	fmt.Println("return \n}")

}

// func optimize(instr []Expr) []Expr {
// 	res := make([]Expr)

// 	return res
// }

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
	env["w1"] = Interval{9, 9}
	env["w2"] = Interval{9, 9}
	env["w3"] = Interval{1, 9}
	env["w4"] = Interval{1, 9}
	env["w5"] = Interval{7, 9}
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
	for _, instr := range instructions {
		fmt.Print(instr)
		env = abstractInterpretationInstr(instr, env)
		fmt.Println(env)
	}
	// instructions = optimize(instructions)
	// genCode(lines)

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
