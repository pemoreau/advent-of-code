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

const (
	north = 0
	east  = 1
	south = 2
	west  = 3
)

type Grid = game2d.MatrixChar

type Guard struct {
	game2d.Pos
	dir int
}

func (g *Guard) front(grid *Grid) (game2d.Pos, uint8, bool) {
	var neighbors = []game2d.Pos{g.Pos.N(), g.Pos.E(), g.Pos.S(), g.Pos.W()}
	var nextPos = neighbors[g.dir]
	if grid.IsValidPos(nextPos) {
		return nextPos, grid.GetPos(nextPos), true
	}
	return nextPos, 0, false
}

func (g *Guard) move(grid *Grid, visited set.Set[Guard]) (inside bool, loop bool) {
	if p, v, inside := g.front(grid); !inside {
		return false, false
	} else if v != '#' && visited.Contains(Guard{p, g.dir}) {
		return true, true
	} else if v == '#' {
		g.dir = (g.dir + 1) % 4
		return true, false
	} else {
		g.Pos = p
		return true, false
	}
}

func (g *Guard) run(grid *Grid) ([]Guard, bool) {
	var path []Guard
	var visited = set.NewSet[Guard]()
	path = append(path, *g)
	visited.Add(*g)
	// returns true if loop
	for {
		var inside, loop bool
		inside, loop = g.move(grid, visited)
		if loop {
			return path, true
		} else if !inside {
			return path, false
		}
		path = append(path, *g)
		visited.Add(*g)
	}
}

func findStart(grid *Grid) Guard {
	if p, ok := grid.Find('^'); ok {
		return Guard{p, north}
	}
	return Guard{}
}

func Part1(input string) int {
	var grid = game2d.BuildMatrixCharFromString(input)
	var guard = findStart(grid)
	var track, _ = guard.run(grid)
	var res = set.NewSet[game2d.Pos]()
	for _, g := range track {
		res.Add(g.Pos)
	}
	return res.Len()
}

func Part2(input string) int {
	var grid = game2d.BuildMatrixCharFromString(input)
	var start = findStart(grid)
	var track, _ = start.run(grid)

	var plot = set.NewSet[game2d.Pos]()
	var tried = set.NewSet[game2d.Pos]()
	var previous = start
	for _, g := range track {
		var p = g.Pos
		var guard = previous
		previous = g
		if v := grid.GetPos(p); v != '#' {
			if tried.Contains(p) {
				continue
			}
			tried.Add(p)
			grid.SetPos(p, '#')
			if _, loop := guard.run(grid); loop {
				plot.Add(p)
			}
			grid.SetPos(p, '.')
		}
	}
	return plot.Len()
}

func main() {
	fmt.Println("--2024 day 06 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
