package game2d

import (
	"fmt"
	"iter"
	"math"
	"strings"
)

// ---------------------------
// Map to uint8 Representation
// ---------------------------

type Grid[T comparable] struct {
	m          map[Pos]T
	minX, maxX int
	minY, maxY int
	ToString   func(c T) string
}

type GridChar = Grid[uint8]

func NewGrid[T comparable](toString func(c T) string) *Grid[T] {
	return &Grid[T]{
		m:        make(map[Pos]T),
		minX:     math.MaxInt,
		maxX:     math.MinInt,
		minY:     math.MaxInt,
		maxY:     math.MinInt,
		ToString: toString,
	}
}

func NewGridChar() *GridChar {
	return NewGrid(func(c uint8) string { return string(c) })
}

func BuildGridFunc[T comparable](lines []string, convert func(c int32) T, toString func(c T) string) *Grid[T] {
	var grid = NewGrid[T](toString)
	for j, l := range lines {
		for i, c := range l {
			grid.Set(i, j, convert(c))
		}
	}
	return grid
}

func BuildGridCharFromString(s string) *GridChar {
	lines := strings.Split(s, "\n")
	var toString = func(c uint8) string { return fmt.Sprintf("%c", c) }
	return BuildGridFunc(lines, func(c int32) uint8 { return uint8(c) }, toString)
}

func DisplayMap(g *GridChar, empty uint8) {
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if v, ok := g.Get(x, y); ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Printf("%c", empty)
			}
		}
		fmt.Println()
	}

}

func (g *Grid[T]) String() string {
	var sb strings.Builder
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if v, ok := g.Get(x, y); ok {
				sb.WriteString(g.ToString(v))
			} else {
				sb.WriteString(".")
			}
		}
		if y < g.maxY {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (g *Grid[T]) Get(x, y int) (T, bool) {
	return g.GetPos(Pos{X: x, Y: y})
}

func (g *Grid[T]) Set(x, y int, value T) {
	g.SetPos(Pos{X: x, Y: y}, value)
}

func (g *Grid[T]) GetPos(pos Pos) (T, bool) {
	v, ok := g.m[pos]
	return v, ok
}

func (g *Grid[T]) ClearPos(pos Pos) {
	delete(g.m, pos)
}

func (g *Grid[T]) SetPos(pos Pos, value T) {
	if _, ok := g.m[pos]; !ok {
		g.minX = min(pos.X, g.minX)
		g.maxX = max(pos.X, g.maxX)
		g.minY = min(pos.Y, g.minY)
		g.maxY = max(pos.Y, g.maxY)
	}
	g.m[pos] = value
}

func (g *Grid[T]) Contains(pos Pos, value T) bool {
	v, ok := g.GetPos(pos)
	return ok && v == value
}

func (g *Grid[T]) GetBounds() (minX, maxX, minY, maxY int) {
	return g.minX, g.maxX, g.minY, g.maxY
}

func (g *Grid[T]) All() iter.Seq2[Pos, T] {
	return func(yield func(Pos, T) bool) {
		for p, v := range g.m {
			if !yield(p, v) {
				return
			}
		}
	}
}

func (g *Grid[T]) AllPos() iter.Seq[Pos] {
	return func(yield func(Pos) bool) {
		for p := range g.m {
			if !yield(p) {
				return
			}
		}
	}
}

func (g *Grid[T]) AllByRow() iter.Seq2[Pos, T] {
	return func(yield func(Pos, T) bool) {
		for j := g.minY; j <= g.maxY; j++ {
			for i := g.minX; i <= g.maxX; i++ {
				if v, ok := g.Get(i, j); ok {
					if !yield(Pos{i, j}, v) {
						return
					}
				}
			}
		}
	}
}
