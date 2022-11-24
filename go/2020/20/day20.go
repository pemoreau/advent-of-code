package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input_test.txt
var input_day string

type ID int
type AbstractTile struct {
	id     ID
	nbBits int
	north  uint
	south  uint
	east   uint
	west   uint
	lines  []string
}

func (t AbstractTile) String() string {
	return fmt.Sprintf("#%d: N:%.3d S:%.3d E:%.3d W:%.3d", t.id, t.north, t.south, t.east, t.west)
}

func reverseString(s string) string {
	var res string
	for _, c := range s {
		res = string(c) + res
	}
	return res
}

func flip(t AbstractTile) AbstractTile {
	lines := make([]string, len(t.lines))
	for i, line := range t.lines {
		lines[i] = reverseString(line)
	}
	return AbstractTile{id: t.id, nbBits: t.nbBits, north: reverseBits(t.north, t.nbBits), south: reverseBits(t.south, t.nbBits), east: t.west, west: t.east, lines: lines}
}

func rot90(t AbstractTile) AbstractTile {
	return AbstractTile{id: t.id, nbBits: t.nbBits, north: t.east, south: t.west, east: reverseBits(t.south, t.nbBits), west: reverseBits(t.north, t.nbBits), lines: rot90lines(t.lines)}
}

func rot90lines(lines []string) []string {
	n := len(lines)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		var line string
		for j := 0; j < n; j++ {
			line += string(lines[j][n-i-1])
		}
		res[i] = line
	}
	return res
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

func allRotations(t AbstractTile) []AbstractTile {
	return []AbstractTile{
		t, rot90(t), rot90(rot90(t)), rot90(rot90(rot90(t))),
		flip(t), rot90(flip(t)), rot90(rot90(flip(t))), rot90(rot90(rot90(flip(t))))}
}

func removeTile(tiles []AbstractTile, id ID) []AbstractTile {
	n := len(tiles)
	for i, t := range tiles {
		if t.id == id {
			return append(tiles[:i], tiles[i+1:]...)
		}
	}
	if len(tiles) != n-1 {
		fmt.Printf("Error: tile %d not found\n", id)
		fmt.Printf("tiles: %s\n", tiles)
		panic("tile not found")
	}
	return tiles
}

func puzzle(board, tiles []AbstractTile, index int, size int, rotations map[ID][]AbstractTile, result *[]int) (bool, []AbstractTile) {
	n := int(math.Sqrt(float64(size)))
	if index >= size {
		fmt.Println("Solution found")
		value := int(board[0].id) * int(board[n-1].id) * int(board[size-n].id) * int(board[size-1].id)
		*result = append(*result, value)
		return true, board
	}
	x := index % n
	y := index / n
	for _, tile := range tiles {
		for _, t := range rotations[tile.id] {
			if x > 0 && board[index-1].east != t.west {
				continue
			}
			if y > 0 && board[index-n].south != t.north {
				continue
			}

			// we make copies to avoid side-effects during recursion
			copyBoard := make([]AbstractTile, len(board), len(board)+1)
			copy(copyBoard, board)
			copyTiles := make([]AbstractTile, len(tiles))
			copy(copyTiles, tiles)

			r, b := puzzle(append(copyBoard, t), removeTile(copyTiles, t.id), index+1, size, rotations, result)
			if r {
				return true, b
			}
		}
	}
	// fmt.Printf("No solution found index=%d\n", index)
	return false, board
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	var tiles []AbstractTile
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		var tileNumber int
		fmt.Sscanf(lines[0], "Tile %d:", &tileNumber)
		tile := AbstractTile{id: ID(tileNumber), nbBits: len(lines[1])}
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
		tile.lines = lines[1:]
		tiles = append(tiles, tile)
	}
	rotations := make(map[ID][]AbstractTile)
	for _, tile := range tiles {
		rotations[tile.id] = allRotations(tile)
	}
	// fmt.Println("All rotations computed")
	// for _, tile := range rotations[ID(2953)] {
	// 	for line := range tile.lines {
	// 		fmt.Println(tile.lines[line])
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println("number of tiles: ", len(tiles))
	board := make([]AbstractTile, 0, len(tiles))
	result := make([]int, 0)
	_, b := puzzle(board, tiles, 0, len(tiles), rotations, &result)
	for _, tile := range b {
		for line := range tile.lines {
			fmt.Println(tile.lines[line])
		}
		fmt.Println()
	}
	return result[0]
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
