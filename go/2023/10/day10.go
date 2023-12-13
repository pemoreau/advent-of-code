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
	EMPTY       = '.'
	UPPER_LEFT  = 'F'
	UPPER_RIGHT = '7'
	LOWER_LEFT  = 'L'
	LOWER_RIGHT = 'J'
	START       = 'S'
	VERTICAL    = '|'
	HORIZONTAL  = '-'
	NORTH       = iota
	SOUTH
	EAST
	WEST
)

func step(grid utils.Matrix[uint8], pos utils.Pos, from int) (newPos utils.Pos, newFrom int, ok bool) {
	if !grid.IsValidPos(pos) {
		return pos, from, false
	}
	tile := grid[pos.Y][pos.X]

	switch tile {
	case EMPTY:
		return pos, from, false
	case START:
		return pos, from, true
	case VERTICAL:
		if from == NORTH {
			return utils.Pos{X: pos.X, Y: pos.Y + 1}, from, true
		} else if from == SOUTH {
			return utils.Pos{X: pos.X, Y: pos.Y - 1}, from, true
		}
	case HORIZONTAL:
		if from == EAST {
			return utils.Pos{X: pos.X - 1, Y: pos.Y}, from, true
		} else if from == WEST {
			return utils.Pos{X: pos.X + 1, Y: pos.Y}, from, true
		}
	case UPPER_LEFT:
		if from == SOUTH {
			return utils.Pos{X: pos.X + 1, Y: pos.Y}, WEST, true
		} else if from == EAST {
			return utils.Pos{X: pos.X, Y: pos.Y + 1}, NORTH, true
		}
	case UPPER_RIGHT:
		if from == SOUTH {
			return utils.Pos{X: pos.X - 1, Y: pos.Y}, EAST, true
		} else if from == WEST {
			return utils.Pos{X: pos.X, Y: pos.Y + 1}, NORTH, true
		}
	case LOWER_LEFT:
		if from == NORTH {
			return utils.Pos{X: pos.X + 1, Y: pos.Y}, WEST, true
		} else if from == EAST {
			return utils.Pos{X: pos.X, Y: pos.Y - 1}, SOUTH, true
		}
	case LOWER_RIGHT:
		if from == NORTH {
			return utils.Pos{X: pos.X - 1, Y: pos.Y}, EAST, true
		} else if from == WEST {
			return utils.Pos{X: pos.X, Y: pos.Y - 1}, SOUTH, true
		}
	}
	return pos, from, false
}

func findLoop(grid utils.Matrix[uint8], pos utils.Pos, from int) (path set.Set[utils.Pos], ok bool) {
	if !grid.IsValidPos(pos) {
		return path, false
	}

	path = make(set.Set[utils.Pos])

	if tile := grid[pos.Y][pos.X]; tile == EMPTY {
		return path, false
	}

	path.Add(pos)
	for {
		var found bool
		pos, from, found = step(grid, pos, from)
		if !found {
			return path, false
		}
		path.Add(pos)

		if tile := grid[pos.Y][pos.X]; tile == START {
			return path, true
		}
	}
}

func findStart(grid utils.Matrix[uint8]) utils.Pos {
	for y, l := range grid {
		for x, c := range l {
			if c == START {
				return utils.Pos{X: x, Y: y}
			}
		}
	}
	panic("unreachable")
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var grid = utils.BuildMatrixChar(lines)
	var start = findStart(grid)

	var neighbors = []utils.Pos{
		{start.X + 1, start.Y},
		{start.X - 1, start.Y},
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
	}
	var froms = []int{WEST, EAST, SOUTH, NORTH}

	for i, n := range neighbors {
		loop, found := findLoop(grid, n, froms[i])
		if found {
			var l = len(loop)
			if l%2 == 0 {
				return l / 2
			}
			return l/2 + 1
		}
	}
	panic("no solution found")
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var grid = utils.BuildMatrixChar(lines)
	var start = findStart(grid)

	var neighbors = []utils.Pos{
		{start.X + 1, start.Y},
		{start.X - 1, start.Y},
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
	}
	var froms = []int{WEST, EAST, SOUTH, NORTH}

	var loopSet set.Set[utils.Pos]
	for i, n := range neighbors {
		var found bool
		loopSet, found = findLoop(grid, n, froms[i])
		if found {
			break
		}
	}

	var res int
	for y := 0; y <= grid.MaxY(); y++ {
		var last uint8
		var cpt = 0
		for x := grid.MaxX(); x >= 0; x-- {
			if !loopSet.Contains(utils.Pos{x, y}) {
				if cpt%2 == 1 {
					res++
				}
			} else {
				tile := grid[y][x]
				if last == UPPER_RIGHT && tile == UPPER_LEFT {
					cpt++
				} else if last == UPPER_RIGHT && tile == LOWER_LEFT {
					// do not count
				} else if last == LOWER_RIGHT && tile == LOWER_LEFT {
					cpt++
				} else if last == LOWER_RIGHT && tile == UPPER_LEFT {
					// do not count
				} else if tile != HORIZONTAL {
					last = tile
					cpt++
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println("--2023 day 10 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
