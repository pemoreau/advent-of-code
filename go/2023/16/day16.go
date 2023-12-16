package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

type state struct {
	pos utils.Pos
	dir int
}

func (s state) String() string {
	var dirs = []string{"UP", "RIGHT", "DOWN", "LEFT"}
	return fmt.Sprintf("pos: %v, dir: %s", s.pos, dirs[s.dir])
}

func neighboors(g utils.Grid, s state) []state {
	var res []state
	var deltaDir = []utils.Pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	c, found := g[s.pos]

	if !found {
		return res
	}

	x := s.pos.X
	y := s.pos.Y

	if c == '.' {
		newPos := utils.Pos{X: x + deltaDir[s.dir].X, Y: y + deltaDir[s.dir].Y}
		res = append(res, state{pos: newPos, dir: s.dir})
		return res
	}
	if c == '-' {
		if s.dir == UP || s.dir == DOWN {
			res = append(res, state{pos: utils.Pos{X: x - 1, Y: y}, dir: LEFT})
			res = append(res, state{pos: utils.Pos{X: x + 1, Y: y}, dir: RIGHT})
			return res
		}
		if s.dir == LEFT {
			res = append(res, state{pos: utils.Pos{X: x - 1, Y: y}, dir: LEFT})
			return res
		}
		if s.dir == RIGHT {
			res = append(res, state{pos: utils.Pos{X: x + 1, Y: y}, dir: RIGHT})
			return res
		}
	}
	if c == '|' {
		if s.dir == LEFT || s.dir == RIGHT {
			res = append(res, state{pos: utils.Pos{X: x, Y: y - 1}, dir: UP})
			res = append(res, state{pos: utils.Pos{X: x, Y: y + 1}, dir: DOWN})
			return res
		}
		if s.dir == UP {
			res = append(res, state{pos: utils.Pos{X: x, Y: y - 1}, dir: UP})
			return res
		}
		if s.dir == DOWN {
			res = append(res, state{pos: utils.Pos{X: x, Y: y + 1}, dir: DOWN})
			return res
		}
	}
	if c == '/' {
		if s.dir == UP {
			res = append(res, state{pos: utils.Pos{X: x + 1, Y: y}, dir: RIGHT})
			return res
		}
		if s.dir == RIGHT {
			res = append(res, state{pos: utils.Pos{X: x, Y: y - 1}, dir: UP})
			return res
		}
		if s.dir == DOWN {
			res = append(res, state{pos: utils.Pos{X: x - 1, Y: y}, dir: LEFT})
			return res
		}
		if s.dir == LEFT {
			res = append(res, state{pos: utils.Pos{X: x, Y: y + 1}, dir: DOWN})
			return res
		}
	}
	if c == '\\' {
		if s.dir == UP {
			res = append(res, state{pos: utils.Pos{X: x - 1, Y: y}, dir: LEFT})
			return res
		}
		if s.dir == RIGHT {
			res = append(res, state{pos: utils.Pos{X: x, Y: y + 1}, dir: DOWN})
			return res
		}
		if s.dir == DOWN {
			res = append(res, state{pos: utils.Pos{X: x + 1, Y: y}, dir: RIGHT})
			return res
		}
		if s.dir == LEFT {
			res = append(res, state{pos: utils.Pos{X: x, Y: y - 1}, dir: UP})
			return res
		}
	}
	return res

}

func solve(input string, current state) int {
	input = strings.TrimSpace(input)
	var lines = strings.Split(input, "\n")
	var grid = utils.BuildGrid(lines)

	var energized = utils.BuildGrid([]string{})

	//var current = state{pos: utils.Pos{X: 0, Y: 0}, dir: RIGHT}

	var todo []state
	var visited = set.NewSet[state]()

	todo = append(todo, current)
	for len(todo) > 0 {
		var s = todo[0]
		todo = todo[1:]
		if visited.Contains(s) {
			continue
		}

		if _, ok := grid[s.pos]; ok {
			visited.Add(s)
			energized[s.pos] = '#'
		} else {
			continue
		}

		//utils.DisplayMap(grid, ' ')
		//fmt.Println(s)
		//fmt.Println(todo)

		for _, n := range neighboors(grid, s) {
			todo = append(todo, n)
		}
	}

	//utils.DisplayMap(grid, ' ')
	//utils.DisplayMap(energized, ' ')

	var res int
	res = len(energized)
	return res
}

func Part1(input string) int {
	return solve(input, state{pos: utils.Pos{X: 0, Y: 0}, dir: RIGHT})
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	var lines = strings.Split(input, "\n")
	var grid = utils.BuildGrid(lines)
	minX, maxX, minY, maxY := utils.GridBounds(grid)

	var res int
	for x := minX; x <= maxX; x++ {
		res = max(res, solve(input, state{pos: utils.Pos{X: x, Y: minY}, dir: DOWN}))
		res = max(res, solve(input, state{pos: utils.Pos{X: x, Y: maxY}, dir: UP}))
	}
	for y := minY; y <= maxY; y++ {
		res = max(res, solve(input, state{pos: utils.Pos{X: minX, Y: y}, dir: RIGHT}))
		res = max(res, solve(input, state{pos: utils.Pos{X: maxX, Y: y}, dir: LEFT}))
	}

	return res
}

func main() {
	fmt.Println("--2023 day 16 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
