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

type state struct {
	pattern string
	numbers string
}

var cache = make(map[state]int)

func setCache(pattern string, numbers []uint8, value int) int {
	cache[state{pattern, string(numbers)}] = value
	return value
}

func count(pattern string, numbers []uint8) int {
	if len(pattern) == 0 && len(numbers) == 0 {
		return 1
	}

	if len(pattern) == 0 {
		return 0
	}

	// test cache
	if value, ok := cache[state{pattern, string(numbers)}]; ok {
		return value
	}

	if pattern[0] == '.' {
		res := count(pattern[1:], numbers)
		return setCache(pattern, numbers, res)
	}

	// cut branches: not very useful
	var sum int
	for _, n := range numbers {
		sum += int(n)
	}
	if len(pattern) < sum+len(numbers)-1 {
		res := 0
		return setCache(pattern, numbers, res)
	}

	if pattern[0] == '?' {
		res := count(pattern[1:], numbers) + count("#"+pattern[1:], numbers)
		return setCache(pattern, numbers, res)
	}

	if pattern[0] == '#' {
		if len(numbers) == 0 {
			res := 0
			return setCache(pattern, numbers, res)
		}

		n := numbers[0]
		indexDot := strings.Index(pattern, ".")
		if indexDot == -1 {
			indexDot = len(pattern)
		}
		if indexDot < int(n) {
			// not enough # or ?
			res := 0
			return setCache(pattern, numbers, res)
		}

		// eat n # or ?
		remaining := pattern[n:]
		if len(remaining) == 0 {
			res := count(remaining, numbers[1:])
			return setCache(pattern, numbers, res)
		}

		if remaining[0] == '#' {
			// fail: it should be a .
			res := 0
			return setCache(pattern, numbers, res)
		}
		// remaining[0] == '.' || remaining[0] == '?': skip first ? since it should be a .
		res := count(remaining[1:], numbers[1:])
		return setCache(pattern, numbers, res)
	}
	panic("unreachable")
}

func unfold(pattern string, numbers []uint8) (string, []uint8) {
	var resPattern = pattern + "?" + pattern + "?" + pattern + "?" + pattern + "?" + pattern
	var resNumbers = append(numbers, numbers...)
	resNumbers = append(resNumbers, resNumbers...)
	resNumbers = append(resNumbers, numbers...)
	return resPattern, resNumbers
}

func solve(input string, part2 bool) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var res int
	for _, line := range lines {
		fields := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == ',' })
		var pattern = fields[0]
		var numbers []uint8
		for _, field := range fields[1:] {
			numbers = append(numbers, uint8(utils.ToInt(field)))
		}
		if part2 {
			res += count(unfold(pattern, numbers))
		} else {
			res += count(pattern, numbers)
		}
	}

	return res
}

func Part1(input string) int {
	clear(cache)
	return solve(input, false)
}

func Part2(input string) int {
	clear(cache)
	return solve(input, true)
}

func main() {
	fmt.Println("--2023 day 12 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
