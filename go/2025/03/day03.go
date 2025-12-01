package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func Part1(input string) int {
	//input = strings.TrimSuffix(input, "\n")
	//var lines = strings.Split(input, "\n")
	var res int
	return res
}

func Part2(input string) int {
	//input = strings.TrimSuffix(input, "\n")
	//var lines = strings.Split(input, "\n")
	var res int
	return res
}

func main() {
	fmt.Println("--2025 day 03 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
