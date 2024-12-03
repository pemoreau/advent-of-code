package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/stack"
	"time"
)

//go:embed input.txt
var inputDay string

func explore(input string) *game2d.GridChar {
	var grid = game2d.NewGridChar()
	var pos = game2d.Pos{X: 0, Y: 0}
	grid.SetPos(pos, 'X')
	var stack = stack.NewStack[game2d.Pos]()
	stack.Push(pos)

	for _, c := range input {
		switch c {
		case '^', '$':
		case 'N':
			pos.Y--
			grid.SetPos(pos, '-')
			pos.Y--
			grid.SetPos(pos, '.')
		case 'S':
			pos.Y++
			grid.SetPos(pos, '-')
			pos.Y++
			grid.SetPos(pos, '.')
		case 'E':
			pos.X++
			grid.SetPos(pos, '|')
			pos.X++
			grid.SetPos(pos, '.')
		case 'W':
			pos.X--
			grid.SetPos(pos, '|')
			pos.X--
			grid.SetPos(pos, '.')
		case '(':
			stack.Push(pos)
		case ')':
			pos, _ = stack.Pop()
		case '|':
			pos, _ = stack.Peek()
		}
	}
	var minX, maxX, minY, maxY = grid.GetBounds()
	for y := minY; y <= maxY; y++ {
		grid.Set(minX-1, y, '#')
		grid.Set(maxX+1, y, '#')
	}
	for x := minX; x <= maxX; x++ {
		grid.Set(x, minY-1, '#')
		grid.Set(x, maxY+1, '#')
	}
	return grid
}

func neighbors(p game2d.Pos, grid *game2d.GridChar) []game2d.Pos {
	var result []game2d.Pos
	var neighbors41 = []game2d.Pos{p.N(), p.S(), p.W(), p.E()}
	var neighbors42 = []game2d.Pos{p.N().N(), p.S().S(), p.W().W(), p.E().E()}
	for i := range len(neighbors41) {
		if v, ok := grid.GetPos(neighbors41[i]); ok && (v == '|' || v == '-') {
			if v, ok := grid.GetPos(neighbors42[i]); ok && v == '.' {
				result = append(result, neighbors42[i])
			}
		}
	}
	return result
}

func solve(input string) (int, int) {
	var grid = explore(input)

	var max = 0
	var count = 0
	var start = game2d.Pos{X: 0, Y: 0}
	for goal, v := range grid.All() {
		if v == '.' {
			neighborsF := func(s game2d.Pos) []game2d.Pos { return neighbors(s, grid) }
			costF := func(from, to game2d.Pos) int { return 1 }
			goalF := func(s game2d.Pos) bool { return s == goal }
			heuristicF := func(s game2d.Pos) int { return game2d.ManhattanDistance(s, goal) / 2 }

			_, cost := utils.Astar[game2d.Pos](start, goalF, neighborsF, costF, heuristicF)
			if cost > max {
				max = cost
			}
			if cost >= 1000 {
				count++
			}
		}
	}

	return max, count
}

func Part1(input string) int {
	max, _ := solve(input)
	return max
}

func Part2(input string) int {
	_, count := solve(input)
	return count
}

func main() {
	fmt.Println("--2018 day 20 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
