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

func symbolInNeighborhood(grid utils.Grid, pos utils.Pos) bool {
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

	var grid = utils.BuildGrid(lines)
	minX, maxX, minY, maxY := utils.GridBounds(grid)

	var res int
	for y := minY; y <= maxY; y++ {
		var current int
		var symbol bool
		for x := minX; x <= maxX; x++ {
			if v, ok := grid[utils.Pos{X: x, Y: y}]; ok && v >= '0' && v <= '9' {
				current = 10*current + int(v-'0')
				symbol = symbol || symbolInNeighborhood(grid, utils.Pos{X: x, Y: y})
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

func starInNeighborhood(grid utils.Grid, pos utils.Pos) []utils.Pos {
	var res []utils.Pos
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

	var grid = utils.BuildGrid(lines)
	minX, maxX, minY, maxY := utils.GridBounds(grid)

	var startMap = make(map[utils.Pos][]int)
	var stars = set.NewSet[utils.Pos]()
	for y := minY; y <= maxY; y++ {
		var current int
		for x := minX; x <= maxX; x++ {
			if v, ok := grid[utils.Pos{X: x, Y: y}]; ok && v >= '0' && v <= '9' {
				current = 10*current + int(v-'0')
				stars.AddAll(starInNeighborhood(grid, utils.Pos{X: x, Y: y})...)
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
