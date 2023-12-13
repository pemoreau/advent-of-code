package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func isVerticalMirror(g utils.Grid, x int) bool {
	minX, maxX, minY, maxY := utils.GridBounds(g)
	diff := min(maxX-(x+1), x-minX)
	for y := minY; y <= maxY; y++ {
		for i := 0; i <= diff; i++ {
			if g[utils.Pos{X: x + 1 + i, Y: y}] != g[utils.Pos{X: x - i, Y: y}] {
				return false
			}
		}
	}
	return true
}

func findVerticalMirror(g utils.Grid, old int) int {
	minX, maxX, _, _ := utils.GridBounds(g)
	for x := minX; x <= maxX-1; x++ {
		if x+1 != old && isVerticalMirror(g, x) {
			return x + 1
		}
	}
	return 0
}

func isHorizontalMirror(g utils.Grid, y int) bool {
	minX, maxX, minY, maxY := utils.GridBounds(g)
	diff := min(maxY-(y+1), y-minY)
	for x := minX; x <= maxX; x++ {
		for i := 0; i <= diff; i++ {
			if g[utils.Pos{Y: y + 1 + i, X: x}] != g[utils.Pos{Y: y - i, X: x}] {
				return false
			}
		}
	}
	return true
}

func findHorizontalMirror(g utils.Grid, old int) int {
	_, _, minY, maxY := utils.GridBounds(g)
	for y := minY; y <= maxY-1; y++ {
		if y+1 != old && isHorizontalMirror(g, y) {
			return y + 1
		}
	}
	return 0
}

func findSmudge(g utils.Grid) int {
	h := findHorizontalMirror(g, -1)
	v := findVerticalMirror(g, -1)
	//fmt.Println("previous h v", h, v)

	minX, maxX, minY, maxY := utils.GridBounds(g)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			var res int
			if g[utils.Pos{X: x, Y: y}] == '#' {
				g[utils.Pos{X: x, Y: y}] = '.'
				nh := findHorizontalMirror(g, h)
				nv := findVerticalMirror(g, v)
				if nh > 0 && nh != h {
					res += 100 * nh
				}
				if nv > 0 && nv != v {
					res += nv
				}
				g[utils.Pos{X: x, Y: y}] = '#'
				if res > 0 {
					//fmt.Printf("smugle at %d,%d --> nh=%d nv=%d --> res=%d\n", x, y, nh, nv, res)
					return res
				}
			} else if g[utils.Pos{X: x, Y: y}] == '.' {
				g[utils.Pos{X: x, Y: y}] = '#'
				nh := findHorizontalMirror(g, h)
				nv := findVerticalMirror(g, v)
				if nh > 0 && nh != h {
					res += 100 * nh
				}
				if nv > 0 && nv != v {
					res += nv
				}
				g[utils.Pos{X: x, Y: y}] = '.'
				if res > 0 {
					//fmt.Printf("smugle at %d,%d --> nh=%d nv=%d --> res=%d\n", x, y, nh, nv, res)
					return res
				}
			} else {
				panic("invalid char")
			}
		}
	}
	//fmt.Println("No smugle found")
	return 0
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")

	var res int
	for _, p := range parts {
		lines := strings.Split(p, "\n")
		grid := utils.BuildGrid(lines)
		h := findHorizontalMirror(grid, -1)
		v := findVerticalMirror(grid, -1)
		res += v + (h * 100)
	}

	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")

	var res int
	for _, p := range parts {
		lines := strings.Split(p, "\n")
		grid := utils.BuildGrid(lines)
		s := findSmudge(grid)
		res += s
	}
	return res
}

func main() {
	fmt.Println("--2023 day 13 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
