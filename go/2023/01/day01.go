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
	l := len(s)
	if l == 0 {
		return false, 0
	}
	c := s[0]
	if c >= '0' && c <= '9' {
		return true, int(c - '0')
	}
	if l <= 2 {
		return false, 0
	}
	// discrimination net to improve performance
	switch c {
	case 'z':
		if l > 3 && s[1] == 'e' && s[2] == 'r' && s[3] == 'o' {
			return true, 0
		}
	case 'o':
		if s[1] == 'n' && s[2] == 'e' {
			return true, 1
		}
	case 't':
		if s[1] == 'w' && s[2] == 'o' {
			return true, 2
		} else if l > 4 && s[1] == 'h' && s[2] == 'r' && s[3] == 'e' && s[4] == 'e' {
			return true, 3
		}
	case 'f':
		if l > 3 && s[1] == 'o' && s[2] == 'u' && s[3] == 'r' {
			return true, 4
		} else if l > 3 && s[1] == 'i' && s[2] == 'v' && s[3] == 'e' {
			return true, 5
		}
	case 's':
		if s[1] == 'i' && s[2] == 'x' {
			return true, 6
		} else if l > 4 && s[1] == 'e' && s[2] == 'v' && s[3] == 'e' && s[4] == 'n' {
			return true, 7
		}
	case 'e':
		if l > 4 && s[1] == 'i' && s[2] == 'g' && s[3] == 'h' && s[4] == 't' {
			return true, 8
		}
	case 'n':
		if l > 3 && s[1] == 'i' && s[2] == 'n' && s[3] == 'e' {
			return true, 9
		}
	}

	// simple approach
	//if isDigit, digit := startWithDigit1(s); isDigit {
	//	return true, digit
	//}
	//var digits = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	//for i, d := range digits {
	//	if strings.HasPrefix(s, d) {
	//		return true, i
	//	}
	//}
	return false, 0
}

func solve(input string, startWithDigit func(string) (bool, int)) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var sum int
	for _, line := range lines {
		var first int = -1 // impossible value
		var last int
		for i := range line {
			if isDigit, digit := startWithDigit(line[i:]); isDigit {
				if first < 0 {
					first = digit
				}
				last = digit
			}
		}
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
