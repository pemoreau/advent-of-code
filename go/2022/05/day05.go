package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

func BuildStacks(part0 string) []utils.Stack[uint8] {
	lines := strings.Split(part0, "\n")
	stackLines := lines[0 : len(lines)-1]
	axe := lines[len(lines)-1]

	// build stacks
	stacks := make([]utils.Stack[uint8], 0)
	for i := 1; i < len(axe); i += 4 {
		stacks = append(stacks, utils.BuildStack[uint8]())
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

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func eval(stacks []utils.Stack[uint8], instructions []string, part2 bool) string {
	for _, instruction := range instructions {
		var n, src, dest int
		fmt.Sscanf(instruction, "move %d from %d to %d", &n, &src, &dest)
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
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
