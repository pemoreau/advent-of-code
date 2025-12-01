package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
)

func redistribution(banks []int) {
	var index int
	var maxValue int
	for i, v := range banks {
		if v > maxValue {
			maxValue = v
			index = i
		}
	}
	banks[index] = 0
	for maxValue > 0 {
		index = (index + 1) % len(banks)
		banks[index]++
		maxValue--
	}
}

func solve(input string, part2 bool) int {
	var fields = strings.Fields(input)
	var banks []int
	var res int
	for _, field := range fields {
		v, _ := strconv.Atoi(field)
		banks = append(banks, v)
	}
	var visited = map[string]int{}
	var s = fmt.Sprintf("%v", banks)
	visited[s] = res
	for {
		redistribution(banks)
		s = fmt.Sprintf("%v", banks)
		res++
		if v, ok := visited[s]; ok {
			if part2 {
				return res - v
			} else {
				return res
			}
		}
		visited[s] = res
	}
}
func Part1(input string) string {
	lines := strings.Split(input, "\n")
	var left []string
	//var right [][]string
	var right = set.NewSet[string]()

	for _, line := range lines {
		if strings.Contains(line, "->") {
			var fields = strings.Fields(line)
			left = append(left, fields[0])
			//right = append(right, fields[3:])
			for _, field := range fields[3:] {
				before, _ := strings.CutSuffix(field, ",")
				right.Add(before)
			}
		}
	}
	for _, l := range left {
		if !right.Contains(l) {
			return l
		}
	}
	return ""
}

func Part2(input string) string {
	return ""
}

func main() {
	fmt.Println("--2017 day 06 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
