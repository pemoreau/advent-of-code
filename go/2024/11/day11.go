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

type config struct {
	v int
	n int
}

func produce(value int, n int, cache map[config]int) int {
	if n == 0 {
		return 1
	}

	if r, ok := cache[config{value, n}]; ok {
		return r
	}

	if value == 0 {
		res := produce(1, n-1, cache)
		cache[config{value, n}] = res
		return res
	}

	if s := strconv.Itoa(value); len(s)%2 == 0 {
		a, _ := strconv.Atoi(s[:len(s)/2])
		b, _ := strconv.Atoi(s[len(s)/2:])
		res := produce(a, n-1, cache) + produce(b, n-1, cache)
		cache[config{value, n}] = res
		return res
	}

	res := produce(value*2024, n-1, cache)
	cache[config{value, n}] = res
	return res
}

func solve(input string, n int) int {
	var cache = make(map[config]int)

	var line = strings.Split(input, " ")
	var res int
	for _, e := range line {
		v, _ := strconv.Atoi(e)
		res += produce(v, n, cache)
	}
	return res
}

func Part1(input string) int {
	return solve(input, 25)
}

func Part2(input string) int {
	return solve(input, 75)
}

func main() {
	fmt.Println("--2024 day 11 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
