package main

import (
	_ "embed"
	"fmt"
	"sort"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	list := utils.LinesToNumbers(input)
	sort.Ints(list)
	diff := map[int]int{1: 0, 2: 0, 3: 1}
	diff[list[0]]++
	for i, v := range list {
		if i > 0 {
			diff[v-list[i-1]] += 1
		}
	}
	return diff[1] * diff[3]
}

func Part2(input string) int {
	list := utils.LinesToNumbers(input)
	sort.Ints(list)
	list = append(list, list[len(list)-1]+3)
	list = append([]int{0}, list...)
	res := 1
	nb1 := 0
	for i, v := range list {
		if i > 0 {
			diff := v - list[i-1]
			if diff == 1 {
				nb1++
			}
			if diff == 3 {
				switch nb1 {
				// could have use tribonacci numbers
				case 4:
					res *= 7
				case 3:
					res *= 4
				case 2:
					res *= 2
				}
				nb1 = 0
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2020 day 10 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
