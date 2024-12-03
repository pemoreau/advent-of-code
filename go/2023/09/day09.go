package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type history []int

func (h history) allZero() bool {
	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}

func derive(h history) []history {
	var res = []history{h}
	var last = h
	for !last.allZero() {
		var next history = make([]int, 0, len(last)-1)
		for i := 1; i < len(last); i++ {
			next = append(next, last[i]-last[i-1])
		}
		res = append(res, next)
		last = next
	}
	return res
}

func nextValue(d []history) int {
	var res int
	for i := len(d) - 1; i >= 0; i-- {
		res += d[i][len(d[i])-1]
	}
	return res
}

func nextValue2(d []history) int {
	var res int
	for i := len(d) - 2; i >= 0; i-- {
		res = d[i][0] - res
	}
	return res
}

func solve(input string, nextFunc func([]history) int) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var h history = make([]int, 0)
	var res int
	for _, line := range lines {
		h = h[:0] // set len to 0
		for _, v := range strings.Fields(line) {
			n, _ := strconv.Atoi(v)
			h = append(h, n)
		}
		res += nextFunc(derive(h))
	}

	return res
}

func Part1(input string) int {
	return solve(input, nextValue)
}

func Part2(input string) int {
	return solve(input, nextValue2)
}

func main() {
	fmt.Println("--2023 day 09 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
