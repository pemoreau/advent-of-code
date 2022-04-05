package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func increment(inp []int, index int, stop int) bool {

	if inp[index] == 9 {
		inp[index] = 1
		return increment(inp, index-1, stop)
	} else {
		inp[index]++
	}
	if index < stop {
		return false
	}
	return true
}

func decrement(inp []int, index int, stop int) bool {

	if inp[index] == 1 {
		inp[index] = 9
		return decrement(inp, index-1, stop)
	} else {
		inp[index]--
	}
	if index < stop {
		return false
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

func makeInput(n int) []int {
	inp := make([]int, 14)
	for i := 0; i < 14; i++ {
		inp[13-i] = n % 10
		if inp[13-i] == 0 {
			return []int{}
		}
		n = n / 10
	}
	return inp
}

func worker(wg *sync.WaitGroup, id int, done chan interface{}, stream <-chan int) {
	defer wg.Done()
	for m := range stream {
		select {
		case <-done:
			return
		default:
			// fmt.Println("received: ", m)
			inp := makeInput(m)
			fmt.Printf("Worker %d received: %v\n", id, inp)
			var cont = len(inp) > 0
			for cont {
				_, _, _, z := Run(inp)
				// fmt.Printf("z=%d inp=%v\n", z, inp)
				// if id == 1 && rand.Uint32()%100000 == 0 {
				// 	fmt.Printf("Worker %d found: %v\n", id, inp)
				// 	done <- true
				// 	// close(done)
				// 	return
				// }
				if z == 0 {
					fmt.Printf("bingo inp=%v\n", inp)
					done <- true
					return
				}
				//cont = decrement(inp, len(inp)-1, 13-7)
				cont = increment(inp, len(inp)-1, 13-7)
			}
			fmt.Printf("Worker %d done:     %v\n", id, inp)

		}
	}
}

const chunkSize = 100000000

func producer(done <-chan interface{}) <-chan int {
	nbThreads := runtime.NumCPU()
	resultStream := make(chan int, nbThreads)
	go func() {
		defer close(resultStream)
		// i := 99999999999999
		//i := 12999999999999
		i := 11811111111111
		for i > 0 {
			select {
			case <-done:
				return
			case resultStream <- i:
				//i = i - chunkSize
				i = i + chunkSize
			}
		}
	}()
	return resultStream
}

func Part1() int {
	nbThreads := runtime.NumCPU()
	done := make(chan interface{})
	defer close(done)
	var wg sync.WaitGroup
	stream := producer(done)
	for w := 1; w <= nbThreads; w++ {
		wg.Add(1)
		go worker(&wg, w, done, stream)
	}
	wg.Wait()
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

// bingo inp=[1 1 8 4 1 7 3 1 6 1 7 1 8 9]
// max
//bingo inp = [1 2 9 9 6 9 9 7 8 2 9 3 9 9]
//12996997829399
// min
//bingo inp = [1 1 8 4 1 6 3 1 5 1 7 1 8 9]
//bingo inp = [1 1 8 4 1 5 3 1 4 1 7 1 8 9]
//bingo inp = [1 1 8 4 1 4 3 1 3 1 7 1 8 9]
//bingo inp = [1 1 8 4 1 3 3 1 2 1 7 1 8 9]
//bingo inp = [1 1 8 4 1 2 3 1 1 1 7 1 8 9]
//11841231117189
