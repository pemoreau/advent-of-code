package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func sum(numbers []int, index int, accu int) (int, int) {
	var nbChild = numbers[index]
	var nbMeta = numbers[index+1]
	index += 2
	for range nbChild {
		index, accu = sum(numbers, index, accu)
	}
	for range nbMeta {
		accu += numbers[index]
		index++
	}
	return index, accu
}

func value(numbers []int, index int) (int, int) {
	var nbChild = numbers[index]
	var nbMeta = numbers[index+1]
	index += 2
	if nbChild == 0 {
		var accu = 0
		for i := 0; i < nbMeta; i++ {
			accu += numbers[index]
			index++
		}
		return index, accu
	}

	var childs []int
	for range nbChild {
		var accu int
		index, accu = value(numbers, index)
		childs = append(childs, accu)
	}

	var accu = 0
	for range nbMeta {
		var meta = numbers[index]
		index++
		if meta > 0 && meta <= nbChild {
			accu += childs[meta-1]
		}
	}

	return index, accu
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, " ")
	var numbers = utils.StringsToNumbers(parts)
	_, res := sum(numbers, 0, 0)
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, " ")
	var numbers = utils.StringsToNumbers(parts)
	_, res := value(numbers, 0)
	return res
}

func main() {
	fmt.Println("--2018 day 08 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
