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

type Piece = game2d.GridChar

func perimeter(piece *Piece) int {
	var res int
	for p := range piece.AllPos() {
		for n := range p.Neighbors4() {
			if !piece.ContainsPos(n) {
				res++
			}
		}
	}
	return res
}

func nbFaces(piece *Piece) int {
	var res int
	for range 4 {
		var minX, maxX, minY, maxY = piece.GetBounds()
		for y := minY - 1; y <= maxY-1; y++ {
			res += northFrontY(piece, minX, maxX, y)
		}
		piece.RotateRight()
	}
	return res
}

func northFrontY(piece *Piece, minX, maxX, y int) int {
	var res int
	var front = false
	for x := minX - 1; x <= maxX; x++ {
		var p = game2d.Pos{x, y}
		var cond = !piece.ContainsPos(p) && piece.ContainsPos(p.S())
		if front == false && cond {
			front = true
			res++
		} else if !cond {
			front = false
		}
	}
	return res
}

func solve(input string, part2 bool) int {
	var grid = game2d.BuildGridCharFromString(input)

	var res int
	var components = grid.ExtractComponents()
	for _, c := range components {
		var area = c.Size()
		if part2 {
			res += area * nbFaces(c)
		} else {
			res += area * perimeter(c)
		}
	}

	return res
}

func Part1(input string) int {
	return solve(input, false)
}

func Part2(input string) int {
	return solve(input, true)
}

func main() {
	fmt.Println("--2024 day 12 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
