package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"time"
)

//go:embed input.txt
var inputDay string

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var deltaDir = []utils.Pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

type state struct {
	pos utils.Pos
	dir int
}

func neighboors(grid utils.MatrixDigit, s state, mini, maxi int) []state {
	var res []state
	dirs := [2]int{(s.dir + 1) % 4, (s.dir + 3) % 4}
	for _, d := range dirs {
		for i := mini; i <= maxi; i++ {
			pos := utils.Pos{s.pos.X + i*deltaDir[d].X, s.pos.Y + i*deltaDir[d].Y}
			if grid.IsValidPos(pos) {
				res = append(res, state{pos, d})
			} else {
				break
			}
		}
	}
	return res
}

func cost(grid utils.MatrixDigit, from, to state) int {
	x1, y1 := from.pos.X, from.pos.Y
	x2, y2 := to.pos.X, to.pos.Y
	var res int
	if x1 == x2 {
		if y1 < y2 {
			for y := y1 + 1; y <= y2; y++ {
				res += int(grid[y][x1])
			}
		} else if y1 > y2 {
			for y := y1 - 1; y >= y2; y-- {
				res += int(grid[y][x1])
			}
		}
		return res
	}
	if y1 == y2 {
		if x1 < x2 {
			for x := x1 + 1; x <= x2; x++ {
				res += int(grid[y1][x])
			}
		} else if x1 > x2 {
			for x := x1 - 1; x >= x2; x-- {
				res += int(grid[y1][x])
			}
		}
		return res
	}
	return res
}

func solve(input string, mini, maxi int) int {
	grid := utils.BuildMatrixDigitFromString(input)

	origin := utils.Pos{0, 0}
	starts := []state{{pos: origin, dir: UP}, {pos: origin, dir: RIGHT}}
	exit := utils.Pos{grid.MaxX(), grid.MaxY()}

	heuristicFunction := func(s state) int { return utils.ManhattanDistance(s.pos, exit) }
	goalFunction := func(s state) bool { return s.pos == exit }
	neighboorsFunction := func(s state) []state { return neighboors(grid, s, mini, maxi) }
	costFunction := func(from, to state) int { return cost(grid, from, to) }

	_, distance := utils.AstarMultipleStart(starts, goalFunction, neighboorsFunction, costFunction, heuristicFunction)
	return distance
}

func Part1(input string) int {
	return solve(input, 1, 3)
}

func Part2(input string) int {
	return solve(input, 4, 10)
}

func main() {
	fmt.Println("--2023 day 17 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
