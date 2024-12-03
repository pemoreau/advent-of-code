package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type ID int16
type Tile struct {
	id     ID
	nbBits int16
	north  int16
	south  int16
	east   int16
	west   int16
	lines  []string
}

func (t Tile) String() string {
	return fmt.Sprintf("#%d: N:%.3d S:%.3d E:%.3d W:%.3d", t.id, t.north, t.south, t.east, t.west)
}

func reverseString(s string) string {
	var res string
	for _, c := range s {
		res = string(c) + res
	}
	return res
}

func flip(t Tile) Tile {
	lines := make([]string, len(t.lines))
	for i, line := range t.lines {
		lines[i] = reverseString(line)
	}
	return Tile{id: t.id, nbBits: t.nbBits, north: reverseBits(t.north, t.nbBits), south: reverseBits(t.south, t.nbBits), east: t.west, west: t.east, lines: lines}
}

func rot90(t Tile) Tile {
	return Tile{id: t.id, nbBits: t.nbBits, north: t.east, south: t.west, east: reverseBits(t.south, t.nbBits), west: reverseBits(t.north, t.nbBits), lines: rot90lines(t.lines)}
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

func toInt(s string) int16 {
	var res int16
	for _, c := range s {
		if c == '#' {
			res = (res << 1) | 1
		} else {
			res = res << 1
		}
	}
	return res
}

func reverseBits(v int16, nbBits int16) int16 {
	var res int16
	for i := int16(0); i < nbBits; i++ {
		res = (res << 1) | (v & 1)
		v = v >> 1
	}
	return res
}

func allRotations(t Tile) []Tile {
	return []Tile{
		t, rot90(t), rot90(rot90(t)), rot90(rot90(rot90(t))),
		flip(t), rot90(flip(t)), rot90(rot90(flip(t))), rot90(rot90(rot90(flip(t))))}
}

func removeTile(tiles []Tile, id ID) []Tile {
	n := len(tiles)
	for i, t := range tiles {
		if t.id == id {
			// we make a copy to not modify the underlying array (needed when backtrack)
			res := make([]Tile, n-1)
			copy(res, tiles[:i])
			copy(res[i:], tiles[i+1:])
			return res
			// return append(tiles[:i], tiles[i+1:]...)
		}
	}
	if len(tiles) != n-1 {
		fmt.Printf("Error: tile %d not found\n", id)
		fmt.Printf("tiles: %s\n", tiles)
		panic("tile not found")
	}
	return tiles
}

func puzzle(board, tiles []Tile, index int16, size int16, rotations map[ID][]Tile) (bool, []Tile) {
	n := int16(math.Sqrt(float64(size)))
	if index >= size {
		// fmt.Println("Solution found")
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

			r, b := puzzle(append(board, t), removeTile(tiles, t.id), index+1, size, rotations)
			if r {
				return true, b
			}
		}
	}
	// fmt.Printf("No solution found index=%d\n", index)
	return false, board
}

func parse(input string) []Tile {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	var tiles []Tile
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		var tileNumber int
		fmt.Sscanf(lines[0], "Tile %d:", &tileNumber)
		tile := Tile{id: ID(tileNumber), nbBits: int16(len(lines[1]))}
		tile.north = toInt(lines[1])
		tile.south = toInt(lines[len(lines)-1])
		var left []byte
		var right []byte
		for _, line := range lines[1:] {
			left = append(left, line[0])
			right = append(right, line[len(line)-1])
		}
		tile.west = toInt(string(left))
		tile.east = toInt(string(right))
		tile.lines = lines[1:]
		tiles = append(tiles, tile)
	}
	return tiles
}

func solve(input string) []Tile {
	tiles := parse(input)
	n := int16(len(tiles))
	rotations := make(map[ID][]Tile)
	for _, tile := range tiles {
		rotations[tile.id] = allRotations(tile)
	}

	board := make([]Tile, 0, n)
	_, board = puzzle(board, tiles, 0, n, rotations)
	return board
}

func Part1(input string) int {
	board := solve(input)
	n := len(board)
	size := int(math.Sqrt(float64(n)))
	return int(board[0].id) * int(board[size-1].id) * int(board[n-size].id) * int(board[n-1].id)
}

type Pos struct{ x, y int16 }
type Image map[Pos]bool

func buildImage(board []Tile) Image {
	n := len(board)
	size := int(math.Sqrt(float64(n)))
	image := make(Image)

	for i, tile := range board {
		x := i % size
		y := i / size
		for j, line := range tile.lines[1 : len(tile.lines)-1] {
			for k, c := range line[1 : len(line)-1] {
				pos := Pos{x: int16(x*8 + k), y: int16(y*8 + j)}
				if c == '#' {
					image[pos] = true
				} else {
					image[pos] = false
				}
			}
		}
	}
	return image
}

func rot90Image(image Image) Image {
	var n int16 = 0
	var m int16 = 0
	for pos := range image {
		if pos.x > n {
			n = pos.x
		}
		if pos.y > m {
			m = pos.y
		}
	}
	res := make(Image)
	for pos, v := range image {
		res[Pos{x: pos.y, y: m - 1 - pos.x}] = v
	}
	return res
}

func flipImage(image Image) Image {
	var n int16 = 0
	var m int16 = 0
	for pos := range image {
		if pos.x > n {
			n = pos.x
		}
		if pos.y > m {
			m = pos.y
		}
	}
	res := make(Image)
	for pos, v := range image {
		res[Pos{x: n - 1 - pos.x, y: pos.y}] = v
	}
	return res
}

func allImageRotation(image Image) []Image {
	return []Image{
		image, rot90Image(image), rot90Image(rot90Image(image)), rot90Image(rot90Image(rot90Image(image))),
		flipImage(image), rot90Image(flipImage(image)), rot90Image(rot90Image(flipImage(image))), rot90Image(rot90Image(rot90Image(flipImage(image))))}
}

func Part2(input string) int {
	board := solve(input)
	image := buildImage(board)

	nbPixels := 0
	for _, v := range image {
		if v {
			nbPixels++
		}
	}

	monster := make(map[Pos]bool)
	monsterPos := []Pos{
		{18, 0},
		{0, 1}, {5, 1}, {6, 1}, {11, 1}, {12, 1}, {17, 1}, {18, 1}, {19, 1},
		{1, 2}, {4, 2}, {7, 2}, {10, 2}, {13, 2}, {16, 2}}
	for _, pos := range monsterPos {
		monster[pos] = true
	}
	allMonsters := allImageRotation(monster)
	nbMonsters := 0
	for _, monster := range allMonsters {
		for pos := range image {
			found := true
			for monsterPos := range monster {
				if !image[Pos{x: pos.x + monsterPos.x, y: pos.y + monsterPos.y}] {
					found = false
					break
				}
			}
			if found {
				nbMonsters++
			}
		}
	}
	return nbPixels - len(monster)*nbMonsters

}

func main() {
	fmt.Println("--2020 day 20 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
