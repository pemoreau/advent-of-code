package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var x, y int
	var polygon = []utils.Pos{{X: x, Y: y}}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		n, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "U":
			y -= n
		case "D":
			y += n
		case "L":
			x -= n
		case "R":
			x += n
		}
		polygon = append(polygon, utils.Pos{X: x, Y: y})
	}
	return utils.PolygonArea(polygon)
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var x, y int
	var polygon = []utils.Pos{{X: x, Y: y}}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		// direction: last digit of part[2]
		d := parts[2][len(parts[2])-2]
		hex := parts[2][2 : len(parts[2])-2]
		n, _ := strconv.ParseUint(hex, 16, 64)
		switch d {
		case '3': //"U":
			y -= int(n)
		case '1': //"D":
			y += int(n)
		case '2': //"L":
			x -= int(n)
		case '0': //"R":
			x += int(n)
		}
		polygon = append(polygon, utils.Pos{X: x, Y: y})
	}
	return utils.PolygonArea(polygon)
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
