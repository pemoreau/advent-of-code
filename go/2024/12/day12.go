package main

import (
	_ "embed"
	"fmt"
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
	fmt.Println("--2024 day 10 solution--")
	//var inputDay = utils.Input()
	var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
