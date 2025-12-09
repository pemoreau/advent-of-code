package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
)

type rectangle struct {
	minX, minY, maxX, maxY int
	area                   int
}

func minMax(a, b game2d.Pos) (minX, minY, maxX, maxY int) {
	minX = min(a.X, b.X)
	minY = min(a.Y, b.Y)
	maxX = max(a.X, b.X)
	maxY = max(a.Y, b.Y)
	return minX, minY, maxX, maxY
}

func parse(input string) (seats []game2d.Pos, rectangles []rectangle) {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")

	for _, line := range lines {
		var a, b int
		fmt.Sscanf(line, "%d,%d", &a, &b)
		seats = append(seats, game2d.Pos{a, b})
	}

	for i := 0; i < len(seats)-1; i++ {
		for j := i + 1; j < len(seats); j++ {
			var a = seats[i]
			var b = seats[j]
			var minX, minY, maxX, maxY = minMax(a, b)
			var area = (1 + maxX - minX) * (1 + maxY - minY)
			rectangles = append(rectangles, rectangle{minX, minY, maxX, maxY, area})
		}
	}
	return seats, rectangles
}

func Part1(input string) int {
	var _, rectangles = parse(input)
	return slices.MaxFunc(rectangles, func(a, b rectangle) int { return cmp.Compare(a.area, b.area) }).area
}

func traversed(r rectangle, seats []game2d.Pos) bool {
	for i := 0; i < len(seats); i++ {
		var minX, minY, maxX, maxY = minMax(seats[i], seats[(i+1)%len(seats)])
		if maxY < r.minY+1 || minY > r.maxY-1 || maxX < r.minX+1 || minX > r.maxX-1 {
			continue
		}
		return true
	}
	return false
}

func Part2(input string) int {
	var seats, rectangles = parse(input)
	var maxArea int
	for _, r := range rectangles {
		if !traversed(r, seats) {
			if r.area > maxArea {
				maxArea = r.area
			}
		}
	}
	return maxArea
}

func main() {
	fmt.Println("--2025 day 09 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
