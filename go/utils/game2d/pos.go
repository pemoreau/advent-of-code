package game2d

import (
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
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

func ShoeslaceFormulae(polygon []Pos) int {
	var sum1, sum2 int
	for i := 0; i < len(polygon)-1; i++ {
		sum1 += polygon[i].X * polygon[i+1].Y
		sum2 += polygon[i].Y * polygon[i+1].X
	}
	sum1 += polygon[len(polygon)-1].X * polygon[0].Y
	sum2 += polygon[len(polygon)-1].Y * polygon[0].X
	return utils.Abs(sum1-sum2) / 2
}

func PolygonArea(polygon []Pos) int {
	// Pick's theorem
	// A = i + b/2 - 1
	// where:
	// - A is the area of the polygon,
	// - i is the number of interior points with integer coordinates,
	// - b is the number of boundary points with integer coordinates.
	//
	// So, i = A - b/2 + 1
	var b int
	for i := 0; i < len(polygon)-1; i++ {
		b += ManhattanDistance(polygon[i], polygon[i+1])
	}
	b += ManhattanDistance(polygon[0], polygon[len(polygon)-1])
	i := ShoeslaceFormulae(polygon) - b/2 + 1
	return i + b
}
