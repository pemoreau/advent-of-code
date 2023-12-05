package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

// ab: [a, b] seeds to transform
// i: [src, src + len-1] interval to transform
func transformer(ab, i utils.Interval, dest int) (transformed []utils.Interval, other []utils.Interval) {
	a := ab.Min
	b := ab.Max
	delta := dest - i.Min
	if ab.Inter(i) == utils.Empty() {
		return transformed, []utils.Interval{ab}
	}
	if a >= i.Min && b <= i.Max {
		return []utils.Interval{{a + delta, b + delta}}, other
	}
	if b >= i.Min && b < i.Max {
		// to avoid empty interval
		if a < i.Min {
			other = append(other, utils.Interval{a, i.Min - 1})
		}
		return []utils.Interval{{i.Min + delta, b + delta}}, other
	}
	if a > i.Min && a <= i.Max {
		if b > i.Max {
			other = append(other, utils.Interval{i.Max + 1, b})
		}
		return []utils.Interval{{a + delta, i.Max + delta}}, other
	}
	if a <= i.Min && i.Max <= b {
		// to avoid empty interval
		if a < i.Min {
			other = append(other, utils.Interval{a, i.Min - 1})
		}
		if i.Max < b {
			other = append(other, utils.Interval{i.Max + 1, b})
		}
		return []utils.Interval{{i.Min + delta, i.Max + delta}}, other
	}
	panic("not implemented")
}

type Rule struct {
	i    utils.Interval
	dest int
}

func applyOneRule(rule Rule, seeds []utils.Interval) (transformed []utils.Interval, other []utils.Interval) {
	for _, seed := range seeds {
		t, o := transformer(seed, rule.i, rule.dest)
		transformed = append(transformed, t...)
		other = append(other, o...)

	}
	return transformed, other
}

func apply(rules []Rule, ab utils.Interval) []utils.Interval {
	var res []utils.Interval
	var todo = []utils.Interval{ab}
	for _, rule := range rules {
		t, o := applyOneRule(rule, todo)
		res = append(res, t...)
		todo = o
	}
	res = append(res, todo...)
	return res
}

func solve(parts []string, seeds []utils.Interval) int {
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		lines = lines[1:]
		var rules []Rule
		for _, line := range lines {
			var dest, src, l int
			fmt.Sscanf(line, "%d %d %d", &dest, &src, &l)
			rules = append(rules, Rule{utils.Interval{src, src + l - 1}, dest})
		}
		var newSeeds []utils.Interval
		for _, seed := range seeds {
			newSeeds = append(newSeeds, apply(rules, seed)...)
		}
		seeds = newSeeds
	}

	var res = math.MaxInt32
	for _, seed := range seeds {
		res = min(res, seed.Min)
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	numbers, _ := strings.CutPrefix(parts[0], "seeds: ")

	var seeds []utils.Interval
	for _, v := range strings.Fields(numbers) {
		n, _ := strconv.Atoi(v)
		seeds = append(seeds, utils.Interval{n, n})
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

	var seeds []utils.Interval
	for i := 0; i < len(values); i += 2 {
		seeds = append(seeds, utils.Interval{values[i], values[i] + values[i+1] - 1})
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
