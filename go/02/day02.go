package main

import (
	"aoc2021/utils"
	"fmt"
	"time"
)

func Part1(filename string) int {
	values := utils.ReadNumbers(filename)
	cpt := 0
	for i := 1; i < len(values); i++ {
		if values[i-1] < values[i] {
			cpt++
		}
	}
	return cpt
}

func Part2(filename string) int {
	values := utils.ReadNumbers(filename)
	cpt := 0
	for i := 3; i < len(values); i++ {
		if values[i-3] < values[i] {
			cpt++
		}
	}
	return cpt
}

func main() {
	filename := "../../inputs/day02.txt"

	start := time.Now()
	fmt.Println("part1: ", Part1(filename))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(filename))
	fmt.Println(time.Since(start))
}
