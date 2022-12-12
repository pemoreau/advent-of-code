package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := utils.BuildIntMatrix(lines)
	from := search('S', m)
	to := search('E', m)
	m[from.Y][from.X] = 'a'
	m[to.Y][to.X] = 'z'
	_, cost := utils.Path(from, to, m, neighbors, cost, heuristic)
	return cost
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := utils.BuildIntMatrix(lines)

	from := search('S', m)
	to := search('E', m)
	m[from.Y][from.X] = 'a'
	m[to.Y][to.X] = 'z'
	_, cost := utils.Path(from, to, m, neighbors2, cost2, heuristic2)
	return cost
}

func main() {
	fmt.Println("--2022 day 12 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}

func search(v int, m utils.IntMatrix) utils.Pos {
	for j, l := range m {
		for i, c := range l {
			if c == v {
				return utils.Pos{i, j}
			}
		}
	}
	return utils.Pos{}
}

func neighbors(m utils.IntMatrix, i, j int) []utils.Pos {
	pos := []utils.Pos{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	res := []utils.Pos{}
	for _, p := range pos {
		if p.Y >= 0 && p.Y < len(m) && p.X >= 0 && p.X < len(m[0]) {
			src := m[j][i]
			dest := m[p.Y][p.X]
			if dest-src <= 1 {
				res = append(res, p)
			}
		}
	}
	return res
}

func cost(from, to utils.Pos, m utils.IntMatrix) int {
	return 1
}

func heuristic(from, to utils.Pos, m utils.IntMatrix) int {
	//return manhattanDistance(from, to)
	return m[to.Y][to.X] - m[from.Y][from.X]
}

func neighbors2(m utils.IntMatrix, i, j int) []utils.Pos {
	n := neighbors(m, i, j)
	if m[j][i] == 'a' {
		a := search('a', m)
		return append(n, a)
	}
	return n
}

func cost2(from, to utils.Pos, m utils.IntMatrix) int {
	if m[from.Y][from.X] == 'a' && m[to.Y][to.X] == 'a' {
		return 0
	}
	return 1
}

func heuristic2(from, to utils.Pos, m utils.IntMatrix) int {
	return m[to.Y][to.X] - m[from.Y][from.X]
}

//func manhattanDistance(from, to utils.Pos) int {
//	absX := from.X - to.X
//	if absX < 0 {
//		absX = -absX
//	}
//	absY := from.Y - to.Y
//	if absY < 0 {
//		absY = -absY
//	}
//	return absX + absY
//}

/*
type Pos struct{ i, j int }
type node struct {
	Pos
	priority int
	index    int
}

type heuristicFunction func(from, to Pos, m matrix) int
type costFunction func(from, to Pos, m matrix) int
type neighborsFunction func(m matrix, i, j int) []Pos

func path(start, to Pos, m matrix, neighbors neighborsFunction, cost costFunction, heuristic heuristicFunction) (path []Pos, distance int) {
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

		for _, neighbor := range neighbors(m, current.i, current.j) {
			newCost := costSoFar[current] + cost(current, neighbor, m)
			if _, ok := costSoFar[neighbor]; !ok || newCost < costSoFar[neighbor] {
				costSoFar[neighbor] = newCost
				priority := newCost + heuristic(neighbor, to, m)
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
*/
