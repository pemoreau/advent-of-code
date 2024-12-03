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

const LUMBER = '#'
const TREE = '|'
const OPEN = '.'

func step(m game2d.MatrixChar) game2d.MatrixChar {
	var newM = game2d.NewMatrix[uint8](m.LenX(), m.LenY())

	for p := range m.AllPos() {
		var countLumber = 0
		var countTree = 0
		for n := range p.Neighbors8() {
			if !m.IsValidPos(n) {
				continue
			}
			if m.GetPos(n) == LUMBER {
				countLumber++
			} else if m.GetPos(n) == TREE {
				countTree++
			}
		}
		var e = m.GetPos(p)
		if e == OPEN && countTree >= 3 {
			newM.SetPos(p, TREE)
		} else if e == TREE && countLumber >= 3 {
			newM.SetPos(p, LUMBER)
		} else if e == LUMBER && (countLumber == 0 || countTree == 0) {
			newM.SetPos(p, OPEN)
		} else {
			newM.SetPos(p, e)
		}
	}
	return newM
}

func countTreeLumber(m game2d.MatrixChar) (int, int) {
	var countLumber = 0
	var countTree = 0
	for p := range m.AllPos() {
		if m.GetPos(p) == LUMBER {
			countLumber++
		} else if m.GetPos(p) == TREE {
			countTree++
		}
	}
	return countTree, countLumber
}

func Part1(input string) int {
	var matrix = game2d.BuildMatrixCharFromString(input)
	for range 10 {
		matrix = step(matrix)
	}
	tree, lumber := countTreeLumber(matrix)

	return lumber * tree
}

func Part2(input string) int {
	var matrix = game2d.BuildMatrixCharFromString(input)
	var N = 1000000000
	var period = 28
	var n = 1000 + (N-1000)%period
	for range n {
		matrix = step(matrix)
	}
	var tree, lumber = countTreeLumber(matrix)
	return lumber * tree
}

func main() {
	fmt.Println("--2018 day 18 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
