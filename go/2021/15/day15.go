package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := BuildMatrix(lines)
	from := Pos{0, 0}
	to := Pos{j: len(m) - 1, i: len(m[0]) - 1}
	_, cost := path(from, to, m)
	return cost
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := buildMegaMatrix(BuildMatrix(lines))
	from := Pos{0, 0}
	to := Pos{j: len(m) - 1, i: len(m[0]) - 1}
	_, cost := path(from, to, m)
	return cost
}

func main() {
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}

type matrix [][]uint8

func BuildMatrix(lines []string) matrix {
	m := make([][]uint8, len(lines))
	for j, l := range lines {
		l = strings.TrimSpace(l)
		m[j] = make([]uint8, len(l))
		for i, c := range l {
			m[j][i] = uint8(c - '0')
		}
	}
	return m
}

func buildMegaMatrix(m matrix) matrix {
	mm := make([][]uint8, 5*len(m))
	for j, l := range m {
		for kj := 0; kj < 5; kj++ {
			new_j := kj*len(m) + j
			mm[new_j] = make([]uint8, 5*len(l))
			for i, risk := range l {
				for ki := 0; ki < 5; ki++ {
					new_i := ki*len(l) + i
					new_risk := risk + uint8(ki) + uint8(kj)
					if new_risk > 9 {
						new_risk = new_risk % 9
					}
					mm[new_j][new_i] = uint8(new_risk)
				}
			}
		}
	}
	return mm
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.i, p.j)
}

func neighboors(m matrix, i, j int) []Pos {
	pos := []struct{ i, j int }{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	res := []Pos{}
	for _, p := range pos {
		if p.j >= 0 && p.j < len(m) && p.i >= 0 && p.i < len(m[0]) {
			res = append(res, p)
		}
	}
	return res
}

func manhattanDistance(from, to Pos) int {
	absX := from.i - to.i
	if absX < 0 {
		absX = -absX
	}
	absY := from.j - to.j
	if absY < 0 {
		absY = -absY
	}
	return absX + absY
}

type Pos struct{ i, j int }
type node struct {
	Pos
	priority int
	index    int
}

func path(start, to Pos, m matrix) (path []Pos, distance int) {
	frontier := &PriorityQueue{}
	heap.Init(frontier)
	heap.Push(frontier, &node{Pos: start, priority: 0})

	cameFrom := map[Pos]Pos{start: start}
	costSoFar := map[Pos]int{start: 0}

	for {
		if frontier.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := heap.Pop(frontier).(*node).Pos
		if current == to {
			// Found a path to the goal.
			path := []Pos{}
			curr := current
			for curr != start {
				path = append(path, curr)
				curr = cameFrom[curr]
			}
			return path, costSoFar[to]
		}

		for _, neighbor := range neighboors(m, current.i, current.j) {
			newCost := costSoFar[current] + int(m[neighbor.j][neighbor.i])
			if _, ok := costSoFar[neighbor]; !ok || newCost < costSoFar[neighbor] {
				costSoFar[neighbor] = newCost
				priority := newCost + manhattanDistance(neighbor, to)
				heap.Push(frontier, &node{Pos: neighbor, priority: priority})
				cameFrom[neighbor] = current
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
