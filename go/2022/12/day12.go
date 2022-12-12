package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := utils.BuildIntMatrix(lines)
	from := search('S', m)
	to := search('E', m)
	m[from.Y][from.X] = 'a'
	m[to.Y][to.X] = 'z'
	_, cost := utils.Path(from, to, m, neighbors, cost, heuristic)
	return cost
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := utils.BuildIntMatrix(lines)

	from := search('S', m)
	to := search('E', m)
	m[from.Y][from.X] = 'a'
	m[to.Y][to.X] = 'z'
	_, cost := utils.Path(from, to, m, neighbors2, cost2, heuristic2)
	return cost
}

func main() {
	fmt.Println("--2022 day 12 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}

func search(v int, m utils.IntMatrix) utils.Pos {
	for j, l := range m {
		for i, c := range l {
			if c == v {
				return utils.Pos{i, j}
			}
		}
	}
	return utils.Pos{}
}

func neighbors(m utils.IntMatrix, i, j int) []utils.Pos {
	pos := []utils.Pos{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	res := []utils.Pos{}
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

func cost(from, to utils.Pos, m utils.IntMatrix) int {
	return 1
}

func heuristic(from, to utils.Pos, m utils.IntMatrix) int {
	//return manhattanDistance(from, to)
	return m[to.Y][to.X] - m[from.Y][from.X]
}

func neighbors2(m utils.IntMatrix, i, j int) []utils.Pos {
	n := neighbors(m, i, j)
	if m[j][i] == 'a' {
		a := search('a', m)
		return append(n, a)
	}
	return n
}

func cost2(from, to utils.Pos, m utils.IntMatrix) int {
	if m[from.Y][from.X] == 'a' && m[to.Y][to.X] == 'a' {
		return 0
	}
	return 1
}

func heuristic2(from, to utils.Pos, m utils.IntMatrix) int {
	return m[to.Y][to.X] - m[from.Y][from.X]
}
