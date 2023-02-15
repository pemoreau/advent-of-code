package main

import (
	"container/ring"
	"fmt"
	"time"
)

func displayRing(current, r *ring.Ring) {
	for i := 0; i < r.Len(); i++ {
		v, ok := r.Value.(int)
		if ok {
			if r == current {
				fmt.Printf("(%d) ", v)
			} else {
				fmt.Printf("%d ", v)
			}
		}
		r = r.Next()
	}
	fmt.Println()
}

func step(current, r *ring.Ring) (*ring.Ring, *ring.Ring) {
	fmt.Print("cups: ")
	displayRing(current, r)

	// remove next 3 elements
	removed := current.Unlink(3)

	fmt.Print("removed: ")
	displayRing(current, removed)
	r = current
	fmt.Print("ring: ")
	displayRing(current, r)

	destination := current.Value.(int) - 1
	if destination == 0 {
		destination = 9
	}
	for {
		found := false
		for i := 0; i < removed.Len(); i++ {
			if removed.Value.(int) == destination {
				found = true
				break
			}
			removed = removed.Next()
		}
		if !found {
			break
		}
		destination--
		if destination == 0 {
			destination = 9
		}
	}
	fmt.Println("destination: ", destination)

	// find destination
	for i := 0; i < r.Len(); i++ {
		if r.Value.(int) == destination {
			break
		}
		r = r.Next()
	}

	fmt.Println("before insertion: ", r.Value)
	displayRing(current, r)

	// insert removed elements
	r.Link(removed)

	fmt.Print("result: ")
	displayRing(current.Next(), r)

	return current.Next(), r
}

func Part1(input string) int {
	r := ring.New(len(input))

	//var cell1 *ring.Ring
	var current = r
	for _, c := range input {
		n := int(c - '0')
		r.Value = n
		if n == 1 {
			//cell1 = r
		}
		r = r.Next()
	}
	displayRing(current, r)
	fmt.Println("current: ", current.Value)

	for i := 0; i < 10; i++ {
		fmt.Println()
		fmt.Println("-- move ", i+1, " --")
		current, r = step(current, r)
	}

	displayRing(current, r)
	fmt.Println("current: ", current.Value)
	return 0
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println("--2020 day 23 solution--")
	//input_day := "156794823"
	input_day := "389125467"
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
