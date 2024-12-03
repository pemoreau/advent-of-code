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

const (
	rocky  = 0
	wet    = 1
	narrow = 2
)

type param struct {
	x, y, targetX, targetY, depth int
}

const (
	neither  = 0
	torch    = 1
	climbing = 2
)

type state struct {
	param
	tool   int
	region int
}

func buildState(s state, x, y int) state {
	region := regionType(erosionLevel(geologicIndex(x, y, s.targetX, s.targetY, s.depth), s.depth))
	return state{param{x, y, s.targetX, s.targetY, s.depth}, s.tool, region}
}

func keepFunc(s state) bool {
	if s.region == rocky && (s.tool == torch || s.tool == climbing) {
		return true
	}
	if s.region == wet && (s.tool == climbing || s.tool == neither) {
		return true
	}
	if s.region == narrow && (s.tool == torch || s.tool == neither) {
		return true
	}
	return false
}

func neighbors(s state) []state {
	var states []state
	var dir = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dir {
		x := s.x + d[0]
		y := s.y + d[1]
		if x < 0 || y < 0 {
			continue
		}
		s := buildState(s, x, y)
		if keepFunc(s) {
			states = append(states, s)
		}
	}

	var switchFunc = func(region int, tool int) int {
		switch {
		case region == rocky && tool == torch:
			return climbing
		case region == rocky && tool == climbing:
			return torch
		case region == wet && tool == climbing:
			return neither
		case region == wet && tool == neither:
			return climbing
		case region == narrow && tool == torch:
			return neither
		case region == narrow && tool == neither:
			return torch
		}
		return -1
	}

	states = append(states, state{s.param, switchFunc(s.region, s.tool), s.region})
	return states
}

var cache = make(map[param]int)

func geologicIndex(x, y, targetX, targetY int, depth int) int {
	if (x == 0 && y == 0) || (x == targetX && y == targetY) {
		return 0
	}
	if y == 0 {
		return x * 16807
	}
	if x == 0 {
		return y * 48271
	}

	if v, ok := cache[param{x, y, targetX, targetY, depth}]; ok {
		return v
	}
	v := erosionLevel(geologicIndex(x-1, y, targetX, targetY, depth), depth) * erosionLevel(geologicIndex(x, y-1, targetX, targetY, depth), depth)
	cache[param{x, y, targetX, targetY, depth}] = v
	return v
}

func erosionLevel(geoIndex int, depth int) int {
	return (geoIndex + depth) % 20183
}

func regionType(geoIndex int) int {
	return geoIndex % 3
}

func Part1(input string) int {
	input = strings.Trim(input, "\n")
	var lines = strings.Split(input, "\n")
	var depth int
	var targetX, targetY int
	fmt.Sscanf(lines[0], "depth: %d", &depth)
	fmt.Sscanf(lines[1], "target: %d,%d", &targetX, &targetY)

	var riskLevel int
	for y := 0; y <= targetY; y++ {
		for x := 0; x <= targetX; x++ {
			region := regionType(erosionLevel(geologicIndex(x, y, targetX, targetY, depth), depth))
			riskLevel += region
		}
	}
	return riskLevel
}

func Part2(input string) int {
	input = strings.Trim(input, "\n")
	var lines = strings.Split(input, "\n")
	var depth int
	var targetX, targetY int
	fmt.Sscanf(lines[0], "depth: %d", &depth)
	fmt.Sscanf(lines[1], "target: %d,%d", &targetX, &targetY)

	var start = state{param{0, 0, targetX, targetY, depth}, torch, rocky}

	neighborsF := func(s state) []state { return neighbors(s) }
	costF := func(from, to state) int {
		if from.tool == to.tool {
			return 1
		}
		return 7
	}

	goalF := func(s state) bool { return s.x == targetX && s.y == targetY && s.tool == torch }
	heuristicF := func(s state) int { return utils.Abs(s.x-targetX) + utils.Abs(s.y-targetY) }

	_, cost := utils.Astar[state](start, goalF, neighborsF, costF, heuristicF)

	return cost
}

func main() {
	fmt.Println("--2018 day 22 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
