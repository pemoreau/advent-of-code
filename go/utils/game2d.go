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
// m[line][column]
// 0 1 2
// 1
// 2
// ---------------------

type Matrix[T any] [][]T
type MatrixChar = Matrix[uint8]

func BuildMatrixChar(lines []string) MatrixChar {
	m := make([][]uint8, len(lines))
	for j, l := range lines {
		m[j] = []uint8(l)
	}
	return m
}

func BuildMatrixCharFromString(s string) MatrixChar {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	return BuildMatrixChar(lines)
}

func BuildMatrixInt[T constraints.Integer](lines []string) Matrix[T] {
	return BuildMatrixFunc[T](lines, func(c int32) T { return T(c) })
}

func BuildMatrixFunc[T any](lines []string, convert func(c int32) T) Matrix[T] {
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

func (m Matrix[any]) LenY() int {
	if len(m) == 0 {
		return 0
	}
	return len(m)
}

func (m Matrix[any]) LenX() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

func (m Matrix[any]) MaxY() int {
	return m.LenY() - 1
}

func (m Matrix[any]) MaxX() int {
	return m.LenX() - 1
}

func (m Matrix[any]) IsValidPos(pos Pos) bool {
	return pos.Y >= 0 && pos.Y < len(m) && pos.X >= 0 && pos.X < len(m[pos.Y])
}

func (m Matrix[any]) RotateLeft() Matrix[any] {
	var m2 = make([][]any, len(m[0]))
	for i := range m2 {
		m2[i] = make([]any, len(m))
	}
	for j, l := range m {
		for i, c := range l {
			m2[m.MaxX()-i][j] = c
		}
	}
	return m2
}

func (m Matrix[any]) RotateRight() Matrix[any] {
	var m2 = make([][]any, len(m[0]))
	for i := range m2 {
		m2[i] = make([]any, len(m))
	}
	for j, l := range m {
		for i, c := range l {
			m2[i][m.MaxY()-j] = c
		}
	}
	return m2
}

func (m Matrix[any]) Transpose() Matrix[any] {
	var m2 = make([][]any, len(m[0]))
	for i := range m2 {
		m2[i] = make([]any, len(m))
	}
	for j, l := range m {
		for i, c := range l {
			m2[i][j] = c
		}
	}
	return m2
}

func (m Matrix[any]) String() string {
	var sb strings.Builder
	for i, l := range m {
		for _, c := range l {
			sb.WriteString(fmt.Sprintf("%c", c))
		}
		if i < len(m)-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
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

func BuildGridUp(lines []string) Grid {
	grid := make(Grid)
	var height = len(lines)
	for j, l := range lines {
		for i, c := range l {
			grid[Pos{X: i, Y: height - j - 1}] = uint8(c)
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
