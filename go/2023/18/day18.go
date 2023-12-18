package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func parseLine1(line string) (uint8, int) {
	d := line[0]
	index := strings.Index(line[2:], " ")
	dec := line[2 : index+2]
	n, _ := strconv.Atoi(dec)
	return d, n
}

func parseLine2(line string) (uint8, int) {
	l := len(line)
	d := line[l-2] // direction: last digit of line
	hex := line[l-7 : l-2]
	n, _ := strconv.ParseUint(hex, 16, 64)
	var dirs = []uint8{'R', 'D', 'L', 'U'}
	return dirs[d-'0'], int(n)
}

func solve(input string, parseLine func(string) (uint8, int)) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var x, y int
	var polygon = []game2d.Pos{{X: x, Y: y}}

	for _, line := range lines {
		dir, n := parseLine(line)
		switch dir {
		case 'U':
			y -= n
		case 'D':
			y += n
		case 'L':
			x -= n
		case 'R':
			x += n
		}
		polygon = append(polygon, game2d.Pos{X: x, Y: y})
	}
	return game2d.PolygonArea(polygon)
}

func Part1(input string) int {
	return solve(input, parseLine1)
}

func Part2(input string) int {
	return solve(input, parseLine2)
}

func main() {
	fmt.Println("--2023 day 18 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
