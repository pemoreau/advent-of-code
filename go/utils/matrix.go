package utils

import (
	"fmt"
	"strings"
)

type DigitMatrix [][]uint8

type IntMatrix [][]int
type Pos struct{ X, Y int }

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
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
