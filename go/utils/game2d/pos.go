package game2d

import (
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"iter"
)

type Pos struct{ X, Y int }

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// 0 ------> X
// |      N
// |      ↑
// |  W <- -> E
// |      ↓
// ↓      S
// Y

func (p Pos) N() Pos { return Pos{X: p.X, Y: p.Y - 1} }
func (p Pos) S() Pos { return Pos{X: p.X, Y: p.Y + 1} }
func (p Pos) W() Pos { return Pos{X: p.X - 1, Y: p.Y} }
func (p Pos) E() Pos { return Pos{X: p.X + 1, Y: p.Y} }

func (p Pos) NW() Pos { return Pos{X: p.X - 1, Y: p.Y - 1} }
func (p Pos) NE() Pos { return Pos{X: p.X + 1, Y: p.Y - 1} }
func (p Pos) SW() Pos { return Pos{X: p.X - 1, Y: p.Y + 1} }
func (p Pos) SE() Pos { return Pos{X: p.X + 1, Y: p.Y + 1} }

func ManhattanDistance(from, to Pos) int {
	return utils.Abs(from.X-to.X) + utils.Abs(from.Y-to.Y)
}

func (p Pos) Add(p2 Pos) Pos {
	return Pos{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (p Pos) Neighbors4() iter.Seq[Pos] {
	return func(yield func(Pos) bool) {
		if !yield(p.N()) {
			return
		}
		if !yield(p.S()) {
			return
		}
		if !yield(p.W()) {
			return
		}
		if !yield(p.E()) {
			return
		}
	}
}

func (p Pos) Neighbors8() iter.Seq[Pos] {
	return func(yield func(Pos) bool) {
		for _, p := range &[...]Pos{p.N(), p.S(), p.W(), p.E(), p.NW(), p.NE(), p.SW(), p.SE()} {
			if !yield(p) {
				return
			}
		}
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
