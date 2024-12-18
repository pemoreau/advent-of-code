package game2d

import (
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/set"
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
type GridInt = Grid[int]

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

func NewGridInt() *GridInt {
	return NewGrid(func(c int) string { return fmt.Sprintf("%d", c) })
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

func (g *Grid[T]) ContainsPos(pos Pos) bool {
	_, ok := g.m[pos]
	return ok
}

func (g *Grid[T]) Contains(pos Pos, value T) bool {
	v, ok := g.m[pos]
	return ok && v == value
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

func (m *Grid[T]) IsValidPos(pos Pos) bool {
	return pos.Y >= 0 && pos.Y <= m.MaxY() && pos.X >= 0 && pos.X <= m.MaxX()
}

func (g *Grid[T]) GetBounds() (minX, maxX, minY, maxY int) {
	return g.minX, g.maxX, g.minY, g.maxY
}

func (g *Grid[T]) MaxX() int {
	return g.maxX
}

func (g *Grid[T]) MaxY() int {
	return g.maxY
}

func (g *Grid[T]) MinX() int {
	return g.minX
}

func (g *Grid[T]) MinY() int {
	return g.minY
}

func (g *Grid[T]) Size() int {
	return len(g.m)
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

// Extract connected components from a grid
// AAAA
// BBCD              BB  C
// BBCC  ==>  [AAAA, BB, CC, D, EEE]
// EEEC                   C

func (g *Grid[T]) ExtractComponents() []*Grid[T] {
	var res []*Grid[T]
	var visited = set.NewSet[Pos]()
	for pos := range g.AllByRow() {
		if visited.Contains(pos) {
			continue
		}
		var piece = g.collectComponent(pos, visited)
		res = append(res, piece)
	}
	return res
}

func (g *Grid[T]) collectComponent(pos Pos, visited set.Set[Pos]) *Grid[T] {
	var firstValue, _ = g.GetPos(pos)
	var res = NewGrid[T](g.ToString)
	var todo = []Pos{pos}
	for len(todo) > 0 {
		var p = todo[0]
		todo = todo[1:]
		if visited.Contains(p) {
			continue
		}
		if g.Contains(p, firstValue) {
			visited.Add(p)
			res.SetPos(p, firstValue)
			for n := range p.Neighbors4() {
				if !visited.Contains(n) {
					todo = append(todo, n)
				}
			}
		}
	}
	return res
}

func (m *Grid[T]) RotateLeft() {
	var data = make(map[Pos]T)
	for p := range m.AllPos() {
		data[Pos{m.maxY - p.Y, p.X}] = m.m[p]
	}
	m.m = data
	m.minX, m.maxX, m.minY, m.maxY = m.maxX-m.maxY, m.maxX-m.minY, m.minX, m.maxX
}

func (m *Grid[T]) RotateRight() {
	var data = make(map[Pos]T)
	for p := range m.AllPos() {
		data[Pos{p.Y, m.maxX - p.X}] = m.m[p]
	}
	m.m = data
	m.minX, m.maxX, m.minY, m.maxY = m.minY, m.maxY, m.maxX-m.maxX, m.maxX-m.minX
}
