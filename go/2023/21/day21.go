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

func findStart(m *game2d.MatrixChar) (game2d.Pos, bool) {
	if p, ok := m.Find('S'); ok {
		return p, true
	}
	return game2d.Pos{}, false
}

func Part1(input string) int {
	grid := game2d.BuildMatrixCharFromString(input)
	start, _ := findStart(grid)

	elves := set.NewSet[game2d.Pos]()
	elves.Add(start)
	var reached = set.NewSet[game2d.Pos]()

	for range 64 {
		for e := range elves {
			for n := range e.Neighbors4() {
				if !grid.IsValidPos(n) {
					continue
				}
				if grid.GetPos(n) == '#' {
					continue
				}
				if reached.Contains(n) {
					continue
				}
				reached.Add(n)
			}
		}
		tmp := elves
		elves = reached
		reached = tmp
		clear(reached)
	}
	return len(elves)
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m % b
}

func posModulo(p game2d.Pos, m *game2d.MatrixChar) game2d.Pos {
	lx := m.LenX()
	ly := m.LenY()
	return game2d.Pos{X: mod(p.X, lx), Y: mod(p.Y, ly)}
}

func Part2(input string) int {
	grid := game2d.BuildMatrixCharFromString(input)
	start, _ := findStart(grid)

	elves := set.NewSet[game2d.Pos]()
	elves.Add(start)
	var reached = set.NewSet[game2d.Pos]()

	var values []int
	var delta1 []int
	var delta2 []int
	n := grid.LenX()

	COMPUTE := 3 * n //1000
	N := COMPUTE - n //500

	for range COMPUTE {
		for e := range elves {
			for n := range e.Neighbors4() {
				m := posModulo(n, grid)
				if !grid.IsValidPos(m) {
					continue
				}
				if grid.GetPos(m) == '#' {
					continue
				}
				if reached.Contains(n) {
					continue
				}
				reached.Add(n)
			}
		}
		values = append(values, len(reached))
		//fmt.Printf("values_%d=%d\t", len(values)-1, values[len(values)-1])

		delta1 = append(delta1, len(reached)-len(elves))
		//fmt.Printf("delta1_%d=%d\t", len(delta1)-1, delta1[len(delta1)-1])

		if len(delta1) > n {
			last := delta1[len(delta1)-1]
			previous := delta1[len(delta1)-1-n]
			diff := last - previous
			delta2 = append(delta2, diff)
		} else {
			delta2 = append(delta2, 0)
		}
		//fmt.Printf("delta2_%d=%d\n", len(delta2)-1, delta2[len(delta2)-1])

		tmp := elves
		elves = reached
		reached = tmp
		clear(reached)
	}
	values = values[:N]
	delta1 = delta1[:N]
	delta2 = delta2[:N]

	value_i := values[len(values)-1]

	//LAST := 5000
	LAST := 26501365
	for i := N + 1; i < LAST+1; i++ {
		delta2_i := delta2[i-n-1]
		delta1_i := delta1[i-n-1] + delta2_i
		value_i += delta1_i

		delta1 = append(delta1, delta1_i)
		delta2 = append(delta2, delta2_i)
	}

	return value_i
}

func main() {
	fmt.Println("--2023 day 21 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
