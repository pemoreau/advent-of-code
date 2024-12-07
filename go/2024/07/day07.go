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

//go:embed sample.txt
var inputTest string

func solve(input string, part2 bool) int {
	var lines = strings.Split(input, "\n")

	var res int
	for _, line := range lines {
		var before, after, _ = strings.Cut(line, ":")
		var goal, _ = strconv.Atoi(before)
		var elements []int
		for _, e := range slices.Backward(strings.Fields(after)) {
			var el, _ = strconv.Atoi(e)
			elements = append(elements, el)
		}
		if check(goal, elements, part2) {
			res += goal
		}
	}
	return res
}

func check(goal int, elements []int, part2 bool) bool {
	if len(elements) == 1 && elements[0] == goal {
		return true
	}
	if len(elements) == 1 {
		return false
	}

	var head = elements[0]
	var tail = elements[1:]
	var subgoal1 = goal - head
	var subgoal2 = goal / head

	if subgoal1 >= 0 && check(subgoal1, tail, part2) {
		return true
	}
	if subgoal2*head == goal && check(subgoal2, tail, part2) {
		return true
	}
	if part2 {
		// p = 10^(len(head))
		var h = head
		var p = 1
		for h > 0 {
			h /= 10
			p *= 10
		}
		var subgoal3 = (goal - head) / p
		if subgoal3*p+head == goal && check(subgoal3, tail, part2) {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	return solve(input, false)
}

func Part2(input string) int {
	return solve(input, true)
}

func main() {
	fmt.Println("--2024 day 07 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
