package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
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

	components := []utils.Set[Pos]{}
	for p := range g {
		if !contains(components, p) {
			newComponent := collectOccupied(p, g)
			components = append(components, newComponent)
		}
	}

	res := 0
	for _, c := range components {
		xmin, xmax, ymin, ymax, zmin, zmax := minmax(c)
		volume := (xmax - xmin + 1) * (ymax - ymin + 1) * (zmax - zmin + 1)
		occupied := len(c)
		componentFree := collectFree(g, xmin, xmax, ymin, ymax, zmin, zmax)
		free := len(componentFree)
		internal := volume - occupied - free
		fmt.Println("x", xmin, xmax, "y", ymin, ymax, "z", zmin, zmax)
		fmt.Println(c)
		fmt.Println("volume", volume, "occupied", occupied, "free", free, "internal", internal)
		res += count(c, componentFree)
	}

	//1975 too low
	return res
}

func count(component utils.Set[Pos], free utils.Set[Pos]) int {
	res := 0
	xmin, xmax, ymin, ymax, zmin, zmax := minmax(component)
	for p := range component {
		for _, q := range p.neighbors() {
			if q.X < xmin || q.X > xmax || q.Y < ymin || q.Y > ymax || q.Z < zmin || q.Z > zmax {
				res++
			} else if !component.Contains(q) && free.Contains(q) {
				res++
			}
		}
	}
	return res
}

func findFree(g utils.Set[Pos], xmin, xmax, ymin, ymax, zmin, zmax int) (Pos, bool) {
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

func contains(components []utils.Set[Pos], p Pos) bool {
	for _, c := range components {
		if c.Contains(p) {
			return true
		}
	}
	return false
}

func minmax(g utils.Set[Pos]) (int, int, int, int, int, int) {
	var xmin, xmax, ymin, ymax, zmin, zmax = math.MaxInt, math.MinInt, math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for p := range g {
		xmin = utils.Min(xmin, p.X)
		xmax = utils.Max(xmax, p.X)
		ymin = utils.Min(ymin, p.Y)
		ymax = utils.Max(ymax, p.Y)
		zmin = utils.Min(zmin, p.Z)
		zmax = utils.Max(zmax, p.Z)
	}
	return xmin, xmax, ymin, ymax, zmin, zmax
}

func collectFree(g utils.Set[Pos], xmin, xmax, ymin, ymax, zmin, zmax int) utils.Set[Pos] {
	todo := utils.Set[Pos]{}
	res := utils.Set[Pos]{}
	p, ok := findFree(g, xmin, xmax, ymin, ymax, zmin, zmax)
	if !ok {
		return res
	}
	todo.Add(p)
	for len(todo) > 0 {
		p = todo.Pop()
		//fmt.Println("p", p)
		if g.Contains(p) || res.Contains(p) {
			//fmt.Println("continue1")
			continue
		}
		if p.X < xmin || p.X > xmax || p.Y < ymin || p.Y > ymax || p.Z < zmin || p.Z > zmax {
			//fmt.Println("continue2")
			continue
		}
		//fmt.Println("adding", p)
		res.Add(p)
		for _, q := range p.neighbors() {
			//if !g.Contains(q) && !res.Contains(q) {
			todo.Add(q)
			//}
		}
	}
	return res
}

func collectOccupied(p Pos, g utils.Set[Pos]) utils.Set[Pos] {
	todo := utils.Set[Pos]{}
	res := utils.Set[Pos]{}
	todo.Add(p)
	//fmt.Println("todo", len(todo))
	for len(todo) > 0 {
		p = todo.Pop()
		if g.Contains(p) && !res.Contains(p) {
			res.Add(p)
			for _, q := range p.neighbors27() {
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
