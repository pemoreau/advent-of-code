package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pemoreau/advent-of-code-2021/go/utils"
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
		var inputName = Reg(fmt.Sprintf("w%d", *index))
		*index += 1
		//return Assign{reg: reg, rhs: inputName}
		return Input{inputName}
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

func value(e Expr, env *Env) int {
	switch exp := e.(type) {
	case Value:
		return int(exp)
	case Reg:
		return env[regIndex(exp)]
	}
	return 0
}

func regIndex(reg Reg) int {
	return int(reg[0] - 'w')
}

type Env [4]int
type State struct {
	env      Env
	min, max int
}
type World []*State

func (w World) String() string {
	var s string
	for _, state := range w {
		s += fmt.Sprintf("%v %d %d, ", state.env, state.min, state.max)
	}
	s += "\n"
	return s
}

func eval(e Instr, remaining []Instr, world World) World {
	switch instr := e.(type) {
	case Input:
		fmt.Printf("BEFORE MERGE = %v\n", len(world))
		world = merge(world)
		fmt.Printf("AFTER MERGE  = %v\n", len(world))

		index := regIndex(instr.reg)

		var wg sync.WaitGroup
		wg.Add(9)
		var tmp [10]World
		for i := 1; i <= 9; i++ {
			go func(i int) {
				defer wg.Done()
				for _, state := range world {
					env := state.env
					env[index] = i

					envInterval := createEnvInterval(env)
					envInterval[index] = utils.Interval{i, i}
					// If reachable(remaining, createEnvInterval(state.env)) {
					if reachable(remaining, envInterval) {
						newState := State{env: env, min: 10*state.min + i, max: 10*state.max + i}
						tmp[i] = append(tmp[i], &newState)
						// fmt.Printf("tmp[%d]=%v\n", i, tmp[i])
					}
				}
			}(i)
		}
		wg.Wait()

		// Merge tmp[i] into world
		var newWorld World
		for i := 1; i <= 9; i++ {
			newWorld = append(newWorld, tmp[i]...)
		}
		world = newWorld
		// fmt.Printf("AFTER INPUT  = %v\n", len(world))
	case Assign:
		switch exp := instr.rhs.(type) {
		case Add:
			index := regIndex(exp.reg)
			for _, state := range world {
				state.env[index] += value(exp.arg, &state.env)
			}
		case Mul:
			index := regIndex(exp.reg)
			for _, state := range world {
				state.env[index] *= value(exp.arg, &state.env)
			}
		case Div:
			index := regIndex(exp.reg)
			for _, state := range world {
				if value(exp.arg, &state.env) == 0 {
					panic("divide by zero")
				}
				state.env[index] /= value(exp.arg, &state.env)
			}
		case Mod:
			index := regIndex(exp.reg)
			for _, state := range world {
				if state.env[regIndex(exp.reg)] < 0 {
					panic("negative modulo")
				} else if value(exp.arg, &state.env) <= 0 {
					panic("modulo by zero or negative")
				}
				state.env[index] %= value(exp.arg, &state.env)
			}
		case Eql:
			index := regIndex(exp.reg)
			for _, state := range world {
				if value(exp.arg, &state.env) == state.env[index] {
					state.env[index] = 1
				} else {
					state.env[index] = 0
				}
			}
		}
	}
	return world
}

func merge(w World) World {
	m := map[Env]*struct{ min, max int }{}
	//fmt.Printf("MERGE %v\n", w)
	for _, state := range w {
		if entry, ok := m[state.env]; ok {
			entry.min = utils.Min(entry.min, state.min)
			entry.max = utils.Max(entry.max, state.max)
			//fmt.Printf("UPDATE %v -> %v\n", state, entry)
		} else {
			m[state.env] = &struct{ min, max int }{min: state.min, max: state.max}
			//fmt.Printf("ADD %v -> %v\n", state, m[state.env])
		}
	}
	res := make(World, 0, len(m))
	for env, minmax := range m {
		res = append(res, &State{env, minmax.min, minmax.max})
	}
	return res
}

var min = math.MaxInt
var max = 0

func Solve(input string) {
	if min < max {
		return
	}
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	instructions := make([]Instr, 0, len(lines))
	index := 1
	for _, line := range lines {
		instructions = append(instructions, parse(line, &index))
	}

	world := World{&State{env: Env{0, 0, 0, 0}, min: 0, max: 0}}
	for i, instr := range instructions {
		fmt.Printf("#%d: %v\n", i, instr)
		world = eval(instr, instructions[i+1:], world)
	}
	for _, state := range world {
		if state.env[regIndex("z")] == 0 {
			min = utils.Min(min, state.min)
			max = utils.Max(max, state.max)
		}
	}

}
func Part1(input string) int {
	Solve(input)
	return max

}

func Part2(input string) int {
	Solve(input)
	return min
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
