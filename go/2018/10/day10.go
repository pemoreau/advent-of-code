package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Point struct {
	x, y   int
	vx, vy int
}

func parseInput(input string) []Point {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var points []Point

	for _, l := range lines {
		var x, y, vx, vy int
		fmt.Sscanf(l, "position=<%d, %d> velocity=<%d, %d>", &x, &y, &vx, &vy)
		points = append(points, Point{x: x, y: y, vx: vx, vy: vy})
	}
	return points
}

func step(points []Point) []Point {
	var res []Point
	for _, p := range points {
		res = append(res, Point{x: p.x + p.vx, y: p.y + p.vy, vx: p.vx, vy: p.vy})
	}
	return res
}

func display(points []Point) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	var grid = set.NewSet[game2d.Pos]()
	for _, p := range points {
		minX = min(minX, p.x)
		minY = min(minY, p.y)
		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)
		grid.Add(game2d.Pos{X: p.x, Y: p.y})
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, ok := grid[game2d.Pos{X: x, Y: y}]; ok {
				fmt.Printf("%c", '#')
			} else {
				fmt.Printf("%c", '.')
			}
		}
		fmt.Println()
	}

}

func detectVertical(points []Point) bool {
	var grid = set.NewSet[game2d.Pos]()
	for _, p := range points {
		grid.Add(game2d.Pos{X: p.x, Y: p.y})
	}
	for _, p := range points {
		var found = true
		for i := range 7 {
			if _, ok := grid[game2d.Pos{X: p.x, Y: p.y + i}]; !ok {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	var points = parseInput(input)
	for {
		points = step(points)
		if detectVertical(points) {
			display(points)
			fmt.Println()
			break
		}
	}
	return 0
}

func Part2(input string) int {
	var points = parseInput(input)
	var i = 0
	for {
		points = step(points)
		i++
		if detectVertical(points) {
			break
		}
	}
	return i
}

func main() {
	fmt.Println("--2018 day 10 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
