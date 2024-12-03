package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"time"
)

//go:embed sample.txt
var inputTest string

type Blizzard [4][][]bool

func newBlizzard(minX, maxX, minY, maxY int) *Blizzard {
	b := Blizzard{}
	lenX := maxX - minX - 1
	lenY := maxY - minY - 1
	b[0] = make([][]bool, lenY)
	b[1] = make([][]bool, lenY)
	b[2] = make([][]bool, lenX)
	b[3] = make([][]bool, lenX)
	for j := 0; j < lenY; j++ {
		b[0][j] = make([]bool, lenX)
		b[1][j] = make([]bool, lenX)
	}
	for i := 0; i < lenX; i++ {
		b[2][i] = make([]bool, lenY)
		b[3][i] = make([]bool, lenY)
	}
	return &b
}

func (b *Blizzard) add(x, y int, dir byte) {
	//0: up (^), 1: down (v), 2: left (<), or 3: right (>).
	switch dir {
	case '^':
		b[0][y][x] = true
	case 'v':
		b[1][y][x] = true
	case '<':
		b[2][x][y] = true
	case '>':
		b[3][x][y] = true
	}
}

func duplicate(s [][]bool) [][]bool {
	duplicate := make([][]bool, len(s))
	for i := range s {
		duplicate[i] = make([]bool, len(s[i]))
		copy(duplicate[i], s[i])
	}
	return duplicate
}

func (b *Blizzard) step() *Blizzard {
	newB := *b
	for i := 0; i < 4; i++ {
		newB[i] = duplicate(b[i])
	}
	newB[0] = append(newB[0][1:], newB[0][0])
	newB[1] = append(newB[1][len(newB[1])-1:], newB[1][:len(newB[1])-1]...)
	newB[2] = append(newB[2][1:], newB[2][0])
	newB[3] = append(newB[3][len(newB[3])-1:], newB[3][:len(newB[3])-1]...)
	return &newB
}

func (b *Blizzard) display(p game2d.Pos) {
	for j := 0; j < len(b[0]); j++ {
		for i := 0; i < len(b[0][j]); i++ {
			cpt := 0
			if b[0][j][i] {
				cpt++
			}
			if b[1][j][i] {
				cpt++
			}
			if b[2][i][j] {
				cpt++
			}
			if b[3][i][j] {
				cpt++
			}
			if p.X == i && p.Y == j {
				fmt.Print("E")
			} else if cpt > 1 {
				fmt.Printf("%d", cpt)
			} else if b[0][j][i] {
				fmt.Print("^")
			} else if b[1][j][i] {
				fmt.Print("v")
			} else if b[2][i][j] {
				fmt.Print("<")
			} else if b[3][i][j] {
				fmt.Print(">")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type State struct {
	pos  game2d.Pos
	time int
}

func (s State) String() string {
	return fmt.Sprintf("== time: %d == pos: %v", s.time, s.pos)
}

func neighbors(s State, blizzards []Blizzard) []State {
	i, j := s.pos.X, s.pos.Y
	explore := []game2d.Pos{{i, j}, {i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	var res []State
	newTime := s.time + 1
	b := blizzards[newTime%len(blizzards)]

	exit1 := game2d.Pos{0, -1}
	exit2 := game2d.Pos{X: len(b[2]) - 1, Y: len(b[0]) - 1 + 1}

	for _, p := range explore {
		if p == exit1 || p == exit2 {
			res = append(res, State{p, newTime})
			continue
		}
		if p.X < 0 || p.Y < 0 || p.X >= len(b[0][0]) || p.Y >= len(b[0]) {
			continue
		}
		if b[0][p.Y][p.X] || b[1][p.Y][p.X] || b[2][p.X][p.Y] || b[3][p.X][p.Y] {
			continue
		}
		res = append(res, State{p, newTime})
	}
	return res
}

func parse(input string) ([]Blizzard, game2d.Pos, game2d.Pos) {
	grid := game2d.BuildGridCharFromString(input)

	minX, maxX, minY, maxY := grid.GetBounds()
	b := newBlizzard(minX, maxX, minY, maxY)
	for p, v := range grid.All() {
		if v == '^' || v == 'v' || v == '<' || v == '>' {
			b.add(p.X-1, p.Y-1, v)
		}
	}

	blizzards := []Blizzard{*b}
	lenX := maxX - minX - 1
	lenY := maxY - minY - 1

	//fmt.Println("step", 0)
	for i := 1; i < utils.LCM(lenX, lenY); i++ {
		b = b.step()
		blizzards = append(blizzards, *b)
		//fmt.Println("step", i)
	}

	start := game2d.Pos{0, -1}
	goal := game2d.Pos{maxX - minX - 2, maxY - minY - 1}
	return blizzards, start, goal
}

func Part1(input string) int {
	blizzards, start, goal := parse(input)

	neighborsF := func(s State) []State { return neighbors(s, blizzards) }
	costF := func(from, to State) int { return 1 }
	goalF := func(s State) bool { return s.pos == goal }
	heuristicF := func(s State) int { return game2d.ManhattanDistance(s.pos, goal) }

	_, cost := utils.Astar[State](State{start, 0}, goalF, neighborsF, costF, heuristicF)
	return cost
}

func Part2(input string) int {
	blizzards, start, goal := parse(input)

	neighborsF := func(s State) []State { return neighbors(s, blizzards) }
	costF := func(from, to State) int {
		return 1
	}

	goal1 := func(s State) bool { return s.pos == goal }
	goal2 := func(s State) bool { return s.pos == start }

	heuristic1 := func(s State) int { return game2d.ManhattanDistance(s.pos, goal) }
	heuristic2 := func(s State) int { return game2d.ManhattanDistance(s.pos, start) }

	_, cost1 := utils.Astar[State](State{start, 0}, goal1, neighborsF, costF, heuristic1)
	_, cost2 := utils.Astar[State](State{goal, cost1}, goal2, neighborsF, costF, heuristic2)
	_, cost3 := utils.Astar[State](State{start, cost1 + cost2}, goal1, neighborsF, costF, heuristic1)

	return cost1 + cost2 + cost3
}

func main() {
	fmt.Println("--2022 day 24 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
