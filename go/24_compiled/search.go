package main

import (
	"fmt"
	"time"
)

func decrement(inp []int, index int, stop int) bool {
	if index < stop {
		return false
	}
	if inp[index] == 1 {
		inp[index] = 9
		return decrement(inp, index-1, stop)
	} else {
		inp[index]--
		if index < 5 {
			fmt.Println(inp)
		}
	}
	return true
}

func search() {
	// inp := []int{9, 9, 6, 8, 7, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	// inp := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	inp := []int{9, 9, 9, 1, 7, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	stop := 0
	c := make([]int, len(inp))
	copy(c, inp)

	fmt.Printf("run %v\n", c)
	var cont = true
	for cont {
		_, _, _, z := Run(c)
		if z == 0 {
			fmt.Printf("bingo inp=%v\n", c)
			return
		}
		cont = decrement(c, len(c)-1, stop)
	}
	// fmt.Printf("finished d0=%d d1=%d\n", d0, d1)

	// return c

}

func Part1() int {

	search()
	// // inp := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 9, 9, 9, 9, 9}
	// for {

	// 	_, _, _, z := Run(inp)

	// 	// fmt.Printf("inp=%v\n", inp)
	// 	// fmt.Printf("w: %d, x: %d, y: %d, z: %d, i: %d\n", w, x, y, z, i)

	// 	if z == 0 {
	// 		fmt.Printf("inp=%v\n", inp)
	// 		break
	// 	}
	// 	inp = decrement(inp, len(inp)-1)
	// }

	return 0
}

func Part2() int {
	return 0

}

func main() {
	fmt.Println("--2021 day 24 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1())
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2())
	fmt.Println(time.Since(start))
}
