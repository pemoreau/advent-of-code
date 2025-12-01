package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var res int
	var dial = 50
	for _, order := range lines {
		v, _ := strconv.Atoi(order[1:])
		v = v % 100
		if order[0] == 'L' {
			v = -v
		}
		dial = (dial + 100 + v) % 100
		if dial == 0 {
			res = res + 1
		}
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var res int
	var dial = 50
	for _, order := range lines {
		v, _ := strconv.Atoi(order[1:])
		if v >= 100 {
			var n = v / 100
			res = res + n
			v = v % 100
		}
		if order[0] == 'L' {
			v = -v
		}
		if dial != 0 && (dial+v <= 0 || dial+v >= 100) {
			res = res + 1
		}
		dial = (dial + 100 + v) % 100
	}
	return res
}

func Part2Naive(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var res int
	var dial = 50
	for _, order := range lines {
		v, _ := strconv.Atoi(order[1:])
		for range v {
			if order[0] == 'L' {
				dial = (dial + 100 - 1) % 100
			} else {
				dial = (dial + 1) % 100
			}
			if dial == 0 {
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2025 day 01 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
