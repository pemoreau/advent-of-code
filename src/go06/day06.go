package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func simulate(values []int, n int) int {
	mult := make([]int, 9)
	for _, v := range values {
		mult[v] = mult[v] + 1
	}

	for i := 0; i < n; i++ {
		// rotate array left
		mult = append(mult[1:], mult[0])
		mult[6] += mult[8]

	}

	return count(mult)
}

func count(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
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
	return simulate(values, 80)
}

func Part2(input string) int {
	values := make([]int, 0)
	for _, in := range strings.Split(strings.TrimSuffix(input, "\n"), ",") {
		v, _ := strconv.Atoi(in)
		values = append(values, v)
	}
	return simulate(values, 256)
}

func main() {
	content, _ := ioutil.ReadFile("../../inputs/day06.txt")

	start := time.Now()
	fmt.Println("part1: ", Part1(string(content)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(content)))
	fmt.Println(time.Since(start))
}
