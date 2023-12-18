package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func symbolInNeighborhood(grid game2d.Grid, pos game2d.Pos) bool {
	for _, neighbor := range pos.Neighbors8() {
		if n, ok := grid[neighbor]; ok && n != '.' && !(n >= '0' && n <= '9') {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var grid = game2d.BuildGrid(lines)
	minX, maxX, minY, maxY := game2d.GridBounds(grid)

	var res int
	for y := minY; y <= maxY; y++ {
		var current int
		var symbol bool
		for x := minX; x <= maxX; x++ {
			if v, ok := grid[game2d.Pos{X: x, Y: y}]; ok && v >= '0' && v <= '9' {
				current = 10*current + int(v-'0')
				symbol = symbol || symbolInNeighborhood(grid, game2d.Pos{X: x, Y: y})
			} else {
				if symbol {
					res += current
				}
				current = 0
				symbol = false
			}
		}
		if symbol {
			res += current
		}
	}
	return res
}

func starInNeighborhood(grid game2d.Grid, pos game2d.Pos) []game2d.Pos {
	var res []game2d.Pos
	for _, neighbor := range pos.Neighbors8() {
		if n, ok := grid[neighbor]; ok && n == '*' {
			res = append(res, neighbor)
		}
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var grid = game2d.BuildGrid(lines)
	minX, maxX, minY, maxY := game2d.GridBounds(grid)

	var startMap = make(map[game2d.Pos][]int)
	var stars = set.NewSet[game2d.Pos]()
	for y := minY; y <= maxY; y++ {
		var current int
		for x := minX; x <= maxX; x++ {
			if v, ok := grid[game2d.Pos{X: x, Y: y}]; ok && v >= '0' && v <= '9' {
				current = 10*current + int(v-'0')
				stars.AddAll(starInNeighborhood(grid, game2d.Pos{X: x, Y: y})...)
			} else {
				if len(stars) > 0 {
					for star := range stars {
						startMap[star] = append(startMap[star], current)
					}
				}
				current = 0
				stars.Clear()
			}
		}
		if len(stars) > 0 {
			for star := range stars {
				startMap[star] = append(startMap[star], current)
			}
		}
		stars.Clear()
	}

	var res int
	for _, v := range startMap {
		if len(v) == 2 {
			res += v[0] * v[1]
		}
	}
	return res
}

func main() {
	fmt.Println("--2023 day 03 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
