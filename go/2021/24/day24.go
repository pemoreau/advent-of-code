package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/interval"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Instr interface {
	isInstr()
}

type Assign struct {
	reg Reg
	rhs Expr
}

type Input struct {
}

func (a Assign) isInstr() {}
func (i Input) isInstr()  {}

func (i Input) String() string {
	return "Input w"
}

func (a Assign) String() string {
	return fmt.Sprintf("%v = %v\n", a.reg, a.rhs)
}

type Expr interface {
	isExpr()
}
type Value int
type Reg uint8 // w:0 x:1 y:2 z:3

func (r Reg) String() string {
	return string('w' + r)
}

func regIndex(name byte) Reg {
	return Reg(name - 'w')
}

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

func parse(input string) (r Instr) {
	cmd := strings.Split(input, " ")

	reg := regIndex(cmd[1][0])
	if cmd[0] == "inp" {
		return Input{}
	}

	var arg Expr
	if cmd[2] == "w" || cmd[2] == "x" || cmd[2] == "y" || cmd[2] == "z" {
		arg = regIndex(cmd[2][0])
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
		return env[exp]
	}
	return 0
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
		// fmt.Printf("BEFORE MERGE = %v\n", len(world))
		world = merge(world)
		// fmt.Printf("AFTER MERGE  = %v\n", len(world))

		wIndex := regIndex('w')

		var wg sync.WaitGroup
		wg.Add(9)
		var tmp [10]World
		for i := 1; i <= 9; i++ {
			go func(i int) {
				defer wg.Done()
				for _, state := range world {
					env := state.env
					env[wIndex] = i

					envInterval := createEnvInterval(env)
					envInterval[wIndex] = interval.Interval{i, i}
					// If reachable(remaining, createEnvInterval(state.env)) {
					if reachable(remaining, envInterval) {
						newState := State{env: env, min: 10*state.min + i, max: 10*state.max + i}
						tmp[i] = append(tmp[i], &newState)
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
	case Assign:
		switch exp := instr.rhs.(type) {
		case Add:
			for _, state := range world {
				state.env[exp.reg] += value(exp.arg, &state.env)
			}
		case Mul:
			for _, state := range world {
				state.env[exp.reg] *= value(exp.arg, &state.env)
			}
		case Div:
			for _, state := range world {
				if value(exp.arg, &state.env) == 0 {
					panic("divide by zero")
				}
				state.env[exp.reg] /= value(exp.arg, &state.env)
			}
		case Mod:
			for _, state := range world {
				if state.env[exp.reg] < 0 {
					panic("negative modulo")
				} else if value(exp.arg, &state.env) <= 0 {
					panic("modulo by zero or negative")
				}
				state.env[exp.reg] %= value(exp.arg, &state.env)
			}
		case Eql:
			for _, state := range world {
				if value(exp.arg, &state.env) == state.env[exp.reg] {
					state.env[exp.reg] = 1
				} else {
					state.env[exp.reg] = 0
				}
			}
		}
	}
	return world
}

func merge(w World) World {
	m := map[Env]*struct{ min, max int }{}
	for _, state := range w {
		if entry, ok := m[state.env]; ok {
			entry.min = min(entry.min, state.min)
			entry.max = max(entry.max, state.max)
		} else {
			m[state.env] = &struct{ min, max int }{min: state.min, max: state.max}
		}
	}
	res := make(World, 0, len(m))
	for env, minmax := range m {
		res = append(res, &State{env, minmax.min, minmax.max})
	}
	return res
}

var minValue = math.MaxInt
var maxValue = 0

func Solve(input string) {
	if minValue < maxValue {
		return
	}
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	instructions := make([]Instr, 0, len(lines))
	for _, line := range lines {
		instructions = append(instructions, parse(line))
	}

	world := World{&State{env: Env{0, 0, 0, 0}, min: 0, max: 0}}
	for i, instr := range instructions {
		// fmt.Printf("#%d: %v\n", i, instr)
		world = eval(instr, instructions[i+1:], world)
	}
	z := regIndex('z')
	for _, state := range world {
		if state.env[z] == 0 {
			minValue = min(minValue, state.min)
			maxValue = max(maxValue, state.max)
		}
	}

}
func Part1(input string) int {
	Solve(input)
	return maxValue

}

func Part2(input string) int {
	Solve(input)
	return minValue
}

func main() {
	fmt.Println("--2021 day 24 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
