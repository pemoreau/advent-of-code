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

//go:embed sample.txt
var inputTest string

func check(goal int, elements []int) int {
	if len(elements) == 1 && elements[0] == goal {
		return 1
	}
	if len(elements) < 2 {
		return 0
	}

	var res1, res2 int
	var head = elements[0]
	var tail = elements[1:]
	var subgoal1 = goal - head
	var subgoal2 = goal / head
	if subgoal1 >= 0 {
		res1 = check(subgoal1, tail)
	}
	if subgoal2*head == goal {
		res2 = check(subgoal2, tail)
	}
	return res1 + res2
}

func reverse(elements []int) []int {
	var res []int
	for i := len(elements) - 1; i >= 0; i-- {
		res = append(res, elements[i])
	}
	return res
}

func Part1(input string) int {
	var lines = strings.Split(input, "\n")

	var res int
	for _, line := range lines {
		var before, after, _ = strings.Cut(line, ":")
		var goal, _ = strconv.Atoi(before)
		var elements []int
		for _, e := range strings.Fields(after) {
			var el, _ = strconv.Atoi(e)
			elements = append(elements, el)
		}
		var n = check(goal, reverse(elements))
		//fmt.Printf("goal: %d, elements: %v: %d\n", goal, elements, n)
		if n > 0 {
			res += goal
		}

	}

	return res
}

func check2(goal int, elements []int) bool {
	if len(elements) == 1 && elements[0] == goal {
		return true
	}
	if len(elements) == 1 {
		return false
	}

	var res1, res2, res3 bool
	var head = elements[0]
	var tail = elements[1:]
	var subgoal1 = goal - head
	var subgoal2 = goal / head
	var s = strconv.Itoa(head)
	var p = int(math.Pow(10, float64(len(s))))
	fmt.Printf("p=%d\n", p)
	var subgoal3 = (goal - head) / p

	if subgoal1 >= 0 {
		res1 = check2(subgoal1, tail)
	}
	if subgoal2*head == goal {
		res2 = check2(subgoal2, tail)
	}
	if subgoal3*p+head == goal {
		res3 = check2(subgoal3, tail)
	}
	return res1 || res2 || res3
}

func Part2(input string) int {
	var lines = strings.Split(input, "\n")

	var res int
	for _, line := range lines {
		var before, after, _ = strings.Cut(line, ":")
		var goal, _ = strconv.Atoi(before)
		var elements []int
		for _, e := range strings.Fields(after) {
			var el, _ = strconv.Atoi(e)
			elements = append(elements, el)
		}
		var n = check2(goal, reverse(elements))
		fmt.Printf("goal: %d, elements: %v: %v\n", goal, elements, n)
		if n {
			res += goal
		}

	}

	return res
}

func main() {
	fmt.Println("--2024 day 07 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
