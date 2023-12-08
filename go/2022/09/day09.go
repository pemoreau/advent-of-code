package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

//go:embed input.txt
var inputDay string

type Pos struct {
	x, y int
}

type State struct {
	n    int
	rope []Pos
	path set.Set[Pos]
}

func NewState(n int) *State {
	res := &State{
		n:    n,
		rope: make([]Pos, n),
		path: set.NewSet[Pos](),
	}
	res.path.Add(res.rope[n-1])
	return res
}

func (s *State) Move(dir byte) {
	switch dir {
	case 'U':
		s.rope[0].y++
	case 'D':
		s.rope[0].y--
	case 'L':
		s.rope[0].x--
	case 'R':
		s.rope[0].x++
	}
}

func (s *State) MoveTail() {
	for i := 1; i < s.n; i++ {
		delta := Pos{s.rope[i-1].x - s.rope[i].x, s.rope[i-1].y - s.rope[i].y}
		if utils.Abs(delta.x) <= 1 && utils.Abs(delta.y) <= 1 {
			return
		}
		if delta.y > 0 {
			s.rope[i].y++
		} else if delta.y < 0 {
			s.rope[i].y--
		}
		if delta.x > 0 {
			s.rope[i].x++
		} else if delta.x < 0 {
			s.rope[i].x--
		}
	}
	s.path.Add(s.rope[s.n-1])
}

func run(input string, n int) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	state := NewState(n)
	for _, line := range lines {
		var dir byte
		var nb int
		fmt.Sscanf(line, "%c %d", &dir, &nb)
		for i := 0; i < nb; i++ {
			state.Move(dir)
			state.MoveTail()
		}
	}
	return state.path.Len()
}

func Part1(input string) int {
	return run(input, 2)
}

func Part2(input string) int {
	return run(input, 10)
}

func main() {
	fmt.Println("--2022 day 09 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
