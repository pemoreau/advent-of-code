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
	x, y   int
	vx, vy int
	w, h   int
}

// move module w and h
func (r *robot) move(n int) {
	r.x += r.vx * n
	r.y += r.vy * n
	r.x = utils.Mod(r.x, r.w)
	r.y = utils.Mod(r.y, r.h)
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

func numberOfRobotsPerQuadrant(robots []*robot) (int, int, int, int) {
	var q1, q2, q3, q4 int
	var midXLeft, midXRight, midYTop, midYBottom int

	if robots[0].w%2 == 0 {
		midXLeft = robots[0].w / 2
		midXRight = midXLeft
	} else {
		midXLeft = (robots[0].w - 1) / 2
		midXRight = midXLeft + 1
	}
	if robots[0].h%2 == 0 {
		midYTop = robots[0].h / 2
		midYBottom = midYTop
	} else {
		midYTop = (robots[0].h - 1) / 2
		midYBottom = midYTop + 1
	}

	for _, r := range robots {
		if r.x < midXLeft && r.y < midYTop {
			q1++
		} else if r.x >= midXRight && r.y < midYTop {
			q2++
		} else if r.x < midXLeft && r.y >= midYBottom {
			q3++
		} else if r.x >= midXRight && r.y >= midYBottom {
			q4++
		}
	}
	return q1, q2, q3, q4
}

func Part1(input string) int {
	var robots = parse(input)

	for _, r := range robots {
		r.move(100)
		//fmt.Printf("x: %d, y: %d\n", r.x, r.y)
	}

	var q1, q2, q3, q4 = numberOfRobotsPerQuadrant(robots)

	return q1 * q2 * q3 * q4
}

func Display(robots []*robot) (string, bool) {
	var grid = game2d.NewGridChar()
	for _, r := range robots {
		grid.Set(r.x, r.y, '#')
	}

	// search for vertical bar of size N and an horizontal bar of size M
	var minX, maxX, minY, maxY = grid.GetBounds()
	var found bool
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if grid.Contains(game2d.Pos{x, y}, '#') {
				var countX, countY int
				var N = 9
				var M = 3
				for j := y; j <= y+N; j++ {
					if grid.Contains(game2d.Pos{x, j}, '#') {
						countY++
					}
				}
				for i := x; i <= x+M; i++ {
					if grid.Contains(game2d.Pos{i, y}, '#') {
						countX++
					}
				}
				if countY == N && countX == M {
					found = true
				}
			}
		}

	}

	return grid.String(), found
}

func Part2(input string) int {
	var robots = parse(input)

	var minq = math.MaxInt
	var index int

	for i := 1; i <= 10000; i++ {
		for _, r := range robots {
			r.move(1)
		}

		var q1, q2, q3, q4 = numberOfRobotsPerQuadrant(robots)
		if q1*q2*q3*q4 < minq {
			minq = q1 * q2 * q3 * q4
			index = i
			//fmt.Println(q1, q2, q3, q4, q1*q2*q3*q4, i)
		}
	}

	return index
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
