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

type Guard struct {
	game2d.Pos
	dir int
}

func front(grid game2d.GridChar, guard Guard) (game2d.Pos, uint8, bool) {
	var nextPos = []game2d.Pos{guard.Pos.N(), guard.Pos.E(), guard.Pos.S(), guard.Pos.W()}
	v, inside := grid.GetPos(nextPos[guard.dir])
	return nextPos[guard.dir], v, inside
}

func move(grid *game2d.GridChar, g Guard, path set.Set[Guard]) (guard Guard, inside bool, loop bool) {
	if p, v, inside := front(*grid, g); !inside {
		return g, false, false
	} else if v != '#' && path.Contains(Guard{p, g.dir}) {
		return g, true, true
	} else if v == '#' {
		return Guard{g.Pos, (g.dir + 1) % 4}, true, false
	} else {
		return Guard{p, g.dir}, true, false
	}
}

func run(grid *game2d.GridChar, g Guard) (set.Set[Guard], bool) {
	var path = set.NewSet[Guard]()
	path.Add(g)
	// returns true if loop
	for {
		var inside, loop bool
		g, inside, loop = move(grid, g, path)
		if loop {
			return path, true
		} else if !inside {
			return path, false
		}
		path.Add(g)
	}
}

func computeTrack(grid *game2d.GridChar, guard Guard) set.Set[game2d.Pos] {
	var plot = set.NewSet[game2d.Pos]()
	path, _ := run(grid, guard)
	for p := range path.All() {
		plot.Add(p.Pos)
	}
	return plot
}

func findStart(grid *game2d.GridChar) Guard {
	for p, e := range grid.All() {
		if e == '^' {
			return Guard{p, north}
		}
	}
	return Guard{}
}

func Part1(input string) int {
	var grid = game2d.BuildGridCharFromString(input)
	var guard = findStart(grid)
	return computeTrack(grid, guard).Len()
}

func Part2(input string) int {
	var grid = game2d.BuildGridCharFromString(input)
	var guard = findStart(grid)
	var track = computeTrack(grid, guard)

	var plot = set.NewSet[game2d.Pos]()
	for p := range track.All() {
		if v, _ := grid.GetPos(p); v == '.' {
			grid.SetPos(p, '#')
			if _, loop := run(grid, guard); loop {
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
