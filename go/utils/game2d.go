package utils

import (
	"fmt"
	"math"
	"strings"
)

type Pos struct{ X, Y int }

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// ---------------------
// Matrix Representation
// ---------------------

type DigitMatrix [][]uint8

type IntMatrix [][]int

func BuildDigitMatrix(lines []string) DigitMatrix {
	m := make([][]uint8, len(lines))
	for j, l := range lines {
		l = strings.TrimSpace(l)
		m[j] = make([]uint8, len(l))
		for i, c := range l {
			m[j][i] = uint8(c - '0')
		}
	}
	return m
}

func BuildIntMatrix(lines []string) IntMatrix {
	m := make([][]int, len(lines))
	for j, l := range lines {
		l = strings.TrimSpace(l)
		m[j] = make([]int, len(l))
		for i, c := range l {
			m[j][i] = int(c)
		}
	}
	return m
}

// ---------------------------
// Map to uint8 Representation
// ---------------------------

type Grid map[Pos]uint8

func BuildGrid(lines []string) Grid {
	grid := make(Grid)
	for j, l := range lines {
		for i, c := range l {
			grid[Pos{X: i, Y: j}] = uint8(c)
		}
	}
	return grid
}

func (p Pos) Neighbors4() []Pos {
	return []Pos{
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
	}
}

func (p Pos) Neighbors8() []Pos {
	return []Pos{
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y - 1},
		{X: p.X + 1, Y: p.Y - 1},
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y + 1},
	}
}

func DisplayMap(grid map[Pos]uint8, empty uint8) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for p := range grid {
		minX = Min(p.X, minX)
		minY = Min(p.Y, minY)
		maxX = Max(p.X, maxX)
		maxY = Max(p.Y, maxY)
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if v, ok := grid[Pos{X: x, Y: y}]; ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Printf("%c", empty)
			}
		}
		fmt.Println()
	}

}
