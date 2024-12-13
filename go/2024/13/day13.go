package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400
func solve(input string, part2 bool) int {
	var parts = strings.Split(input, "\n\n")
	var res int
	for _, part := range parts {
		var lines = strings.Split(part, "\n")
		var ax, ay, bx, by, px, py int
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d\n", &ax, &ay)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d\n", &bx, &by)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d\n", &px, &py)

		// solve:
		// ax*A+bx*B=px
		// ay*A+by*B=py
		if part2 {
			var C = 10000000000000
			px += C
			py += C
		}

		var A = (px*by - py*bx) / (ax*by - ay*bx)
		var B = (px*ay - py*ax) / (bx*ay - by*ax)
		var cond = part2 || (A <= 100 && B <= 100)
		if (px*by-py*bx)%(ax*by-ay*bx) == 0 && (px*ay-py*ax)%(bx*ay-by*ax) == 0 && A >= 0 && B >= 0 && cond {
			res += A*3 + B*1
		}
	}
	return res
}

func Part1(input string) int {
	return solve(input, false)
}

func Part2(input string) int {
	return solve(input, true)
}

func main() {
	fmt.Println("--2024 day 13 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
