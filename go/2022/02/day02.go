package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func play1(i, j int8) int8 {
	if i == j { // drawn
		return 3 + (j + 1)
	}
	if j == (i+1)%3 { // win
		return 6 + (j + 1)
	}
	// loose
	return j + 1
}

func play2(i, j int8) int8 {
	switch j {
	case 0: // loose
		//return play1(i, (3+i-1)%3)
		return 1 + (3+i-1)%3
	case 1: // draw
		//return play1(i, i)
		return 3 + (i + 1)
	default: // win
		//return play1(i, (i+1)%3)
		return 6 + 1 + (i+1)%3
	}
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	var table [3][3]int8
	for i := int8(0); i < 3; i++ {
		for j := int8(0); j < 3; j++ {
			table[i][j] = play1(i, j)
		}
	}
	score := 0
	for _, line := range lines {
		//score += play1(line[0]-'A', line[2]-'X')
		score += int(table[line[0]-'A'][line[2]-'X'])
	}
	return score
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	var table [3][3]int16
	for j := int16(0); j < 3; j++ {
		for i := int16(0); i < 3; i++ {
			switch j {
			case 0: // loose
				table[i][j] = 1 + (3+i-1)%3
			case 1: // draw
				table[i][j] = 3 + (i + 1)
			default: // win
				table[i][j] = 6 + 1 + (i+1)%3
			}
		}
	}
	score := 0
	for _, line := range lines {
		score += int(table[line[0]-'A'][line[2]-'X'])
	}
	return score
}

func main() {
	fmt.Println("--2022 day 02 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
