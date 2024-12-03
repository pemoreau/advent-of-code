package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	values := utils.LinesToNumbers(input)
	cpt := 0
	for i := 1; i < len(values); i++ {
		if values[i-1] < values[i] {
			cpt++
		}
	}
	return cpt
}

func Part2(input string) int {
	values := utils.LinesToNumbers(input)
	cpt := 0
	for i := 3; i < len(values); i++ {
		if values[i-3] < values[i] {
			cpt++
		}
	}
	return cpt
}

func main() {
	fmt.Println("--2021 day 01 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
