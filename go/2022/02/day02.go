package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

func play1(i, j uint8) int {
	if i == j {
		return int(3 + (j + 1))
	}
	if j == (i+1)%3 {
		return int(6 + (j + 1))
	}
	return int(j + 1)
}

func play2(i, j uint8) int {
	if j == 1 { // drawn
		return play1(i, i)
	}
	if j == 0 { // loose
		return play1(i, (3+i-1)%3)
	}
	// win
	return play1(i, (i+1)%3)
}

func Part(input string, play func(uint8, uint8) int) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	score := 0
	for _, line := range lines {
		s := strings.Split(line, " ")
		score += play(s[0][0]-'A', s[1][0]-'X')
	}
	return score
}

func Part1(input string) int {
	return Part(input, play1)
}

func Part2(input string) int {
	return Part(input, play2)
}

func main() {
	fmt.Println("--2022 day 02 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
