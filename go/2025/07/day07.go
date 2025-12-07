package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
)

func Part1(input string) int {
	var grid = game2d.BuildMatrixCharFromString(input)
	var start, _ = grid.Find('S')

	var res int
	var todo = []game2d.Pos{start}
	var visited = set.NewSet[game2d.Pos]()
	for len(todo) > 0 {
		var p = todo[0]
		todo = todo[1:]
		if visited.Contains(p) || !grid.IsValidPos(p) {
			continue
		}
		visited.Add(p)
		var c = grid.GetPos(p)
		if c == '.' || c == 'S' {
			todo = append(todo, p.S())
		} else if c == '^' {
			todo = append(todo, p.E(), p.W())
			res++
		}
	}

	return res
}

func count(grid *game2d.MatrixChar, p game2d.Pos, memo map[game2d.Pos]int) int {
	if v, ok := memo[p]; ok {
		return v
	}
	if !grid.IsValidPos(p) {
		return 1
	}

	if c := grid.GetPos(p); c == '.' || c == 'S' {
		var res = count(grid, p.S(), memo)
		memo[p] = res
		return res
	} else if c == '^' {
		var res = count(grid, p.E(), memo) + count(grid, p.W(), memo)
		memo[p] = res
		return res
	}
	return 0
}

func Part2(input string) int {
	var grid = game2d.BuildMatrixCharFromString(input)
	var start, _ = grid.Find('S')
	var memo = make(map[game2d.Pos]int)
	return count(grid, start, memo)
}

func main() {
	fmt.Println("--2025 day 07 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
