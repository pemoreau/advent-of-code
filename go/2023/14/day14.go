package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"time"
)

//go:embed input.txt
var inputDay string

func moveNorth(m utils.Matrix[uint8]) {
	for x := 0; x <= m.MaxX(); x++ {
		var last int
		for y := 0; y <= m.MaxY(); y++ {
			switch m[y][x] {
			case '.':
				// do nothing
			case '#':
				last = y + 1
			case 'O':
				if y > last {
					m[y][x] = '.'
					m[last][x] = 'O'
				}
				last = last + 1
			}
		}
	}
}

func totalLoad(m utils.Matrix[uint8]) int {
	var res int
	for y, l := range m {
		for _, c := range l {
			if c == 'O' {
				res += m.MaxY() - y + 1
			}
		}
	}
	return res
}

func cycle(m utils.Matrix[uint8]) utils.Matrix[uint8] {
	moveNorth(m)
	m = m.RotateRight()
	moveNorth(m)
	m = m.RotateRight()
	moveNorth(m)
	m = m.RotateRight()
	moveNorth(m)
	m = m.RotateRight()
	return m

	//moveNorth(m)
	//moveWest(m)
	//moveSouth(m)
	//moveEast(m)
}

func Part1(input string) int {
	m := utils.BuildMatrixCharFromString(input)
	moveNorth(m)
	return totalLoad(m)
}

func repeatWithCycle(m utils.Matrix[uint8], n int) utils.Matrix[uint8] {
	var valueToIndex = make(map[string]int)
	var indexToValue []string
	for i := 0; i < n; i++ {
		m = cycle(m)
		s := m.String() // use as a key
		indexToValue = append(indexToValue, s)
		if _, ok := valueToIndex[s]; ok {
			// cycle found
			indexForN := valueToIndex[s] + (n-1-i)%(i-valueToIndex[s])
			s2 := indexToValue[indexForN]
			m2 := utils.BuildMatrixCharFromString(s2)
			return m2
		} else {
			valueToIndex[s] = i
		}
	}
	return m
}

func Part2(input string) int {
	m := utils.BuildMatrixCharFromString(input)
	return totalLoad(repeatWithCycle(m, 1000000000))
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
//func moveColumn(m utils.Matrix[uint8], x int, dir int) {
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
//func moveRow(m utils.Matrix[uint8], y int, dir int) {
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
//func moveSouth(m utils.Matrix[uint8]) {
//	for x := 0; x <= m.MaxX(); x++ {
//		moveColumn(m, x, SOUTH)
//	}
//}
//func moveWest(m utils.Matrix[uint8]) {
//	for y := 0; y <= m.MaxY(); y++ {
//		moveRow(m, y, WEST)
//	}
//}
//func moveEast(m utils.Matrix[uint8]) {
//	for y := 0; y <= m.MaxY(); y++ {
//		moveRow(m, y, EAST)
//	}
//}
