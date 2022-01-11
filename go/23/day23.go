package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code-2021/go/utils"
)

//go:embed input.txt
var input_day string

// kind of field
const (
	Hallway byte = iota
	Room
	Door
	Wall
)

const hallwayY int = 1

var hallwayPos []Pos = []Pos{{1, 1}, {2, 1}, {4, 1}, {6, 1}, {8, 1}, {10, 1}, {11, 1}}
var roomX = map[byte]int{'A': 3, 'B': 5, 'C': 7, 'D': 9}
var costMove = map[byte]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}

type Pos struct {
	x, y int
}

type MoveCost struct {
	src  Pos
	dest Pos
	cost int
}

// occupant of field can be 0, 'A', 'B', 'C', 'D'
const empty byte = 0

type Field struct {
	kind     byte
	occupant byte
	atHome   bool
}

type World struct {
	maxX, maxY int
	grid       map[Pos]Field
}

func (f Field) String() string {
	return fmt.Sprintf("%v %c", f.kind, f.occupant)
}

func (w World) String() string {
	var sb strings.Builder
	for y := 0; y < w.maxY; y++ {
		for x := 0; x < w.maxX; x++ {
			pos := Pos{x: x, y: y}
			f, ok := w.grid[pos]
			if ok {
				switch f.kind {
				case Wall:
					sb.WriteByte('#')
				case Door:
					sb.WriteByte('_')
				case Hallway, Room:
					if f.occupant != empty {
						sb.WriteByte(f.occupant)
					} else {
						sb.WriteByte('.')
					}
				}
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// #############
// #...........#
// ###D#A#B#C###
//   #B#A#D#C#
//   #########
func createWorld(lines []string) World {
	world := World{
		grid: make(map[Pos]Field),
		maxX: len(lines[0]),
		maxY: len(lines),
	}
	door := map[int]struct{}{3: {}, 5: {}, 7: {}, 9: {}}
	for y, line := range lines {
		for x, char := range line {
			pos := Pos{x: x, y: y}
			switch char {
			case '#':
				world.grid[pos] = Field{kind: Wall}
			case '.':
				if _, ok := door[x]; ok {
					world.grid[pos] = Field{kind: Door}
				} else {
					world.grid[pos] = Field{kind: Hallway}
				}
			case 'A', 'B', 'C', 'D':
				world.grid[pos] = Field{kind: Room, occupant: byte(char)}
			}
		}
	}
	return world
}

func createGoal(w World) World {
	res := World{
		grid: make(map[Pos]Field, len(w.grid)),
		maxX: w.maxX,
		maxY: w.maxY,
	}
	for p, f := range w.grid {
		if f.kind == Room && p.y >= 2 {
			switch p.x {
			case 3:
				f.occupant = 'A'
			case 5:
				f.occupant = 'B'
			case 7:
				f.occupant = 'C'
			case 9:
				f.occupant = 'D'
			}
		}
		res.grid[p] = f
	}
	return res
}

func (w World) moveHome(src, dest Pos) World {
	res := w.move(src, dest)
	f := res.grid[dest]
	f.atHome = true
	res.grid[dest] = f
	return res
}

func (w World) move(src, dest Pos) World {
	res := World{
		grid: make(map[Pos]Field, len(w.grid)),
		maxX: w.maxX,
		maxY: w.maxY,
	}
	for p, f := range w.grid {
		if p == src {
			f.occupant = empty
		} else if p == dest {
			f.occupant = w.occupant(src)
		}
		res.grid[p] = f
	}
	return res
}

func (w World) atHome(p Pos) bool {
	return w.grid[p].atHome
}

func (w World) occupied(p Pos) bool {
	field, ok := w.grid[p]
	if !ok {
		return false
	}
	return field.occupant != empty
}

func (w World) occupant(p Pos) byte {
	field, ok := w.grid[p]
	if !ok {
		panic("invalid pos")
	}
	return field.occupant
}

func (w World) accessibleHallway(srcX, destX int) bool {
	if srcX == destX {
		return true
	}

	if srcX < destX {
		for x := srcX + 1; x <= destX; x++ {
			if w.occupied(Pos{x, hallwayY}) {
				return false
			}
		}
	}
	for x := srcX - 1; x >= destX; x-- {
		if w.occupied(Pos{x, hallwayY}) {
			return false
		}
	}

	return true
}

func (w World) blockedHallway() bool {
	for x1 := 4; x1 <= 8; x1 += 2 {
		occupant1 := w.occupant(Pos{x: x1, y: hallwayY})
		for x2 := 4; x2 <= 8; x2 += 2 {
			occupant2 := w.occupant(Pos{x: x2, y: hallwayY})
			if x1 != x2 && occupant1 != empty && occupant2 != empty && roomX[occupant1] > x2 && roomX[occupant2] < x1 {
				return true
			}
		}
	}
	return false
}

func (w World) freeHomeY(roomX int) (int, bool) {
	for y := w.maxY - 2; y >= 2; y -= 1 {
		if !w.occupied(Pos{roomX, y}) {
			return y, true
		}
		if !w.atHome(Pos{roomX, y}) {
			return y, false
		}
	}
	return 0, false
}

func (w World) moveHallwayToHome() (World, int) {
	cost := 0
	stop := false
	for !stop {
		stop = true
		for _, p := range hallwayPos {
			if w.occupied(p) {
				occupant := w.occupant(p)
				homeX := roomX[occupant]
				if w.accessibleHallway(p.x, homeX) {
					if homeY, ok := w.freeHomeY(homeX); ok {
						home := Pos{homeX, homeY}
						cost += manhattanDistance(p, home) * costMove[occupant]
						w = w.moveHome(p, home)
						stop = false
					}
				}
			}
		}
	}
	return w, cost
}

func (w World) moveRoomToHome(x int) (World, int) {
	cost := 0
	for roomY := 2; roomY <= w.maxY-2; roomY++ {
		p := Pos{x, roomY}
		if w.occupied(p) {
			if w.atHome(p) {
				return w, cost
			}

			occupant := w.occupant(p)
			homeX := roomX[occupant]
			homeY, ok := w.freeHomeY(homeX)
			if ok && w.accessibleHallway(x, homeX) {
				distance := utils.Abs(homeX-p.x) + (homeY - hallwayY) + (p.y - hallwayY)
				cost = cost + distance*costMove[occupant]
				w = w.moveHome(p, Pos{homeX, homeY})
			} else {
				return w, cost
			}
		}
	}
	return w, cost
}

func (w World) moveRoomToHallway(roomX int) []MoveCost {
	var res []MoveCost
	for roomY := 2; roomY <= w.maxY-2; roomY++ {
		p := Pos{roomX, roomY}
		if w.occupied(p) {
			if !w.atHome(p) {
				occupant := w.occupant(p)
				for _, h := range hallwayPos {
					if w.accessibleHallway(p.x, h.x) {
						cost := manhattanDistance(p, h) * costMove[occupant]
						res = append(res, MoveCost{src: p, dest: h, cost: cost})
						// fmt.Printf("%c %v can reach hallway %v with cost: %d\n", occupant, p, Pos{x, hallwayY}, cost)
					}
				}
			}
			return res
		}
	}
	return res
}

type State struct {
	world World
	cost  int
}

func (w World) step() []State {
	var res []State
	var cost int
	var c int

	if w.blockedHallway() {
		return res
	}

	w, c = w.moveHallwayToHome()
	cost += c

	// this is an optimization, this should not be necessary
	for roomX := 3; roomX <= 9; roomX += 2 {
		w, c = w.moveRoomToHome(roomX)
		cost += c
	}

	for roomX := 3; roomX <= 9; roomX += 2 {
		for _, m := range w.moveRoomToHallway(roomX) {
			res = append(res, State{world: w.move(m.src, m.dest), cost: cost + m.cost})
		}
	}

	if len(res) == 0 && cost > 0 {
		res = append(res, State{w, cost})
	}

	return res
}

type node struct {
	World
	priority int
	index    int
}

func manhattanDistance(from, to Pos) int {
	absX := from.x - to.x
	if absX < 0 {
		absX = -absX
	}
	absY := from.y - to.y
	if absY < 0 {
		absY = -absY
	}
	return absX + absY
}

func heuristicDistance(w World) int {
	var res int
	cpt := map[byte]int{}

	for x := 3; x <= 9; x += 2 {
		for y := w.maxX - 2; y >= 2; y-- {
			p := Pos{x, y}
			if w.occupied(p) && !w.atHome(p) {
				occupant := w.occupant(p)
				cpt[occupant] = cpt[occupant] + 1
				distance := cpt[occupant] + manhattanDistance(p, Pos{roomX[occupant], hallwayY})
				res += distance * costMove[occupant]
			}
		}
	}
	for _, p := range hallwayPos {
		if w.occupied(p) {
			occupant := w.occupant(p)
			cpt[occupant] = cpt[occupant] + 1
			distance := cpt[occupant] + manhattanDistance(p, Pos{roomX[occupant], hallwayY})
			res += distance * costMove[occupant]
		}
	}

	// fmt.Printf("heuristic distance: %d\n%v\n", res, w)
	return res
}

func signature(w World) string {
	var sb strings.Builder
	for _, p := range hallwayPos {
		sb.WriteByte(w.occupant(p))
	}
	for y := 2; y <= w.maxY-2; y++ {
		for x := 3; x <= 9; x += 2 {
			sb.WriteByte(w.occupant(Pos{x, y}))
		}
	}
	return sb.String()
}

func path(start, to World) (path []World, distance int) {
	toSignature := signature(to)
	startSignature := signature(start)

	frontier := &PriorityQueue{}
	heap.Init(frontier)
	heap.Push(frontier, &node{World: start, priority: 0})

	cameFrom := map[string]World{startSignature: start}
	costSoFar := map[string]int{startSignature: 0}

	for {
		if frontier.Len() == 0 {
			// There's no path, return found false.
			return
		}
		var current World = heap.Pop(frontier).(*node).World
		var currentSignature = signature(current)

		// fmt.Printf("signature: %s\ncurrent:\n%v\n", currentSignature, current)
		if currentSignature == toSignature {
			// Found a path to the goal.
			var path []World
			currentSignature = signature(current)
			for currentSignature != startSignature {
				path = append(path, current)
				current = cameFrom[currentSignature]
				currentSignature = signature(current)
			}
			return path, costSoFar[toSignature]
		}

		var next []State = current.step()
		for _, neighbor := range next {
			newCost := costSoFar[currentSignature] + neighbor.cost
			neighborSignature := signature(neighbor.world)
			if _, ok := costSoFar[neighborSignature]; !ok || newCost < costSoFar[neighborSignature] {
				costSoFar[neighborSignature] = newCost
				priority := newCost + heuristicDistance(neighbor.world)
				heap.Push(frontier, &node{World: neighbor.world, priority: priority})
				cameFrom[neighborSignature] = current
			}
		}
	}
}

// A PriorityQueue implements heap.Interface and holds Items.
// Code copied from https://pkg.go.dev/container/heap@go1.17.5
type PriorityQueue []*node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func Part1(input string) int {
	var d int
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	w := createWorld(lines)
	_, d = path(w, createGoal(w))
	return d
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	l := strings.Split(input, "\n")
	l1 := "  #D#C#B#A#  "
	l2 := "  #D#B#A#C#  "
	lines := []string{l[0], l[1], l[2], l1, l2, l[3], l[4]}
	w := createWorld(lines)
	_, d := path(w, createGoal(w))
	return d
	// return 0
}

func main() {
	fmt.Println("--2021 day 23 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}
