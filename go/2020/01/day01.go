package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	values := utils.LinesToNumbers(input)
	for i1, v1 := range values {
		for i2, v2 := range values {
			if i1 != i2 && v1+v2 == 2020 {
				return v1 * v2
			}
		}
	}
	return 0
}

func Part2(input string) int {
	values := utils.LinesToNumbers(input)
	for i1, v1 := range values {
		for i2, v2 := range values {
			for i3, v3 := range values {
				if i1 != i2 && i2 != i3 && i1 != i3 && v1+v2+v3 == 2020 {
					return v1 * v2 * v3
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println("--2020 day 01 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part1: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
