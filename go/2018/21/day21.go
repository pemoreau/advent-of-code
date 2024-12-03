package main

import (
	_ "embed"
	"fmt"
	. "github.com/pemoreau/advent-of-code/go/2018/device"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	var m = CreateMachine(input)

	//m.registers[0] = 11513432
	for m.Run(false) {
		if m.Ip() == 28 {
			return m.Register(5)
		}
	}

	return 0
}

func Part2(input string) int {
	var m = CreateMachine(input)

	var previous int
	var history = set.NewSet[int]()
	for m.Run(false) {
		if m.Ip() == 28 {
			if history.Contains(m.Register(5)) {
				return previous
			}
			previous = m.Register(5)
			history.Add(previous)
		}
		//fmt.Println(m.registers[0])
	}

	return 0
}

func main() {
	fmt.Println("--2018 day 21 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
