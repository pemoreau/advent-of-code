package main

import (
	"container/ring"
	_ "embed"
	"fmt"
	"slices"
	"time"
)

//go:embed input.txt
var inputDay string

func displayRing(r *ring.Ring) {
	for i := 0; i < r.Len(); i++ {
		v, ok := r.Value.(int)
		if ok {
			fmt.Printf("%d ", v)
		} else {
			fmt.Printf(". ")
		}

		r = r.Next()
	}
	fmt.Println()
}

func Solve(nbPlayer, nbMarble int) int {
	var players = make([]int, nbPlayer+1)
	var toPlay = 1

	r := ring.New(1)
	r.Value = 0

	for i := 1; i <= nbMarble; i++ {
		r = r.Next()

		if i%23 == 0 {
			players[toPlay] += i
			r = r.Move(-9)
			var removed = r.Unlink(1)
			players[toPlay] += removed.Value.(int)
			r = r.Next()
		} else {
			n := ring.New(1)
			n.Value = i
			r = r.Link(n)
			r = r.Prev()
		}

		toPlay = 1 + (toPlay)%nbPlayer
	}

	return slices.Max(players)
}
func Part1(input string) int {
	var nbPlayer, nbMarble int
	fmt.Sscanf(input, "%d players; last marble is worth %d points", &nbPlayer, &nbMarble)
	return Solve(nbPlayer, nbMarble)
}

func Part2(input string) int {
	var nbPlayer, nbMarble int
	fmt.Sscanf(input, "%d players; last marble is worth %d points", &nbPlayer, &nbMarble)
	return Solve(nbPlayer, 100*nbMarble)
}

func main() {
	fmt.Println("--2018 day 09 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
