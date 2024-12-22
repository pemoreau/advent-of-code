package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"math"
	"slices"
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

type Path = []uint8

func (pad *Pad) allPaths(start, end game2d.Pos) []Path {
	if start == end {
		return []Path{{}}
	}

	var neighbors []game2d.Pos
	var dist = game2d.ManhattanDistance(start, end)
	for p := range start.Neighbors4() {
		if _, ok := pad.layout[p]; ok && game2d.ManhattanDistance(p, end) < dist {
			neighbors = append(neighbors, p)
		}
	}

	var res []Path
	for _, n := range neighbors {
		var all = pad.allPaths(n, end)
		var allNewPaths []Path
		for _, path := range all {
			var move uint8
			if n.X > start.X {
				move = '>'
			} else if n.X < start.X {
				move = '<'
			} else if n.Y > start.Y {
				move = 'v'
			} else if n.Y < start.Y {
				move = '^'
			}
			var newPath = append([]uint8{move}, path...)
			allNewPaths = append(allNewPaths, newPath)
		}
		res = append(res, allNewPaths...)
	}
	return res
}

func StringPath(p []Path) string {
	var res string
	res += "["
	for i, path := range p {
		if i > 0 {
			res += ", "
		}
		res += string(path)
	}
	res += "]"
	return res
}

func (pad *Pad) Reach(button uint8) []Path {
	var target = pad.buttons[button]
	var paths = pad.allPaths(pad.Pos, target)
	//fmt.Printf("Reach %c from: %c: %v\n", button, pad.layout[pad.Pos], StringPath(paths))
	pad.Pos = target

	return paths
}

//func (p *Pad) Reach(button uint8) []uint8 {
//	var res []uint8
//	var target = p.buttons[button]
//	var x, y = p.Pos.X, p.Pos.Y
//	for x != target.X || y != target.Y {
//		if x < target.X {
//			res = append(res, '>')
//			x++
//		} else if y > target.Y {
//			res = append(res, '^')
//			y--
//		} else if y < target.Y {
//			res = append(res, 'v')
//			y++
//		} else if x > target.X {
//			res = append(res, '<')
//			x--
//		}
//	}
//	p.Pos = target
//	return res
//}

//	func (p *Pad) ComputeOrders(compose string) []uint8 {
//		var res []uint8
//		for _, c := range compose {
//			res = append(res, p.Reach(uint8(c))...)
//			res = append(res, uint8('A'))
//		}
//		return res
//	}

func product(a []Path, b []Path) []Path {
	var res []Path
	if len(a) == 0 {
		return b
	}
	for _, pa := range a {
		for _, pb := range b {
			var newPath = append(pa, pb...)
			res = append(res, slices.Clone(newPath))
		}
	}
	//fmt.Printf("product: %v * %v -> %v\n", StringPath(a), StringPath(b), StringPath(res))

	return res
}

//func product(a []Path, b []Path) []Path {
//	var res []Path
//	if len(a) == 0 {
//		return b
//	}
//	var head = a[0]
//	var tail = a[1:]
//	var paths = product(tail, b)
//	for _, p := range paths {
//		var newPath = append(head, p...)
//		//fmt.Printf("head: %s p: %s -> %s\n", string(head), string(p), string(newPath))
//		res = append(res, slices.Clone(newPath))
//		//fmt.Printf("res: %s\n", StringPath(res))
//	}
//	fmt.Printf("product: %v * %v -> %v\n", StringPath(a), StringPath(b), StringPath(res))
//	return res
//}

func (p *Pad) ComputeOrders(compose []uint8) []Path {
	var res [][]uint8
	for _, c := range compose {
		var orders = p.Reach(c)
		orders = product(orders, []Path{[]uint8{'A'}})
		res = product(res, orders)
	}
	//fmt.Printf("compose: %s orders: %v\n", string(compose), StringPath(res))
	return res
}

func (p *Pad) ComputeMultiOrders(multiCompose [][]uint8) []Path {
	var res []Path
	for _, compose := range multiCompose {
		var orders = p.ComputeOrders(compose)
		res = append(res, orders...)
	}
	return res
}

func (p *Pad) Move(direction uint8) {
	var newPos = p.Pos
	switch direction {
	case '>':
		newPos.X++
	case '<':
		newPos.X--
	case '^':
		newPos.Y--
	case 'v':
		newPos.Y++
	}
	if _, ok := p.layout[newPos]; ok {
		p.Pos = newPos
	}
}

func removeDuplicates(paths []Path) []Path {
	var s = set.NewSet[string]()
	for _, path := range paths {
		s.Add(string(path))
	}
	var res []Path
	for path := range s.All() {
		res = append(res, []uint8(path))
	}
	return res
}

func Part1(input string) int {
	var lines = strings.Split(input, "\n")

	var numericPad = NewNumericPad("numeric")
	var command1 = NewDirectionalPad("command-1")
	var command2 = NewDirectionalPad("command-2")

	var res int
	for _, line := range lines {
		//fmt.Printf("line: %s\n", line)
		var orders1 = numericPad.ComputeMultiOrders([][]uint8{[]uint8(line)})
		orders1 = removeDuplicates(orders1)
		//for _, order := range orders1 {
		//	fmt.Printf("    orders1: %s\n", string(order))
		//}
		var orders2 = command1.ComputeMultiOrders(orders1)
		orders2 = removeDuplicates(orders2)
		//for _, order := range orders2 {
		//	fmt.Printf("    orders2: %s\n", string(order))
		//}
		var orders3 = command2.ComputeMultiOrders(orders2)
		orders3 = removeDuplicates(orders3)
		//for _, order := range orders3 {
		//	fmt.Printf("    orders3: %s\n", string(order))
		//}

		var num, _ = strconv.Atoi(line[:len(line)-1])
		var minorder = math.MaxInt
		for _, order := range orders3 {
			if len(order) < minorder {
				minorder = len(order)
			}
		}
		fmt.Printf("    num: %d * len: %d\n", num, minorder)
		res += num * minorder
	}

	return res
}

func Part2(input string) int {
	//var lines = strings.Split(input, "\n")
	//
	//var pads []*Pad
	//pads = append(pads, NewNumericPad("numeric"))
	//for i := 0; i < 25; i++ {
	//	pads = append(pads, NewDirectionalPad(fmt.Sprintf("command-%d", i)))
	//}
	//
	//var res int
	//for _, line := range lines {
	//	var orders = [][]uint8{[]uint8(line)}
	//	for _, pad := range pads {
	//		orders = pad.ComputeMultiOrders(orders)
	//		orders = removeDuplicates(orders)
	//	}
	//
	//	var num, _ = strconv.Atoi(line[:len(line)-1])
	//	var minorder = math.MaxInt
	//	for _, order := range orders {
	//		if len(order) < minorder {
	//			minorder = len(order)
	//		}
	//	}
	//	fmt.Printf("    num: %d * len: %d\n", num, minorder)
	//	res += num * minorder
	//}
	//return res

	// +-+-+-+
	// | |^|A|
	// +-+-+-+
	// |<|v|>|
	// +-+-+-+
	var transisionTable = make(map[string]map[string][]string)
	transisionTable["^"] = map[string][]string{"^": {"A"}, "v": {"vA"}, "<": {"v<A"}, ">": {"v>A", ">vA"}, "A": {">A"}}
	transisionTable["v"] = map[string][]string{"^": {"^A"}, "v": {"A"}, "<": {"<A"}, ">": {">A"}, "A": {">^A", "^>A"}}
	transisionTable["<"] = map[string][]string{"^": {">^A"}, "v": {">A"}, "<": {"A"}, ">": {">>A"}, "A": {">>^A", ">^>A"}}
	transisionTable[">"] = map[string][]string{"^": {"<^A", "^<A"}, "v": {"<A"}, "<": {"<<A"}, ">": {"A"}, "A": {"^A"}}
	transisionTable["A"] = map[string][]string{"^": {"<A"}, "v": {"v<A", "<vA"}, "<": {"v<<A", "<v<A"}, ">": {"vA"}, "A": {"A"}}

	//var transisionTable = make(map[string]map[string][]string)
	//transisionTable["^"] = map[string][]string{"^": {}, "v": {"v"}, "<": {"v<"}, ">": {"v>", ">v"}, "A": {">"}}
	//transisionTable["v"] = map[string][]string{"^": {"^"}, "v": {}, "<": {"<"}, ">": {">"}, "A": {">^", "^>"}}
	//transisionTable["<"] = map[string][]string{"^": {">^"}, "v": {">"}, "<": {}, ">": {">>"}, "A": {">>^", ">^>"}}
	//transisionTable[">"] = map[string][]string{"^": {"^<", "<^"}, "v": {"<"}, "<": {"<<"}, ">": {}, "A": {"^"}}
	//transisionTable["A"] = map[string][]string{"^": {"<"}, "v": {"v<", "<v"}, "<": {"<v<", "v<<"}, ">": {"v"}, "A": {}}

	var lines = strings.Split(input, "\n")
	var res int
	for _, line := range lines {
		var num, _ = strconv.Atoi(line[:len(line)-1])
		var numericPad = NewNumericPad("numeric")
		var orders = numericPad.ComputeMultiOrders([][]uint8{[]uint8(line)})

		var minLine = math.MaxInt
		for _, order := range orders {
			var soup = make(map[string]int)

			fillMap(string(order), 1, soup)
			fmt.Printf("map: %v\n", soup)
			soup = mapStep(soup, transisionTable)

			//var toEncode = string(order)
			fmt.Println("order: ", string(order))
			for i := 1; i <= 25; i++ {
				soup = mapStep(soup, transisionTable)
				//toEncode = encodeOneLevel('A', toEncode, transisionTable)
				//fmt.Printf("level %d: %d\n", i, len(toEncode))
				//fmt.Printf("level %d: %d %s\n", i, len(toEncode), toEncode)
				fmt.Printf("level %d: %d %v\n", i, lenMap(soup), soup)
				//res += num * lenMap(soup)

			}
			minLine = min(minLine, lenMap(soup))
			fmt.Println("minLine: ", minLine)

			//var encoded = encodeString('A', string(order), 2, transisionTable)
		}
		res += num * minLine
	}

	return res
}

func sumDist(s string, table map[string]map[string][]string) int {
	var last = 'A'
	var res int
	for _, c := range s {
		res += len(table[string(last)][string(c)][0])
		last = c
	}
	return res
}

var cache = make(map[string]string)

func orderToMap(order string) map[string]int {
	var res = make(map[string]int)

	for _, c := range strings.SplitAfter(order, "A") {
		if len(c) == 0 {
			continue
		}
		if _, ok := res[string(c)]; ok {
			res[string(c)]++
		} else {
			res[string(c)] = 1
		}
	}
	return res
}

func fillMap(order string, mult int, m map[string]int) {
	for _, c := range strings.SplitAfter(order, "A") {
		if len(c) == 0 {
			continue
		}
		if _, ok := m[string(c)]; ok {
			m[string(c)] += mult
		} else {
			m[string(c)] = mult
		}
	}
}

func lenMap(m map[string]int) int {
	var res int
	for _, v := range m {
		res += v
	}
	return res
}

func mapStep(m map[string]int, transisionTable map[string]map[string][]string) map[string]int {
	var res = make(map[string]int)
	for k, v := range m {
		s := encodeOneLevel('A', k, transisionTable)
		fillMap(s, v, res)
	}
	return res
}

func encodeOneLevel(prev byte, s string, table map[string]map[string][]string) string {
	if len(s) == 0 {
		return s
	}

	//if c, ok := cache[s]; ok {
	//	return c
	//}

	head := s[0]
	tail := s[1:]

	var encodedTail = encodeOneLevel(head, tail, table)

	var list = table[string(prev)][string(head)]
	//fmt.Printf("(%c,%c) -> %v\n", prev, head, list[0])

	//var best = math.MaxInt
	var bestString string
	//for _, s := range list {
	//	if len(s+encodedTail) < best {
	//		bestString = s + encodedTail
	//		best = len(bestString)
	//	}
	//}

	bestString = list[0] + encodedTail

	//cache[s] = bestString
	return bestString
}

//func encode(from, to byte, level int, table map[string]map[string][]string) string {
//	var list = table[string(from)][string(to)]
//	if len(list) == 0 {
//		return encodeString("A", level-1, table)
//	}
//	if len(list) == 1 {
//		return encodeString(list[0]+"A", level-1, table)
//	}
//	var s1 = encodeString(list[0]+"A", level-1, table)
//	var s2 = encodeString(list[1]+"A", level-1, table)
//	if len(s1) < len(s2) {
//		return s1
//	}
//	return s2
//}

func encodeString(prev byte, s string, level int, table map[string]map[string][]string) string {
	if level <= 0 {
		return s
	}
	if len(s) == 0 {
		return s
	}

	head := s[0]
	tail := s[1:]

	var list = table[string(prev)][string(head)]
	var encodedTail = encodeString(head, tail, level, table)
	var best = math.MaxInt
	var bestString string
	for _, s := range list {
		//fmt.Printf("s: %s encodedTail: %s\n", s, encodedTail)
		var encoded = encodeString('A', s+encodedTail, level-1, table)
		if len(encoded) < best {
			best = len(encoded)
			bestString = encoded
		}
	}
	return bestString
}

// 396092315483086 too high
func main() {
	fmt.Println("--2024 day 21 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	//fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
