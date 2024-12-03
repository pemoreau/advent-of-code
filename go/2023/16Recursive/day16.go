package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"time"
)

//go:embed input.txt
var inputDay string

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
	UPDOWN
	LEFTRIGHT
)

func nextDir(dir int, c uint8) int {
	switch dir {
	case RIGHT:
		switch c {
		case '.', '-':
			return RIGHT
		case '|':
			return UPDOWN
		case '/':
			return UP
		case '\\':
			return DOWN
		}
	case LEFT:
		switch c {
		case '.', '-':
			return LEFT
		case '|':
			return UPDOWN
		case '/':
			return DOWN
		case '\\':
			return UP
		}
	case UP:
		switch c {
		case '.', '|':
			return UP
		case '-':
			return LEFTRIGHT
		case '/':
			return RIGHT
		case '\\':
			return LEFT
		}
	case DOWN:
		switch c {
		case '.', '|':
			return DOWN
		case '-':
			return LEFTRIGHT
		case '/':
			return LEFT
		case '\\':
			return RIGHT
		}
	}
	panic("invalid state")
}

type state struct {
	pos game2d.Pos
	dir int
}

func solve(grid game2d.MatrixChar, visited set.Set[state], energized set.Set[game2d.Pos], s state) {

	if visited.Contains(s) || !grid.IsValidPos(s.pos) {
		return
	}

	visited.Add(s)
	energized.Add(s.pos)

	switch nextDir(s.dir, grid.GetPos(s.pos)) {
	case UP:
		solve(grid, visited, energized, state{pos: s.pos.N(), dir: UP})
	case RIGHT:
		solve(grid, visited, energized, state{pos: s.pos.E(), dir: RIGHT})
	case DOWN:
		solve(grid, visited, energized, state{pos: s.pos.S(), dir: DOWN})
	case LEFT:
		solve(grid, visited, energized, state{pos: s.pos.W(), dir: LEFT})
	case UPDOWN:
		solve(grid, visited, energized, state{pos: s.pos.N(), dir: UP})
		solve(grid, visited, energized, state{pos: s.pos.S(), dir: DOWN})
	case LEFTRIGHT:
		solve(grid, visited, energized, state{pos: s.pos.W(), dir: LEFT})
		solve(grid, visited, energized, state{pos: s.pos.E(), dir: RIGHT})
	}
}

func Part1(input string) int {
	grid := game2d.BuildMatrixCharFromString(input)
	var visited = set.NewSet[state]()
	var energized = set.NewSet[game2d.Pos]()
	var start = state{pos: game2d.Pos{X: 0, Y: 0}, dir: RIGHT}
	solve(grid, visited, energized, start)
	return len(energized)
}

func Part2(input string) int {
	grid := game2d.BuildMatrixCharFromString(input)
	var visited = set.NewSet[state]()
	var energized = set.NewSet[game2d.Pos]()

	var starts []state
	for x := 0; x <= grid.MaxX(); x++ {
		starts = append(starts, state{pos: game2d.Pos{X: x, Y: 0}, dir: DOWN})
		starts = append(starts, state{pos: game2d.Pos{X: x, Y: grid.MaxY()}, dir: UP})
	}
	for y := 0; y <= grid.MaxY(); y++ {
		starts = append(starts, state{pos: game2d.Pos{X: 0, Y: y}, dir: RIGHT})
		starts = append(starts, state{pos: game2d.Pos{X: grid.MaxX(), Y: y}, dir: LEFT})
	}

	var res int
	for _, start := range starts {
		clear(visited)
		clear(energized)
		solve(grid, visited, energized, start)
		res = max(res, len(energized))
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
