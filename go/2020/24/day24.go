package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Hex struct {
	r, s, q int
}

// use axial coordinates (https://www.redblobgames.com/grids/hexagons/)
func hexNeighbors(h Hex) []Hex {
	// directions: sw, w, nw, ne, e, se
	return []Hex{
		{h.r + 1, h.s, h.q - 1},
		{h.r, h.s + 1, h.q - 1},
		{h.r - 1, h.s + 1, h.q},
		{h.r - 1, h.s, h.q + 1},
		{h.r, h.s - 1, h.q + 1},
		{h.r + 1, h.s - 1, h.q},
	}
}

func toDirections(s string) []int {
	s = strings.ReplaceAll(s, "sw", "0")
	s = strings.ReplaceAll(s, "nw", "2")
	s = strings.ReplaceAll(s, "ne", "3")
	s = strings.ReplaceAll(s, "se", "5")
	s = strings.ReplaceAll(s, "w", "1")
	s = strings.ReplaceAll(s, "e", "4")
	res := make([]int, len(s))
	for i, c := range s {
		res[i] = int(c - '0')
	}
	return res
}

func parse(input string) set.Set[Hex] {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	blackTiles := set.Set[Hex]{}
	for _, line := range lines {
		hex := Hex{0, 0, 0}
		for _, dir := range toDirections(line) {
			hex = hexNeighbors(hex)[dir]
		}
		if blackTiles.Contains(hex) {
			blackTiles.Remove(hex)
		} else {
			blackTiles.Add(hex)
		}
	}

	return blackTiles
}

func Part1(input string) int {
	return len(parse(input))
}

func step(blackTiles set.Set[Hex]) set.Set[Hex] {
	// number of black blackNeighbors
	blackNeighbors := make(map[Hex]int)
	for hex := range blackTiles {
		for _, n := range hexNeighbors(hex) {
			blackNeighbors[n]++
		}
	}

	// apply rules
	res := set.Set[Hex]{}
	for hex := range blackNeighbors {
		if blackTiles.Contains(hex) {
			if !(blackNeighbors[hex] == 0 || blackNeighbors[hex] > 2) {
				res.Add(hex)
			}
		} else {
			if blackNeighbors[hex] == 2 {
				res.Add(hex)
			}
		}
	}

	return res
}

func Part2(input string) int {
	blackTiles := parse(input)
	for i := 0; i < 100; i++ {
		blackTiles = step(blackTiles)
	}
	return blackTiles.Len()
}

func main() {
	fmt.Println("--2020 day 24 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
