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

var roomX = map[byte]int{'A': 3, 'B': 5, 'C': 7, 'D': 9}
var costMove = map[byte]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}

type Pos struct {
	x, y int
}

type PosCost struct {
	Pos
	cost int
}

// occupant of field can be 0, 'A', 'B', 'C', 'D'
type Field struct {
	kind     byte
	occupant byte
}

type World struct {
	maxX, maxY int
	grid       map[Pos]Field
}

type State struct {
	world World
	cost  int
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
					if f.occupant != 0 {
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
	door := map[int]struct{}{3: struct{}{}, 5: struct{}{}, 7: struct{}{}, 9: struct{}{}}
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

func (w World) move(src, dest Pos) World {
	res := World{
		grid: make(map[Pos]Field, len(w.grid)),
		maxX: w.maxX,
		maxY: w.maxY,
	}
	for p, f := range w.grid {
		if p == src {
			f.occupant = 0
		} else if p == dest {
			f.occupant = w.grid[src].occupant
		}
		res.grid[p] = f
	}
	return res
}

func correctRoom(occupant byte, x int) bool {
	targetX, ok := roomX[occupant]
	if !ok {
		fmt.Printf("invalid occupant %c x=%d\n", occupant, x)
		panic("invalid occupant")
	}
	return targetX == x
}

// In the correct room and not blocking anyone
func (w World) atHome(p Pos) bool {
	f := w.grid[p]
	if f.kind == Room {
		if correctRoom(f.occupant, p.x) {
			for y := p.y + 1; y < w.maxY-1; y++ {
				if !w.atHome(Pos{x: p.x, y: y}) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (w World) occupied(p Pos) bool {
	return w.grid[p].occupant != 0
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

func (w World) reachable(p Pos) []PosCost {
	var res []PosCost
	occupant := w.grid[p].occupant
	if occupant == 0 {
		panic("no occupant")
	}

	if w.atHome(p) {
		// fmt.Printf("%c %v at home: do not move\n", occupant, p)
		return res
	}
	if w.grid[p].kind == Room && p.y > 2 && w.occupied(Pos{p.x, p.y - 1}) {
		// fmt.Printf("%c %v cannot move up\n", occupant, p)
		return res
	}
	if w.grid[p].kind == Hallway {
		targetRoom := roomX[occupant]
		if !w.accessibleHallway(p.x, targetRoom) {
			// fmt.Printf("%c %v cannot access door %d\n", occupant, p, targetRoom)
			return res
		}
		// search for a free slot in my Room
		y := w.maxY - 2
		for w.occupied(Pos{targetRoom, y}) {
			y--
		}
		if y < 1 {
			// fmt.Printf("%c %v did not find free slot in room %d\n", occupant, p, targetRoom)
			return res
		}
		if y+1 < w.maxY-1 && !w.atHome(Pos{targetRoom, y + 1}) {
			// fmt.Printf("%c %v blocked because %c %v is not at home\n", occupant, p, w.grid[Pos{targetRoom, y + 1}].occupant, Pos{targetRoom, y + 1})
			return res
		}
		distance := utils.Abs(targetRoom-p.x) + y - hallwayY
		cost := distance * costMove[occupant]
		res = append(res, PosCost{Pos{targetRoom, y}, cost})
		// fmt.Printf("%c %v can reach room %v with cost: %d\n", occupant, p, Pos{targetRoom, y}, cost)
		return res
	}
	if w.grid[p].kind == Room {
		// collect all possible fields in hallway
		for x := 1; x <= 11; x++ {
			if w.grid[Pos{x, hallwayY}].kind == Hallway && w.accessibleHallway(p.x, x) {
				distance := utils.Abs(x-p.x) + p.y - hallwayY
				cost := distance * costMove[occupant]
				res = append(res, PosCost{Pos{x, hallwayY}, cost})
				// fmt.Printf("%c %v can reach hallway %v with cost: %d\n", occupant, p, Pos{x, hallwayY}, cost)
			}
		}
	}

	return res
}

func (w World) step() []State {
	var res []State
	var src []Pos
	for p, f := range w.grid {
		if f.occupant != 0 {
			src = append(src, p)
		}
	}
	for _, s := range src {
		for _, d := range w.reachable(s) {
			res = append(res, State{world: w.move(s, d.Pos), cost: d.cost})
		}
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	w := createWorld(lines)
	_, d := path(w, createGoal(w))
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
	// for _, w := range p {
	// 	fmt.Printf("%v\n", w)
	// }
	return d
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
	for p, f := range w.grid {
		if f.occupant != 0 && p.x != roomX[f.occupant] {
			res += (manhattanDistance(p, Pos{x: roomX[f.occupant], y: hallwayY}) * costMove[f.occupant])
		}
	}
	// fmt.Printf("heuristic distance: %d\n%v\n", res, w)
	return res
}

func signature(w World) string {
	return w.String()
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
