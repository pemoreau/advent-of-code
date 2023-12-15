package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func hash(input string) int {
	var res int
	for _, c := range input {
		if c != '\n' {
			res += int(c)
			res *= 17
			res %= 256
		}
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	var parts = strings.Split(input, ",")
	var res int
	for _, part := range parts {
		fmt.Println(part, hash(part))
		res += hash(part)
	}
	return res
}

type Box []Lens

type Lens struct {
	name  string
	value int
}

func index(b Box, name string) int {
	for i, l := range b {
		if l.name == name {
			return i
		}
	}
	return -1
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	var parts = strings.Split(input, ",")
	var boxes = make([]Box, 256)
	for _, part := range parts {
		if before, after, found := strings.Cut(part, "="); found {
			fmt.Println("=", before, after, hash(before))
			value, _ := strconv.Atoi(after)
			h := hash(before)
			box := boxes[h]
			if i := index(box, before); i != -1 {
				boxes[h] = append(box[:i], Lens{before, value})
				boxes[h] = append(boxes[h], box[i+1:]...)
			} else {
				boxes[h] = append(box, Lens{before, value})
			}
		} else if before, _, found := strings.Cut(part, "-"); found {
			fmt.Println("-", before, hash(before))
			h := hash(before)
			box := boxes[h]
			if i := index(box, before); i != -1 {
				boxes[h] = append(box[:i], box[i+1:]...)
			}
		}
	}

	var res int
	for i, box := range boxes {
		for j, lens := range box {
			fmt.Printf("%d: %d: %s: %d\n", i, j, lens.name, lens.value)
			res += (1 + i) * (j + 1) * lens.value
		}
	}

	return res
}

func main() {
	fmt.Println("--2023 day 14 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}

//const (
//	NORTH = 1
//	SOUTH = -1
//)
//
//func moveColumn(m utils.MatrixChar, x int, dir int) {
//	var start, end int
//	if dir == NORTH {
//		start = 0
//		end = m.MaxY() + dir
//	} else if dir == SOUTH {
//		start = m.MaxY()
//		end = 0 + dir
//	}
//	var last = start
//	for y := start; y != end; y = y + dir {
//		v := m[y][x]
//		if v == '.' {
//			// do nothing
//		} else if v == '#' {
//			last = y + dir
//		} else if v == 'O' {
//			if (dir == NORTH && y > last) || (dir == SOUTH && y < last) {
//				m[y][x] = '.'
//				m[last][x] = 'O'
//				last = last + dir
//			} else {
//				last = last + dir
//			}
//		} else {
//			panic("unknown char")
//		}
//	}
//}
//
//const (
//	WEST = 1
//	EAST = -1
//)
//
//func moveRow(m utils.MatrixChar, y int, dir int) {
//	var start, end int
//	if dir == WEST {
//		start = 0
//		end = m.MaxX() + dir
//	} else if dir == EAST {
//		start = m.MaxX()
//		end = 0 + dir
//	}
//	var last = start
//	for x := start; x != end; x = x + dir {
//		v := m[y][x]
//		if v == '.' {
//			// do nothing
//		} else if v == '#' {
//			last = x + dir
//		} else if v == 'O' {
//			if (dir == WEST && x > last) || (dir == EAST && x < last) {
//				m[y][x] = '.'
//				m[y][last] = 'O'
//				last = last + dir
//			} else {
//				last = last + dir
//			}
//		} else {
//			panic("unknown char")
//		}
//	}
//}
//
//func moveSouth(m utils.MatrixChar) {
//	for x := 0; x <= m.MaxX(); x++ {
//		moveColumn(m, x, SOUTH)
//	}
//}
//func moveWest(m utils.MatrixChar) {
//	for y := 0; y <= m.MaxY(); y++ {
//		moveRow(m, y, WEST)
//	}
//}
//func moveEast(m utils.MatrixChar) {
//	for y := 0; y <= m.MaxY(); y++ {
//		moveRow(m, y, EAST)
//	}
//}
