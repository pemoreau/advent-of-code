package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"time"
)

//go:embed sample.txt
var inputTest string

func simulate(values []int, n int) int {
	mult := make([]int, 9)
	for _, v := range values {
		mult[v] = mult[v] + 1
	}

	for i := 0; i < n; i++ {
		// rotate array left
		mult = append(mult[1:], mult[0])
		mult[6] += mult[8]
	}

	return count(mult)
}

func count(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func Part1(input string) int {
	values := utils.CommaSeparatedToNumbers(input)
	return simulate(values, 80)
}

func Part2(input string) int {
	values := utils.CommaSeparatedToNumbers(input)
	return simulate(values, 256)
}

func main() {
	fmt.Println("--2021 day 06 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
