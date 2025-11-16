package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func run(jumps []int, part2 bool) int {
	var step int
	var maxIndex = len(jumps) - 1
	var index int
	for index >= 0 && index <= maxIndex {
		//print(jumps, index)
		var newIndex = index + jumps[index]
		if part2 && jumps[index] >= 3 {
			jumps[index]--
		} else {
			jumps[index]++
		}
		index = newIndex
		step++
	}
	//print(jumps, index)
	return step
}

func print(jumps []int, index int) {
	for i, jump := range jumps {
		if i == index {
			fmt.Printf("(%d) ", jump)
		} else {
			fmt.Printf("%d ", jump)
		}
		if i < len(jumps)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func solve(input string, part2 bool) int {
	input = strings.TrimSpace(input)
	var lines = strings.Split(input, "\n")
	var jumps []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		v, _ := strconv.Atoi(line)
		jumps = append(jumps, v)
	}
	return run(jumps, part2)
}

func Part1(input string) int {
	return solve(input, false)
}

func Part2(input string) int {
	return solve(input, true)
}

func main() {
	fmt.Println("--2017 day 05 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
