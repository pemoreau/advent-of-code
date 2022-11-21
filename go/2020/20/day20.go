package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input_test.txt
var input_day string

type AbstractTile struct {
	id     int
	nbBits int
	north  uint
	south  uint
	east   uint
	west   uint
}

func (t AbstractTile) String() string {
	return fmt.Sprintf("#%d: N:%.3d S:%.3d E:%.3d W:%.3d", t.id, t.north, t.south, t.east, t.west)
}

func leftToRight(t AbstractTile) AbstractTile {
	return AbstractTile{id: t.id, north: reverseBits(t.north, t.nbBits), south: reverseBits(t.south, t.nbBits), east: t.west, west: t.east}
}

func topToBottom(t AbstractTile) AbstractTile {
	return AbstractTile{id: t.id, north: t.south, south: t.north, east: reverseBits(t.east, t.nbBits), west: reverseBits(t.west, t.nbBits)}
}

func rot90(t AbstractTile) AbstractTile {
	return AbstractTile{id: t.id, north: t.east, south: t.west, east: t.south, west: t.north}
}

func rot180(t AbstractTile) AbstractTile {
	return AbstractTile{id: t.id, north: t.south, south: t.north, east: t.west, west: t.east}
}

func rot270(t AbstractTile) AbstractTile {
	return AbstractTile{id: t.id, north: t.west, south: t.east, east: t.north, west: t.south}
}

func toInt(s string) uint {
	var res uint
	for _, c := range s {
		if c == '#' {
			res = (res << 1) | 1
		} else {
			res = res << 1
		}
	}
	return res
}

func reverseBits(v uint, nbBits int) uint {
	var res uint
	for i := 0; i < nbBits; i++ {
		res = (res << 1) | (v & 1)
		v = v >> 1
	}
	return res
}

func Part1(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	tiles := make(map[int]AbstractTile)
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		var tileNumber int
		fmt.Sscanf(lines[0], "Tile %d:", &tileNumber)
		tile := AbstractTile{id: tileNumber, nbBits: len(lines[1])}
		tile.north = toInt(lines[1])
		tile.south = toInt(lines[len(lines)-1])
		var left = make([]byte, 0)
		var right = make([]byte, 0)
		for _, line := range lines[1:] {
			left = append(left, line[0])
			right = append(right, line[len(line)-1])
		}
		tile.west = toInt(string(left))
		tile.east = toInt(string(right))
		tiles[tileNumber] = tile
	}
	fmt.Println("number of tiles: ", len(tiles))

	return 0
}

func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	return 0

}

func main() {
	fmt.Println("--2020 day 20 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}
