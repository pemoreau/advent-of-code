package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
)

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var operators = strings.Fields(lines[len(lines)-1])

	lines = lines[:len(lines)-1]

	var toString = func(c int) string { return fmt.Sprintf("%d", c) }
	var maxX = len(strings.Fields(lines[0]))
	var matrix = game2d.NewMatrix[int](maxX, len(lines), toString)
	for j, l := range lines {
		l = strings.TrimSpace(l)
		for i, c := range strings.Fields(l) {
			v, _ := strconv.Atoi(c)
			matrix.Set(i, j, v)
		}
	}

	var res int
	for i := range matrix.LenX() {
		var accu int
		if operators[i] == "*" {
			accu = 1
		}
		for j := range matrix.LenY() {
			if operators[i] == "+" {
				accu += matrix.Get(i, j)
			} else {
				accu *= matrix.Get(i, j)
			}
		}
		res += accu
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")

	var toString = func(c uint8) string { return fmt.Sprintf("%c", c) }
	var matrix = game2d.BuildMatrixFunc(lines, func(c int32) uint8 { return uint8(c) }, toString)

	var res int
	var accu int
	var numbers []int
	for i := matrix.MaxX(); i >= 0; i-- {
		var n = 0
		for j := 0; j <= matrix.MaxY()-1; j++ {
			var c = matrix.Get(i, j)
			if c == ' ' || c == 0 {
			} else if matrix.Get(i, j) >= '0' && matrix.Get(i, j) <= '9' {
				n = 10*n + int(matrix.Get(i, j)-'0')
			} else {
				fmt.Printf("strange: i = %d j = %d c = '%c'\n", i, j, matrix.Get(i, j))
			}
		}
		numbers = append(numbers, n)
		if matrix.Get(i, matrix.MaxY()) == '+' {
			accu = 0
			for _, n := range numbers {
				accu += n
			}
		} else if matrix.Get(i, matrix.MaxY()) == '*' {
			accu = 1
			for _, n := range numbers {
				accu *= n
			}
		}
		if accu > 0 {
			res += accu
			i--
			numbers = numbers[:0]
			accu = 0
		}
	}
	return res
}

func main() {
	fmt.Println("--2025 day 06 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
