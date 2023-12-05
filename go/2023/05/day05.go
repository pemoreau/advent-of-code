package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Interval struct{ Min, Max int }

var EmptyInterval = Interval{0, -1}

// ab: [a, b] seeds to transform
// i: [src, src + len-1] interval to transform
func transformer(ab, i Interval, dest int) (transformed Interval, other []Interval) {
	a := ab.Min
	b := ab.Max
	delta := dest - i.Min
	if b < i.Min || a > i.Max {
		return EmptyInterval, []Interval{ab}
	}
	if a >= i.Min && b <= i.Max {
		return Interval{a + delta, b + delta}, other
	}
	if b >= i.Min && b < i.Max {
		// to avoid empty interval
		if a < i.Min {
			other = append(other, Interval{a, i.Min - 1})
		}
		return Interval{i.Min + delta, b + delta}, other
	}
	if a > i.Min && a <= i.Max {
		if b > i.Max {
			other = append(other, Interval{i.Max + 1, b})
		}
		return Interval{a + delta, i.Max + delta}, other
	}
	if a <= i.Min && i.Max <= b {
		// to avoid empty interval
		if a < i.Min {
			other = append(other, Interval{a, i.Min - 1})
		}
		if i.Max < b {
			other = append(other, Interval{i.Max + 1, b})
		}
		return Interval{i.Min + delta, i.Max + delta}, other
	}
	panic("not implemented")
}

type Rule struct {
	i    Interval
	dest int
}

func applyOneRule(rule Rule, seeds []Interval) (transformed []Interval, other []Interval) {
	for _, seed := range seeds {
		t, o := transformer(seed, rule.i, rule.dest)
		if t != EmptyInterval {
			transformed = append(transformed, t)
		}
		other = append(other, o...)
	}
	return transformed, other
}

func apply(rules []Rule, ab Interval) []Interval {
	var res []Interval
	var todo = []Interval{ab}
	for _, rule := range rules {
		t, o := applyOneRule(rule, todo)
		res = append(res, t...)
		todo = o
		if len(todo) == 0 {
			return res
		}
	}
	res = append(res, todo...)
	return res
}

func solve(parts []string, seeds []Interval) int {
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		lines = lines[1:]
		var rules []Rule
		for _, line := range lines {
			var dest, src, l int
			fmt.Sscanf(line, "%d %d %d", &dest, &src, &l)
			rules = append(rules, Rule{Interval{src, src + l - 1}, dest})
		}
		var newSeeds []Interval
		for _, seed := range seeds {
			newSeeds = append(newSeeds, apply(rules, seed)...)
		}
		seeds = newSeeds
	}

	var res = math.MaxInt64
	for _, seed := range seeds {
		res = min(res, seed.Min)
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	numbers, _ := strings.CutPrefix(parts[0], "seeds: ")

	var seeds []Interval
	for _, v := range strings.Fields(numbers) {
		n, _ := strconv.Atoi(v)
		seeds = append(seeds, Interval{n, n})
	}
	return solve(parts[1:], seeds)
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	numbers, _ := strings.CutPrefix(parts[0], "seeds: ")

	var values []int
	for _, v := range strings.Fields(numbers) {
		if n, err := strconv.Atoi(v); err == nil {
			values = append(values, n)
		}
	}

	var seeds []Interval
	for i := 0; i < len(values); i += 2 {
		seeds = append(seeds, Interval{values[i], values[i] + values[i+1] - 1})
	}
	return solve(parts[1:], seeds)
}

func main() {
	fmt.Println("--2023 day 05 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
