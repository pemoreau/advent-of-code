package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"sort"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func neighbors(p game2d.Pos, grid *game2d.GridInt, t int) []game2d.Pos {
	var res []game2d.Pos
	var minX, maxX, minY, maxY = grid.GetBounds()
	for n := range p.Neighbors4() {
		if n.X >= minX && n.X <= maxX && n.Y >= minY && n.Y <= maxY {
			if v, ok := grid.GetPos(n); !ok || v > t {
				res = append(res, n)
			}
		}
	}
	return res
}

func parse(input string) (grid *game2d.GridInt, positions []game2d.Pos) {
	var lines = strings.Split(input, "\n")
	grid = game2d.NewGridInt()
	for i, line := range lines {
		var p game2d.Pos
		fmt.Sscanf(line, "%d,%d", &p.X, &p.Y)
		grid.Set(p.X, p.Y, i+1)
		positions = append(positions, p)
	}
	return
}

func minimumCost(grid *game2d.GridInt, start, goal game2d.Pos, t int) int {
	neighborsF := func(s game2d.Pos) []game2d.Pos { return neighbors(s, grid, t) }
	costF := func(from, to game2d.Pos) int { return 1 }
	goalF := func(s game2d.Pos) bool { return s == goal }
	heuristicF := func(s game2d.Pos) int { return 0 }

	_, cost := utils.Astar[game2d.Pos](start, goalF, neighborsF, costF, heuristicF)
	return cost
}

func Part1(input string) int {
	var grid, _ = parse(input)
	var start = game2d.Pos{0, 0}
	var goal = game2d.Pos{grid.MaxX(), grid.MaxY()}
	return minimumCost(grid, start, goal, 1024)
}

func Part2(input string) string {
	var grid, positions = parse(input)

	var start = game2d.Pos{0, 0}
	var goal = game2d.Pos{grid.MaxX(), grid.MaxY()}

	// find the smallest t such that connected is false using binary search
	var res = sort.Search(len(positions), func(t int) bool { return minimumCost(grid, start, goal, t) == 0 })
	return fmt.Sprintf("%d,%d", positions[res-1].X, positions[res-1].Y)

	//var t = 1024
	//for {
	//	if !connected(grid, start, goal, t) {
	//		return fmt.Sprintf("%d,%d", positions[t-1].X, positions[t-1].Y)
	//	}
	//	t++
	//}
}

func main() {
	fmt.Println("--2024 day 18 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
