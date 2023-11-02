package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

//go:embed input.txt
var input_day string

type Pos struct {
	i, j int
}

func neighboors(m utils.Matrix[uint8], i, j int) []Pos {
	res := make([]Pos, 0, 4)
	pos := []Pos{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	for _, p := range pos {
		if p.j >= 0 && p.j < len(m) && p.i >= 0 && p.i < len(m[p.j]) && m[p.j][p.i] != 9 {
			// collect neighboors different from 9
			res = append(res, p)
		}
	}
	return res
}

func collectNeighboors(p Pos, m utils.Matrix[uint8]) int {
	toVisit := []Pos{p}
	collected := 0
	for len(toVisit) > 0 {
		p := toVisit[0]
		toVisit = toVisit[1:]
		if m[p.j][p.i] == 9 {
			// already visited: skip
		} else {
			m[p.j][p.i] = 9
			collected += 1
			toVisit = append(toVisit, neighboors(m, p.i, p.j)...)
		}
	}
	return collected
}

func smallerThanNeighboors(m utils.Matrix[uint8], i, j int) bool {
	pos := []Pos{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	for _, p := range pos {
		if p.j >= 0 && p.j < len(m) && p.i >= 0 && p.i < len(m[p.j]) && !(m[j][i] < m[p.j][p.i]) {
			return false
		}
	}
	return true
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := utils.BuildConvertMatrix[uint8](lines, func(c int32) uint8 { return uint8(c - '0') })
	res := 0
	for j := range m {
		for i := range m[j] {
			if smallerThanNeighboors(m, i, j) {
				res += int(m[j][i] + 1)
			}
		}
	}
	return res
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := utils.BuildConvertMatrix[uint8](lines, func(c int32) uint8 { return uint8(c - '0') })

	var sizes []int
	for j := range m {
		for i := range m[j] {
			if m[j][i] != 9 {
				size := collectNeighboors(Pos{i, j}, m)
				sizes = append(sizes, size)
			}
		}
	}

	sort.Ints(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func main() {

	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
