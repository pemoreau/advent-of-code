package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type choice struct {
	left  string
	right string
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var commands = lines[0]
	var split = func(r rune) bool {
		return r == '=' || r == '(' || r == ')' || r == ',' || r == ' '
	}

	var network = make(map[string]choice)
	for _, line := range lines[2:] {
		var fields = strings.FieldsFunc(line, split)
		network[fields[0]] = choice{fields[1], fields[2]}
	}

	var current = "AAA"
	var res int
	for current != "ZZZ" {
		var index = res % len(commands)
		var command = commands[index]
		var choice = network[current]
		if command == 'L' {
			current = choice.left
		} else {
			current = choice.right
		}
		res++
	}

	return res
}

func length(commands string, network map[string]choice, current string) int {
	var res int
	for current[2] != 'Z' {
		var index = res % len(commands)
		var command = commands[index]
		var choice = network[current]
		if command == 'L' {
			current = choice.left
		} else {
			current = choice.right
		}
		res++
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var commands = lines[0]
	var split = func(r rune) bool {
		return r == '=' || r == '(' || r == ')' || r == ',' || r == ' '
	}

	var network = make(map[string]choice)
	var starts []string

	for _, line := range lines[2:] {
		var fields = strings.FieldsFunc(line, split)
		network[fields[0]] = choice{fields[1], fields[2]}
		if fields[0][2] == 'A' {
			starts = append(starts, fields[0])
		}
	}

	var lengths []int
	for _, start := range starts {
		lengths = append(lengths, length(commands, network, start))
	}

	var res = utils.LCM(lengths[0], lengths[1], lengths...)

	return res
}

func main() {
	fmt.Println("--2023 day 08 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
