package utils

import "strings"

type DigitMatrix [][]uint8

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
