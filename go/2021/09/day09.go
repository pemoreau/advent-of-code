package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"sort"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type matrix [][]uint8
type Pos struct {
	i, j int
}

func buildMatrix(lines []string) matrix {
	m := make([][]uint8, len(lines))
	for i, l := range lines {
		l = strings.TrimSpace(l)
		m[i] = make([]uint8, len(l))
		for j, c := range l {
			m[i][j] = uint8(c - '0')
		}
	}
	return m
}

func neighbors(m matrix, i, j int) []Pos {
	var res []Pos
	pos := []Pos{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	for _, p := range pos {
		if p.i >= 0 && p.i < len(m) && p.j >= 0 && p.j < len(m[p.i]) {
			res = append(res, p)
		}
	}
	return res
}

func smallerThanNeighbors(m matrix, i, j int) bool {
	for _, p := range neighbors(m, i, j) {
		if !(m[i][j] < m[p.i][p.j]) {
			return false
		}
	}
	return true
}

func explore(m matrix) []set.Set[Pos] {
	var collectedBasin []set.Set[Pos]
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 9 {
				// already visited: skip
			} else {
				newBasin := set.NewSet[Pos]()
				collectNeighbors(Pos{i, j}, m, newBasin)
				collectedBasin = append(collectedBasin, newBasin)
			}
		}
	}
	return collectedBasin
}

func collectNeighbors(p Pos, m matrix, collected set.Set[Pos]) {
	if collected.Contains(p) {
		return
	}
	if m[p.i][p.j] == 9 {
		return
	}
	collected.Add(p)
	m[p.i][p.j] = 9 // mark as visited
	for _, n := range neighbors(m, p.i, p.j) {
		collectNeighbors(n, m, collected)
	}
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := buildMatrix(lines)
	res := 0
	for i := range m {
		for j := range m[i] {
			if smallerThanNeighbors(m, i, j) {
				res += int(m[i][j] + 1)
			}
		}
	}
	return res
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := buildMatrix(lines)

	collectedBasin := explore(m) // m is modified here
	sizes := make([]int, 0, len(collectedBasin))
	for _, s := range collectedBasin {
		sizes = append(sizes, s.Len())
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
