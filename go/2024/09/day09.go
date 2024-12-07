package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"time"
)

//go:embed sample.txt
var inputTest string

func Part1(input string) int {
	//var lines = strings.Split(input, "\n")

	var res int
	return res
}

func Part2(input string) int {
	//var lines = strings.Split(input, "\n")

	var res int
	return res
}

func main() {
	fmt.Println("--2024 day 09 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
