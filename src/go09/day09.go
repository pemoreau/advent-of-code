package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

type matrix [][]uint8
type Pos struct {
	i, j int
}
type set map[Pos]struct{}

func BuildSet() set {
	return make(map[Pos]struct{})
}

func (s set) Add(value Pos) {
	s[value] = struct{}{}
}

func (s set) Contains(value Pos) bool {
	_, ok := s[value]
	return ok
}
func (s set) Len() int {
	return len(s)
}

func BuildMatrix(lines []string) matrix {
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

func neighboors(m matrix, i, j int) []Pos {
	res := []Pos{}
	pos := []Pos{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	for _, p := range pos {
		if p.i >= 0 && p.i < len(m) && p.j >= 0 && p.j < len(m[p.i]) {
			res = append(res, p)
		}
	}
	return res
}

func smallerThanNeighboors(m matrix, i, j int) bool {
	for _, p := range neighboors(m, i, j) {
		if !(m[i][j] < m[p.i][p.j]) {
			return false
		}
	}
	return true
}

func explore(m matrix) [](set) {
	collectedBassin := [](set){}
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 9 {
				// already visited: skip
			} else {
				newBasin := BuildSet()
				collectNeighboors(Pos{i, j}, m, newBasin)
				collectedBassin = append(collectedBassin, newBasin)
			}
		}
	}
	return collectedBassin
}

func collectNeighboors(p Pos, m matrix, collected set) {
	if collected.Contains(p) {
		return
	}
	if m[p.i][p.j] == 9 {
		return
	}
	collected.Add(p)
	m[p.i][p.j] = 9 // mark as visited
	for _, n := range neighboors(m, p.i, p.j) {
		collectNeighboors(n, m, collected)
	}
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := BuildMatrix(lines)
	res := 0
	for i := range m {
		for j := range m[i] {
			if smallerThanNeighboors(m, i, j) {
				res += int(m[i][j] + 1)
			}
		}
	}
	return res
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := BuildMatrix(lines)
	duplicate := make(matrix, len(m))
	for i := range m {
		duplicate[i] = make([]uint8, len(m[i]))
		copy(duplicate[i], m[i])
	}

	collectedBassin := explore(duplicate)
	sizes := []int{}
	for _, s := range collectedBassin {
		sizes = append(sizes, s.Len())
	}
	sort.Ints(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func main() {
	content, _ := ioutil.ReadFile("../../inputs/day09.txt")

	start := time.Now()
	fmt.Println("part1: ", Part1(string(content)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(content)))
	fmt.Println(time.Since(start))
}
