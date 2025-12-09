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
	a, b game2d.Pos
	area int
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
			var area = (1 + utils.Abs(b.X-a.X)) * (1 + utils.Abs(b.Y-a.Y))
			rectangles = append(rectangles, rectangle{a, b, area})
		}
	}
	return seats, rectangles
}

func Part1(input string) int {
	var _, rectangles = parse(input)
	slices.SortFunc(rectangles, func(a, b rectangle) int { return cmp.Compare(a.area, b.area) })
	return rectangles[len(rectangles)-1].area
}

func traverse(r rectangle, seats []game2d.Pos) bool {
	var minX = min(r.a.X, r.b.X) + 1
	var minY = min(r.a.Y, r.b.Y) + 1
	var maxX = max(r.a.X, r.b.X) - 1
	var maxY = max(r.a.Y, r.b.Y) - 1
	for i := 0; i < len(seats); i++ {
		var s1 = seats[i]
		var s2 = seats[(i+1)%len(seats)]
		var miX = min(s1.X, s2.X)
		var miY = min(s1.Y, s2.Y)
		var maX = max(s1.X, s2.X)
		var maY = max(s1.Y, s2.Y)
		if maY < minY || miY > maxY {
			continue
		}
		if maX < minX || miX > maxX {
			continue
		}
		return true
	}

	return false
}

func Part2(input string) int {
	var seats, rectangles = parse(input)
	var filtered []rectangle
	for _, r := range rectangles {
		if !traverse(r, seats) {
			filtered = append(filtered, r)
		}
	}
	slices.SortFunc(filtered, func(a, b rectangle) int { return cmp.Compare(a.area, b.area) })
	return filtered[len(filtered)-1].area
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
