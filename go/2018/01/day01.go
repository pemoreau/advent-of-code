package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	var values = utils.LinesToNumbers(input)
	var total = 0
	for _, v := range values {
		total += v
	}
	return total
}

func Part2(input string) int {
	var values = utils.LinesToNumbers(input)
	var visited = set.NewSet[int]()
	visited.Add(0)
	var index = 0
	var total = 0
	for {
		total += values[index]
		index = (index + 1) % len(values)
		if visited.Contains(total) {
			return total
		}
		visited.Add(total)
	}
}

func main() {
	fmt.Println("--2018 day 01 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part1: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
