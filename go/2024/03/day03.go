package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

func parse(input string, part2 bool) int {
	var res int
	var x, y int
	var enable = true
	for i := 0; i < len(input); i++ {
		if part2 && strings.HasPrefix(input[i:], "do()") {
			enable = true
		} else if part2 && strings.HasPrefix(input[i:], "don't()") {
			enable = false
		} else if strings.HasPrefix(input[i:], "mul(") {
			if _, err := fmt.Sscanf(input[i+4:], "%d,%d)", &x, &y); err == nil && enable {
				res += x * y
			}
		}
	}
	return res
}

func Part1(input string) int {
	return parse(input, false)
}

func Part2(input string) int {
	return parse(input, true)
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
