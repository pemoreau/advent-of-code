package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type Pad struct {
	id string
	game2d.Pos
	layout  map[game2d.Pos]uint8
	buttons map[uint8]game2d.Pos
}

func (p *Pad) String() string {
	var key, _ = p.layout[p.Pos]
	return fmt.Sprintf("Pad %s at key: %v pos:%v", p.id, key, p.Pos)
}

func NewNumericPad(id string) *Pad {
	pad := Pad{
		id:      id,
		layout:  make(map[game2d.Pos]uint8),
		buttons: make(map[uint8]game2d.Pos),
	}
	pad.layout[game2d.Pos{0, 0}] = '7'
	pad.layout[game2d.Pos{1, 0}] = '8'
	pad.layout[game2d.Pos{2, 0}] = '9'
	pad.layout[game2d.Pos{0, 1}] = '4'
	pad.layout[game2d.Pos{1, 1}] = '5'
	pad.layout[game2d.Pos{2, 1}] = '6'
	pad.layout[game2d.Pos{0, 2}] = '1'
	pad.layout[game2d.Pos{1, 2}] = '2'
	pad.layout[game2d.Pos{2, 2}] = '3'
	pad.layout[game2d.Pos{1, 3}] = '0'
	pad.layout[game2d.Pos{2, 3}] = 'A'

	for pos, val := range pad.layout {
		pad.buttons[val] = pos
	}
	pad.Pos = pad.buttons['A']
	return &pad
}

func NewDirectionalPad(id string) *Pad {
	pad := Pad{
		id:      id,
		layout:  make(map[game2d.Pos]uint8),
		buttons: make(map[uint8]game2d.Pos),
	}
	pad.layout[game2d.Pos{1, 0}] = '^'
	pad.layout[game2d.Pos{2, 0}] = 'A'
	pad.layout[game2d.Pos{0, 1}] = '<'
	pad.layout[game2d.Pos{1, 1}] = 'v'
	pad.layout[game2d.Pos{2, 1}] = '>'

	for pos, val := range pad.layout {
		pad.buttons[val] = pos
	}
	pad.Pos = pad.buttons['A']
	return &pad
}

func (p *Pad) allPaths(start, end game2d.Pos) []string {
	//fmt.Printf("allPaths start: %v end: %v\n", start, end)
	if start == end {
		return []string{""}
	}

	var neighbors []game2d.Pos
	var dist = game2d.ManhattanDistance(start, end)
	for n := range start.Neighbors4() {
		if _, ok := p.layout[n]; ok && game2d.ManhattanDistance(n, end) < dist {
			neighbors = append(neighbors, n)
		}
	}

	var res []string
	for _, n := range neighbors {
		var all = p.allPaths(n, end)
		for _, path := range all {
			var move string
			if n.X > start.X {
				move = ">"
			} else if n.X < start.X {
				move = "<"
			} else if n.Y > start.Y {
				move = "v"
			} else if n.Y < start.Y {
				move = "^"
			}
			res = append(res, move+path)
		}
	}
	return res
}

func StringPath(p []string) string {
	return "[" + strings.Join(p, ", ") + "]"
}

func (p *Pad) Reach(button uint8) []string {
	var target = p.buttons[button]
	var paths = p.allPaths(p.Pos, target)
	p.Pos = target
	return paths
}

type Entry struct {
	order string
	level int
}

func length(order string, level int, transitionTable map[string]map[string][]string, cache map[Entry]int) int {
	if level == 0 {
		return len(order)
	}

	if v, ok := cache[Entry{order, level}]; ok {
		return v
	}

	var res int
	var from = "A"
	for _, c := range order {
		var list = transitionTable[from][string(c)]
		if len(list) == 1 {
			res += length(list[0], level-1, transitionTable, cache)
		} else if len(list) == 2 {
			l1 := length(list[0], level-1, transitionTable, cache)
			l2 := length(list[1], level-1, transitionTable, cache)
			res += min(l1, l2)
		}
		from = string(c)
	}
	cache[Entry{order, level}] = res
	return res
}

func productString(a []string, b []string) []string {
	if len(a) == 0 {
		return b
	}
	var res []string
	for _, pa := range a {
		for _, pb := range b {
			res = append(res, pa+pb)
		}
	}
	return res
}

func solve(input string, n int) int {
	var lines = strings.Split(input, "\n")

	var numericPad = NewNumericPad("numeric")
	var transisionTable = make(map[string]map[string][]string)
	transisionTable["^"] = map[string][]string{"^": {"A"}, "v": {"vA"}, "<": {"v<A"}, ">": {">vA", "v>A"}, "A": {">A"}}
	transisionTable["v"] = map[string][]string{"^": {"^A"}, "v": {"A"}, "<": {"<A"}, ">": {">A"}, "A": {">^A", "^>A"}}
	transisionTable["<"] = map[string][]string{"^": {">^A"}, "v": {">A"}, "<": {"A"}, ">": {">>A"}, "A": {">>^A", ">^>A"}}
	transisionTable[">"] = map[string][]string{"^": {"<^A", "^<A"}, "v": {"<A"}, "<": {"<<A"}, ">": {"A"}, "A": {"^A"}}
	transisionTable["A"] = map[string][]string{"^": {"<A"}, "v": {"<vA", "v<A"}, "<": {"v<<A", "<v<A"}, ">": {"vA"}, "A": {"A"}}

	var cache = make(map[Entry]int)

	var res int
	for _, line := range lines {
		var num, _ = strconv.Atoi(line[:len(line)-1])
		var s int
		for _, c := range line {
			var orders = numericPad.Reach(uint8(c))
			orders = productString(orders, []string{"A"})
			var m = math.MaxInt
			for _, order := range orders {
				m = min(m, length(order, n, transisionTable, cache))
			}
			s += m
		}
		res += s * num
	}
	return res
}

func Part1(input string) int {
	return solve(input, 2)
}

func Part2(input string) int {
	return solve(input, 25)
}

func main() {
	fmt.Println("--2024 day 21 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
