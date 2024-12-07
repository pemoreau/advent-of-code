package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"sort"
	"time"
)

//go:embed sample.txt
var inputTest string

func neighboors(m *game2d.MatrixDigit, p game2d.Pos) []game2d.Pos {
	var res []game2d.Pos
	for n := range p.Neighbors4() {
		if m.IsValidPos(n) && m.GetPos(n) != 9 {
			// collect neighboors different from 9
			res = append(res, n)
		}
	}
	return res
}

func collectNeighboors(p game2d.Pos, m *game2d.MatrixDigit) int {
	toVisit := []game2d.Pos{p}
	collected := 0
	for len(toVisit) > 0 {
		p := toVisit[0]
		toVisit = toVisit[1:]
		if m.GetPos(p) == 9 {
			// already visited: skip
		} else {
			m.SetPos(p, 9)
			collected += 1
			toVisit = append(toVisit, neighboors(m, p)...)
		}
	}
	return collected
}

func smallerThanNeighboors(m *game2d.MatrixDigit, p game2d.Pos) bool {
	for n := range p.Neighbors4() {
		if m.IsValidPos(n) && !(m.GetPos(p) < m.GetPos(n)) {
			return false
		}
	}
	return true
}

func Part1(input string) int {
	var m = game2d.BuildMatrixDigitFromString(input)
	var res int
	for p, v := range m.All() {
		if smallerThanNeighboors(m, p) {
			res = res + int(v+1)
		}
	}
	return res
}

func Part2(input string) int {
	var m = game2d.BuildMatrixDigitFromString(input)

	var sizes []int
	for p, v := range m.All() {
		if v != 9 {
			size := collectNeighboors(p, m)
			sizes = append(sizes, size)
		}
	}

	sort.Ints(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func main() {

	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
