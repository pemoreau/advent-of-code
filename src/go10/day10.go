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

func parseCorrupted(line string) (int, error) {
	score := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	stack := BuildStack()
	for _, c := range line {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack.Push(c)
		} else {
			p, err := stack.Pop()
			if err != nil {
				return 0, errors.New("empty stack")
			}
			if closing[p] != c {
				return score[c], nil
			}
		}
	}
	return 0, errors.New("complete line")
}

func parseAutocomplete(line string) (int, error) {
	score := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	stack := BuildStack()
	for _, c := range line {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack.Push(c)
		} else {
			p, err := stack.Pop()
			if err != nil {
				return 0, errors.New("empty stack")
			}
			if closing[p] != c {
				return 0, errors.New("corrupted line: " + line)
			}
		}
	}
	if !stack.IsEmpty() {
		res := 0
		for !stack.IsEmpty() {
			p, _ := stack.Pop()
			res = (5 * res) + score[p]
		}
		return res, nil
	}
	return 0, errors.New("complete line:" + line)
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	res := 0
	for line := range lines {
		score, err := parseCorrupted(lines[line])
		if err == nil {
			res += score
		}
	}
	return res
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	scores := []int{}
	for line := range lines {
		score, err := parseAutocomplete(lines[line])
		if err == nil {
			scores = append(scores, score)
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
