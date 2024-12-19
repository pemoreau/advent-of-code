package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

//go:embed sample2.txt
var inputTest2 string

func parse(input string) ([]instruction, int, int, int) {
	input = strings.Trim(input, "\n")
	var parts = strings.Split(input, "\n\n")
	var a, b, c int
	fmt.Sscanf(parts[0], "Register A: %d\nRegister B: %d\nRegister C: %d", &a, &b, &c)
	var _, after, _ = strings.Cut(parts[1], " ")
	var inst []instruction
	for _, e := range strings.Split(after, ",") {
		var v, _ = strconv.Atoi(e)
		inst = append(inst, v)
	}
	return inst, a, b, c
}

func findN(m *Machine, start int, end int, goal []int) int {
	for i := start; i <= end; i++ {
		o := m.Run(i, 0, 0, false)
		if slices.Equal(o.out, goal) {
			//fmt.Printf("found a=%d out=%v goal=%v\n", i, o.out, goal)
			return i
		} else {
			//fmt.Printf("try a=%d out=%v goal=%v\n", i, o.out, goal)
		}
	}
	return -1
}

func Part1(input string) string {
	var inst, a, b, c = parse(input)
	var m = CreateMachine(inst)
	var o = m.Run(a, b, c, false)
	var res string
	for i, e := range o.out {
		if i > 0 {
			res = res + ","
		}
		res = res + strconv.Itoa(e)
	}
	return res
}

func solveSlow(m *Machine, expected []int, start int) int {
	var res int
	for i := len(expected) - 1; i >= 0; i-- {
		start = findN(m, start, 8*start+1, expected[i:])
		if start <= 0 {
			return 0
		}
		res = start
		start = 8 * start
	}
	return res
}

func solveFast(m *Machine, a int, index int, expected []int) int {
	//fmt.Println("index", index, "start", start)
	if index < 0 {
		return a
	}
	for i := 0; i < 8; i++ {
		var nextA = 8*a + i
		if o := m.Run(nextA, 0, 0, false); slices.Equal(o.out, expected[index:]) {
			if v := solveFast(m, nextA, index-1, expected); v >= 0 {
				return v
			}
		}
	}
	return -1
}

func solveFast2(m *Machine, a int, index int, expected []int) int {
	var o = m.Run(a, 0, 0, false)
	if slices.Equal(o.out, expected) {
		return a
	}
	if index == 0 || slices.Equal(o.out, expected[len(expected)-index:]) {
		for digit := 0; digit < 8; digit++ {
			if a := solveFast2(m, 8*a+digit, index+1, expected); a >= 0 {
				return a
			}
		}
	}
	return -1
}

func Part2(input string) int {
	var inst, _, _, _ = parse(input)
	var m = CreateMachine(inst)
	//return solveSlow(m, inst, 1)
	return solveFast(m, 0, len(inst)-1, inst)
	//return solveFast2(m, 0, 0, inst)
}

func main() {
	fmt.Println("--2024 day 17 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
