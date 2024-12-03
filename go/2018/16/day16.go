package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

var opcodes = []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"}
var decode = []int{13, 6, 0, 11, 3, 10, 2, 4, 7, 14, 15, 5, 8, 12, 1, 9}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func run(instruction Instruction, registers *Registers) {
	var opcode = opcodes[decode[instruction[0]]]
	var regA = registers[instruction[1]]
	var regB = registers[instruction[2]]
	var valA = instruction[1]
	var valB = instruction[2]
	var C int

	//fmt.Printf("opcode: %s, regA: %d, regB: %d, valA: %d, valB: %d\n", opcode, regA, regB, valA, valB)
	//fmt.Printf("input:  %v\n", registers)

	switch opcode {
	case "addr":
		C = regA + regB
	case "addi":
		C = regA + valB
	case "mulr":
		C = regA * regB
	case "muli":
		C = regA * valB
	case "banr":
		C = regA & regB
	case "bani":
		C = regA & valB
	case "borr":
		C = regA | regB
	case "bori":
		C = regA | valB
	case "setr":
		C = regA
	case "seti":
		C = valA
	case "gtir":
		C = bool2int(valA > regB)
	case "gtri":
		C = bool2int(regA > valB)
	case "gtrr":
		C = bool2int(regA > regB)
	case "eqir":
		C = bool2int(valA == regB)
	case "eqri":
		C = bool2int(regA == valB)
	case "eqrr":
		C = bool2int(regA == regB)
	}
	registers[instruction[3]] = C
}

type Instruction [4]int
type Registers [4]int

func Part1(input string) int {
	input = strings.Trim(input, "\n")
	var parts = strings.Split(input, "\n\n\n\n")
	var samples = strings.Split(parts[0], "\n\n")
	var res int
	for _, sample := range samples {
		var regBefore Registers
		var regAfter Registers
		var instruction Instruction
		var lines = strings.Split(sample, "\n")
		fmt.Sscanf(lines[0], "Before: [%d, %d, %d, %d]", &regBefore[0], &regBefore[1], &regBefore[2], &regBefore[3])
		fmt.Sscanf(lines[1], "%d %d %d %d", &instruction[0], &instruction[1], &instruction[2], &instruction[3])
		fmt.Sscanf(lines[2], "After: [%d, %d, %d, %d]", &regAfter[0], &regAfter[1], &regAfter[2], &regAfter[3])
		var count = 0
		for i, _ := range opcodes {
			var registers = regBefore
			instruction[0] = i
			run(instruction, &registers)
			if registers == regAfter {
				count++
			}
		}
		if count >= 3 {
			res++
		}
	}

	return res
}

func Part2(input string) int {
	input = strings.Trim(input, "\n")
	var parts = strings.Split(input, "\n\n\n\n")
	var program = strings.Split(parts[1], "\n")

	var registers Registers
	for _, line := range program {
		var instruction Instruction
		fmt.Sscanf(line, "%d %d %d %d", &instruction[0], &instruction[1], &instruction[2], &instruction[3])
		run(instruction, &registers)
	}

	return registers[0]
}

func main() {
	fmt.Println("--2018 day 16 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
