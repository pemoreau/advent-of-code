package game2d

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"iter"
	"strings"
)

// ---------------------
// Matrix Representation
// m[line][column]
// 0 1 2
// 1
// 2
// ---------------------

type Matrix[T comparable] struct {
	width, height int
	data          []T
	toString      func(c T) string
}

type MatrixChar = Matrix[uint8]
type MatrixDigit = Matrix[uint8]

func NewMatrix[T comparable](width, height int) *Matrix[T] {
	return &Matrix[T]{width: width, height: height, data: make([]T, width*height)}
}

func Clone[T comparable](m Matrix[T]) *Matrix[T] {
	var m2 = NewMatrix[T](m.width, m.height)
	m2.toString = m.toString
	copy(m2.data, m.data)
	return m2
}

func BuildMatrixFunc[T comparable](lines []string, convert func(c int32) T, toString func(c T) string) *Matrix[T] {
	if len(lines) == 0 {
		panic("matrix: empty input")
	}
	var m = NewMatrix[T](len(lines[0]), len(lines))
	m.toString = toString
	for j, l := range lines {
		l = strings.TrimSpace(l)
		for i, c := range l {
			m.Set(i, j, convert(c))
		}
	}
	return m
}

func BuildMatrixInt[T constraints.Integer](lines []string) *Matrix[T] {
	var toString = func(c T) string { return fmt.Sprintf("%d", c) }
	return BuildMatrixFunc[T](lines, func(c int32) T { return T(c) }, toString)
}

func BuildMatrixCharFromString(s string) *MatrixChar {
	//s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	var toString = func(c uint8) string { return fmt.Sprintf("%c", c) }
	return BuildMatrixFunc(lines, func(c int32) uint8 { return uint8(c) }, toString)
}

func BuildMatrixDigitFromString(s string) *MatrixDigit {
	//s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	var toString = func(c uint8) string { return string('0' + c) }
	return BuildMatrixFunc(lines, func(c int32) uint8 { return uint8(c - '0') }, toString)
}

func (m *Matrix[T]) Get(x, y int) T {
	if x < 0 || x >= m.width {
		panic(fmt.Sprintf("matrix: x index out of range [%d] with width %d", x, m.width))
	}
	if y < 0 || y >= m.height {
		panic(fmt.Sprintf("matrix: y index out of range [%d] with height %d", y, m.height))
	}
	return m.getUnchecked(x, y)
}

func (m *Matrix[T]) getUnchecked(x, y int) T {
	return m.data[x+y*m.width]
}

func (m *Matrix[T]) Set(x, y int, value T) {
	if x < 0 || x >= m.width {
		panic(fmt.Sprintf("matrix: x index out of range [%d] with width %d", x, m.width))
	}
	if y < 0 || y >= m.height {
		panic(fmt.Sprintf("matrix: y index out of range [%d] with height %d", y, m.height))
	}
	m.setUnchecked(x, y, value)
}

func (m *Matrix[T]) setUnchecked(x, y int, value T) {
	m.data[x+y*m.width] = value
}

func (m *Matrix[T]) GetPos(pos Pos) T {
	return m.Get(pos.X, pos.Y)
}

func (m *Matrix[T]) SetPos(pos Pos, value T) {
	m.Set(pos.X, pos.Y, value)
}

func (m *Matrix[T]) LenY() int {
	return m.height
}

func (m *Matrix[T]) LenX() int {
	return m.width
}

func (m *Matrix[T]) MaxY() int {
	return m.LenY() - 1
}

func (m *Matrix[T]) MaxX() int {
	return m.LenX() - 1
}

func (m *Matrix[T]) IsValidPos(pos Pos) bool {
	return pos.Y >= 0 && pos.Y < m.LenY() && pos.X >= 0 && pos.X < m.LenX()
}

func (m *Matrix[T]) RotateLeft() {
	//var m2 = Matrix[T]{width: m.height, height: m.width, data: make([]T, len(m.data)), toString: m.toString}
	//for j := range m.LenY() {
	//	for i := range m.LenX() {
	//		m2.setUnchecked(j, m.width-i-1, m.getUnchecked(i, j))
	//	}
	//}
	//return m2
	var data = make([]T, len(m.data))
	for j := range m.LenY() {
		for i := range m.LenX() {
			data[j+m.height*(m.width-i-1)] = m.getUnchecked(i, j)
		}
	}
	m.width, m.height = m.height, m.width
	m.data = data
}

func (m *Matrix[T]) RotateRight() {
	//var m2 = Matrix[T]{width: m.height, height: m.width, data: make([]T, len(m.data)), toString: m.toString}
	//for j := range m.LenY() {
	//	for i := range m.LenX() {
	//		m2.setUnchecked(m.height-j-1, i, m.getUnchecked(i, j))
	//	}
	//}
	//return m2
	var data = make([]T, len(m.data))
	for j := range m.LenY() {
		for i := range m.LenX() {
			data[m.height-j-1+m.height*i] = m.getUnchecked(i, j)
		}
	}
	m.width, m.height = m.height, m.width
	m.data = data
}

func (m *Matrix[T]) Transpose() {
	//var m2 = Matrix[any]{width: m.height, height: m.width, data: make([]any, len(m.data)), toString: m.toString}
	//for j := 0; j < m.height; j++ {
	//	for i := 0; i < m.width; i++ {
	//		m2.setUnchecked(j, i, m.getUnchecked(i, j))
	//	}
	//}
	var data = m.data
	m.data = make([]T, len(m.data))
	m.height, m.width = m.width, m.height
	rx := 0
	for _, e := range data {
		m.data[rx] = e
		rx += m.width
		if rx >= len(m.data) {
			rx -= len(m.data) - 1
		}
	}
}

func (m *Matrix[T]) String() string {
	var sb strings.Builder
	for j := range m.LenY() {
		for i := range m.LenX() {
			//sb.WriteString(fmt.Sprintf("%c", m.getUnchecked(i, j)))
			sb.WriteString(m.toString(m.getUnchecked(i, j)))
		}
		if j < m.MaxY() {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (m *Matrix[T]) All() iter.Seq2[Pos, T] {
	return func(yield func(Pos, T) bool) {
		for j := range m.LenY() {
			for i := range m.LenX() {
				if !yield(Pos{i, j}, m.getUnchecked(i, j)) {
					return
				}
			}
		}
	}
}

func (m *Matrix[T]) AllPos() iter.Seq[Pos] {
	return func(yield func(Pos) bool) {
		for j := range m.LenY() {
			for i := range m.LenX() {
				if !yield(Pos{i, j}) {
					return
				}
			}
		}
	}
}

func (m *Matrix[T]) Find(value T) (Pos, bool) {
	for pos, e := range m.All() {
		if e == value {
			return pos, true
		}
	}
	return Pos{}, false
}
