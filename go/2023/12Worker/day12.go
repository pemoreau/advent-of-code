package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"runtime"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type state struct {
	pattern string
	numbers string
}

func (s state) String() string {
	var n string
	for _, c := range s.numbers {
		n += fmt.Sprintf("%d ", c)
	}
	return fmt.Sprintf("%s [%s]", s.pattern, n)
}

type cache map[state]int

//var cache = make(map[state]int)

func (c cache) setCache(pattern string, numbers []uint8, value int) int {
	c[state{pattern, string(numbers)}] = value
	return value
}

func count(c cache, pattern string, numbers []uint8) int {
	if len(pattern) == 0 && len(numbers) == 0 {
		return 1
	}

	if len(pattern) == 0 {
		return 0
	}

	// test cache
	if value, ok := c[state{pattern, string(numbers)}]; ok {
		return value
	}

	if pattern[0] == '.' {
		res := count(c, pattern[1:], numbers)
		return c.setCache(pattern, numbers, res)
	}

	// cut branches: not very useful
	var sum int
	for _, n := range numbers {
		sum += int(n)
	}
	if len(pattern) < sum+len(numbers)-1 {
		res := 0
		return c.setCache(pattern, numbers, res)
	}

	if pattern[0] == '?' {
		res := count(c, pattern[1:], numbers) + count(c, "#"+pattern[1:], numbers)
		return c.setCache(pattern, numbers, res)
	}

	if pattern[0] == '#' {
		if len(numbers) == 0 {
			res := 0
			return c.setCache(pattern, numbers, res)
		}

		n := numbers[0]
		indexDot := strings.Index(pattern, ".")
		if indexDot == -1 {
			indexDot = len(pattern)
		}
		if indexDot < int(n) {
			// not enough # or ?
			res := 0
			return c.setCache(pattern, numbers, res)
		}

		// eat n # or ?
		remaining := pattern[n:]
		if len(remaining) == 0 {
			res := count(c, remaining, numbers[1:])
			return c.setCache(pattern, numbers, res)
		}

		if remaining[0] == '#' {
			// fail: it should be a .
			res := 0
			return c.setCache(pattern, numbers, res)
		}
		// remaining[0] == '.' || remaining[0] == '?': skip first ? since it should be a .
		res := count(c, remaining[1:], numbers[1:])
		return c.setCache(pattern, numbers, res)
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

func solve(c cache, input string, part2 bool) int {
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
			pattern, numbers = unfold(pattern, numbers)
		}
		res += count(c, pattern, numbers)
	}

	return res
}

func Part1(input string) int {
	var c = make(cache)
	return solve(c, input, false)
}

func producer(input string, part2 bool) (<-chan state, int) {
	nbThreads := runtime.NumCPU()
	resultStream := make(chan state, nbThreads)
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	go func() {
		defer close(resultStream)

		for _, line := range lines {
			fields := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == ',' })
			var pattern = fields[0]
			var numbers []uint8
			for _, field := range fields[1:] {
				numbers = append(numbers, uint8(utils.ToInt(field)))
			}
			if part2 {
				pattern, numbers = unfold(pattern, numbers)
			}

			s := state{pattern, string(numbers)}
			select {
			case resultStream <- s:
			}
		}
	}()
	return resultStream, len(lines)
}

func worker(id int, stream <-chan state, results chan<- int) {
	var c = make(cache)

	for m := range stream {
		select {
		default:
			//fmt.Printf("Worker %d received: %v\n", id, m)
			res := count(c, m.pattern, []uint8(m.numbers))
			//fmt.Printf("Worker %d send:     %d\n", id, res)
			results <- res
			//fmt.Printf("Worker %d done:     %v\n", id, m)
		}
	}
	//fmt.Println("stream EMPTY", id)
}

func Part2(input string) int {
	nbThreads := runtime.NumCPU()
	results := make(chan int, nbThreads)
	defer close(results)

	stream, n := producer(input, true)
	for w := 1; w <= nbThreads; w++ {
		go worker(w, stream, results)
	}

	var res int
	for n > 0 {
		res += <-results
		n--
	}
	return res
}

func main() {
	fmt.Println("--2023 day 12 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
