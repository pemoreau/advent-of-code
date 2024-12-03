package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"slices"
	"strconv"
	"strings"
	"time"
)

func parse(input string) (left []int, right []int) {
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var before, after, _ = strings.Cut(line, " ")
		var l, _ = strconv.Atoi(before)
		var r, _ = strconv.Atoi(strings.Trim(after, " "))
		left = append(left, l)
		right = append(right, r)
	}
	return
}

func Part1(input string) int {
	var left, right = parse(input)
	slices.Sort(left)
	slices.Sort(right)
	var res int
	for i := 0; i < len(left); i++ {
		res += utils.Abs(left[i] - right[i])
	}
	return res
}

func Part2(input string) int {
	var left, right = parse(input)
	var occurence = make(map[int]int)
	for _, v := range right {
		occurence[v]++
	}
	var res int
	for _, v := range left {
		n, ok := occurence[v]
		if ok && n > 0 {
			res += n * v
		}
	}
	return res
}

func main() {
	fmt.Println("--2024 day 01 solution--")

	var inputDay = utils.Input()

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
