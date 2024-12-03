package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"iter"
	"slices"
	"strings"
	"time"
)

func parse(input string) [][]int {
	var lines = strings.Split(input, "\n")
	var res [][]int
	for _, line := range lines {
		var parts = strings.Split(line, " ")
		var reports []int
		for _, part := range parts {
			var report int
			fmt.Sscanf(part, "%d", &report)
			reports = append(reports, report)
		}
		res = append(res, reports)
	}
	return res
}

func safeit(it iter.Seq2[int, int]) bool {
	var previous int
	var i = 0
	for _, e := range it {
		if i == 0 {
			previous = e
			i++
			continue
		}
		if e-previous > 3 || e-previous < 1 {
			return false
		}
		previous = e
		i++
	}
	return true

	//var dir = reports[1] - reports[0]
	//for i := 1; i < len(reports); i++ {
	//	var diff int
	//	if dir > 0 {
	//		diff = reports[i] - reports[i-1]
	//	} else {
	//		diff = reports[i-1] - reports[i]
	//	}
	//	if diff >= 1 && diff <= 3 {
	//		continue
	//	}
	//	return false
	//}
	//return true
}

func filterIndex(iter iter.Seq2[int, int], index int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for k, e := range iter {
			if k != index && !yield(k, e) {
				return
			}
		}
	}
}

func solve(input string, safe func([]int) bool) int {
	var reports = parse(input)
	var res int
	for _, report := range reports {
		if safe(report) {
			res++
		}
	}
	return res
}

func Part1(input string) int {
	safe := func(reports []int) bool {
		return safeit(slices.All(reports)) || safeit(slices.Backward(reports))
	}
	return solve(input, safe)
}

func Part2(input string) int {
	safe := func(reports []int) bool {
		for i := 0; i < len(reports); i++ {
			if safeit(filterIndex(slices.All(reports), i)) ||
				safeit(filterIndex(slices.Backward(reports), i)) {
				return true
			}
		}
		return false
	}

	return solve(input, safe)
}

func main() {
	fmt.Println("--2024 day 02 solution--")

	var inputDay = utils.Input()

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))

}
