package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"slices"
	"time"
)

//go:embed sample.txt
var inputTest string

type PosCost struct {
	pos  game2d.Pos
	cost int
}

func exploreSinglePath(grid *game2d.GridChar, previous game2d.Pos, current game2d.Pos, cost int, part2 bool) (PosCost, bool) {
	if c, ok := grid.GetPos(current); ok && c != '#' {
		var cpt int
		for ne := range current.Neighbors4() {
			if c, ok := grid.GetPos(ne); ok && c != '#' {
				cpt++
			}
		}
		if cpt > 2 {
			return PosCost{pos: current, cost: cost}, true
		}
	}

	if !part2 {
		// cut branches in part1
		if c, ok := grid.GetPos(current); ok && c != '.' {
			if current.X > previous.X && c != '>' ||
				current.X < previous.X && c != '<' ||
				current.Y > previous.Y && c != 'v' ||
				current.Y < previous.Y && c != '^' {
				return PosCost{}, false
			}
		}
	}

	for n := range current.Neighbors4() {
		if c, ok := grid.GetPos(n); ok && c != '#' && n != previous {
			return exploreSinglePath(grid, current, n, cost+1, part2)
		}
	}

	return PosCost{pos: current, cost: cost}, true
}

func explore(neighbors Graph, p, goal game2d.Pos, visited map[game2d.Pos]bool, cost int, maxCost int) int {
	if p == goal {
		if cost > maxCost {
			maxCost = cost
		}
		return maxCost
	}

	visited[p] = true
	for _, pc := range neighbors[p] {
		if !visited[pc.pos] {
			maxCost = explore(neighbors, pc.pos, goal, visited, cost+pc.cost, maxCost)
		}
	}
	visited[p] = false
	return maxCost
}

type Graph map[game2d.Pos][]PosCost

func buildGraph(grid *game2d.GridChar, start game2d.Pos, part2 bool) Graph {
	var res = make(map[game2d.Pos][]PosCost)

	var todo []game2d.Pos
	todo = append(todo, start)

	for len(todo) > 0 {
		p := todo[0]
		todo = todo[1:]
		if c, ok := grid.GetPos(p); !ok || c == '#' {
			continue
		}
		for n := range p.Neighbors4() {
			if c, ok := grid.GetPos(n); !ok || c == '#' {
				continue
			}
			pc, ok := exploreSinglePath(grid, p, n, 1, part2)
			if ok && !slices.Contains(res[p], pc) {
				res[p] = append(res[p], pc)
				todo = append(todo, pc.pos)
			}
		}
	}

	return res
}

func solve(input string, part2 bool) int {
	grid := game2d.BuildGridCharFromString(input)

	minX, maxX, minY, maxY := grid.GetBounds()
	start := game2d.Pos{X: minX + 1, Y: minY}
	end := game2d.Pos{X: maxX - 1, Y: maxY}

	neighbors := buildGraph(grid, start, part2)

	var goal = end
	var path = 0

	if part2 && len(neighbors[end]) > 0 {
		// skip last path
		goal = neighbors[end][0].pos
		path = neighbors[end][0].cost
	}

	visited := make(map[game2d.Pos]bool)
	return explore(neighbors, start, goal, visited, path, 0)
}

func Part1(input string) int {
	return solve(input, false)
}

func Part2(input string) int {
	return solve(input, true)
}

func main() {
	fmt.Println("--2023 day 23 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
