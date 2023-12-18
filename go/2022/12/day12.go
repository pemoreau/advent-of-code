package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := game2d.BuildMatrixInt[int](lines)
	from := search('S', m)
	to := search('E', m)
	m[from.Y][from.X] = 'a'
	m[to.Y][to.X] = 'z'

	neighborsF := func(s game2d.Pos) []game2d.Pos { return neighbors(m, s.X, s.Y) }
	costF := func(from, to game2d.Pos) int { return 1 }
	goalF := func(s game2d.Pos) bool { return s == to }
	heuristicF := func(s game2d.Pos) int { return m[to.Y][to.X] - m[from.Y][from.X] }
	_, cost := utils.Astar[game2d.Pos](from, goalF, neighborsF, costF, heuristicF)

	return cost
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := game2d.BuildMatrixInt[int](lines)

	from := search('S', m)
	to := search('E', m)
	m[from.Y][from.X] = 'a'
	m[to.Y][to.X] = 'z'

	neighborsF := func(s game2d.Pos) []game2d.Pos { return neighbors2(m, s.X, s.Y) }
	costF := func(from, to game2d.Pos) int { return cost2(from, to, m) }
	goalF := func(s game2d.Pos) bool { return s == to }
	heuristicF := func(s game2d.Pos) int { return m[to.Y][to.X] - m[from.Y][from.X] }
	_, cost := utils.Astar[game2d.Pos](from, goalF, neighborsF, costF, heuristicF)

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

func search(v int, m game2d.Matrix[int]) game2d.Pos {
	for j, l := range m {
		for i, c := range l {
			if c == v {
				return game2d.Pos{i, j}
			}
		}
	}
	return game2d.Pos{}
}

func neighbors(m game2d.Matrix[int], i, j int) []game2d.Pos {
	pos := []game2d.Pos{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	var res []game2d.Pos
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

func neighbors2(m game2d.Matrix[int], i, j int) []game2d.Pos {
	n := neighbors(m, i, j)
	if m[j][i] == 'a' {
		a := search('a', m)
		return append(n, a)
	}
	return n
}

func cost2(from, to game2d.Pos, m game2d.Matrix[int]) int {
	if m[from.Y][from.X] == 'a' && m[to.Y][to.X] == 'a' {
		return 0
	}
	return 1
}
