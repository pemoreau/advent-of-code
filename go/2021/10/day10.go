package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/stack"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

var closing = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var corruptedScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autoScore = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func parseLine(line string) (corrupted int, auto int) {
	corrupted = 0
	auto = 0
	stack := stack.NewStack[rune]()
	for _, c := range line {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack.Push(c)
		} else {
			p, err := stack.Pop()
			if err != nil {
				return corrupted, auto
			}
			if closing[p] != c {
				return corruptedScore[c], auto
			}
		}
	}
	if !stack.IsEmpty() {
		for !stack.IsEmpty() {
			p, _ := stack.Pop()
			auto = (5 * auto) + autoScore[p]
		}
	}
	return corrupted, auto
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	res := 0
	for line := range lines {
		corrupted, _ := parseLine(lines[line])
		res += corrupted
	}
	return res
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var scores []int
	for line := range lines {
		_, auto := parseLine(lines[line])
		if auto > 0 {
			scores = append(scores, auto)
		}
	}
	sort.Ints(scores)
	return scores[(len(scores)-1)/2]
}

func main() {
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
