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
}

type I int

func (i I) isExpr() {}

type Reg byte

func (r Reg) isExpr() {}

type Inp struct{ reg Reg }

func (i Inp) isExpr() {}

type Add struct {
	reg Reg
	arg Expr
}

func (a Add) isExpr() {}

type Mul struct {
	reg Reg
	arg Expr
}

func (m Mul) isExpr() {}

type Div struct {
	reg Reg
	arg Expr
}

func (d Div) isExpr() {}

type Mod struct {
	reg Reg
	arg Expr
}

func (m Mod) isExpr() {}

type Eql struct {
	reg Reg
	arg Expr
}

func (e Eql) isExpr() {}

func parse(input string) (r Expr) {
	cmd := strings.Split(input, " ")

	reg := Reg(cmd[1][0] - 'w')
	if cmd[0] == "inp" {
		return Inp{reg}
	}

	var arg Expr
	if cmd[2][0] == 'w' || cmd[2][0] == 'x' || cmd[2][0] == 'y' || cmd[2][0] == 'z' {
		arg = Reg(cmd[2][0] - 'w')
	} else {
		num, _ := strconv.Atoi(cmd[2])
		arg = I(num)
	}

	switch cmd[0] {
	case "add":
		return Add{reg, arg}
	case "mul":
		return Mul{reg, arg}
	case "div":
		return Div{reg, arg}
	case "mod":
		return Mod{reg, arg}
	case "eql":
		return Eql{reg, arg}
	}
	return
}

func (i I) String() string {
	return fmt.Sprintf("%d", i)
}
func (r Reg) String() string {
	return fmt.Sprintf("%c", 'w'+byte(r))
}
func (i Inp) String() string {
	return fmt.Sprintf("inp %s\n", i.reg)
}
func (a Add) String() string {
	return fmt.Sprintf("add %s %s\n", a.reg, a.arg)
}
func (m Mul) String() string {
	return fmt.Sprintf("mul %s %s\n", m.reg, m.arg)
}
func (m Mod) String() string {
	return fmt.Sprintf("mod %s %s\n", m.reg, m.arg)
}
func (d Div) String() string {
	return fmt.Sprintf("div %s %s\n", d.reg, d.arg)
}
func (e Eql) String() string {
	return fmt.Sprintf("eql %s %s\n", e.reg, e.arg)
}

func value(e Expr, env []int) int {
	switch exp := e.(type) {
	case I:
		return int(exp)
	case Reg:
		return env[exp]
	}
	return 0
}

func eval(e Expr, env []int, input *[]int) {
	switch exp := e.(type) {
	case Inp:
		env[exp.reg] = (*input)[0]
		*input = (*input)[1:]
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

func increment(inp []int, index int) []int {
	if index < 0 {
		return []int{}
	}
	if inp[index] == 9 {
		inp[index] = 1
		increment(inp, index-1)
	} else {
		inp[index]++
	}
	return inp
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

func compile(input string) string {
	cmd := strings.Split(input, " ")
	switch cmd[0] {
	case "inp":
		return fmt.Sprintf("%v = inp[i]\ni++", cmd[1])
	case "add":
		return fmt.Sprintf("%v += %v", cmd[1], cmd[2])
	case "mul":
		return fmt.Sprintf("%v *= %v", cmd[1], cmd[2])
	case "div":
		return fmt.Sprintf("%v /= %v", cmd[1], cmd[2])
	case "mod":
		return fmt.Sprintf("%v=mod(%v,%v)", cmd[1], cmd[1], cmd[2])
	case "eql":
		return fmt.Sprintf("%v=eql(%v,%v)", cmd[1], cmd[1], cmd[2])
	}
	return ""
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	instructions := make([]Expr, 0, len(lines))
	for _, line := range lines {
		instructions = append(instructions, parse(line))
		fmt.Println(compile(line))
	}

	// inp := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	// for {
	// 	env := make([]int, 4)
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
	fmt.Println("--2021 day 24 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}
