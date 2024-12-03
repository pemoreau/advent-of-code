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

func symbolInNeighborhood(grid *game2d.GridChar, pos game2d.Pos) bool {
	for neighbor := range pos.Neighbors8() {
		if n, ok := grid.GetPos(neighbor); ok && n != '.' && !(n >= '0' && n <= '9') {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	var grid = game2d.BuildGridCharFromString(input)
	minX, maxX, minY, maxY := grid.GetBounds()

	var res int
	for y := minY; y <= maxY; y++ {
		var current int
		var symbol bool
		for x := minX; x <= maxX; x++ {
			if v, ok := grid.Get(x, y); ok && v >= '0' && v <= '9' {
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

func starInNeighborhood(grid *game2d.GridChar, pos game2d.Pos) []game2d.Pos {
	var res []game2d.Pos
	for neighbor := range pos.Neighbors8() {
		if n, ok := grid.GetPos(neighbor); ok && n == '*' {
			res = append(res, neighbor)
		}
	}
	return res
}

func Part2(input string) int {
	var grid = game2d.BuildGridCharFromString(input)
	minX, maxX, minY, maxY := grid.GetBounds()

	var startMap = make(map[game2d.Pos][]int)
	var stars = set.NewSet[game2d.Pos]()
	for y := minY; y <= maxY; y++ {
		var current int
		for x := minX; x <= maxX; x++ {
			if v, ok := grid.Get(x, y); ok && v >= '0' && v <= '9' {
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
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
