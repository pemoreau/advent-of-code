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

func virtualPos(pos utils.Pos, emptyLines []int, emptyColumns []int, factor int) utils.Pos {
	var addX, addY int
	for _, l := range emptyLines {
		if l < pos.Y {
			addY += factor - 1
		}
	}
	for _, c := range emptyColumns {
		if c < pos.X {
			addX += factor - 1
		}
	}
	return utils.Pos{X: pos.X + addX, Y: pos.Y + addY}
}

func solve(input string, factor int) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var galaxies []utils.Pos
	var emptyLines []int
	var occupiedColumn []bool = make([]bool, len(lines[0]))
	for j, l := range lines {
		var emptyLine = true
		for i, c := range l {
			if c == '#' {
				galaxies = append(galaxies, utils.Pos{X: i, Y: j})
				emptyLine = false
				occupiedColumn[i] = true
			}
		}
		if emptyLine {
			emptyLines = append(emptyLines, j)
		}
	}

	var emptyColumns []int
	for i, c := range occupiedColumn {
		if !c {
			emptyColumns = append(emptyColumns, i)
		}
	}

	var expandedGalaxies = make([]utils.Pos, 0, len(galaxies))
	for _, g := range galaxies {
		expandedGalaxies = append(expandedGalaxies, virtualPos(g, emptyLines, emptyColumns, factor))
	}

	var res int
	for i := 0; i < len(expandedGalaxies); i++ {
		for j := i + 1; j < len(expandedGalaxies); j++ {
			res += utils.ManhattanDistance(expandedGalaxies[i], expandedGalaxies[j])
		}
	}
	return res
}

func Part1(input string) int {
	return solve(input, 2)
}
func Part2(input string) int {
	return solve(input, 1000000)
}

func main() {
	fmt.Println("--2023 day 11 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
