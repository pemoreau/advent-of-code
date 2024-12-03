package main

import (
	"container/ring"
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"time"
)

//go:embed sample.txt
var inputTest string

func displayRing(r *ring.Ring) {
	for i := 0; i < r.Len(); i++ {
		v, ok := r.Value.(int)
		if ok {
			fmt.Printf("%d ", v)
		}
		r = r.Next()
	}
	fmt.Println()
}

func step(cell *ring.Ring) *ring.Ring {
	r := cell.Prev()
	removed := r.Unlink(1)
	removedValue := removed.Value.(int)
	r = r.Move(removedValue % r.Len())
	r.Link(removed)
	return r
}

func Part1(input string) int {
	numbers := utils.LinesToNumbers(input)

	r := ring.New(len(numbers))
	index := map[int]*ring.Ring{}

	var cell0 *ring.Ring
	for i, n := range numbers {
		r.Value = n
		index[i] = r
		if n == 0 {
			cell0 = r
		}
		r = r.Next()
	}
	//displayRing(r)
	for i := 0; i < len(index); i++ {
		r = step(index[i])
	}

	cell0 = cell0.Move(1000)
	v0 := cell0.Value.(int)
	cell0 = cell0.Move(1000)
	v1 := cell0.Value.(int)
	cell0 = cell0.Move(1000)
	v2 := cell0.Value.(int)
	return v0 + v1 + v2
}

func Part2(input string) int {
	numbers := utils.LinesToNumbers(input)
	key := 811589153

	r := ring.New(len(numbers))
	index := map[int]*ring.Ring{}

	var cell0 *ring.Ring
	for i, n := range numbers {
		r.Value = n * key
		index[i] = r
		if n == 0 {
			cell0 = r
		}
		r = r.Next()
	}
	for j := 0; j < 10; j++ {
		for i := 0; i < len(index); i++ {
			r = step(index[i])
		}
	}
	cell0 = cell0.Move(1000)
	v0 := cell0.Value.(int)
	cell0 = cell0.Move(1000)
	v1 := cell0.Value.(int)
	cell0 = cell0.Move(1000)
	v2 := cell0.Value.(int)
	return v0 + v1 + v2
}

func main() {
	fmt.Println("--2022 day 20 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
