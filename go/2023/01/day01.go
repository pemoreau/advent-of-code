package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func startWithDigit1(s string) (bool, int) {
	if len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
		return true, int(s[0] - '0')
	}
	return false, 0
}

func startWithDigit2(s string) (bool, int) {
	var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, d := range digits {
		if strings.HasPrefix(s, d) {
			if i > 9 {
				i -= 10
			}
			return true, i
		}
	}
	return false, 0
}

func solve(input string, startWithDigit func(string) (bool, int)) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var sum int
	for _, line := range lines {
		var first int = 10
		var last int
		//fmt.Println(line)
		for i := range line {
			//fmt.Println(i, line[i:])
			isDigit, digit := startWithDigit(line[i:])
			if isDigit {
				if first > 9 {
					first = digit
				}
				last = digit
			}
		}
		//fmt.Println(first, last)
		sum += 10*first + last
	}
	return sum
}

func Part1(input string) int {
	return solve(input, startWithDigit1)
}

func Part2(input string) int {
	return solve(input, startWithDigit2)
}

func main() {
	fmt.Println("--2023 day 01 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
