package utils

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
	"strings"
)

type Pos struct{ X, Y int }

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func ManhattanDistance(from, to Pos) int {
	absX := from.X - to.X
	if absX < 0 {
		absX = -absX
	}
	absY := from.Y - to.Y
	if absY < 0 {
		absY = -absY
	}
	return absX + absY
}

// ---------------------
// Matrix Representation
// ---------------------

type Matrix[T any] [][]T

func BuildMatrix[T constraints.Integer](lines []string) Matrix[T] {
	return BuildConvertMatrix[T](lines, func(c int32) T { return T(c) })
}

func BuildConvertMatrix[T constraints.Integer](lines []string, convert func(c int32) T) Matrix[T] {
	m := make([][]T, len(lines))
	for j, l := range lines {
		l = strings.TrimSpace(l)
		m[j] = make([]T, len(l))
		for i, c := range l {
			m[j][i] = convert(c)
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
	minX, maxX, minY, maxY := GridBounds(grid)
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

func GridBounds(grid map[Pos]uint8) (minX, maxX, minY, maxY int) {
	minX, minY, maxX, maxY = math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for p := range grid {
		minX = min(p.X, minX)
		minY = min(p.Y, minY)
		maxX = max(p.X, maxX)
		maxY = max(p.Y, maxY)
	}
	return minX, maxX, minY, maxY
}
