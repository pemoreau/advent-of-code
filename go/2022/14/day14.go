package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Pos struct {
	X, Y int
}

func buildPair(pair string) Pos {
	xy := strings.Split(pair, ",")
	X, _ := strconv.Atoi(xy[0])
	Y, _ := strconv.Atoi(xy[1])
	return Pos{X: X, Y: Y}
}

type Grid map[Pos]uint8

func buildGrid(input string) (Grid, int) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	maxY := math.MinInt
	grid := map[Pos]uint8{}
	for _, line := range lines {
		pairs := strings.Split(line, " -> ")
		for i := 0; i < len(pairs)-1; i++ {
			p1 := buildPair(pairs[i])
			p2 := buildPair(pairs[i+1])
			for x := min(p1.X, p2.X); x <= max(p1.X, p2.X); x++ {
				for y := min(p1.Y, p2.Y); y <= max(p1.Y, p2.Y); y++ {
					grid[Pos{X: x, Y: y}] = '#'
				}
				maxY = max(p1.Y, maxY)
				maxY = max(p2.Y, maxY)
			}
		}
	}
	return grid, maxY
}

// return true when the grid is stable
func step(grid map[Pos]uint8, p *Pos) bool {
	explore := []Pos{{X: p.X, Y: p.Y + 1}, {X: p.X - 1, Y: p.Y + 1}, {X: p.X + 1, Y: p.Y + 1}}
	for _, e := range explore {
		_, ok := grid[e]
		if !ok {
			*p = e
			return false
		}
	}
	grid[*p] = 'o'
	return true
}

func display(grid map[Pos]uint8, p Pos) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for p := range grid {
		minX = min(p.X, minX)
		minY = min(p.Y, minY)
		maxX = max(p.X, maxX)
		maxY = max(p.Y, maxY)
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if p.X == x && p.Y == y {
				fmt.Print("+")
			} else {
				if v, ok := grid[Pos{X: x, Y: y}]; ok {
					fmt.Printf("%c", v)
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}

}

func Part1(input string) int {
	grid, maxY := buildGrid(input)
	unit := 0
	stable := false
	for !stable {
		n := len(grid)
		rest := false
		p := Pos{X: 500, Y: 0}
		for !rest {
			rest = step(grid, &p)
			//fmt.Println("units", unit)
			if p.Y > maxY {
				//fmt.Println("falling")
				return unit
			}
		}
		if len(grid) > n {
			unit++
			stable = false
		}
	}

	return unit
}

func Part2(input string) int {
	grid, maxY := buildGrid(input)
	unit := 0
	stable := false
	for !stable {
		n := len(grid)
		rest := false
		p := Pos{X: 500, Y: 0}
		if _, ok := grid[p]; ok {
			//fmt.Println("blocked")
			return unit
		}
		for !rest {
			rest = step(grid, &p)
			//fmt.Println("units", unit)
			if p.Y == maxY+1 {
				//fmt.Println("stopped")
				grid[p] = 'o'
				rest = true
			}
		}

		if len(grid) > n {
			unit++
			stable = false
		}
	}

	return unit
}

func main() {
	fmt.Println("--2022 day 14 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
