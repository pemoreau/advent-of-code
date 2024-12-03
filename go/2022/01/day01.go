package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"sort"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func parse(input string) []int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	var elves []int
	for _, part := range parts {
		sum := 0
		for _, line := range strings.Split(part, "\n") {
			sum += utils.ToInt(line)
		}
		elves = append(elves, sum)
	}
	return elves
}

func Part1(input string) int {
	elves := parse(input)
	max := math.MinInt
	for _, elf := range elves {
		if elf > max {
			max = elf
		}
	}
	return max
}

func Part2(input string) int {
	elves := parse(input)
	sort.Ints(elves)
	return elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
}

func main() {
	fmt.Println("--2022 day 01 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
