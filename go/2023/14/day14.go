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

func dropColumn(m utils.Matrix[uint8], x int) int {
	var last int
	var res int
	for y := 0; y <= m.MaxY(); y++ {
		v := m[y][x]
		if v == '.' {
			// do nothing
		} else if v == '#' {
			last = y + 1
			fmt.Printf("x=%d, y=%d, v=%c set last=%d\n", x, y, v, last)
		} else if v == 'O' {
			if y > last {
				m[y][x] = '.'
				m[last][x] = 'O'
				fmt.Printf("x=%d, y=%d, v=%c drop last=%d new last=%d\n", x, y, v, last, last+1)
				last++
			} else {
				last++
				fmt.Printf("x=%d, y=%d, v=%c do not move increment last=%d\n", x, y, v, last)
			}
			fmt.Println("add to res", last, m.MaxY()-last+2)
			res += m.MaxY() - last + 2

		} else {
			panic("unknown char")
		}
	}
	fmt.Println("----------")

	return res
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	m := utils.BuildMatrixChar(lines)
	fmt.Println(m)
	fmt.Println()
	var res int
	for x := 0; x <= m.MaxX(); x++ {
		res += dropColumn(m, x)
		fmt.Println(m)
	}
	return res
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println("--2023 day 13 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
