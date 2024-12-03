package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

const CLAY = '#'
const SAND = '|'
const WATER = '~'

func N(p game2d.Pos) game2d.Pos {
	return game2d.Pos{X: p.X, Y: p.Y - 1}
}
func S(p game2d.Pos) game2d.Pos {
	return game2d.Pos{X: p.X, Y: p.Y + 1}
}
func E(p game2d.Pos) game2d.Pos {
	return game2d.Pos{X: p.X + 1, Y: p.Y}
}
func W(p game2d.Pos) game2d.Pos {
	return game2d.Pos{X: p.X - 1, Y: p.Y}
}

func DisplayMapDebug(grid *game2d.GridChar, empty uint8, pos game2d.Pos) {
	minX, maxX, minY, maxY := grid.GetBounds()
	minX = max(minX, pos.X-30)
	maxX = min(maxX, pos.X+30)
	minY = max(minY, pos.Y-10)
	maxY = min(maxY, pos.Y+10)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if pos.X == x && pos.Y == y {
				fmt.Printf("+")
			} else if v, ok := grid.Get(x, y); ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Printf("%c", empty)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func parseInput(input string) *game2d.GridChar {
	input = strings.Trim(input, "\n")
	var lines = strings.Split(input, "\n")
	var grid = game2d.NewGridChar()
	for _, line := range lines {
		var c1, c2 uint8
		var a, b, c int
		fmt.Sscanf(line, "%c=%d, %c=%d..%d", &c1, &a, &c2, &b, &c)
		if c1 == 'x' {
			for y := b; y <= c; y++ {
				grid.Set(a, y, CLAY)
			}
		} else {
			for x := b; x <= c; x++ {
				grid.Set(x, a, CLAY)
			}
		}
	}
	return grid
}

func isBlocked(grid *game2d.GridChar, pos game2d.Pos, dir func(pos2 game2d.Pos) game2d.Pos) (bool, int) {
	var spreadX int
	for {
		if s, ok := grid.GetPos(S(pos)); !ok || s == SAND {
			return false, pos.X
		}
		if s, ok := grid.GetPos(dir(pos)); ok && (s == CLAY || s == WATER) {
			spreadX = pos.X
			break
		}
		pos = dir(pos)
	}
	return true, spreadX
}

func spread(grid *game2d.GridChar, pos game2d.Pos, bound int, c uint8) {
	for x := min(pos.X, bound); x <= max(pos.X, bound); x++ {
		grid.Set(x, pos.Y, c)
	}
}

func countSandWater(grid *game2d.GridChar, minY int) (int, int) {
	var sand, water int
	for p, v := range grid.All() {
		if p.Y >= minY {
			if v == SAND {
				sand++
			} else if v == WATER {
				water++
			}
		}
	}
	return sand, water
}

func solve(input string) (*game2d.GridChar, int) {
	var grid = parseInput(input)
	var _, _, minY, maxY = grid.GetBounds()
	var todo = []game2d.Pos{{X: 500, Y: 1}} // spring

	for len(todo) > 0 {
		var p = todo[0]
		todo = todo[1:]
		//DisplayMapDebug(grid, '.', p)

		if p.Y > maxY || p.Y < 1 { // out of bound
			continue
		}

		if grid.Contains(p, WATER) { // already water
			continue
		}

		if _, ok := grid.GetPos(p); !ok {
			grid.SetPos(p, SAND)
		}

		if s, ok := grid.GetPos(S(p)); !ok || s == SAND { // fall down
			todo = append(todo, S(p))
			continue
		}

		if s, ok := grid.GetPos(S(p)); ok && (s == CLAY || s == WATER) {
			var blockedLeft, leftX = isBlocked(grid, p, W)
			var blockedRight, rightX = isBlocked(grid, p, E)
			var fallLeft = game2d.Pos{X: leftX, Y: p.Y}
			var fallRight = game2d.Pos{X: rightX, Y: p.Y}

			if blockedLeft && blockedRight {
				spread(grid, p, leftX, WATER)
				spread(grid, p, rightX, WATER)
				todo = append(todo, N(p))
			} else if blockedLeft {
				spread(grid, fallRight, leftX, SAND)
				todo = append(todo, S(fallRight))
			} else if blockedRight {
				spread(grid, fallLeft, rightX, SAND)
				todo = append(todo, S(fallLeft))
			} else {
				spread(grid, fallLeft, rightX, SAND)
				spread(grid, fallRight, leftX, SAND)
				todo = append(todo, S(fallLeft))
				todo = append(todo, S(fallRight))
			}
		}
	}
	return grid, minY
}

func Part1(input string) int {
	grid, minY := solve(input)
	sand, water := countSandWater(grid, minY)
	return sand + water
}

func Part2(input string) int {
	grid, minY := solve(input)
	_, water := countSandWater(grid, minY)
	return water
}

func main() {
	fmt.Println("--2018 day 17 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
