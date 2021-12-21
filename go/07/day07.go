package main

import (
	_ "embed"
	"fmt"
	"math"
	"time"

	"github.com/pemoreau/advent-of-code-2021/go/utils"
)

//go:embed input.txt
var input_day string

type fn func(int, int) int

func search(values []int, cost fn) int {
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
		for _, value := range values {
			sum += cost(h, value)
		}
		if sum < minSum {
			minSum = sum
		}
	}
	return minSum
}

func Part1(input string) int {
	values := utils.CommaSeparatedToNumbers(input)
	return search(values, func(a, b int) int {
		return abs(a - b)
	})
}

func Part2(input string) int {
	values := utils.CommaSeparatedToNumbers(input)
	return search(values, func(a, b int) int {
		n := abs(a - b)
		return n * (n + 1) / 2
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("--2021 day 07 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}
