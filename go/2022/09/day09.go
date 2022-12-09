package main

import (
	_ "embed"
	"fmt"
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
	head Pos
	tail Pos
	path utils.Set[Pos]
}

func NewState() *State {
	res := &State{
		head: Pos{0, 0},
		tail: Pos{0, 0},
		path: utils.BuildSet[Pos](),
	}
	res.path.Add(res.tail)
	return res
}

func (s *State) String() string {
	res := ""
	for y := 4; y >= 0; y-- {
		for x := 0; x < 6; x++ {
			if s.head.x == x && s.head.y == y {
				res += "H"
			} else if s.tail.x == x && s.tail.y == y {
				res += "T"
			} else if s.path.Contains(Pos{x, y}) {
				res += "#"
			} else {
				res += "."
			}
		}
		res += "\n"
	}
	return res
}

func (s *State) Move(dir byte) {
	switch dir {
	case 'U':
		s.head.y++
	case 'D':
		s.head.y--
	case 'L':
		s.head.x--
	case 'R':
		s.head.x++
	}
}

var step = map[Pos]Pos{
	Pos{-2, 0}:  Pos{-1, 0},
	Pos{2, 0}:   Pos{1, 0},
	Pos{0, -2}:  Pos{0, -1},
	Pos{0, 2}:   Pos{0, 1},
	Pos{-1, -2}: Pos{0, -1},
	Pos{-1, 2}:  Pos{0, 1},
	Pos{1, -2}:  Pos{0, -1},
	Pos{1, 2}:   Pos{0, 1},
	Pos{-2, -1}: Pos{-1, 0},
	Pos{-2, 1}:  Pos{-1, 0},
	Pos{2, -1}:  Pos{1, 0},
	Pos{2, 1}:   Pos{1, 0},
}

func (s *State) MoveTail() {
	delta := Pos{s.head.x - s.tail.x, s.head.y - s.tail.y}
	d, ok := step[delta]
	if ok {
		s.tail.x = s.head.x - d.x
		s.tail.y = s.head.y - d.y
	} else {
		fmt.Println("no step for delta", delta)
	}
	// if utils.Abs(delta.x) <= 1 && utils.Abs(delta.y) <= 1 {
	// 	return
	// }
	// if utils.Abs(delta.x) > 1 || utils.Abs(delta.y) > 1 {
	// 	if delta.y > 0 {
	// 		s.tail.y++
	// 	} else if delta.y < 0 {
	// 		s.tail.y--
	// 	}
	// 	if delta.x > 0 {
	// 		s.tail.x++
	// 	} else if delta.x < 0 {
	// 		s.tail.x--
	// 	}
	// }
	s.path.Add(s.tail)
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	state := NewState()
	for _, line := range lines {
		var dir byte
		var nb int
		fmt.Sscanf(line, "%c %d", &dir, &nb)
		for i := 0; i < nb; i++ {
			state.Move(dir)
			state.MoveTail()
			fmt.Printf("dir %c head: %v, tail: %v\n", dir, state.head, state.tail)
			// fmt.Println(state)
		}
	}
	return state.path.Len()
}

func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	return 0

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
