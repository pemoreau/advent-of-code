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
		//fmt.Printf("check %d,%d\n", X, Y)
		if X < 0 || X > 6 || Y < 0 {
			//fmt.Printf("out of bounds %d,%d\n", X, Y)
			return false
		}
		if _, ok := g[Pos{X: X, Y: Y}]; ok {
			//fmt.Printf("occupied %d,%d\n", X, Y)
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
		//fmt.Printf("gaz [index=%d] push %c\n", *index, pattern[*index])
		pos, falling = g.move(pos, rock, pattern[*index])
		*index = (*index + 1) % len(pattern)
		//g.display()
		//fmt.Println("fall 1 unit")
		pos, falling = g.move(pos, rock, 'D')
		//g.display()
	}
	return pos
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	fmt.Println("pattern:", len(input), input)
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
		//fmt.Println("added rock", i)
		//g.display()
		pos := g.fall(start, r, input, &index)
		maxY = utils.Max(maxY, pos.Y+addY[rockIndex])
	}
	//fmt.Println("after rock")
	//g.display()
	return maxY
}

type state struct {
	rockIndex int
	x         int
	diff      int
}

var cache = map[[40]int]int{}

func findCycle(input string) (int, int, int, int, []int, []int) {
	input = strings.TrimSuffix(input, "\n")
	g := Grid{}
	rocks := buildRocks()
	addY := []int{1, 3, 3, 4, 2}
	values := []int{}
	//indexes := []int{}
	STEP := len(input)
	index := 0
	i := 0
	maxY := 0
	for i < 10*STEP {
		rockIndex := i % len(rocks)
		r := rocks[rockIndex]
		start := Pos{X: 2, Y: maxY + 3}
		g.add(start, r)
		pos := g.fall(start, r, input, &index)
		oldMaxY := maxY
		maxY = utils.Max(maxY, pos.Y+addY[rockIndex])
		values = append(values, maxY-oldMaxY)
		//indexes = append(indexes, index)
		i++
	}
	i, j := findRecurringElement(prefixes(values, len(input)))
	//indexStart, indexCycle := findRecurringElement(prefixes(indexes, len(input)))
	fmt.Println(values)
	//fmt.Println(indexes)
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

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	//g := Grid{}
	//rocks := buildRocks()
	//addY := []int{1, 3, 3, 4, 2}
	N := 1000000000000
	//index := 0
	//acc := 0
	//i := 0

	prefix, cycle, sumStart, sumCycle, prefixValues, cycleValues := findCycle(input)
	fmt.Println("prefix", prefix, "cycle", cycle, "sumStart", sumStart, "sumCycle", sumCycle)
	fmt.Println("prefixValues", prefixValues)
	fmt.Println("cycleValues", cycleValues)

	quotient := (N - prefix) / cycle
	remainder := (N - prefix) % cycle
	//i = N - remainder
	fmt.Println("quotient", quotient, "remainder", remainder)
	//i = (prefix + quotient*cycle)
	//fmt.Println("i", i)

	//maxY := sumStart
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

	//l := []int{0, 66, 59, 59, 64, 60, 61, 62, 59, 59, 59, 64, 60, 61, 62, 59, 59, 59, 64, 60, 61, 62}
	//i, j := findRecurringElement(prefixes(l, 5))
	//fmt.Println("recurring ", i, j)
}

func findRecurringElement(l [][]int) (int, int) {
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if reflect.DeepEqual(l[i], l[j]) {
				fmt.Println("found recurring", i, j, l[i], l[j])
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
