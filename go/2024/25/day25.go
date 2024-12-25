package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"slices"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type pins [5]int

var cmpKeys = func(a, b pins) int {
	for i := 0; i < 5; i++ {
		if a[i] < b[i] {
			return 1
		} else if a[i] > b[i] {
			return -1
		}
	}
	return 0
}

func isValidKey(lock, key pins) bool {
	for i := 0; i < 5; i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

// returns true if there exists a key such that:
// lock[i]+sortedKeys[i] <=5 for all i
func searchKey(lock pins, sortedKeys []pins) int {
	var res int
	var perfectKey = pins{5 - lock[0], 5 - lock[1], 5 - lock[2], 5 - lock[3], 5 - lock[4]}
	var i, found = slices.BinarySearchFunc(sortedKeys, perfectKey, cmpKeys)
	fmt.Printf("pefectKey: %v, found: %v i:%d\n", perfectKey, found, i)
	for i < len(sortedKeys) {
		if isValidKey(lock, sortedKeys[i]) {
			res++
		}
		i++
	}
	return res
}

func Part1(input string) int {
	var parts = strings.Split(input, "\n\n")
	var keys, locks []pins
	for _, part := range parts {
		var lines = strings.Split(part, "\n")
		var factor int
		var p pins
		if lines[0] == "#####" {
			factor = 1
		} else {
			factor = -1
			p = pins{5, 5, 5, 5, 5}
		}
		for _, line := range lines[1:] {
			for i, c := range line {
				if c == '#' && factor == 1 {
					p[i]++
				} else if c == '.' && factor == -1 {
					p[i]--
				}
			}
		}
		if factor == 1 {
			locks = append(locks, p)
		} else {
			keys = append(keys, p)
		}
	}
	fmt.Printf("%d locks: %v\n", len(locks), locks)

	slices.SortFunc(keys, cmpKeys)
	fmt.Printf("%d keys: %v\n", len(keys), keys)

	var res int
	for _, lock := range locks {
		res += searchKey(lock, keys)
	}

	return res
}

//func Part2(input string) int {
//	//var lines = strings.Split(input, "\n")
//
//	var res int
//	return res
//}

func main() {
	fmt.Println("--2024 day 25 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	//start = time.Now()
	//fmt.Println("part2: ", Part2(inputDay))
	//fmt.Println(time.Since(start))
}
