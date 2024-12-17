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
		m.init()
		m.SetRegister(A, i)
		m.SetRegister(B, 0)
		m.SetRegister(C, 0)
		var o = Output{}
		for m.Run(&o, false) {
		}
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
	var m = CreateMachine(inst, a, b, c)
	var o = Output{}
	for m.Run(&o, false) {
	}
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

func solveFast(m *Machine, index int, expected []int, start int) int {
	//fmt.Println("index", index, "start", start)
	if index < 0 {
		return start
	}
	for i := 0; i < 8; i++ {
		var a = 8*start + i
		var res = findN(m, a, a, expected[index:])
		if res >= 0 {
			if v := solveFast(m, index-1, expected, a); v >= 0 {
				return v
			}
		}
	}
	return -1
}

func Part2(input string) int {
	var inst, a, b, c = parse(input)
	var m = CreateMachine(inst, a, b, c)
	//return solveSlow(m, inst, 1)
	return solveFast(m, len(inst)-1, inst, 0)
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
