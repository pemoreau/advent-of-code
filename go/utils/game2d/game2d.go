package game2d

import (
	"fmt"
	"math"
)

// ---------------------------
// Map to uint8 Representation
// ---------------------------

type Grid map[Pos]uint8

func BuildGrid(lines []string) Grid {
	grid := make(Grid)
	for j, l := range lines {
		for i, c := range l {
			grid[Pos{X: i, Y: j}] = uint8(c)
		}
	}
	return grid
}

func BuildGridUp(lines []string) Grid {
	grid := make(Grid)
	var height = len(lines)
	for j, l := range lines {
		for i, c := range l {
			grid[Pos{X: i, Y: height - j - 1}] = uint8(c)
		}
	}
	return grid
}

func DisplayMap(grid map[Pos]uint8, empty uint8) {
	minX, maxX, minY, maxY := GridBounds(grid)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if v, ok := grid[Pos{X: x, Y: y}]; ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Printf("%c", empty)
			}
		}
		fmt.Println()
	}

}

func GridBounds(grid map[Pos]uint8) (minX, maxX, minY, maxY int) {
	minX, maxX = math.MaxInt, math.MinInt
	minY, maxY = math.MaxInt, math.MinInt
	for p := range grid {
		minX = min(p.X, minX)
		maxX = max(p.X, maxX)
		minY = min(p.Y, minY)
		maxY = max(p.Y, maxY)
	}
	return minX, maxX, minY, maxY
}
