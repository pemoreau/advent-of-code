package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		var a, b, c, d int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &c, &d)
		// check inclusion
		if (a <= c && d <= b) || (c <= a && b <= d) {
			res++
		}
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		var a, b, c, d int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &c, &d)
		// check overlap
		if !(b < c || d < a) {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println("--2022 day 04 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
