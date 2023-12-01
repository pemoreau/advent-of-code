package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func isDigit(s string, i int) (bool, int) {
	if i >= len(s) {
		return false, 0
	}
	if s[i] >= '0' && s[i] <= '9' {
		return true, int(s[i] - '0')
	}
	return false, 0
}

func isDigitLetter(s string, i int) (bool, int) {
	l := len(s) - i
	if l == 0 {
		return false, 0
	}
	c := s[i]
	if c >= '0' && c <= '9' {
		return true, int(c - '0')
	}
	if l <= 2 {
		return false, 0
	}
	// discrimination net to improve performance
	// zero one two three four five six seven eight nine
	i1 := i + 1
	i2 := i + 2
	i3 := i + 3
	i4 := i + 4
	switch c {
	case 'z':
		if l > 3 && s[i1] == 'e' && s[i2] == 'r' && s[i3] == 'o' {
			return true, 0
		}
	case 'o':
		if s[i1] == 'n' && s[i2] == 'e' {
			return true, 1
		}
	case 't':
		if s[i1] == 'w' {
			if s[i2] == 'o' {
				return true, 2
			}
		} else if l > 4 && s[i1] == 'h' && s[i2] == 'r' && s[i3] == 'e' && s[i4] == 'e' {
			return true, 3
		}
	case 'f':
		if l > 3 {
			if s[i1] == 'o' {
				if s[i2] == 'u' {
					if s[i3] == 'r' {
						return true, 4
					}
				}
			} else if s[i1] == 'i' && s[i2] == 'v' && s[i3] == 'e' {
				return true, 5
			}
		}
	case 's':
		if s[i1] == 'i' {
			if s[i2] == 'x' {
				return true, 6
			}
		} else if l > 4 && s[i1] == 'e' && s[i2] == 'v' && s[i3] == 'e' && s[i4] == 'n' {
			return true, 7
		}
	case 'e':
		if l > 4 && s[i1] == 'i' && s[i2] == 'g' && s[i3] == 'h' && s[i4] == 't' {
			return true, 8
		}
	case 'n':
		if l > 3 && s[i1] == 'i' && s[i2] == 'n' && s[i3] == 'e' {
			return true, 9
		}
	}

	// simple approach
	//if isDigit, digit := isDigit(s); isDigit {
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

func isDigitReverseLetter(s string, i int) (bool, int) {
	l := len(s)
	if l == 0 {
		return false, 0
	}
	c := s[i]
	if c >= '0' && c <= '9' {
		return true, int(c - '0')
	}
	if l <= 2 {
		return false, 0
	}
	// discrimination net to improve performance
	// zero one two three four five six seven eight nine
	// orez eno owt eerht ruof evif xis neves thgie enin
	// orez owt eno eerht evif enin ruof xis neves thgie
	i1 := i - 1
	i2 := i - 2
	i3 := i - 3
	i4 := i - 4

	switch c {
	case 'o':
		if l > 3 && s[i1] == 'r' && s[i2] == 'e' && s[i3] == 'z' {
			return true, 0
		} else if l > 2 && s[i1] == 'w' && s[i2] == 't' {
			return true, 2
		}
	case 'e':
		if l > 2 && s[i1] == 'n' && s[i2] == 'o' {
			return true, 1
		} else if l > 4 && s[i1] == 'e' && s[i2] == 'r' && s[i3] == 'h' && s[i4] == 't' {
			return true, 3
		} else if l > 3 && s[i1] == 'v' && s[i2] == 'i' && s[i3] == 'f' {
			return true, 5
		} else if l > 3 && s[i1] == 'n' && s[i2] == 'i' && s[i3] == 'n' {
			return true, 9
		}
	case 'r':
		if l > 3 && s[i1] == 'u' && s[i2] == 'o' && s[i3] == 'f' {
			return true, 4
		}
	case 'x':
		if l > 2 && s[i1] == 'i' && s[i2] == 's' {
			return true, 6
		}
	case 'n':
		if l > 4 && s[i1] == 'e' && s[i2] == 'v' && s[i3] == 'e' && s[i4] == 's' {
			return true, 7
		}
	case 't':
		if l > 4 && s[i1] == 'h' && s[i2] == 'g' && s[i3] == 'i' && s[i4] == 'e' {
			return true, 8
		}
	}

	// simple approach
	//if isDigit, digit := isDigit(s); isDigit {
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

type isDigitFunc func(string, int) (bool, int)

func solve(input string, isDigit isDigitFunc, isDigitReverse isDigitFunc) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var sum int
	for _, line := range lines {
		var first int
		for i := range line {
			if ok, digit := isDigit(line, i); ok {
				first = digit
				break
			}
		}
		var last int
		for i := len(line) - 1; i >= 0; i-- {
			if ok, digit := isDigitReverse(line, i); ok {
				last = digit
				break
			}
		}
		sum += 10*first + last
	}
	return sum
}

func Part1(input string) int {
	return solve(input, isDigit, isDigit)
}

func Part2(input string) int {
	return solve(input, isDigitLetter, isDigitReverseLetter)
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
