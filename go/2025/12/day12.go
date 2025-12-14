package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
)

type Present = *game2d.MatrixChar

func parse(input string) (presents []Present, puzzles []puzzle) {
	input = strings.TrimSuffix(input, "\n")
	var parts = strings.Split(input, "\n\n")
	for _, part := range parts[:len(parts)-1] {
		var lines = strings.Split(part, "\n")
		lines = lines[1:]
		var toString = func(c uint8) string { return fmt.Sprintf("%c", c) }
		var present = game2d.BuildMatrixFunc(lines, func(c int32) uint8 { return uint8(c) }, toString)
		presents = append(presents, present)
	}

	for _, line := range strings.Split(parts[len(parts)-1], "\n") {
		var x, y, a, b, c, d, e, f int
		fmt.Sscanf(line, "%dx%d: %d %d %d %d %d %d", &x, &y, &a, &b, &c, &d, &e, &f)
		puzzles = append(puzzles, puzzle{x, y, []int{a, b, c, d, e, f}})
	}
	return
}

type puzzle struct {
	sizeX, sizeY int
	quantity     []int
}

// rotations/flip
// 0 init
// 1, 2, 3 left
// 4 flip
// 5, 6, 7 flip, left

type board struct {
	sizeX, sizeY int
}

func nbPixel(p *game2d.MatrixChar) int {
	var res int
	for _, c := range p.All() {
		if c == '#' {
			res++
		}
	}
	return res
}

func filterPuzzle(puzzles []puzzle, presents []*game2d.MatrixChar) []puzzle {
	var res []puzzle
	for _, p := range puzzles {
		var nb int
		for i, qt := range p.quantity {
			nb += nbPixel(presents[i]) * qt
		}
		if nb <= p.sizeX*p.sizeY {
			res = append(res, p)
			fmt.Printf("diff = %d\n", p.sizeX*p.sizeY-nb)
		}
	}
	return res
}

func Part1(input string) int {
	var presents, puzzles = parse(input)
	fmt.Printf("#puzzles: %v\n", len(puzzles))
	puzzles = filterPuzzle(puzzles, presents)
	fmt.Printf("#puzzles: %v\n", len(puzzles))
	return len(puzzles)
}

func Part2(input string) int {
	//input = strings.TrimSuffix(input, "\n")
	//var lines = strings.Split(input, "\n")
	var res int
	return res
}

func main() {
	fmt.Println("--2025 day 12 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
