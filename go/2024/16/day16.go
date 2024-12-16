package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"time"
)

//go:embed sample.txt
var inputTest string

type state struct {
	game2d.Pos
	dir int
}

const (
	N = iota
	E
	S
	W
)

func neighbors(m *game2d.MatrixChar, s state) []state {
	var delta = []game2d.Pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	var res []state
	res = append(res, state{s.Pos, (s.dir + 1) % 4})
	res = append(res, state{s.Pos, (s.dir + 3) % 4})
	var nextPos = s.Pos.Add(delta[s.dir])
	if v := m.GetPos(nextPos); v == '.' || v == 'S' || v == 'E' {
		res = append(res, state{nextPos, s.dir})
	}
	return res
}

func cost(from, to state) int {
	if to.dir == (from.dir+1)%4 || to.dir == (from.dir+3)%4 {
		return 1000
	}
	if game2d.ManhattanDistance(from.Pos, to.Pos) == 1 {
		return 1
	}
	return 0
}

func Part1(input string) int {
	m := game2d.BuildMatrixCharFromString(input)
	from, _ := m.Find('S')
	to, _ := m.Find('E')
	//m.Set(from.X, from.Y, 'a')
	//m.Set(to.X, to.Y, 'z')

	var start = state{from, E}

	neighborsF := func(s state) []state { return neighbors(m, s) }
	costF := func(from, to state) int { return cost(from, to) }
	goalF := func(s state) bool { return s.Pos == to }
	heuristicF := func(s state) int { return 0 }
	_, cost := utils.Astar[state](start, goalF, neighborsF, costF, heuristicF)

	return cost
}

func Part2(input string) int {
	m := game2d.BuildMatrixCharFromString(input)
	from, _ := m.Find('S')
	to, _ := m.Find('E')
	var start = state{from, E}

	neighborsF := func(s state) []state { return neighbors(m, s) }
	costF := func(from, to state) int { return cost(from, to) }
	goalF := func(s state) bool { return s.Pos == to }
	heuristicF := func(s state) int { return 0 }
	_, best := utils.Astar[state](start, goalF, neighborsF, costF, heuristicF)

	var res = set.Set[game2d.Pos]{}
	for p := range m.AllPos() {
		//fmt.Println("p: ", p)
		if p == from || m.GetPos(p) == '#' {
			continue
		}

		path1, cost1 := utils.Astar[state](start, func(s state) bool { return s.Pos == p }, neighborsF, costF, heuristicF)
		last := path1[0]
		fmt.Println("last: ", last)
		fmt.Println(p, cost1)
		_, cost2 := utils.Astar[state](state{p, last.dir}, func(s state) bool { return s.Pos == to }, neighborsF, costF, heuristicF)
		if cost1+cost2 == best {
			res.Add(p)
		}

	}

	return res.Len() + 1
}

// 588 too low
func main() {
	fmt.Println("--2024 day 16 solution--")
	//var inputDay = utils.Input()
	var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
