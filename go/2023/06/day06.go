package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func buildArray(line string) []int {
	var res []int
	for _, field := range strings.Fields(line) {
		if v, err := strconv.Atoi(strings.TrimSpace(field)); err == nil {
			res = append(res, v)
		}
	}
	return res
}

func calcDist(hold, maxTime int) int {
	return (maxTime - hold) * hold
}

func solve(time, dist int) int {
	res := 0
	for i := 0; i < time; i++ {
		d := calcDist(i, time)
		if d > dist {
			res++
		}
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var times = buildArray(lines[0])
	var dist = buildArray(lines[1])

	res := 1
	for i, t := range times {
		res *= solve(t, dist[i])
	}
	return res
}

func buildConcat(line string) int {
	line = strings.ReplaceAll(line, " ", "")
	_, line, _ = strings.Cut(line, ":")
	res, _ := strconv.Atoi(line)
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	return solve(buildConcat(lines[0]), buildConcat(lines[1]))
}

func main() {
	fmt.Println("--2023 day 06 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
