package main

import (
	"container/ring"
	"fmt"
	"strconv"
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

func step(current, r *ring.Ring, index []*ring.Ring) (*ring.Ring, *ring.Ring) {
	//fmt.Print("cups: ")
	//displayRing(current, r)

	// remove next 3 elements
	removed := current.Unlink(3)

	//fmt.Print("removed: ")
	//displayRing(current, removed)
	r = current
	//fmt.Print("ring: ")
	//displayRing(current, r)

	destination := current.Value.(int) - 1

	a := removed.Value.(int)
	b := removed.Next().Value.(int)
	c := removed.Next().Next().Value.(int)
	if destination == 0 {
		destination = len(index) - 1
	}
	for destination == a || destination == b || destination == c {
		destination--
		if destination == 0 {
			destination = len(index) - 1
		}
	}
	//fmt.Println("destination: ", destination)

	// find destination
	r = index[destination]

	//fmt.Println("before insertion: ", r.Value)
	//displayRing(current, r)

	// insert removed elements
	r.Link(removed)

	//fmt.Print("result: ")
	//displayRing(current.Next(), r)

	return current.Next(), r
}

func toResult(r *ring.Ring) string {
	var res string
	// search for value 1
	for i := 0; i < r.Len(); i++ {
		if r.Value.(int) == 1 {
			break
		}
		r = r.Next()
	}

	r = r.Next()
	for i := 0; i < r.Len()-1; i++ {
		res += fmt.Sprintf("%d", r.Value.(int))
		r = r.Next()
	}

	return res
}

func Part1(input string) int {
	r := ring.New(len(input))
	index := make([]*ring.Ring, len(input)+1)
	var current = r
	for _, c := range input {
		n := int(c - '0')
		r.Value = n
		index[n] = r
		r = r.Next()
	}
	//displayRing(current, r)
	//fmt.Println("current: ", current.Value)

	for i := 0; i < 100; i++ {
		//fmt.Println()
		//fmt.Println("-- move ", i+1, " --")
		current, r = step(current, r, index)
	}

	//displayRing(current, r)
	//fmt.Println("current: ", current.Value)

	res, _ := strconv.Atoi(toResult(r))
	return res
}

func Part2(input string) int {
	r := ring.New(1000000)
	index := make([]*ring.Ring, 1000000+1)

	var current = r
	for _, c := range input {
		n := int(c - '0')
		r.Value = n
		index[n] = r
		r = r.Next()
	}
	for i := len(input) + 1; i <= 1000000; i++ {
		r.Value = i
		index[i] = r
		r = r.Next()
	}

	//displayRing(current, r)
	//fmt.Println("current: ", current.Value)

	for i := 0; i < 10000000; i++ {
		//fmt.Println()
		//fmt.Println("-- move ", i+1, " --")
		current, r = step(current, r, index)
	}

	//displayRing(current, r)
	//fmt.Println("current: ", current.Value)

	a := index[1].Next().Value.(int)
	b := index[1].Next().Next().Value.(int)
	return a * b
}

func main() {
	fmt.Println("--2020 day 23 solution--")
	inputDay := "156794823"
	//inputDay := "389125467" // test
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
