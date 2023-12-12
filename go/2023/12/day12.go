package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type state struct {
	pattern string
	numbers string
}

var cache = make(map[state]int)

func setCache(pattern string, numbers []int, value int) int {
	cache[state{pattern, fmt.Sprint(numbers)}] = value
	return value
}

func count(pattern string, numbers []int) int {
	if len(pattern) == 0 && len(numbers) == 0 {
		return 1
	}

	if len(pattern) == 0 {
		return 0
	}

	// test cache
	if value, ok := cache[state{pattern, fmt.Sprint(numbers)}]; ok {
		return value
	}

	if pattern[0] == '.' {
		res := count(pattern[1:], numbers)
		return setCache(pattern, numbers, res)
	}

	var sum int
	for _, n := range numbers {
		sum += n
	}
	if len(pattern) < sum {
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
		if indexDot < n {
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
			// fail
			res := 0
			return setCache(pattern, numbers, res)
		}
		if remaining[0] == '.' {
			res := count(remaining[1:], numbers[1:])
			return setCache(pattern, numbers, res)
		}
		if remaining[0] == '?' {
			// eat first ? since it should be a .
			res := count(remaining[1:], numbers[1:])
			return setCache(pattern, numbers, res)
		}
	}
	panic("unreachable")
}

func solve(input string, unfold bool) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var res int
	for _, line := range lines {
		fields := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == ',' })
		var pattern = fields[0]
		var numbers []int
		for _, field := range fields[1:] {
			numbers = append(numbers, utils.ToInt(field))
		}
		if unfold {
			pattern = unfoldPattern(pattern)
			numbers = unfoldNumbers(numbers)
		}
		v := count(pattern, numbers)
		//fmt.Println(pattern, v)
		res += v
	}

	return res
}

func unfoldPattern(pattern string) string {
	var res = pattern
	for i := 0; i < 4; i++ {
		res = res + "?" + pattern
	}
	return res
}

func unfoldNumbers(numbers []int) []int {
	var res []int
	for i := 0; i < 5; i++ {
		res = append(res, numbers...)
	}
	return res
}

func Part1(input string) int {
	return solve(input, false)
}

func Part2(input string) int {
	return solve(input, true)
	//input = strings.TrimSuffix(input, "\n")
	//lines := strings.Split(input, "\n")
	//
	//var res int
	//for _, line := range lines {
	//	clear(cache)
	//	fields := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == ',' })
	//	var pattern = fields[0]
	//	var numbers []int
	//	for _, field := range fields[1:] {
	//		numbers = append(numbers, utils.ToInt(field))
	//	}
	//	v := count(unfoldPattern(pattern), unfoldNumbers(numbers))
	//	fmt.Println(unfoldPattern(pattern), unfoldNumbers(numbers), v)
	//	res += v
	//}
	//
	//return res
}

func main() {
	fmt.Println("--2023 day 11 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
