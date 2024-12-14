package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"math"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type robot struct {
	x, y     int
	vx, vy   int
	w, h     int
	quadrant int
}

// move module w and h
func (r *robot) move(n int) {
	r.x += r.vx * n
	r.y += r.vy * n
	r.x = utils.Mod(r.x, r.w)
	r.y = utils.Mod(r.y, r.h)
	if r.x == r.w-1-r.x || r.y == r.h-1-r.y {
		r.quadrant = 0
	} else {
		if r.x < r.w-1-r.x {
			r.quadrant = 1
		} else {
			r.quadrant = 2
		}
		if r.y >= r.h-1-r.y {
			r.quadrant += 2
		}
	}
}

func parse(input string) []*robot {
	var lines = strings.Split(input, "\n")
	var robots []*robot

	var minX, minY = math.MaxInt, math.MaxInt
	var maxX, maxY = math.MinInt, math.MinInt
	for _, line := range lines {
		var r *robot = &robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.x, &r.y, &r.vx, &r.vy)
		robots = append(robots, r)
		minX = min(minX, r.x)
		minY = min(minY, r.y)
		maxX = max(maxX, r.x)
		maxY = max(maxY, r.y)
	}
	var width, height = maxX - minX + 1, maxY - minY + 1
	for _, r := range robots {
		r.w = width
		r.h = height
	}
	return robots
}

func Part1(input string) int {
	var robots = parse(input)

	for _, r := range robots {
		r.move(100)
	}
	var quandrants [5]int
	for _, r := range robots {
		quandrants[r.quadrant]++
	}
	return quandrants[1] * quandrants[2] * quandrants[3] * quandrants[4]
}

func Display(robots []*robot) string {
	var grid = game2d.NewGridChar()
	for _, r := range robots {
		grid.Set(r.x, r.y, '#')
	}
	return grid.String()
}

func Part2(input string) int {
	var robots = parse(input)

	var minq = math.MaxInt
	var index int
	var quandrants [5]int
	for _, r := range robots {
		quandrants[r.quadrant]++
	}

	for i := 1; i <= 10000; i++ {
		for _, r := range robots {
			quandrants[r.quadrant]--
			r.move(1)
			quandrants[r.quadrant]++
		}

		if quandrants[1]*quandrants[2]*quandrants[3]*quandrants[4] < minq {
			minq = quandrants[1] * quandrants[2] * quandrants[3] * quandrants[4]
			index = i
		}
	}

	return index
}

func main() {
	fmt.Println("--2024 day 14 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
