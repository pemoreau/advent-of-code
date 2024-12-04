package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"time"
)

//go:embed sample.txt
var inputTest string

func xmas(grid *game2d.GridChar, p game2d.Pos, dir game2d.Pos) bool {
	var letters = []uint8("MAS") // X already checked
	for _, l := range letters {
		if !grid.Contains(p, l) {
			return false
		}
		p = p.Add(dir)
	}
	return true
}

func check2(grid *game2d.GridChar, config []game2d.Pos) bool {
	var letters = []uint8("AMSMS")
	for i, c := range config {
		if !grid.Contains(c, letters[i]) {
			return false
		}
	}
	return true
}

func Part1(input string) int {
	var grid = game2d.BuildGridCharFromString(input)
	var res int
	for p := range grid.All() {
		if grid.Contains(p, 'X') {
			for n := range p.Neighbors8() {
				dir := n.Sub(p)
				if xmas(grid, n, dir) {
					res++
				}
			}
		}
	}
	return res
}

func Part2(input string) int {
	var grid = game2d.BuildGridCharFromString(input)
	var res int
	for p := range grid.All() {
		var configs = [][]game2d.Pos{
			{p, p.NW(), p.SE(), p.NE(), p.SW()},
			{p, p.NW(), p.SE(), p.SW(), p.NE()},
			{p, p.SE(), p.NW(), p.NE(), p.SW()},
			{p, p.SE(), p.NW(), p.SW(), p.NE()}}
		for _, config := range configs {
			if check2(grid, config) {
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2024 day 04 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
