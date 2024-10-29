package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func parseInput(input string) []game2d.Pos {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var coords []game2d.Pos
	for _, line := range lines {
		coord := game2d.Pos{}
		fmt.Sscanf(line, "%d, %d", &coord.X, &coord.Y)
		coords = append(coords, coord)
	}
	return coords
}

func minmax(coords []game2d.Pos) (minX, maxX, minY, maxY int) {
	minX, maxX = math.MaxInt, math.MinInt
	minY, maxY = math.MaxInt, math.MinInt
	for _, p := range coords {
		minX = min(p.X, minX)
		maxX = max(p.X, maxX)
		minY = min(p.Y, minY)
		maxY = max(p.Y, maxY)
	}
	return
}

func Part1(input string) int {
	var coords = parseInput(input)
	var minX, maxX, minY, maxY = minmax(coords)

	var area = make(map[game2d.Pos]int)
	var infinite = make(map[game2d.Pos]bool)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			p := game2d.Pos{X: x, Y: y}
			minDist := math.MaxInt
			var closest game2d.Pos
			var multiple bool
			for _, c := range coords {
				dist := game2d.ManhattanDistance(p, c)
				if dist < minDist {
					minDist = dist
					closest = c
					multiple = false
				} else if dist == minDist {
					closest = game2d.Pos{X: -1, Y: -1}
					multiple = true
				}
			}
			if !multiple {
				area[closest]++
			}
			if x == minX || x == maxX || y == minY || y == maxY {
				infinite[closest] = true
			}
		}
	}

	var maxArea int
	for a, c := range area {
		if infinite[a] {
			continue
		}
		if c > maxArea {
			maxArea = c
		}
	}

	return maxArea
}

func Part2(input string) int {
	var coords = parseInput(input)
	var minX, maxX, minY, maxY = minmax(coords)

	var res int
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			p := game2d.Pos{X: x, Y: y}
			var total int
			for _, c := range coords {
				total = total + game2d.ManhattanDistance(p, c)
			}
			if total < 10000 {
				res = res + 1
			}
		}
	}
	return res
}

func main() {
	fmt.Println("--2018 day 06 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
