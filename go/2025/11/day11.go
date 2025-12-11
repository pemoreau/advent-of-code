package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func parse(input string) map[string][]string {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var graph = make(map[string][]string)

	for _, line := range lines {
		var s = strings.Split(line, " ")
		var key = s[0][:len(s[0])-1]
		graph[key] = s[1:]
	}
	return graph
}

type pair struct {
	a, b string
}

func solve(graph map[string][]string, start string, stop string, memo map[pair]int) int {
	if v, ok := memo[pair{start, stop}]; ok {
		return v
	}
	if start == stop {
		return 1
	}

	var res int
	for _, next := range graph[start] {
		res += solve(graph, next, stop, memo)
	}
	memo[pair{start, stop}] = res
	return res
}
func Part1(input string) int {
	var graph = parse(input)
	var memo = make(map[pair]int)
	return solve(graph, "you", "out", memo)
}

func Part2(input string) int {
	var graph = parse(input)
	var memo = make(map[pair]int)

	if b := solve(graph, "dac", "fft", memo); b != 0 {
		var a = solve(graph, "svr", "dac", memo)
		var c = solve(graph, "fft", "out", memo)
		return a * b * c
	}
	var a = solve(graph, "svr", "fft", memo)
	var b = solve(graph, "fft", "dac", memo)
	var c = solve(graph, "dac", "out", memo)
	return a * b * c
}

func main() {
	fmt.Println("--2025 day 11 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
