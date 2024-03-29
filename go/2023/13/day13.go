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

func isVerticalMirror(m game2d.MatrixChar, x int) bool {
	diff := min(m.MaxX()-(x+1), x)
	for y := 0; y <= m.MaxY(); y++ {
		for i := 0; i <= diff; i++ {
			if m[y][x+1+i] != m[y][x-i] {
				return false
			}
		}
	}
	return true
}

func findVerticalMirror(m game2d.MatrixChar, old int) int {
	for x := 0; x <= m.MaxX()-1; x++ {
		if x+1 != old && isVerticalMirror(m, x) {
			return x + 1
		}
	}
	return 0
}

func isHorizontalMirror(m game2d.MatrixChar, y int) bool {
	diff := min(m.MaxY()-(y+1), y)
	for x := 0; x <= m.MaxX(); x++ {
		for i := 0; i <= diff; i++ {
			if m[y+1+i][x] != m[y-i][x] {
				return false
			}
		}
	}
	return true
}

func findHorizontalMirror(m game2d.MatrixChar, old int) int {
	for y := 0; y <= m.MaxY()-1; y++ {
		if y+1 != old && isHorizontalMirror(m, y) {
			return y + 1
		}
	}
	return 0
}

func computeScore(m game2d.MatrixChar) int {
	var h = findHorizontalMirror(m, -1)
	var v = findVerticalMirror(m, -1)
	return v + (h * 100)
}

func trySwap(m game2d.MatrixChar, x, y int, c uint8, h, v int) int {
	old := m[y][x]
	m[y][x] = c
	if nh := findHorizontalMirror(m, h); nh > 0 {
		m[y][x] = old
		return 100 * nh
	} else if nv := findVerticalMirror(m, v); nv > 0 {
		m[y][x] = old
		return nv
	}
	m[y][x] = old
	return 0
}

func findSmudge(g game2d.MatrixChar) int {
	h := findHorizontalMirror(g, -1)
	v := findVerticalMirror(g, -1)

	for y, l := range g {
		for x, c := range l {
			switch c {
			case '#':
				if r := trySwap(g, x, y, '.', h, v); r > 0 {
					return r
				}
			case '.':
				if r := trySwap(g, x, y, '#', h, v); r > 0 {
					return r
				}
			default:
				panic("invalid char")
			}
		}
	}
	panic("no smudge found")
}

func solve(input string, score func(matrix game2d.MatrixChar) int) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")

	var res int
	for _, p := range parts {
		m := game2d.BuildMatrixCharFromString(p)
		res += score(m)
	}

	return res
}

func Part1(input string) int {
	return solve(input, computeScore)
}

func Part2(input string) int {
	return solve(input, findSmudge)
}

func main() {
	fmt.Println("--2023 day 13 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
