package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input_test.txt
var inputTest string

func parse(lines []string, part2 bool) int {
	var res int
	var x, y int
	var enable = true
	for _, line := range lines {
		for i := 0; i < len(line)-7; i++ {
			if part2 && line[i:i+4] == "do()" {
				enable = true
			} else if part2 && line[i:i+7] == "don't()" {
				enable = false
			} else if _, err := fmt.Sscanf(line[i:], "mul(%d,%d)", &x, &y); err == nil && enable {
				res += x * y
			}
		}
	}

	return res
}

func Part1(input string) int {
	var lines = strings.Split(input, "\n")
	return parse(lines, false)
}

func Part2(input string) int {
	var lines = strings.Split(input, "\n")
	return parse(lines, true)
}

func main() {
	fmt.Println("--2024 day 03 solution--")

	var inputDay = utils.Input()
	//var inputDay = inputTest

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
