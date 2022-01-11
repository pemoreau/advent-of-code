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
	door := map[int]struct{}{3: {}, 5: {}, 7: {}, 9: {}}
	for y, line := range lines {
		for x, char := range line {
			pos := Pos{x: x, y: y}
			switch char {
			case '#':
				// world.grid[pos] = Field{kind: Wall}
			case '.':
				if _, ok := door[x]; ok {
					// world.grid[pos] = Field{kind: Door}
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
	if correctRoom(f.occupant, p.x) {
		for roomY := w.maxY - 2; roomY > p.y; roomY-- {
			occupant := w.grid[Pos{p.x, roomY}].occupant
			if occupant == 0 || !correctRoom(occupant, p.x) {
				return false
			}
		}
		return true
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

func (w World) blockedHallway() bool {
	for x1 := 4; x1 <= 8; x1 += 2 {
		occupant1 := w.grid[Pos{x: x1, y: hallwayY}].occupant
		for x2 := 4; x2 <= 8; x2 += 2 {
			occupant2 := w.grid[Pos{x: x2, y: hallwayY}].occupant
			if x1 != x2 && occupant1 != 0 && occupant2 != 0 && roomX[occupant1] > x2 && roomX[occupant2] < x1 {
				return true
			}
		}
	}
	return false
}

func (w World) freeHomeY(roomX int) (int, bool) {
	for y := w.maxY - 2; y >= 2; y -= 1 {
		occupant := w.grid[Pos{roomX, y}].occupant
		if occupant == 0 {
			return y, true
		} else if !correctRoom(occupant, roomX) {
			return 0, false
		}
	}
	return 0, false
}

func (w World) moveHallwayToHome() (World, int) {
	for x := 1; x <= 11; x++ {
		p := Pos{x, hallwayY}
		if w.occupied(p) {
			occupant := w.grid[p].occupant
			targetRoom := roomX[occupant]
			if w.accessibleHallway(x, targetRoom) {
				y, ok := w.freeHomeY(targetRoom)
				if ok {
					distance := utils.Abs(targetRoom-p.x) + y - hallwayY
					cost := distance * costMove[occupant]
					// fmt.Printf("move %c %v to %d,%d cost = %d\n", occupant, p, targetRoom, y, cost)
					res := w.move(p, Pos{targetRoom, y})
					// fmt.Println(res)
					return res, cost
				}
			}
		}
	}
	return w, 0
}

func (w World) moveRoomToHome() (World, int) {
	for x := 3; x <= 9; x += 2 {
		first := true
		for roomY := 2; roomY <= w.maxY-2 && first; roomY++ {
			p := Pos{x, roomY}
			if w.occupied(p) {
				first = false
				occupant := w.grid[p].occupant
				homeX := roomX[occupant]
				if x != homeX && w.accessibleHallway(x, homeX) {
					homeY, ok := w.freeHomeY(homeX)
					if ok {
						distance := utils.Abs(homeX-p.x) + (homeY - hallwayY) + (roomY - hallwayY)
						cost := distance * costMove[occupant]
						// fmt.Printf("move %c %v to %d,%d cost = %d\n", occupant, p, homeX, homeY, cost)
						res := w.move(p, Pos{homeX, homeY})
						// fmt.Println(res)
						return res, cost
					}
				}
			}
		}
	}
	return w, 0
}

func (w World) moveRoomToHallway() []MoveCost {
	var res []MoveCost
	for x := 3; x <= 9; x += 2 {
		first := true
		for roomY := 2; roomY <= w.maxY-2 && first; roomY++ {
			p := Pos{x, roomY}
			if w.occupied(p) {
				first = false
				occupant := w.grid[p].occupant
				if !w.atHome(p) {
					// for hallwayX := 1; hallwayX <= 11; hallwayX++ {
					// 	if w.grid[Pos{hallwayX, hallwayY}].kind == Hallway && w.accessibleHallway(p.x, hallwayX) {
					for _, h := range hallwayPos {
						if w.grid[h].kind == Hallway && w.accessibleHallway(p.x, h.x) {
							distance := utils.Abs(h.x-p.x) + p.y - hallwayY
							cost := distance * costMove[occupant]
							res = append(res, MoveCost{src: p, dest: Pos{h.x, hallwayY}, cost: cost})
							// fmt.Printf("%c %v can reach hallway %v with cost: %d\n", occupant, p, Pos{x, hallwayY}, cost)
						}
					}
				}
			}
		}
	}
	return res
}

func (w World) step() []State {
	var res []State

	if w.blockedHallway() {
		return res
	}
	var cost, c int
	w, c = w.moveHallwayToHome()
	for c > 0 {
		cost += c
		w, c = w.moveHallwayToHome()
	}

	w, c = w.moveRoomToHome()
	for c > 0 {
		cost += c
		w, c = w.moveRoomToHome()
	}

	// optim
	// for p, f := range w.grid {
	// 	if f.occupant != 0 {
	// 		for _, d := range w.reachable(p) {
	// 			res = append(res, State{world: w.move(p, d.Pos), cost: cost + d.cost})
	// 		}
	// 	}
	// }

	// new code
	for _, m := range w.moveRoomToHallway() {
		res = append(res, State{world: w.move(m.src, m.dest), cost: cost + m.cost})
	}

	if len(res) == 0 && cost > 0 {
		res = append(res, State{w, cost})
	}

	return res
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

	var cpt int
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
			cpt++
			// fmt.Printf("cpt = %d\n", cpt)
			newCost := costSoFar[currentSignature] + neighbor.cost
			neighborSignature := signature(neighbor.world)
			if _, ok := costSoFar[neighborSignature]; !ok || newCost < costSoFar[neighborSignature] {
				costSoFar[neighborSignature] = newCost
				priority := newCost //+ heuristicDistance(neighbor.world)
				heap.Push(frontier, &node{World: neighbor.world, priority: priority})
				cameFrom[neighborSignature] = current
				// fmt.Printf("current:\n%v", current)
				// fmt.Printf("len: %d\n", len(*frontier))
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
