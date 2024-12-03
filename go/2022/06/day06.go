package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"time"
)

//go:embed sample.txt
var inputTest string

func allDifferent(s string) bool {
	var tab [26]bool
	for i := 0; i < len(s); i++ {
		if tab[s[i]-'a'] {
			return false
		}
		tab[s[i]-'a'] = true
	}
	return true
}

func solve(input string, n int) int {
	for i := 0; i < len(input)-n; i++ {
		s := input[i : i+n]
		if allDifferent(s) {
			return i + n
		}
	}
	return -1
}

func Part1(input string) int {
	return solve(input, 4)
}

func Part2(input string) int {
	return solve(input, 14)
}

func main() {
	fmt.Println("--2022 day 06 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
