package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

type stack []rune

func BuildStack() stack {
	return make([]rune, 0)
}

func (s *stack) Push(c rune) {
	*s = append(*s, c)
}

func (s *stack) Pop() (rune, error) {
	l := len(*s)
	if l == 0 {
		return 0, errors.New("stack is empty")
	}
	top := (*s)[l-1]
	*s = (*s)[:l-1]
	return top, nil
}

func (s *stack) Peek() (rune, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

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
	stack := BuildStack()
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
	scores := []int{}
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
	content, _ := ioutil.ReadFile("../../inputs/day10.txt")

	start := time.Now()
	fmt.Println("part1: ", Part1(string(content)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(content)))
	fmt.Println(time.Since(start))
}
