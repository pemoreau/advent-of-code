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
	//v, inside := grid.GetPos(nextPos)
	var x, y = nextPos.X, nextPos.Y
	if inside := x >= 0 && x <= grid.MaxX() && y >= 0 && y <= grid.MaxY(); inside {
		return nextPos, grid.Get(x, y), true
	}
	return nextPos, 0, false
}

func (g *Guard) move(grid *Grid, path set.Set[Guard]) (inside bool, loop bool) {
	if p, v, inside := g.front(grid); !inside {
		return false, false
	} else if v != '#' && path.Contains(Guard{p, g.dir}) {
		return true, true
	} else if v == '#' {
		g.dir = (g.dir + 1) % 4
		return true, false
	} else {
		g.Pos = p
		return true, false
	}
}

func (g *Guard) run(grid *Grid) (set.Set[Guard], bool) {
	var path = set.NewSet[Guard]()
	path.Add(*g)
	// returns true if loop
	for {
		var inside, loop bool
		inside, loop = g.move(grid, path)
		if loop {
			return path, true
		} else if !inside {
			return path, false
		}
		path.Add(*g)
	}
}

func (g *Guard) computeTrack(grid *Grid) set.Set[game2d.Pos] {
	var plot = set.NewSet[game2d.Pos]()
	path, _ := g.run(grid)
	for p := range path.All() {
		plot.Add(p.Pos)
	}
	return plot
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
	return guard.computeTrack(grid).Len()
}

func Part2(input string) int {
	var grid = game2d.BuildMatrixCharFromString(input)
	var start = findStart(grid)
	var guard = start
	var track = guard.computeTrack(grid)

	var plot = set.NewSet[game2d.Pos]()
	var tried = set.NewSet[game2d.Pos]()
	for p := range track.All() {
		var guard = start
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
