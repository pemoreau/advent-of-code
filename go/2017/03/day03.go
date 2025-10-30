package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"iter"
	"strconv"
	"strings"
	"time"
)

// iterSpiral produit une séquence infinie de positions en spirale
// en partant de (0,0) puis (1,0), (1,-1), (0,-1), (-1,-1), ...
// Ordre des directions : Est, Nord, Ouest, Sud, avec des longueurs 1,1,2,2,3,3,...
func iterSpiral() iter.Seq[game2d.Pos] {
	return func(yield func(game2d.Pos) bool) {
		p := game2d.Pos{X: 0, Y: 0}
		if !yield(p) {
			return
		}

		step := 1
		for {
			// E de 'step' cases
			for i := 0; i < step; i++ {
				p = p.E()
				if !yield(p) {
					return
				}
			}
			// N de 'step' cases
			for i := 0; i < step; i++ {
				p = p.N()
				if !yield(p) {
					return
				}
			}
			step++

			// W de 'step' cases
			for i := 0; i < step; i++ {
				p = p.W()
				if !yield(p) {
					return
				}
			}
			// S de 'step' cases
			for i := 0; i < step; i++ {
				p = p.S()
				if !yield(p) {
					return
				}
			}
			step++
		}
	}
}

func findDistance2(n int) int {
	i := 0
	for p := range iterSpiral() {
		i++
		if i == n {
			return utils.Abs(p.X) + utils.Abs(p.Y)
		}
	}
	return 0
}

func findDistance(n int) int {
	//fmt.Println("search n =", n)
	if n == 1 {
		return 0
	}
	size := 1
	lower, upper := 1, 1
	i := 0
	for {
		//fmt.Printf("i = %d --> lower = %d, upper = %d\n", i, lower, upper)
		if lower <= n && n <= upper {
			delta := ((size - 1) / 2) - 1
			incr := -1
			for j := lower; j <= upper; j++ {
				if j == n {
					//fmt.Printf("delta = %d found %d at distance = %d\n", delta, n, i+delta)
					return i + delta
				}
				//fmt.Printf("j = %d delta = %d\n", j, delta)
				delta += incr
				if delta == 0 || delta == (size-1)/2 {
					incr = -incr
				}
			}
		}

		size += 2
		lower = upper + 1
		upper = size * size
		i++
	}
}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	n, _ := strconv.Atoi(input)
	return findDistance2(n)
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	n, _ := strconv.Atoi(input)
	grid := make(map[game2d.Pos]int)

	grid[game2d.Pos{X: 0, Y: 0}] = 1
	for p := range iterSpiral() {
		sum := 0
		//fmt.Printf("p = %v\n", p)
		for q := range p.Neighbors8() {
			if v, ok := grid[q]; ok {
				sum += v
			}
		}
		//fmt.Printf("p = %v sum = %d\n", p, sum)
		if sum > 0 {
			grid[p] = sum
		}
		if sum > n {
			return sum
		}
	}
	return 0
}

func main() {
	fmt.Println("--2017 day 03 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
