package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/interval"
	"github.com/pemoreau/advent-of-code/go/utils/set"
)

func isValid(n int, list []interval.Interval) bool {
	for _, i := range list {
		if i.Contains(n) {
			return true
		}
	}
	return false
}

func parse(input string) ([]interval.Interval, int) {
	var ranges = strings.Split(input, ",")
	var intervals []interval.Interval
	var maxB int
	for _, r := range ranges {
		var a, b int
		_, _ = fmt.Sscanf(r, "%d-%d", &a, &b)
		intervals = append(intervals, interval.Interval{a, b})
		maxB = max(maxB, b)
	}
	return intervals, maxB
}

func shiftLeft(number int, n int) int {
	for range n {
		number *= 10
	}
	return number
}

func nbDigit(number int) int {
	if number == 0 {
		return 1
	}
	return 1 + int(math.Log10(float64(number)))
}

func concat(a, b int) int {
	return shiftLeft(a, nbDigit(b)) + b
}

func repeat(number int, maxDigit int) []int {
	var res []int
	var n = number
	var nbDigitNumber = nbDigit(number)
	var nbDigit = nbDigitNumber
	for nbDigit <= maxDigit {
		n = concat(n, number)
		nbDigit += nbDigitNumber
		if nbDigit <= maxDigit {
			res = append(res, n)
		}
	}
	return res
}

func Part1(input string) int {
	intervals, maxB := parse(input)
	var maxDigit = nbDigit(maxB)

	var res int
	var upper = shiftLeft(1, maxDigit/2)
	for i := range upper {
		if n := concat(i, i); isValid(n, intervals) {
			res = res + n
		}
	}
	return res
}

func Part2(input string) int {
	intervals, maxB := parse(input)
	var maxDigit = nbDigit(maxB)

	var upper = shiftLeft(1, maxDigit/2)
	var numbers = set.NewSet[int]()
	for i := range upper {
		var l = repeat(i, maxDigit)
		numbers.AddAll(l...)
	}

	var res int
	for i := range numbers {
		if isValid(i, intervals) {
			res += i
		}
	}

	return res
}

func main() {
	fmt.Println("--2025 day 02 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
