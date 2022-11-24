package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
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

func flip(t AbstractTile) AbstractTile {
	return AbstractTile{id: t.id, nbBits: t.nbBits, north: reverseBits(t.north, t.nbBits), south: reverseBits(t.south, t.nbBits), east: t.west, west: t.east}
}

func rot90(t AbstractTile) AbstractTile {
	return AbstractTile{id: t.id, nbBits: t.nbBits, north: t.east, south: t.west, east: reverseBits(t.south, t.nbBits), west: reverseBits(t.north, t.nbBits)}
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

func removeTile(tiles []AbstractTile, id int) []AbstractTile {
	n := len(tiles)
	for i, t := range tiles {
		if t.id == id {
			return append(tiles[:i], tiles[i+1:]...)
		}
	}
	if len(tiles) != n-1 {
		fmt.Printf("Error: tile %d not found\n", id)
		fmt.Printf("tiles: %d\n", tiles)
		panic("tile not found")
	}
	return tiles
}

func ids(tiles []AbstractTile) []int {
	var res []int
	for _, t := range tiles {
		res = append(res, t.id)
	}
	return res
}

func puzzle(board, tiles []AbstractTile, index int, size int, rotations map[AbstractTile][]AbstractTile, result *[]int) bool {
	n := int(math.Sqrt(float64(size)))
	if index >= size {
		fmt.Println("Solution found")
		// fmt.Println("index: ", index)
		// fmt.Println("tiles: ", len(tiles))
		// for _, t := range board {
		// 	fmt.Println(t)
		// }
		value := board[0].id * board[n-1].id * board[size-n].id * board[size-1].id
		fmt.Println("value: ", value)
		*result = append(*result, value)
		return true
	}
	// fmt.Printf("index: %d, board: %d tiles: %d\n", index, len(board), len(tiles))
	// fmt.Printf("\tboard: %d\n", board)
	// fmt.Printf("\ttiles: %d\n", ids(tiles))
	x := index % n
	y := index / n
	// fmt.Printf("x: %d, y: %d\n", x, y)
	for _, tile := range tiles {
		for _, t := range rotations[tile] {
			if x > 0 && board[index-1].east != t.west {
				continue
			}
			if y > 0 && board[index-n].south != t.north {
				continue
			}

			copyBoard := make([]AbstractTile, len(board))
			copy(copyBoard, board)
			copyTiles := make([]AbstractTile, len(tiles))
			copy(copyTiles, tiles)

			// fmt.Printf("placer: %d\n", t.id)
			// fmt.Printf("tiles: %d\n", ids(copyTiles))
			r := puzzle(append(copyBoard, t), removeTile(copyTiles, t.id), index+1, size, rotations, result)
			if r {
				return true
			}
			// fmt.Printf("retirer: %d\n", t.id)
			// fmt.Printf("tiles: %d\n", ids(copyTiles))

		}
	}
	// fmt.Printf("No solution found index=%d\n", index)
	return false
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	var tiles []AbstractTile
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
		tiles = append(tiles, tile)
	}
	rotations := make(map[AbstractTile][]AbstractTile)
	for _, tile := range tiles {
		rotations[tile] = allRotations(tile)
	}
	fmt.Println("number of tiles: ", len(tiles))
	board := make([]AbstractTile, 0)
	result := make([]int, 0)
	puzzle(board, tiles, 0, len(tiles), rotations, &result)
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
