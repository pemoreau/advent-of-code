package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type point struct {
	x, y, z, t int
}

func distance(p1, p2 point) int {
	return utils.Abs(p1.x-p2.x) + utils.Abs(p1.y-p2.y) + utils.Abs(p1.z-p2.z) + utils.Abs(p1.t-p2.t)
}

func parseInput(input string) []point {
	input = strings.TrimSuffix(input, "\n")
	var lines []string = strings.Split(input, "\n")
	var points []point
	for _, line := range lines {
		var p point
		fmt.Sscanf(line, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.t)
		points = append(points, p)
	}
	return points
}

func computeNeighbours(points []point) map[point][]point {
	var result = make(map[point][]point)
	for _, p := range points {
		for _, p2 := range points {
			if distance(p, p2) <= 3 {
				result[p] = append(result[p], p2)
			}
		}
	}
	return result
}

func hasNext(visited set.Set[point], points []point) (point, bool) {
	for _, p := range points {
		if !visited.Contains(p) {
			return p, true
		}
	}
	return point{}, false
}

func Part1(input string) int {
	var points = parseInput(input)
	var neighbours = computeNeighbours(points)

	var todo []point
	var visited = set.NewSet[point]()

	var res int
	for next, ok := hasNext(visited, points); ok; next, ok = hasNext(visited, points) {
		todo = append(todo, next)
		visited.Add(next)
		res++

		for len(todo) > 0 {
			var current = todo[0]
			todo = todo[1:]
			for _, neighbour := range neighbours[current] {
				if !visited.Contains(neighbour) {
					todo = append(todo, neighbour)
					visited.Add(neighbour)
				}
			}
		}

	}

	return res
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println("--2018 day 25 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
