package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

var score = map[string]int{"X": 1, "Y": 2, "Z": 3}

func play1(a, b string) int {
	i := a[0] - 'A'
	j := b[0] - 'X'
	if i == j {
		return 3 + score[b]
	}
	if j == (i+1)%3 {
		return 6 + score[b]
	}
	return score[b]
}

func play2(a, b string) int {
	i := a[0] - 'A'
	if b == "Y" { // drawn
		return play1(a, string('X'+i))
	}
	if b == "X" { // loose
		return play1(a, string('X'+(3+i-1)%3))
	}
	// win
	return play1(a, string('X'+(i+1)%3))
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	score := 0
	for _, line := range lines {
		s := strings.Split(line, " ")
		score += play1(s[0], s[1])
	}
	return score
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	score := 0
	for _, line := range lines {
		s := strings.Split(line, " ")
		p := play2(s[0], s[1])
		fmt.Println("score:", p)
		score += p
	}
	return score

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
