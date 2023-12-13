package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := utils.BuildMatrixInt[int](lines)
	from := search('S', m)
	to := search('E', m)
	m[from.Y][from.X] = 'a'
	m[to.Y][to.X] = 'z'

	neighborsF := func(s utils.Pos) []utils.Pos { return neighbors(m, s.X, s.Y) }
	costF := func(from, to utils.Pos) int { return 1 }
	goalF := func(s utils.Pos) bool { return s == to }
	heuristicF := func(s utils.Pos) int { return m[to.Y][to.X] - m[from.Y][from.X] }
	_, cost := utils.Astar[utils.Pos](from, goalF, neighborsF, costF, heuristicF)

	return cost
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := utils.BuildMatrixInt[int](lines)

	from := search('S', m)
	to := search('E', m)
	m[from.Y][from.X] = 'a'
	m[to.Y][to.X] = 'z'

	neighborsF := func(s utils.Pos) []utils.Pos { return neighbors2(m, s.X, s.Y) }
	costF := func(from, to utils.Pos) int { return cost2(from, to, m) }
	goalF := func(s utils.Pos) bool { return s == to }
	heuristicF := func(s utils.Pos) int { return m[to.Y][to.X] - m[from.Y][from.X] }
	_, cost := utils.Astar[utils.Pos](from, goalF, neighborsF, costF, heuristicF)

	return cost
}

func main() {
	fmt.Println("--2022 day 12 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}

func search(v int, m utils.Matrix[int]) utils.Pos {
	for j, l := range m {
		for i, c := range l {
			if c == v {
				return utils.Pos{i, j}
			}
		}
	}
	return utils.Pos{}
}

func neighbors(m utils.Matrix[int], i, j int) []utils.Pos {
	pos := []utils.Pos{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	var res []utils.Pos
	for _, p := range pos {
		if p.Y >= 0 && p.Y < len(m) && p.X >= 0 && p.X < len(m[0]) {
			src := m[j][i]
			dest := m[p.Y][p.X]
			if dest-src <= 1 {
				res = append(res, p)
			}
		}
	}
	return res
}

func neighbors2(m utils.Matrix[int], i, j int) []utils.Pos {
	n := neighbors(m, i, j)
	if m[j][i] == 'a' {
		a := search('a', m)
		return append(n, a)
	}
	return n
}

func cost2(from, to utils.Pos, m utils.Matrix[int]) int {
	if m[from.Y][from.X] == 'a' && m[to.Y][to.X] == 'a' {
		return 0
	}
	return 1
}
