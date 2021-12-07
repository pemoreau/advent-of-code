package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

type fn func(int, int) int

func search(values []int, cost fn) int {
	minValue, maxValue := math.MaxInt, 0
	for _, value := range values {
		if value < minValue {
			minValue = value
		}
		if value > maxValue {
			maxValue = value
		}
	}

	minSum := math.MaxInt

	for h := minValue; h <= maxValue; h++ {
		sum := 0
		for _, value := range values {
			sum += cost(h, value)
		}
		if sum < minSum {
			minSum = sum
		}
	}
	return minSum
}

func Part1(input string) int {
	values := make([]int, 0)
	for _, in := range strings.Split(strings.TrimSuffix(input, "\n"), ",") {
		v, err := strconv.Atoi(in)
		if err != nil {
			panic(err)
		}
		values = append(values, v)
	}
	return search(values, func(a, b int) int {
		return abs(a - b)
	})
}

func Part2(input string) int {
	values := make([]int, 0)
	for _, in := range strings.Split(strings.TrimSuffix(input, "\n"), ",") {
		v, _ := strconv.Atoi(in)
		values = append(values, v)
	}
	return search(values, func(a, b int) int {
		n := abs(a - b)
		return n * (n + 1) / 2
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	content, _ := ioutil.ReadFile("../../inputs/day07.txt")

	start := time.Now()
	fmt.Println("part1: ", Part1(string(content)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(content)))
	fmt.Println(time.Since(start))
}
