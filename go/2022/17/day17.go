package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"reflect"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

type Pos struct {
	X, Y int
}

type Grid map[Pos]uint8

type Rock []Pos

func buildRocks() []Rock {
	rocks := []Rock{
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},
		{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}},
		{{X: 2, Y: 2}, {X: 2, Y: 1}, {X: 2, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 0}},
		{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}},
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
	}
	return rocks
}

func (g Grid) add(place Pos, rock Rock) {
	for _, p := range rock {
		g[Pos{X: place.X + p.X, Y: place.Y + p.Y}] = '#'
	}
}

func (g Grid) remove(place Pos, rock Rock) {
	for _, p := range rock {
		delete(g, Pos{X: place.X + p.X, Y: place.Y + p.Y})
	}
}

func (g Grid) free(place Pos, rock Rock) bool {
	for _, p := range rock {
		X := place.X + p.X
		Y := place.Y + p.Y
		if X < 0 || X > 6 || Y < 0 {
			return false
		}
		if _, ok := g[Pos{X: X, Y: Y}]; ok {
			return false
		}
	}
	return true
}

func (g Grid) maxY() int {
	max := 0
	for p := range g {
		max = utils.Max(p.Y+1, max)
	}
	return max
}

func (g Grid) display() {
	maxY := g.maxY()
	for y := maxY - 1; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x < 7; x++ {
			if v, ok := g[Pos{X: x, Y: y}]; ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|")
		fmt.Println()
	}
	fmt.Println("+-------+")

}

func (g Grid) move(pos Pos, rock Rock, dir uint8) (Pos, bool) {
	var newPos Pos
	switch dir {
	case 'D':
		newPos = Pos{X: pos.X, Y: pos.Y - 1}
	case '<':
		newPos = Pos{X: pos.X - 1, Y: pos.Y}
	case '>':
		newPos = Pos{X: pos.X + 1, Y: pos.Y}
	}
	g.remove(pos, rock)
	if g.free(newPos, rock) {
		g.add(newPos, rock)
		return newPos, true
	} else {
		g.add(pos, rock)
		return pos, false
	}
}

func (g Grid) fall(pos Pos, rock Rock, pattern string, index *int) Pos {
	falling := true
	for falling {
		pos, falling = g.move(pos, rock, pattern[*index])
		*index = (*index + 1) % len(pattern)
		pos, falling = g.move(pos, rock, 'D')
	}
	return pos
}

func findRecurringElement(l [][]int) (int, int) {
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if reflect.DeepEqual(l[i], l[j]) {
				return i, j
			}
		}
	}
	return -1, -1
}

func prefixes(l []int, n int) [][]int {
	var res [][]int
	for i := 0; i < len(l)-n; i++ {
		res = append(res, l[i:i+n])
	}
	return res
}

func findCycle(input string) (int, int, int, int, []int, []int) {
	input = strings.TrimSuffix(input, "\n")
	g := Grid{}
	rocks := buildRocks()
	addY := []int{1, 3, 3, 4, 2}
	values := []int{}
	STEP := len(input)
	index := 0
	maxY := 0
	for i := 0; i < 3*STEP; i++ {
		rockIndex := i % len(rocks)
		r := rocks[rockIndex]
		start := Pos{X: 2, Y: maxY + 3}
		g.add(start, r)
		pos := g.fall(start, r, input, &index)
		oldMaxY := maxY
		maxY = utils.Max(maxY, pos.Y+addY[rockIndex])
		values = append(values, maxY-oldMaxY)
	}
	i, j := findRecurringElement(prefixes(values, len(input)))
	sumStart := 0
	sumCycle := 0
	for k := 0; k < i; k++ {
		sumStart += values[k]
	}
	for k := i; k < j; k++ {
		sumCycle += values[k]
	}
	return i, j - i, sumStart, sumCycle, values[:i], values[i:j]
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	g := Grid{}
	rocks := buildRocks()
	addY := []int{1, 3, 3, 4, 2}
	index := 0
	maxY := 0
	for i := 0; i < 2022; i++ {
		rockIndex := i % len(rocks)
		r := rocks[rockIndex]
		start := Pos{X: 2, Y: maxY + 3}
		g.add(start, r)
		pos := g.fall(start, r, input, &index)
		maxY = utils.Max(maxY, pos.Y+addY[rockIndex])
	}
	return maxY
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	N := 1000000000000

	prefix, cycle, sumStart, sumCycle, _, cycleValues := findCycle(input)
	quotient := (N - prefix) / cycle
	remainder := (N - prefix) % cycle
	maxY := sumStart + quotient*sumCycle
	for i := 0; i < remainder; i++ {
		maxY += cycleValues[i]
	}
	return maxY
}

func main() {
	fmt.Println("--2022 day 17 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
