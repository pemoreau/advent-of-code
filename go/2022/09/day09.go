package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

//go:embed input.txt
var input_day string

// 2906 too low

type Pos struct {
	x, y int
}

type State struct {
	n    int
	rope []Pos
	path utils.Set[Pos]
}

func NewState(n int) *State {
	res := &State{
		n:    n,
		rope: make([]Pos, n),
		path: utils.BuildSet[Pos](),
	}
	for i := 0; i < n; i++ {
		res.rope[i] = Pos{0, 0}
	}
	res.path.Add(res.rope[n-1])
	return res
}

func (s *State) String() string {
	res := ""
	for y := 15; y >= -6; y-- {
		for x := -10; x < 10; x++ {
			output := false
			for i := 0; i < s.n; i++ {
				if s.rope[i].x == x && s.rope[i].y == y {
					if i == 0 {
						res += "H"
					} else if i == s.n-1 {
						res += "T"
					} else {
						res += strconv.Itoa(i)
					}
					output = true
					break
				}
			}
			if !output {
				if s.path.Contains(Pos{x, y}) {
					res += "#"
				} else {
					res += "."
				}
			}
		}
		res += "\n"
	}
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

// var step = map[Pos]Pos{
// 	{-2, 0}:  {-1, 0},
// 	{2, 0}:   {1, 0},
// 	{0, -2}:  {0, -1},
// 	{0, 2}:   {0, 1},
// 	{-1, -2}: {0, -1},
// 	{-1, 2}:  {0, 1},
// 	{1, -2}:  {0, -1},
// 	{1, 2}:   {0, 1},
// 	{-2, -1}: {-1, 0},
// 	{-2, 1}:  {-1, 0},
// 	{2, -1}:  {1, 0},
// 	{2, 1}:   {1, 0},
// 	{-2, -2}: {-1, -1},
// 	{-2, 2}:  {-1, 1},
// 	{2, -2}:  {1, -1},
// 	{2, 2}:   {1, 1},
// }

func (s *State) MoveTail() {
	// for i := 1; i < s.n; i++ {
	// 	delta := Pos{s.rope[i-1].x - s.rope[i].x, s.rope[i-1].y - s.rope[i].y}
	// 	d, ok := step[delta]
	// 	if ok {
	// 		s.rope[i].x = s.rope[i-1].x - d.x
	// 		s.rope[i].y = s.rope[i-1].y - d.y
	// 	}
	// }

	for i := 1; i < s.n; i++ {
		delta := Pos{s.rope[i-1].x - s.rope[i].x, s.rope[i-1].y - s.rope[i].y}
		if utils.Abs(delta.x) <= 1 && utils.Abs(delta.y) <= 1 {
			return
		}
		if utils.Abs(delta.x) > 1 || utils.Abs(delta.y) > 1 {
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
			// fmt.Printf("dir %c head: %v, tail: %v\n", dir, state.head, state.tail)
			// fmt.Println(state)
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
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
