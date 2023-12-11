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

func BuildGrid(lines []string) (utils.Grid, []int) {
	grid := make(utils.Grid)
	var emptyLines []int
	for j, l := range lines {
		var emptyLine = true
		for i, c := range l {
			if c == '#' {
				grid[utils.Pos{X: i, Y: j}] = uint8(c)
				emptyLine = false
			}
		}
		if emptyLine {
			emptyLines = append(emptyLines, j)
		}
	}
	return grid, emptyLines
}

func virtualPos(pos utils.Pos, emptyLines []int, emptyColumns []int, factor int) utils.Pos {
	var x = pos.X
	var y = pos.Y
	var addX, addY int
	for _, l := range emptyLines {
		if l < y {
			addY += factor - 1
		}
	}
	for _, c := range emptyColumns {
		if c < x {
			addX += factor - 1
		}
	}
	return utils.Pos{X: x + addX, Y: y + addY}
}

func solve(input string, factor int) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var grid, emptyLines = BuildGrid(lines)
	var emptyColumns []int
	minX, maxX, minY, maxY := utils.GridBounds(grid)

	for i := minX; i <= maxX; i++ {
		var emptyColumn = true
		for j := minY; j <= maxY; j++ {
			if c, ok := grid[utils.Pos{X: i, Y: j}]; ok && c == '#' {
				emptyColumn = false
			}
		}
		if emptyColumn {
			emptyColumns = append(emptyColumns, i)
		}
	}

	var g = make([]utils.Pos, 0, len(grid))
	for k := range grid {
		g = append(g, virtualPos(k, emptyLines, emptyColumns, factor))
	}

	var res int
	for i := 0; i < len(g); i++ {
		for j := i + 1; j < len(g); j++ {
			res += utils.ManhattanDistance(g[i], g[j])
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
