package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"text/scanner"
	"time"
)

//go:embed input.txt
var input_day string

func BuildStacks(part0 string) []utils.Stack[uint8] {
	lines := strings.Split(part0, "\n")
	stackLines := lines[0 : len(lines)-1]
	axe := lines[len(lines)-1]
	columns := make([]int, 0)

	// parse axe
	var s scanner.Scanner
	s.Init(strings.NewReader(axe))
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		columns = append(columns, s.Position.Column-1)
	}

	// build stacks
	stacks := make([]utils.Stack[uint8], 0)
	for i := 0; i < len(columns); i++ {
		stacks = append(stacks, utils.BuildStack[uint8]())
	}
	// parse stacks
	for i := len(stackLines) - 1; i >= 0; i-- {
		for index, j := range columns {
			if j <= len(stackLines[i]) && stackLines[i][j] != ' ' {
				stacks[index].Push(stackLines[i][j])
			}
		}
	}
	return stacks
}

func eval(stacks []utils.Stack[uint8], instructions []string, part2 bool) string {
	tmpStack := utils.BuildStack[uint8]()
	for _, instruction := range instructions {
		var n, src, dest int
		fmt.Sscanf(instruction, "move %d from %d to %d", &n, &src, &dest)
		if !part2 {
			for i := 0; i < n; i++ {
				e, _ := stacks[src-1].Pop()
				stacks[dest-1].Push(e)
			}
		} else {
			for i := 0; i < n; i++ {
				e, _ := stacks[src-1].Pop()
				tmpStack.Push(e)
			}
			for i := 0; i < n; i++ {
				e, _ := tmpStack.Pop()
				stacks[dest-1].Push(e)
			}
		}
	}
	res := ""
	for _, stack := range stacks {
		e, _ := stack.Peek()
		res += string(e)
	}
	return res
}

func Part1(input string) string {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	stacks := BuildStacks(parts[0])
	instructions := strings.Split(parts[1], "\n")
	return eval(stacks, instructions, false)
}

func Part2(input string) string {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	stacks := BuildStacks(parts[0])
	instructions := strings.Split(parts[1], "\n")
	return eval(stacks, instructions, true)
}

func main() {
	fmt.Println("--2022 day 05 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
