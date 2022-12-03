package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
	"unicode"
)

//go:embed input.txt
var input_day string

func set(s string) utils.Set[uint8] {
	res := utils.BuildSet[uint8]()
	for _, c := range s {
		res.Add(uint8(c))
	}
	return res
}

func score(i utils.Set[uint8]) int {
	res := 0
	for c := range i {
		if unicode.IsLower(rune(c)) {
			res += int(1 + c - 'a')
		} else {
			res += int(27 + c - 'A')
		}
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		i := set(line[:len(line)/2]).Intersect(set(line[len(line)/2:]))
		res += score(i)
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for l := 0; l < len(lines); l += 3 {
		i := set(lines[l]).Intersect(set(lines[l+1])).Intersect(set(lines[l+2]))
		res += score(i)
	}
	return res
}

func main() {
	fmt.Println("--2022 day 03 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
