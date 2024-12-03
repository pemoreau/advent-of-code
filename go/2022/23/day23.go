package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"math"
	"time"
)

//go:embed input.txt
var inputDay string

type State struct {
	directions []string
}

func step(g *game2d.GridChar, s *State) (*game2d.GridChar, int) {
	project := map[game2d.Pos][]game2d.Pos{} // Pos -> []Pos who want to go there

	var isElve = func(p game2d.Pos) bool {
		v, ok := g.GetPos(p)
		return ok && v == '#'
	}

	for p, v := range g.All() {
		if v != '#' {
			continue
		}
		move := false
		for n := range p.Neighbors8() {
			if isElve(n) {
				move = true
				break
			}
		}
		if move {
			explore := map[string][]game2d.Pos{
				"N": {{p.X, p.Y - 1}, {p.X + 1, p.Y - 1}, {p.X - 1, p.Y - 1}},
				"S": {{p.X, p.Y + 1}, {p.X + 1, p.Y + 1}, {p.X - 1, p.Y + 1}},
				"W": {{p.X - 1, p.Y}, {p.X - 1, p.Y - 1}, {p.X - 1, p.Y + 1}},
				"E": {{p.X + 1, p.Y}, {p.X + 1, p.Y - 1}, {p.X + 1, p.Y + 1}},
			}
			for _, d := range s.directions {
				move = false
				if !isElve(explore[d][0]) && !isElve(explore[d][1]) && !isElve(explore[d][2]) {
					//fmt.Println("project", p, d)
					project[explore[d][0]] = append(project[explore[d][0]], p)
					move = true
					break
				}
			}
			if !move {
				project[p] = append(project[p], p)
			}
		} else {
			//fmt.Println("do not move", p)
			project[p] = append(project[p], p)
		}
	}

	cpt := 0
	newGrid := game2d.NewGrid[uint8](g.ToString)
	for p, ps := range project {
		if len(ps) == 1 {
			from := ps[0]
			v, _ := g.GetPos(from)
			newGrid.SetPos(p, v)
			if p != from {
				cpt++
			}
		} else {
			for _, from := range ps {
				// nobody move
				v, _ := g.GetPos(from)
				newGrid.SetPos(from, v)
			}
		}
	}
	s.directions = append(s.directions[1:], s.directions[0])
	return newGrid, cpt
}

func Part1(input string) int {
	grid := game2d.BuildGridCharFromString(input)
	state := State{directions: []string{"N", "S", "W", "E"}}
	cpt := 1
	for i := 0; i < 10; i++ {
		grid, cpt = step(grid, &state)
	}

	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	cpt = 0
	for p, v := range grid.All() {
		if v == '#' {
			minX = min(p.X, minX)
			minY = min(p.Y, minY)
			maxX = max(p.X, maxX)
			maxY = max(p.Y, maxY)
			cpt++
		}
	}
	return (maxX-minX+1)*(maxY-minY+1) - cpt
}

func Part2(input string) int {
	grid := game2d.BuildGridCharFromString(input)
	state := State{directions: []string{"N", "S", "W", "E"}}
	round := 0
	cpt := 1
	for cpt > 0 {
		round++
		grid, cpt = step(grid, &state)
	}
	return round
}

func main() {
	fmt.Println("--2022 day 23 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
