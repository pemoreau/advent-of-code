package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strconv"
	"strings"
	"time"
)

type Instr struct {
	op  string
	arg int
}

func parseLine(line string) Instr {
	parts := strings.Split(line, " ")
	op := strings.TrimSpace(parts[0])
	arg, _ := strconv.Atoi(parts[1])
	return Instr{op, arg}
}

func readFile(input string) []Instr {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	res := []Instr{}
	for _, line := range lines {
		res = append(res, parseLine(line))
	}
	return res
}

type Env struct {
	prg []Instr
	acc int
	pc  int
}

func step(env *Env) {
	if env.pc < len(env.prg) {
		inst := env.prg[env.pc]
		if inst.op == "acc" {
			// fmt.Printf("ACC: pc=%d acc=%d\n", env.pc, env.acc)
			env.acc += inst.arg
			env.pc++
		} else if inst.op == "jmp" {
			// fmt.Printf("JMP: pc=%d arg=%d\n", env.pc, inst.arg)
			env.pc += inst.arg
		} else if inst.op == "nop" {
			// fmt.Printf("NOP: pc=%d\n", env.pc)
			env.pc++
		}
	}
}

func terminates(env Env) (bool, int) {
	visited := set.NewSet[int]()
	for !visited.Contains(env.pc) {
		if env.pc >= len(env.prg) {
			return true, env.acc
		}
		visited.Add(env.pc)
		step(&env)
	}
	return false, env.acc
}

func Part1(input string) int {
	prg := readFile(input)
	_, acc := terminates(Env{prg: prg})
	return acc
}

func Part2(input string) int {
	prg := readFile(input)
	env := Env{prg: prg}

	for line, inst := range prg {
		if inst.op == "nop" {
			env.prg[line].op = "jmp"
			if t, acc := terminates(env); t {
				return acc
			}
			env.prg[line].op = "nop"
		}
	}
	for line, inst := range prg {
		if inst.op == "jmp" {
			env.prg[line].op = "nop"
			if t, acc := terminates(env); t {
				return acc
			}
			env.prg[line].op = "jmp"
		}
	}

	return 0
}

func main() {
	fmt.Println("--2020 day 08 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
