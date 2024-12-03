package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func isVerticalMirror(m game2d.MatrixChar, x int) bool {
	diff := min(m.MaxX()-(x+1), x)
	for y := 0; y <= m.MaxY(); y++ {
		for i := 0; i <= diff; i++ {
			if m.Get(x+1+i, y) != m.Get(x-i, y) {
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
			if m.Get(x, y+1+i) != m.Get(x, y-i) {
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
	old := m.Get(x, y)
	m.Set(x, y, c)
	if nh := findHorizontalMirror(m, h); nh > 0 {
		m.Set(x, y, old)
		return 100 * nh
	} else if nv := findVerticalMirror(m, v); nv > 0 {
		m.Set(x, y, old)
		return nv
	}
	m.Set(x, y, old)
	return 0
}

func findSmudge(g game2d.MatrixChar) int {
	h := findHorizontalMirror(g, -1)
	v := findVerticalMirror(g, -1)

	for y := range g.LenY() {
		for x := range g.LenX() {
			switch g.Get(x, y) {
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
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
