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

func cheatN(grid *game2d.MatrixChar, pos game2d.Pos, n int) set.Set[game2d.Pos] {
	var res = set.NewSet[game2d.Pos]()

	for deltaX := 0; deltaX <= n; deltaX++ {
		for deltaY := 0; deltaY <= n-deltaX; deltaY++ {
			if deltaX+deltaY >= 2 && deltaX+deltaY <= n {
				var positions = []game2d.Pos{
					{X: pos.X + deltaX, Y: pos.Y + deltaY},
					{X: pos.X - deltaX, Y: pos.Y - deltaY},
					{X: pos.X + deltaX, Y: pos.Y - deltaY},
					{X: pos.X - deltaX, Y: pos.Y + deltaY},
				}
				for _, p := range positions {
					if grid.IsValidPos(p) && grid.GetPos(p) != '#' {
						res.Add(p)
					}
				}
			}
		}
	}
	return res
}

func solve(input string, n int, m int) int {
	var grid = game2d.BuildMatrixCharFromString(input)
	var start, _ = grid.Find('S')
	var end, _ = grid.Find('E')

	var path []game2d.Pos
	var todo []game2d.Pos
	var visited = make(map[game2d.Pos]int)
	todo = append(todo, start)
	for len(todo) > 0 {
		var pos = todo[0]
		todo = todo[1:]
		if _, ok := visited[pos]; ok {
			continue
		}

		visited[pos] = len(path)
		path = append(path, pos)

		if pos == end {
			break
		}
		for neighbour := range pos.Neighbors4() {
			if grid.GetPos(neighbour) != '#' {
				todo = append(todo, neighbour)
			}
		}
	}

	var res int
	//for i, pos := range path[:len(path)-m] {
	//	//for c := range cheatN(grid, pos, n) {
	//	for _, c := range path[i+m:] {
	//		if game2d.ManhattanDistance(pos, c) > n {
	//			continue
	//		}
	//		var index2, _ = visited[c]
	//		var win = index2 - i - game2d.ManhattanDistance(pos, c)
	//		if win >= m {
	//			res++
	//		}
	//	}
	//}
	for i, pos := range path[:len(path)-m] {
		for j, c := range path[i+m:] {
			var dist = game2d.ManhattanDistance(pos, c)
			if dist > n {
				continue
			}
			var index2 = i + m + j
			var win = index2 - i - dist
			if win >= m {
				res++
			}
		}
	}

	return res
}

func Part1(input string) int {
	return solve(input, 2, 100)
}

func Part2(input string) int {
	return solve(input, 20, 100)
}

func main() {
	fmt.Println("--2024 day 20 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
