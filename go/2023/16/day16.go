package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
)

//go:embed sample.txt
var inputTest string

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
	fmt.Printf("dir=%d c='%c'\n", dir, c)
	panic("invalid state")
}

type state struct {
	pos game2d.Pos
	dir int
}

func solve(grid *game2d.MatrixChar, current state) int {
	var todo []state
	var visited = set.NewSet[state]()
	var energized = set.NewSet[game2d.Pos]()

	todo = append(todo, current)
	for len(todo) > 0 {
		var s = todo[0]
		todo = todo[1:]
		if visited.Contains(s) || !grid.IsValidPos(s.pos) {
			continue
		}
		visited.Add(s)
		energized.Add(s.pos)
		c := grid.GetPos(s.pos)
		switch nextDir(s.dir, c) {
		case UP:
			todo = append(todo, state{pos: s.pos.N(), dir: UP})
		case RIGHT:
			todo = append(todo, state{pos: s.pos.E(), dir: RIGHT})
		case DOWN:
			todo = append(todo, state{pos: s.pos.S(), dir: DOWN})
		case LEFT:
			todo = append(todo, state{pos: s.pos.W(), dir: LEFT})
		case UPDOWN:
			todo = append(todo, state{pos: s.pos.N(), dir: UP})
			todo = append(todo, state{pos: s.pos.S(), dir: DOWN})
		case LEFTRIGHT:
			todo = append(todo, state{pos: s.pos.W(), dir: LEFT})
			todo = append(todo, state{pos: s.pos.E(), dir: RIGHT})
		}
	}
	return len(energized)
}

func Part1(input string) int {
	grid := game2d.BuildMatrixCharFromString(input)
	return solve(grid, state{pos: game2d.Pos{X: 0, Y: 0}, dir: RIGHT})
}

func Part2(input string) int {
	grid := game2d.BuildMatrixCharFromString(input)

	var res int
	for x := 0; x <= grid.MaxX(); x++ {
		res = max(res, solve(grid, state{pos: game2d.Pos{X: x, Y: 0}, dir: DOWN}))
		res = max(res, solve(grid, state{pos: game2d.Pos{X: x, Y: grid.MaxY()}, dir: UP}))
	}
	for y := 0; y <= grid.MaxY(); y++ {
		res = max(res, solve(grid, state{pos: game2d.Pos{X: 0, Y: y}, dir: RIGHT}))
		res = max(res, solve(grid, state{pos: game2d.Pos{X: grid.MaxX(), Y: y}, dir: LEFT}))
	}

	return res
}

func main() {
	fmt.Println("--2023 day 16 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
