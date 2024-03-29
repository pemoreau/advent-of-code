package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type maxHeight struct {
	N, S, E, W int8
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	matrix := game2d.BuildMatrixInt[int](lines)

	maxY := len(matrix)
	maxX := len(matrix[0])

	m := make([][]maxHeight, maxY)
	for j := 0; j < maxY; j++ {
		m[j] = make([]maxHeight, maxX)
		for i := 0; i < maxX; i++ {
			N := int8(-1)
			if j > 0 {
				N = int8(max(int(m[j-1][i].N), matrix[j-1][i]))
			}
			W := int8(-1)
			if i > 0 {
				W = int8(max(int(m[j][i-1].W), matrix[j][i-1]))
			}
			m[j][i] = maxHeight{N: N, W: W}
		}
	}
	for j := maxY - 1; j >= 0; j-- {
		for i := maxX - 1; i >= 0; i-- {
			S := int8(-1)
			if j < maxY-1 {
				S = int8(max(int(m[j+1][i].S), matrix[j+1][i]))
			}
			E := int8(-1)
			if i < maxX-1 {
				E = int8(max(int(m[j][i+1].E), matrix[j][i+1]))
			}
			m[j][i] = maxHeight{S: S, E: E, N: m[j][i].N, W: m[j][i].W}
		}
	}

	res := 0
	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			v := int8(matrix[j][i])
			if v > m[j][i].N || v > m[j][i].S || v > m[j][i].E || v > m[j][i].W {
				res++
			}
		}
	}

	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	matrix := game2d.BuildMatrixInt[int](lines)

	res := 0
	maxY := len(matrix)
	maxX := len(matrix[0])
	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			v := matrix[j][i]
			n, s, e, w := 0, 0, 0, 0
			for k := j - 1; k >= 0; k-- {
				n++
				if v <= matrix[k][i] {
					break
				}
			}
			for k := j + 1; k < maxY; k++ {
				s++
				if v <= matrix[k][i] {
					break
				}
			}
			for k := i - 1; k >= 0; k-- {
				w++
				if v <= matrix[j][k] {
					break
				}
			}
			for k := i + 1; k < maxX; k++ {
				e++
				if v <= matrix[j][k] {
					break
				}
			}
			res = max(res, n*s*e*w)
		}
	}
	return res

}

func main() {
	fmt.Println("--2022 day 08 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
