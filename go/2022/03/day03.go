package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	set2 "github.com/pemoreau/advent-of-code/go/utils/set"
	"strings"
	"time"
	"unicode"
)

//go:embed sample.txt
var inputTest string

func set(s string) set2.Set[uint8] {
	res := make(set2.Set[uint8], len(s))
	for _, c := range s {
		res.Add(uint8(c))
	}
	return res
}

func score(i set2.Set[uint8]) int {
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

func Part1_slow(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		i := set(line[:len(line)/2]).Intersect(set(line[len(line)/2:]))
		res += score(i)
	}
	return res
}

func Part2_slow(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for l := 0; l < len(lines); l += 3 {
		i := set(lines[l]).Intersect(set(lines[l+1])).Intersect(set(lines[l+2]))
		res += score(i)
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	res := 0
	for _, line := range lines {
		tab := [123]bool{}
		n := len(line)
		for i := 0; i < n/2; i++ {
			tab[line[i]] = true
		}
	loop:
		for i := n / 2; i < n; i++ {
			c := line[i]
			if tab[c] {
				if unicode.IsLower(rune(c)) {
					res += int(1 + c - 'a')
				} else {
					res += int(27 + c - 'A')
				}
				break loop
			}
		}
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for l := 0; l < len(lines); l += 3 {
		tab0 := [123]bool{}
		tab1 := [123]bool{}
		line0 := lines[l]
		line1 := lines[l+1]
		for _, c := range line0 {
			tab0[c] = true
		}
		for _, c := range line1 {
			tab1[c] = true
		}

	loop:
		for _, c := range lines[l+2] {
			if tab0[c] && tab1[c] {
				if unicode.IsLower(c) {
					res += int(1 + c - 'a')
				} else {
					res += int(27 + c - 'A')
				}
				break loop
			}
		}

		//i := set(lines[l]).Intersect(set(lines[l+1])).Intersect(set(lines[l+2]))
		//res += score(i)
	}
	return res
}

func main() {
	fmt.Println("--2022 day 03 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
