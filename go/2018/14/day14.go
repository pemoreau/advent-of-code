package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"time"
)

func computeRecipes(n int) []int {
	recipes := []int{3, 7}
	elf1, elf2 := 0, 1

	for len(recipes) < n+10 {
		sum := recipes[elf1] + recipes[elf2]
		if sum >= 10 {
			recipes = append(recipes, sum/10, sum%10)
		} else {
			recipes = append(recipes, sum)
		}

		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}

	return recipes[n : n+10]
}

func Part1(input string) string {
	var res = computeRecipes(640441)
	return fmt.Sprintf(("%c%c%c%c%c%c%c%c%c%c"), res[0]+'0', res[1]+'0', res[2]+'0', res[3]+'0', res[4]+'0', res[5]+'0', res[6]+'0', res[7]+'0', res[8]+'0', res[9]+'0')
}

func Part2(input string) int {
	var target = []int{6, 4, 0, 4, 4, 1}
	//var target = []int{5, 9, 4, 1, 4}
	recipes := []int{3, 7}
	elf1, elf2 := 0, 1

	for i := 0; ; i++ {
		sum := recipes[elf1] + recipes[elf2]
		if sum >= 10 {
			recipes = append(recipes, sum/10, sum%10)
		} else {
			recipes = append(recipes, sum)
		}
		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)

		if recipes[i] == target[0] && recipes[i+1] == target[1] && recipes[i+2] == target[2] && recipes[i+3] == target[3] && recipes[i+4] == target[4] {
			return i
		}

	}
}

func main() {
	fmt.Println("--2018 day 14 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
