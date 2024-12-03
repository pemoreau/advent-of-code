package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func parseInput(input string) (string, []string) {
	input = strings.Trim(input, "\n")
	var lines = strings.Split(input, "\n")
	var rules []string
	for _, rule := range lines[2:] {
		var s = fmt.Sprintf("%s%s", rule[0:5], rule[9:10])
		rules = append(rules, s)
	}
	return lines[0][15:], rules
}

type State struct {
	rshift int
	rubber []uint8
}

func (s State) String() string {
	var b strings.Builder
	for i := 0; i < s.rshift; i++ {
		b.WriteByte(' ')
	}
	b.WriteByte('v')
	b.WriteByte('\n')
	for _, r := range s.rubber {
		b.WriteByte(r)
	}
	return b.String()
}

func step(state State, rules []string) State {
	var newState State
	newState.rshift = state.rshift
	newState.rubber = make([]uint8, len(state.rubber))
	newState.rubber[0] = '.'
	newState.rubber[1] = '.'
	newState.rubber[len(state.rubber)-1] = '.'
	newState.rubber[len(state.rubber)-2] = '.'
	for i := 2; i < len(state.rubber)-2; i++ {
		var subject = string(state.rubber[i-2 : i+3])
		var applied = false
		for _, rule := range rules {
			if subject == rule[0:5] {
				newState.rubber[i] = rule[5]
				applied = true
				break
			}
		}
		if !applied {
			newState.rubber[i] = '.'
		}
	}
	return newState
}

func padding(state State) State {
	var emptyLeft, emptyRight = 0, 0
	var padLeft, padRight = 0, 0
	for i := 0; i < len(state.rubber); i++ {
		if state.rubber[i] == '#' {
			emptyLeft = i
			break
		}
	}
	if emptyLeft < 4 {
		padLeft = 4 - emptyLeft
	}
	for i := len(state.rubber) - 1; i >= 0; i-- {
		if state.rubber[i] == '#' {
			emptyRight = len(state.rubber) - i - 1
			break
		}
	}
	if emptyRight < 4 {
		padRight = 4 - emptyRight
	}
	var newState State
	newState.rshift = state.rshift + padLeft
	newState.rubber = make([]uint8, len(state.rubber)+padLeft+padRight)
	for i := 0; i < padLeft; i++ {
		newState.rubber[i] = '.'
	}
	for i := 0; i < len(state.rubber); i++ {
		newState.rubber[i+padLeft] = state.rubber[i]
	}
	for i := 0; i < padRight; i++ {
		newState.rubber[len(state.rubber)+padLeft+i] = '.'
	}
	return newState
}

func sum(state State) int {
	var sum int
	for x, c := range state.rubber {
		if c == '#' {
			sum += (x - state.rshift)
		}
	}
	return sum
}

func Part1(input string) int {
	var start, rules = parseInput(input)
	var state State
	state.rshift = 0
	state.rubber = make([]uint8, len(start))
	for i, c := range start {
		state.rubber[i] = uint8(c)
	}
	state = padding(state)
	for i := 1; i < 21; i++ {
		state = step(state, rules)
		state = padding(state)
	}
	return sum(state)
}

func difference(values []int) []int {
	var diffs []int
	for i := 1; i < len(values); i++ {
		diffs = append(diffs, values[i]-values[i-1])
	}
	return diffs
}

func Part2(input string) int {
	var start, rules = parseInput(input)
	var state State
	state.rshift = 0
	state.rubber = make([]uint8, len(start))
	for i, c := range start {
		state.rubber[i] = uint8(c)
	}
	state = padding(state)

	var values []int
	for i := 1; i <= 500; i++ {
		state = step(state, rules)
		state = padding(state)
		values = append(values, sum(state))
	}
	var n = 500
	var value = values[n-1]
	var N = 50000000000
	var diff = 50 // difference between values[n] and values[n-1]
	return value + (N-n)*diff
}

func main() {
	fmt.Println("--2018 day 12 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
