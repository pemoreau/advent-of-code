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
	if i == j { // drawn
		return int(3 + (j + 1))
	}
	if j == (i+1)%3 { // win
		return int(6 + (j + 1))
	}
	// loose
	return int(j + 1)
}

func play2(i, j uint8) int {
	switch j {
	case 0: // loose
		//return play1(i, (3+i-1)%3)
		return int(1 + (3+i-1)%3)
	case 1: // draw
		//return play1(i, i)
		return int(3 + (i + 1))
	default: // win
		//return play1(i, (i+1)%3)
		return int(6 + 1 + (i+1)%3)
	}
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	score := 0
	for _, line := range lines {
		score += play1(line[0]-'A', line[2]-'X')
	}
	return score
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	score := 0
	for _, line := range lines {
		i := line[0] - 'A'
		j := line[2] - 'X'
		switch j {
		case 0: // loose
			score += int(1 + (3+i-1)%3)
		case 1: // draw
			score += int(3 + (i + 1))
		default: // win
			score += int(6 + 1 + (i+1)%3)
		}
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
