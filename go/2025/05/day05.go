package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/interval"
)

func isValid(n int, list []interval.Interval) bool {
	for _, i := range list {
		if i.Contains(n) {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var parts = strings.Split(input, "\n\n")
	var intervals []interval.Interval
	for _, r := range strings.Split(parts[0], "\n") {
		var a, b int
		_, _ = fmt.Sscanf(r, "%d-%d", &a, &b)
		intervals = append(intervals, interval.Interval{a, b})
	}
	var res int
	for _, r := range strings.Split(parts[1], "\n") {
		v, _ := strconv.Atoi(r)
		if isValid(v, intervals) {
			res++
		}
	}

	return res
}
func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var parts = strings.Split(input, "\n\n")
	var freespace = interval.FreeSpace{}
	for _, r := range strings.Split(parts[0], "\n") {
		var a, b int
		_, _ = fmt.Sscanf(r, "%d-%d", &a, &b)
		freespace.Add(interval.Interval{a, b})
	}
	freespace.Merge()
	return freespace.Cardinality()
}

func main() {
	fmt.Println("--2025 day 05 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
