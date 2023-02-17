package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

type Hex struct {
	r, s, q int
}

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

func countBlack(tiles map[Hex]bool) int {
	count := 0
	for _, v := range tiles {
		if v {
			count++
		}
	}
	return count
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	tiles := make(map[Hex]bool)
	for _, line := range lines {
		hex := Hex{0, 0, 0}
		for _, dir := range toDirections(line) {
			hex = hexNeighbors(hex)[dir]
		}
		tiles[hex] = !tiles[hex]

	}

	return countBlack(tiles)
}

func step(tiles map[Hex]bool) map[Hex]bool {

	// count black neighbors
	neighbors := make(map[Hex]int)
	for hex, isBlack := range tiles {
		if isBlack {
			for _, n := range hexNeighbors(hex) {
				neighbors[n]++
			}
		}
	}

	res := make(map[Hex]bool)
	// apply rules
	for hex := range neighbors {
		if tiles[hex] {
			if neighbors[hex] == 0 || neighbors[hex] > 2 {
				res[hex] = false
			} else {
				res[hex] = true
			}
		} else {
			if neighbors[hex] == 2 {
				res[hex] = true
			}
		}
	}

	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	tiles := make(map[Hex]bool)
	for _, line := range lines {
		hex := Hex{0, 0, 0}
		for _, dir := range toDirections(line) {
			hex = hexNeighbors(hex)[dir]
		}
		if _, ok := tiles[hex]; !ok {
			tiles[hex] = true
		} else {
			delete(tiles, hex)
		}
	}

	for i := 0; i < 100; i++ {
		tiles = step(tiles)
		//fmt.Println("day", i+1, ":", countBlack(tiles))
	}
	return countBlack(tiles)
}

func main() {
	fmt.Println("--2020 day 24 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
