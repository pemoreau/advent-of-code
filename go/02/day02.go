package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input_day string

func Part1(input string) int {
	return 0
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println("--2021 day 02 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}
