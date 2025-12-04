package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
)

func step(grid *game2d.GridChar) int {
	var res int
	for pos := range grid.AllPos() {
		var free int
		if !grid.Contains(pos, '@') {
			continue
		}
		for n := range pos.Neighbors8() {
			if grid.Contains(n, '@') || grid.Contains(n, 'x') {

				free++
			}
		}
		if free < 4 {
			res++
			grid.SetPos(pos, 'x')
		}
	}
	return res
}

func Part1(input string) int {
	var grid = game2d.BuildGridCharFromString(input)
	return step(grid)
}

func Part2(input string) int {
	var res int
	var grid = game2d.BuildGridCharFromString(input)

	var moved = true
	for moved {
		moved = false
		res += step(grid)
		for pos := range grid.AllPos() {
			if grid.Contains(pos, 'x') {
				grid.SetPos(pos, '.')
				moved = true
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2025 day 04 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
