package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input_test.txt
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
	fmt.Println("pattern:", input)
	g := Grid{}
	rocks := buildRocks()
	index := 0
	for i := 0; i < 2022; i++ {
		r := rocks[i%len(rocks)]
		start := Pos{X: 2, Y: g.maxY() + 3}
		if g.free(start, r) {
			g.add(start, r)
			fmt.Println("added rock", i)
			//g.display()
			g.fall(start, r, input, &index)
		}
	}
	//fmt.Println("after rock")
	//g.display()
	// lines := strings.Split(input, "\n")
	return g.maxY()
}

type state struct {
	rockIndex int
	x         int
}

var cache = map[state]int{}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	fmt.Println("pattern:", input)
	g := Grid{}
	//rocks := buildRocks()
	//index := 0
	//for i := 0; i < 1000000000000; i++ {
	//	rockIndex := i % len(rocks)
	//	r := rocks[rockIndex]
	//	if i%10000 == 0 {
	//		fmt.Println("rock", i)
	//	}
	//	start := Pos{X: 2, Y: g.maxY() + 3}
	//	if g.free(start, r) {
	//		g.add(start, r)
	//		//fmt.Println("added rock", i)
	//		//g.display()
	//		pos := g.fall(start, r, input, &index)
	//		if _, ok := cache[state{rockIndex: rockIndex, x: pos.X}]; ok {
	//			fmt.Println("found loop", i, rockIndex, pos)
	//		} else {
	//			cache[state{rockIndex: rockIndex, x: pos.X}] = pos.Y
	//		}
	//	}
	//}
	return g.maxY()
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
