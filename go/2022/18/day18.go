package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Pos struct {
	X, Y, Z int
}

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
func (p Pos) neighbors27() []Pos {
	return []Pos{
		{p.X + 1, p.Y, p.Z},
		{p.X - 1, p.Y, p.Z},
		{p.X, p.Y + 1, p.Z},
		{p.X, p.Y - 1, p.Z},
		{p.X, p.Y, p.Z + 1},
		{p.X, p.Y, p.Z - 1},
		{p.X + 1, p.Y + 1, p.Z},
		{p.X + 1, p.Y - 1, p.Z},
		{p.X - 1, p.Y + 1, p.Z},
		{p.X - 1, p.Y - 1, p.Z},
		{p.X + 1, p.Y, p.Z + 1},
		{p.X + 1, p.Y, p.Z - 1},
		{p.X - 1, p.Y, p.Z + 1},
		{p.X - 1, p.Y, p.Z - 1},
		{p.X, p.Y + 1, p.Z + 1},
		{p.X, p.Y + 1, p.Z - 1},
		{p.X, p.Y - 1, p.Z + 1},
		{p.X, p.Y - 1, p.Z - 1},
		{p.X + 1, p.Y + 1, p.Z + 1},
		{p.X + 1, p.Y + 1, p.Z - 1},
		{p.X + 1, p.Y - 1, p.Z + 1},
		{p.X + 1, p.Y - 1, p.Z - 1},
		{p.X - 1, p.Y + 1, p.Z + 1},
		{p.X - 1, p.Y + 1, p.Z - 1},
		{p.X - 1, p.Y - 1, p.Z + 1},
		{p.X - 1, p.Y - 1, p.Z - 1},
	}
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	g := set.Set[Pos]{}
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
	g := set.Set[Pos]{}
	var x, y, z int
	for _, line := range lines {
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		g.Add(Pos{x, y, z})
	}

	xmin, xmax, ymin, ymax, zmin, zmax := minmax(g)
	freeCells := collectFree(g, xmin, xmax, ymin, ymax, zmin, zmax)
	res := count(g, freeCells)
	return res
}

func count(component set.Set[Pos], free set.Set[Pos]) int {
	res := 0
	for p := range component {
		for _, q := range p.neighbors() {
			if !component.Contains(q) && free.Contains(q) {
				res++
			}
		}
	}
	return res
}

func findFree(g set.Set[Pos], xmin, xmax, ymin, ymax, zmin, zmax int) (Pos, bool) {
	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			for z := zmin; z <= zmax; z++ {
				p := Pos{x, y, z}
				if !g.Contains(p) {
					return p, true
				}
			}
		}
	}
	return Pos{}, false
}

func minmax(g set.Set[Pos]) (int, int, int, int, int, int) {
	var xmin, xmax, ymin, ymax, zmin, zmax = math.MaxInt, math.MinInt, math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for p := range g {
		xmin = min(xmin, p.X)
		xmax = max(xmax, p.X)
		ymin = min(ymin, p.Y)
		ymax = max(ymax, p.Y)
		zmin = min(zmin, p.Z)
		zmax = max(zmax, p.Z)
	}
	return xmin, xmax, ymin, ymax, zmin, zmax
}

func collectFree(g set.Set[Pos], xmin, xmax, ymin, ymax, zmin, zmax int) set.Set[Pos] {
	todo := set.Set[Pos]{}
	res := set.Set[Pos]{}
	p, ok := findFree(g, xmin, xmax, ymin, ymax, zmin, zmax)
	if !ok {
		return res
	}
	todo.Add(p)
	for len(todo) > 0 {
		p = todo.Pop()
		//fmt.Println("p", p)
		if g.Contains(p) || res.Contains(p) {
			continue
		}
		if p.X < xmin-1 || p.X > xmax+1 || p.Y < ymin-1 || p.Y > ymax+1 || p.Z < zmin-1 || p.Z > zmax+1 {
			continue
		}
		res.Add(p)
		for _, q := range p.neighbors() {
			todo.Add(q)
		}
	}
	return res
}

func main() {
	fmt.Println("--2022 day 18 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
