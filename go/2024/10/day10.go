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

func countTrails(grid *game2d.MatrixDigit, p game2d.Pos) (int, int) {
	var res = make(map[game2d.Pos]int)

	var todo []game2d.Pos
	todo = append(todo, p)
	for len(todo) > 0 {
		p := todo[0]
		todo = todo[1:]
		var v = grid.GetPos(p)
		if v == 9 {
			res[p]++
			continue
		}
		for n := range p.Neighbors4() {
			if grid.IsValidPos(n) && grid.GetPos(n) == v+1 {
				todo = append(todo, n)
			}
		}
	}

	var sum int
	for _, v := range res {
		sum += v
	}

	return len(res), sum
}

func solve(input string) (int, int) {
	var grid = game2d.BuildMatrixDigitFromString(input)
	var nb, sum int
	for p, v := range grid.All() {
		if v == 0 {
			c, s := countTrails(grid, p)
			//fmt.Println("count: ", c)
			nb += c
			sum += s
		}
	}
	return nb, sum
}

func Part1(input string) int {
	nb, _ := solve(input)
	return nb
}

func Part2(input string) int {
	_, sum := solve(input)
	return sum
}

func main() {
	fmt.Println("--2024 day 10 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
