package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

//go:embed sample.txt
var inputTest string

func getList(line string) []int {
	parts := strings.Fields(line)
	var list []int
	for _, s := range parts {
		n, _ := strconv.Atoi(s)
		list = append(list, n)
	}
	return list
}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		list := getList(line)
		minValue := math.MaxInt
		maxValue := math.MinInt
		for _, n := range list {
			if n < minValue {
				minValue = n
			}
			if n > maxValue {
				maxValue = n
			}
		}
		res += maxValue - minValue
	}
	return res
}

func isDivisor(n int, list []int) int {
	for _, v := range list {
		if v%n == 0 && v != n {
			return v / n
		}
	}
	return 0
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		list := getList(line)
		for _, n := range list {
			d := isDivisor(n, list)
			if d > 0 {
				res += d
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2017 day 02 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
