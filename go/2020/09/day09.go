package main

import (
	_ "embed"
	"fmt"
	"math"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

//go:embed input.txt
var inputDay string

func check25(list []int, start int, end int, value int) bool {
	for i := start; i < end; i++ {
		for j := i + 1; j < end; j++ {
			if list[i]+list[j] == value {
				return true
			}
		}
	}
	return false
}

func searchSum(list []int, value int) int {
	for i := 0; i < len(list); i++ {
		sum := 0
		min := math.MaxInt64
		max := 0
		for j := i; j < len(list); j++ {
			v := list[j]
			sum += v
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
			if sum == value {
				return min + max
			} else if sum > value {
				break
			}
		}
	}
	return 0
}

func Part1(input string) int {
	list := utils.LinesToNumbers(input)
	for i := 0; i < len(list); i++ {
		if !check25(list, i, i+25, list[i+25]) {
			return list[i+25]
		}
	}
	return 0
}

func Part2(input string) int {
	num := Part1(input)
	list := utils.LinesToNumbers(input)
	return searchSum(list, num)
}

func main() {
	fmt.Println("--2020 day 09 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
