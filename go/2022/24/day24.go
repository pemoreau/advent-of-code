package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input_test.txt
var input_day string

const (
	DIR = 4
	X   = 6
	Y   = 4
)

type Blizzard [DIR][][]bool

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

func (b *Blizzard) display(p utils.Pos) {
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
	pos  utils.Pos
	time int
}

func (s State) String() string {
	return fmt.Sprintf("%v", s.pos)
}

//func solve(s State, neighbors func(s State) []State, goal func(s State) bool) int {
//	todo := []State{s}
//
//	var min int = math.MaxInt
//	for len(todo) > 0 {
//		s = todo[0]
//
//		todo = todo[1:]
//		next := neighbors(s)
//
//		fmt.Printf("time = %d -- neighborsF (%v) --> %v\n", s.time, s.pos, next)
//		s.blizzard.display(s.pos)
//		fmt.Println("len(todo) = ", len(todo))
//
//		//for len(todo) > 0 && todo[0].time == s.time {
//		//	s = todo[0]
//		//	todo = todo[1:]
//		//	next = append(next, neighbors(s)...)
//		//}
//		//fmt.Println("len(next) = ", len(next))
//		//
//		//if len(todo) > 0 {
//		//	panic("should not happen")
//		//}
//
//		//next = removeDuplicates(next)
//
//		for _, n := range next {
//			if goal(n) {
//				if n.time < min {
//					min = n.time
//				}
//			} else {
//				todo = append(todo, n)
//			}
//		}
//	}
//	return int(min)
//}

func neighbors(s State, blizzards []Blizzard) []State {
	i, j := s.pos.X, s.pos.Y
	explore := []utils.Pos{{i, j}, {i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	res := []State{}
	newTime := s.time + 1
	b := blizzards[newTime%len(blizzards)]
	start := utils.Pos{0, -1}
	goal := utils.Pos{X: len(b[2]) - 1, Y: len(b[0]) - 1 + 1}
	for _, p := range explore {
		if p == start || p == goal {
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
	//fmt.Printf("time = %d -- neighborsF (%d,%d) --> %v\n", s.time, i, j, res)
	//newB.display(p)
	return res
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	grid := utils.BuildGrid(lines)

	//utils.DisplayMap(grid, '.')
	minX, maxX, minY, maxY := utils.GridBounds(grid)
	//fmt.Println(maxX-minX-1, maxY-minY-1)

	b := newBlizzard(minX, maxX, minY, maxY)
	for p, v := range grid {
		if v == '^' || v == 'v' || v == '<' || v == '>' {
			b.add(p.X-1, p.Y-1, v)
		}
	}

	blizzards := []Blizzard{*b}
	lenX := maxX - minX - 1
	lenY := maxY - minY - 1

	b.display(utils.Pos{0, -1})
	for i := 0; i < utils.LCM(lenX, lenY); i++ {
		//fmt.Println("step", i)
		b = b.step()
		blizzards = append(blizzards, *b)
		//b.display(utils.Pos{0, -1})
	}

	start := utils.Pos{0, -1}
	goal := utils.Pos{maxX - minX - 2, maxY - minY - 1}

	neighborsF := func(s State) []State {
		return neighbors(s, blizzards)
	}

	goalF := func(s State) bool {
		return s.pos == goal
	}

	costF := func(from, to State) int {
		return 1
	}

	heuristicF := func(s State) int {
		return utils.ManhattanDistance(s.pos, goal)
	}

	path, cost := utils.Path[State](State{start, 0}, goalF, neighborsF, costF, heuristicF)
	//res := solve(State{start, *b, 0}, neighborsF, goalF)
	//fmt.Println(res)
	fmt.Println("path", path)
	fmt.Println("cost", cost)
	return 0
}

func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	return 0

}

func main() {
	fmt.Println("--2022 day 24 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
