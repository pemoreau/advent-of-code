package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func validDesign(design string, patterns []string, cache map[string]int) int {
	if len(design) == 0 {
		return 1
	}

	if v, ok := cache[design]; ok {
		return v
	}

	var res int
	for _, p := range patterns {
		if strings.HasPrefix(design, p) {
			res = res + validDesign(design[len(p):], patterns, cache)
		}
	}
	cache[design] = res
	return res
}

func solve(input string) (int, int) {
	var parts = strings.Split(input, "\n\n")
	var patterns []string
	for _, p := range strings.Split(parts[0], ",") {
		patterns = append(patterns, strings.TrimSpace(p))
	}
	var designs []string
	for _, t := range strings.Split(parts[1], "\n") {
		designs = append(designs, t)
	}
	var res1, res2 int
	var cache = make(map[string]int)

	for _, design := range designs {
		var v = validDesign(design, patterns, cache)
		if v > 0 {
			res1++
		}
		res2 += v
	}
	return res1, res2
}

func Part1(input string) int {
	var res, _ = solve(input)
	return res
}

func Part2(input string) int {
	var _, res = solve(input)
	return res
}

func main() {
	fmt.Println("--2024 day 19 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
