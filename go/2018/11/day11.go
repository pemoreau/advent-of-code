package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strconv"
	"time"
)

func powerLevel(x, y, serial int) int {
	var rackID = x + 10
	var res = ((y * rackID) + serial) * rackID
	if res < 100 {
		res = 0
	} else {
		var s = strconv.Itoa(res)
		var d = s[len(s)-3]
		res = int(d - '0')
	}

	res -= 5
	return res
}

// https://en.wikipedia.org/wiki/Summed-area_table
func summedAreaTable(serial int) [][]int {
	var table = make([][]int, 301)
	for i := range table {
		table[i] = make([]int, 301)
	}

	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			table[i][j] = powerLevel(i, j, serial) + table[i-1][j] + table[i][j-1] - table[i-1][j-1]
		}
	}

	return table
}

func sumSquare(table [][]int, x, y, size int) int {
	return table[x-1+size][y-1+size] - table[x-1][y-1+size] - table[x-1+size][y-1] + table[x-1][y-1]
}

func searchMax(serial int) (int, int) {
	var res int
	var x, y int
	var table = summedAreaTable(serial)
	for i := 1; i < 300-2; i++ {
		for j := 1; j < 300-2; j++ {
			sum := sumSquare(table, i, j, 3)
			if sum > res {
				res = sum
				x = i
				y = j
			}
		}
	}

	return x, y
}

func searchMax2(serial int) (int, int, int) {
	var res int
	var x, y, z int
	var table = summedAreaTable(serial)
	for size := 1; size < 300; size++ {
		for i := 1; i < 300-(size-1); i++ {
			for j := 1; j < 300-(size-1); j++ {
				sum := sumSquare(table, i, j, size)
				if sum > res {
					res = sum
					x = i
					y = j
					z = size
				}
			}
		}
	}

	return x, y, z
}

func Part1(input string) string {
	var serial, _ = strconv.Atoi(input)
	var x, y = searchMax(serial)
	return fmt.Sprintf("%d,%d", x, y)
}

func Part2(input string) string {
	var serial, _ = strconv.Atoi(input)
	var x, y, z = searchMax2(serial)
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func main() {
	fmt.Println("--2018 day 11 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
