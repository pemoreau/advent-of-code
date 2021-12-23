package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

func ParseCuboid(s string) Cuboid {
	var xmin, xmax, ymin, ymax, zmin, zmax int
	fmt.Sscanf(s, "x=%d..%d,y=%d..%d,z=%d..%d", &xmin, &xmax, &ymin, &ymax, &zmin, &zmax)
	return CreateCuboid(xmin, xmax+1, ymin, ymax+1, zmin, zmax+1)
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
			box50 := CreateCuboid(-50, 51, -50, 51, -50, 51)
			if i, ok := Intersection(box50, c); ok {
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
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}
