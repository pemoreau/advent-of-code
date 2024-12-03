package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/interval"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func ParseCuboid(s string) interval.Cuboid {
	var xmin, xmax, ymin, ymax, zmin, zmax int
	fmt.Sscanf(s, "x=%d..%d,y=%d..%d,z=%d..%d", &xmin, &xmax, &ymin, &ymax, &zmin, &zmax)
	return interval.CreateCuboid(xmin, xmax+1, ymin, ymax+1, zmin, zmax+1)
}

func solve(input string, part int) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	world := World{}
	for _, line := range lines {
		commands := strings.Split(line, " ")
		info := uint8(2)
		if commands[0] == "off" {
			info = uint8(1)
		}
		c := ParseCuboid(commands[1])
		if part == 1 {
			box50 := interval.CreateCuboid(-50, 51, -50, 51, -50, 51)
			if i, ok := interval.Intersection(box50, c); ok {
				world.Add(i, info)
			}
		} else {
			world.Add(c, info)
		}
	}
	return world.Count(2)
}

func Part1(input string) int {
	return solve(input, 1)
}

func Part2(input string) int {
	return solve(input, 2)
}

func main() {
	fmt.Println("--2021 day 22 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
