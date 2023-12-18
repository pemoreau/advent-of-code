package game2d

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
)

// ---------------------
// Matrix Representation
// m[line][column]
// 0 1 2
// 1
// 2
// ---------------------

type Matrix[T any] [][]T
type MatrixChar = Matrix[uint8]
type MatrixDigit = Matrix[uint8]

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

func BuildMatrixDigitFromString(s string) MatrixDigit {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	return BuildMatrixFunc(lines, func(c int32) uint8 { return uint8(c - '0') })
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
