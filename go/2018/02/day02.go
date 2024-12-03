package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func check2(s string) bool {
	for _, b := range check(s) {
		if b == 2 {
			// fmt.Printf("check2: %#U occurs %d\n", a, b)
			return true
		}
	}
	return false
}
func check3(s string) bool {
	for _, b := range check(s) {
		if b == 3 {
			// fmt.Printf("check3: %#U occurs %d\n", a, b)
			return true
		}
	}
	return false
}

func check(s string) map[rune]int {
	var visited = make(map[rune]int)
	for _, b := range s {
		visited[b] += 1
	}
	return visited
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	nb2 := 0
	nb3 := 0
	for _, v := range lines {
		if check2(v) {
			nb2 += 1
		}
		if check3(v) {
			nb3 += 1
		}
	}
	return nb2 * nb3
}

func differByOne(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diff := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			diff += 1
		}
	}
	return diff == 1
}

func Part2(input string) string {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	for i, v := range lines {
		for _, w := range lines[i+1:] {
			if differByOne(v, w) {
				res := ""
				for i := range v {
					if v[i] == w[i] {
						res += string(v[i])
					}
				}
				return res
			}
		}
	}
	return ""
}

func main() {
	fmt.Println("--2018 day 02 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part1: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
