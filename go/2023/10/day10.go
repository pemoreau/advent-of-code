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

func tileToString(tile uint8) string {
	switch tile {
	case EMPTY:
		return "EMPTY"
	case UPPER_LEFT:
		return "UPPER_LEFT"
	case UPPER_RIGHT:
		return "UPPER_RIGHT"
	case LOWER_LEFT:
		return "LOWER_LEFT"
	case LOWER_RIGHT:
		return "LOWER_RIGHT"
	case START:
		return "START"
	case VERTICAL:
		return "VERTICAL"
	case HORIZONTAL:
		return "HORIZONTAL"
	}
	panic("unreachable")
}

func dirToString(dir int) string {
	switch dir {
	case NORTH:
		return "NORTH"
	case SOUTH:
		return "SOUTH"
	case EAST:
		return "EAST"
	case WEST:
		return "WEST"
	}
	panic("unreachable")
}

func step(grid utils.Grid, pos utils.Pos, from int) (newPos utils.Pos, newFrom int, ok bool) {
	tile, found := grid[pos]

	if !found {
		return pos, from, false
	}
	if tile == EMPTY {
		return pos, from, false
	}
	if tile == START {
		return pos, from, true
	}

	//fmt.Printf("pos %v tile %s from %s\n", pos, tileToString(tile), dirToString(from))

	if tile == VERTICAL {
		if from == NORTH {
			return utils.Pos{X: pos.X, Y: pos.Y + 1}, from, true
		}
		if from == SOUTH {
			return utils.Pos{X: pos.X, Y: pos.Y - 1}, from, true
		}
		return pos, from, false
	}
	if tile == HORIZONTAL {
		if from == EAST {
			return utils.Pos{X: pos.X - 1, Y: pos.Y}, from, true
		}
		if from == WEST {
			return utils.Pos{X: pos.X + 1, Y: pos.Y}, from, true
		}
		return pos, from, false
	}
	if tile == UPPER_LEFT {
		if from == SOUTH {
			return utils.Pos{X: pos.X + 1, Y: pos.Y}, WEST, true
		}
		if from == EAST {
			return utils.Pos{X: pos.X, Y: pos.Y + 1}, NORTH, true
		}
		return pos, from, false
	}
	if tile == UPPER_RIGHT {
		if from == SOUTH {
			return utils.Pos{X: pos.X - 1, Y: pos.Y}, EAST, true
		}
		if from == WEST {
			return utils.Pos{X: pos.X, Y: pos.Y + 1}, NORTH, true
		}
		return pos, from, false
	}
	if tile == LOWER_LEFT {
		if from == NORTH {
			return utils.Pos{X: pos.X + 1, Y: pos.Y}, WEST, true
		}
		if from == EAST {
			return utils.Pos{X: pos.X, Y: pos.Y - 1}, SOUTH, true
		}
		return pos, from, false
	}
	if tile == LOWER_RIGHT {
		if from == NORTH {
			return utils.Pos{X: pos.X - 1, Y: pos.Y}, EAST, true
		}
		if from == WEST {
			return utils.Pos{X: pos.X, Y: pos.Y - 1}, SOUTH, true
		}
		return pos, from, false
	}
	panic("unreachable")
}

func findLoop(grid utils.Grid, pos utils.Pos, from int) (path []utils.Pos, ok bool) {
	var tile, found = grid[pos]
	if found {
		path = append(path, pos)
	}
	for found {
		pos, from, found = step(grid, pos, from)
		if !found {
			return path, false
		}
		tile, found = grid[pos]
		if !found || tile == EMPTY {
			return path, false
		}
		path = append(path, pos)
		if tile == START {
			return path, true
		}
	}
	return path, false
}

func findStart(grid utils.Grid) (pos utils.Pos) {
	for pos, tile := range grid {
		if tile == START {
			return pos
		}
	}
	panic("unreachable")
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var grid = utils.BuildGrid(lines)
	var start = findStart(grid)

	//utils.DisplayMap(grid, EMPTY)

	var neighbors = []utils.Pos{
		{start.X + 1, start.Y},
		{start.X - 1, start.Y},
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
	}
	var froms = []int{WEST, EAST, SOUTH, NORTH}

	var res int
	for i, n := range neighbors {
		fmt.Println("findLoop", i)
		loop, found := findLoop(grid, n, froms[i])
		if found {
			res = max(res, len(loop))
			break
		} else {
			continue
		}
	}
	if res%2 == 0 {
		return res / 2
	}
	return res/2 + 1
}

func numberIntersections(p utils.Pos, grid utils.Grid, maxX int, path set.Set[utils.Pos]) int {
	var res int
	fmt.Print("numberIntersections", p)
	if path.Contains(p) {
		return 0
	}
	var last uint8
	for x := p.X + 1; x <= maxX; x++ {
		if path.Contains(utils.Pos{X: x, Y: p.Y}) {
			tile, _ := grid[utils.Pos{X: x, Y: p.Y}]
			if last == UPPER_LEFT && tile == UPPER_RIGHT {
				res += 2
			} else if last == UPPER_LEFT && tile == LOWER_RIGHT {
				res++
			} else if last == LOWER_LEFT && tile == LOWER_RIGHT {
				res += 2
			} else if last == LOWER_LEFT && tile == UPPER_RIGHT {
				res++
			}
			if tile != HORIZONTAL {
				res++
				last = tile
			}
		}
	}
	fmt.Println(" --> ", res)
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var grid = utils.BuildGrid(lines)
	var start = findStart(grid)

	//utils.DisplayMap(grid, EMPTY)

	var neighbors = []utils.Pos{
		{start.X + 1, start.Y},
		{start.X - 1, start.Y},
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
	}
	var froms = []int{WEST, EAST, SOUTH, NORTH}

	var loop []utils.Pos
	for i, n := range neighbors {
		fmt.Println("findLoop", i)
		var found bool
		loop, found = findLoop(grid, n, froms[i])
		if found {
			break
		} else {
			fmt.Println("loop not found")
			continue
		}
	}

	var loopSet = make(set.Set[utils.Pos])
	for _, p := range loop {
		loopSet.Add(p)
	}
	var res int
	minX, maxX, minY, maxY := utils.GridBounds(grid)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := utils.Pos{X: x, Y: y}
			if loopSet.Contains(p) {
				continue
			}
			n := numberIntersections(p, grid, maxX, loopSet)
			if n%2 == 1 {
				fmt.Println(x, y)
				res++
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
