package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"strings"
	"time"
)

//go:embed input_test.txt
var input_day string

type Pos struct {
	X, Y, Z int
}

// type Grid map[Pos]uint8
func (p Pos) neighbors() []Pos {
	return []Pos{
		{p.X + 1, p.Y, p.Z},
		{p.X - 1, p.Y, p.Z},
		{p.X, p.Y + 1, p.Z},
		{p.X, p.Y - 1, p.Z},
		{p.X, p.Y, p.Z + 1},
		{p.X, p.Y, p.Z - 1},
	}
}

//type Grid utils.Set[Pos]

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	g := utils.Set[Pos]{}
	var x, y, z int
	for _, line := range lines {
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		g.Add(Pos{x, y, z})
	}

	var count int
	for p := range g {
		n := 0
		for _, q := range p.neighbors() {
			if g.Contains(q) {
				n++
			}
		}
		count += (6 - n)
	}
	return count
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	g := utils.Set[Pos]{}
	var x, y, z int
	for _, line := range lines {
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		g.Add(Pos{x, y, z})
	}

	var xmin, xmax, ymin, ymax, zmin, zmax = math.MaxInt, math.MinInt, math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for p := range g {
		xmin = utils.Min(xmin, p.X) - 1
		xmax = utils.Max(xmax, p.X) + 1
		ymin = utils.Min(ymin, p.Y) - 1
		ymax = utils.Max(ymax, p.Y) + 1
		zmin = utils.Min(zmin, p.Z) - 1
		zmax = utils.Max(zmax, p.Z) + 1
	}

	components := []utils.Set[Pos]{}
	for p := range g {
		if !contains(components, p) {
			newComponent := collectOccupied(p, g)
			components = append(components, newComponent)
		}
	}

	for _, c := range components {
		fmt.Println("component size", len(c))
	}

	return 0
}

func contains(components []utils.Set[Pos], p Pos) bool {
	for _, c := range components {
		if c.Contains(p) {
			return true
		}
	}
	return false
}

func collectFree(p Pos, g utils.Set[Pos]) utils.Set[Pos] {
	todo := utils.Set[Pos]{}
	res := utils.Set[Pos]{}
	todo.Add(p)
	for len(todo) > 0 {
		p = todo.Element()
		if !g.Contains(p) && !res.Contains(p) {
			res.Add(p)
			for _, q := range p.neighbors() {
				if !g.Contains(q) && !res.Contains(q) {
					todo.Add(q)
				}
			}
		}
	}
	return res
}

func collectOccupied(p Pos, g utils.Set[Pos]) utils.Set[Pos] {
	todo := utils.Set[Pos]{}
	res := utils.Set[Pos]{}
	todo.Add(p)
	for len(todo) > 0 {
		p = todo.Element()
		if g.Contains(p) && !res.Contains(p) {
			res.Add(p)
			for _, q := range p.neighbors() {
				if g.Contains(q) && !res.Contains(q) {
					todo.Add(q)
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2022 day 18 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
