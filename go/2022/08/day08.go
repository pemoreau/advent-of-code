package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

type max struct {
	N, S, E, W int8
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	matrix := utils.BuildDigitMatrix(lines)

	maxY := len(matrix)
	maxX := len(matrix[0])

	maxmatrix := make([][]max, maxY)
	for j := 0; j < maxY; j++ {
		maxmatrix[j] = make([]max, maxX)
		for i := 0; i < maxX; i++ {
			N := int8(-1)
			if j > 0 {
				N = int8(utils.Max(int(maxmatrix[j-1][i].N), int(matrix[j-1][i])))
			}
			W := int8(-1)
			if i > 0 {
				W = int8(utils.Max(int(maxmatrix[j][i-1].W), int(matrix[j][i-1])))
			}
			maxmatrix[j][i] = max{N: N, W: W}
		}
	}
	for j := maxY - 1; j >= 0; j-- {
		for i := maxX - 1; i >= 0; i-- {
			S := int8(-1)
			if j < maxY-1 {
				S = int8(utils.Max(int(maxmatrix[j+1][i].S), int(matrix[j+1][i])))
			}
			E := int8(-1)
			if i < maxX-1 {
				E = int8(utils.Max(int(maxmatrix[j][i+1].E), int(matrix[j][i+1])))
			}
			maxmatrix[j][i] = max{S: S, E: E, N: maxmatrix[j][i].N, W: maxmatrix[j][i].W}
		}
	}

	//for j := 0; j < maxY; j++ {
	//	for i := 0; i < maxX; i++ {
	//		if int8(matrix[j][i]) > maxmatrix[j][i].N || int8(matrix[j][i]) > maxmatrix[j][i].S || int8(matrix[j][i]) > maxmatrix[j][i].E || int8(matrix[j][i]) > maxmatrix[j][i].W {
	//			fmt.Printf("[%d N:%d S%d E:%d W:%d] | ", matrix[j][i],
	//				maxmatrix[j][i].N, maxmatrix[j][i].S, maxmatrix[j][i].E, maxmatrix[j][i].W)
	//		} else {
	//			fmt.Printf(" %d N:%d S%d E:%d W:%d  | ", matrix[j][i],
	//				maxmatrix[j][i].N, maxmatrix[j][i].S, maxmatrix[j][i].E, maxmatrix[j][i].W)
	//		}
	//	}
	//	fmt.Println()
	//}

	res := 0
	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			if int8(matrix[j][i]) > maxmatrix[j][i].N || int8(matrix[j][i]) > maxmatrix[j][i].S || int8(matrix[j][i]) > maxmatrix[j][i].E || int8(matrix[j][i]) > maxmatrix[j][i].W {
				res++
			}
		}
	}

	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	matrix := utils.BuildDigitMatrix(lines)

	res := 0
	maxY := len(matrix)
	maxX := len(matrix[0])
	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			v := matrix[j][i]
			n, s, e, w := 0, 0, 0, 0
			for k := j - 1; k >= 0; k-- {
				n++
				if v <= matrix[k][i] {
					break
				}
			}
			for k := j + 1; k < maxY; k++ {
				s++
				if v <= matrix[k][i] {
					break
				}
			}
			for k := i - 1; k >= 0; k-- {
				w++
				if v <= matrix[j][k] {
					break
				}
			}
			for k := i + 1; k < maxX; k++ {
				e++
				if v <= matrix[j][k] {
					break
				}
			}
			res = utils.Max(res, n*s*e*w)
		}
	}
	return res

}

func main() {
	fmt.Println("--2022 day 08 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
