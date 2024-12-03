package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"slices"
	"strings"
	"time"
)

type Claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func parseClaim(line string) Claim {
	id, x, y, w, h := 0, 0, 0, 0, 0
	// #1 @ 1,3: 4x4
	res := new(Claim)
	fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
	res.id = id
	res.x = x
	res.y = y
	res.width = w
	res.height = h
	return *res
}

func totalWidth(claim Claim) int {
	return 2*claim.x + claim.width
}

func totalHeight(claim Claim) int {
	return 2*claim.y + claim.height
}

func maxClaims(claims []Claim) (int, int) {
	var maxH = make([]int, len(claims))
	var maxW = make([]int, len(claims))
	for i := range len(claims) {
		maxH[i] = totalHeight(claims[i])
		maxW[i] = totalWidth(claims[i])
	}
	return slices.Max(maxH), slices.Max(maxW)
}

func fillMatrix(list []Claim) [][]int {
	maxH, maxW := maxClaims(list)
	matrix := make([][]int, maxW)
	for i := 0; i < maxW; i++ {
		matrix[i] = make([]int, maxH)
	}
	for _, c := range list {
		for i := range c.width {
			for j := range c.height {
				matrix[c.x+i][c.y+j]++
			}
		}
	}
	return matrix
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	claims := make([]Claim, len(lines))
	for i, line := range lines {
		claims[i] = parseClaim(line)
	}
	matrix := fillMatrix(claims)
	maxW := len(matrix)
	maxH := len(matrix[0])
	sum := 0
	for i := 0; i < maxW; i++ {
		for j := 0; j < maxH; j++ {
			if matrix[i][j] > 1 {
				sum++
			}
		}
	}

	return sum
}

func overlap(c Claim, matrix *[][]int) bool {
	for i := 0; i < c.width; i++ {
		for j := 0; j < c.height; j++ {
			if (*matrix)[c.x+i][c.y+j] > 1 {
				return true
			}
		}
	}
	return false
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	claims := make([]Claim, len(lines))
	for i, line := range lines {
		claims[i] = parseClaim(line)
	}
	matrix := fillMatrix(claims)
	for _, c := range claims {
		if !overlap(c, &matrix) {
			return c.id
		}
	}

	return -1
}

func main() {
	fmt.Println("--2018 day 03 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part1: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
