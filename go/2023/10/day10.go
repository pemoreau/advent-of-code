package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"time"
)

//go:embed sample.txt
var inputTest string

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

func step(grid *game2d.MatrixChar, pos game2d.Pos, from int) (newPos game2d.Pos, newFrom int, ok bool) {
	if !grid.IsValidPos(pos) {
		return pos, from, false
	}

	switch grid.GetPos(pos) {
	case EMPTY:
		return pos, from, false
	case START:
		return pos, from, true
	case VERTICAL:
		if from == NORTH {
			return pos.S(), from, true
		} else if from == SOUTH {
			return pos.N(), from, true
		}
	case HORIZONTAL:
		if from == EAST {
			return pos.W(), from, true
		} else if from == WEST {
			return pos.E(), from, true
		}
	case UPPER_LEFT:
		if from == SOUTH {
			return pos.E(), WEST, true
		} else if from == EAST {
			return pos.S(), NORTH, true
		}
	case UPPER_RIGHT:
		if from == SOUTH {
			return pos.W(), EAST, true
		} else if from == WEST {
			return pos.S(), NORTH, true
		}
	case LOWER_LEFT:
		if from == NORTH {
			return pos.E(), WEST, true
		} else if from == EAST {
			return pos.N(), SOUTH, true
		}
	case LOWER_RIGHT:
		if from == NORTH {
			return pos.W(), EAST, true
		} else if from == WEST {
			return pos.N(), SOUTH, true
		}
	}
	return pos, from, false
}

func findLoop(grid *game2d.MatrixChar, pos game2d.Pos, from int) (path set.Set[game2d.Pos], ok bool) {
	if !grid.IsValidPos(pos) {
		return path, false
	}

	path = make(set.Set[game2d.Pos])

	if grid.GetPos(pos) == EMPTY {
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

		if grid.GetPos(pos) == START {
			return path, true
		}
	}
}

func findStart(grid *game2d.MatrixChar) game2d.Pos {
	if start, ok := grid.Find(START); ok {
		return start
	}
	panic("unreachable")
}

func Part1(input string) int {
	var grid = game2d.BuildMatrixCharFromString(input)
	var start = findStart(grid)

	var neighbors = []game2d.Pos{start.E(), start.W(), start.N(), start.S()}
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
	var grid = game2d.BuildMatrixCharFromString(input)
	var start = findStart(grid)

	var neighbors = []game2d.Pos{start.E(), start.W(), start.N(), start.S()}
	var froms = []int{WEST, EAST, SOUTH, NORTH}

	var loopSet set.Set[game2d.Pos]
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
			if !loopSet.Contains(game2d.Pos{x, y}) {
				if cpt%2 == 1 {
					res++
				}
			} else {
				tile := grid.Get(x, y)
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
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
