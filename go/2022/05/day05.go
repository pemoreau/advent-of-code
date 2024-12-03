package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/stack"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func BuildStacks(part0 string) []stack.Stack[uint8] {
	lines := strings.Split(part0, "\n")
	stackLines := lines[0 : len(lines)-1]
	axe := lines[len(lines)-1]

	// build stacks
	stacks := make([]stack.Stack[uint8], 0)
	for i := 1; i < len(axe); i += 4 {
		stacks = append(stacks, stack.NewStack[uint8]())
	}

	// parse stacks
	reverse(stackLines)
	for _, line := range stackLines {
		for i, j := 0, 1; j < len(line); i, j = i+1, j+4 {
			if line[j] != ' ' {
				stacks[i].Push(line[j])
			}
		}
	}
	return stacks
}

func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func scan(s string, index int) (int, int) {
	res := 0
	for i := index; i < len(s); i++ {
		if s[i] == ' ' {
			return res, i
		}
		res = res*10 + int(s[i]-'0')
	}
	return res, len(s)
}

func parseInstruction(instruction string) (n, src, dest int) {
	i := 5
	n, i = scan(instruction, i)
	i += 6
	src, i = scan(instruction, i)
	i += 4
	dest, i = scan(instruction, i)
	return
}

func eval(stacks []stack.Stack[uint8], instructions []string, part2 bool) string {
	for _, instruction := range instructions {

		n, src, dest := parseInstruction(instruction)

		if !part2 {
			e, _ := stacks[src-1].PopN(n)
			reverse(e)
			stacks[dest-1].PushN(e)
		} else {
			e, _ := stacks[src-1].PopN(n)
			stacks[dest-1].PushN(e)
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
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
