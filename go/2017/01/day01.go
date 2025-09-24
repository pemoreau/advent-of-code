package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

//go:embed sample.txt
var inputTest string

func solve(input string, n int) int {
	input = strings.TrimSuffix(input, "\n")
	var res = 0
	for i, v := range input {
		next := input[(i+n)%len(input)]
		d1 := uint8(v - '0')
		d2 := uint8(next - '0')
		if d1 == d2 {
			res += int(d1)
		}
	}
	return res
}

func Part1(input string) int {
	return solve(input, 1)
}

func Part2(input string) int {
	n := len(input) / 2
	return solve(input, n)
}

func main() {
	fmt.Println("--2017 day 01 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
