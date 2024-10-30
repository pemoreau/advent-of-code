package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"slices"
	"strings"
	"time"
)

//go:embed input_test.txt
var inputDay string

type Graph map[uint8]struct {
	previous []uint8
	next     []uint8
}

func parseInput(input string) (map[uint8][]uint8, uint8) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	var res = make(map[uint8][]uint8)
	var appearsRhs = set.NewSet[uint8]()
	for _, line := range lines {
		// Step C must be finished before step A can begin.
		var step, nt uint8
		fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &step, &nt)
		res[nt] = append(res[nt], step)
		appearsRhs.Add(step)
	}
	var root uint8
	for k, v := range res {
		res[k] = sortAndRemoveDouble(v)
		if !appearsRhs.Contains(k) {
			root = k
		}
	}
	return res, root
}

func sortAndRemoveDouble(steps []uint8) []uint8 {
	var res = slices.Clone(steps)
	slices.Sort(res)
	for i := 0; i < len(res)-1; {
		if res[i] == res[i+1] {
			res = append(res[:i], res[i+1:]...)
		} else {
			i++
		}
	}
	return res
}

func generate(grammar map[uint8][]uint8, root uint8) string {
	var res = ""
	var todo = make([]uint8, 0)
	var visited = set.NewSet[uint8]()
	todo = append(todo, root)
	for len(todo) > 0 {
		var newTodo = make([]uint8, 0)
		for i := len(todo) - 1; i >= 0; i-- {
			if visited.Contains(todo[i]) {
				continue
			}
			res += string(todo[i])
			visited.Add(todo[i])
			newTodo = slices.Concat(newTodo, grammar[todo[i]])
		}
		fmt.Println(todo, newTodo)
		todo = sortAndRemoveDouble(newTodo)
	}
	return res
}

func Part1(input string) int {
	var grammar, root = parseInput(input)

	fmt.Println(grammar, root)
	fmt.Println(generate(grammar, root))
	return 0
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println("--2018 day 07 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
