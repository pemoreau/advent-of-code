package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"time"
)

//go:embed sample.txt
var inputTest string

func search(values []int, part1 bool) int {
	minValue, maxValue := math.MaxInt, 0
	for _, value := range values {
		if value < minValue {
			minValue = value
		}
		if value > maxValue {
			maxValue = value
		}
	}

	minSum := math.MaxInt

	for h := minValue; h <= maxValue; h++ {
		sum := 0
		if part1 {
			for _, value := range values {
				sum += utils.Abs(h - value)
			}

		} else {
			for _, value := range values {
				n := utils.Abs(h - value)
				sum += n * (n + 1) / 2
			}
		}
		if sum < minSum {
			minSum = sum
		}
	}
	return minSum
}
func Part1(input string) int {
	values := utils.CommaSeparatedToNumbers(input)
	return search(values, true)
}

func Part2(input string) int {
	values := utils.CommaSeparatedToNumbers(input)
	return search(values, false)
}

func main() {
	fmt.Println("--2021 day 07 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
