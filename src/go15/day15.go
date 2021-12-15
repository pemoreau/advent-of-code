package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var input_test string

func Part1(input string) int {
	// lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	// m := BuildMatrix(lines)
	// _, cost, _ := path(m[0][0], m[len(m)-1][len(m[0])-1], m)
	// return int(cost)
	return 0
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := BuildMatrix(lines)
	mm := buildMegaMatrix(m)
	from := mm[0][0]
	to := mm[len(mm)-1][len(mm[0])-1]
	fmt.Printf("from=%v, to=%v\n", from, to)
	_, cost, _ := path(from, to, mm)
	// fmt.Println(p, cost, b)
	return int(cost)
}

func main() {
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_test)))
	fmt.Println(time.Since(start))
}

type matrix [][]*node

func BuildMatrix(lines []string) matrix {
	m := make([][]*node, len(lines))
	for j, l := range lines {
		l = strings.TrimSpace(l)
		m[j] = make([]*node, len(l))
		for i, c := range l {
			m[j][i] = &node{
				i:    i,
				j:    j,
				risk: uint8(c - '0'),
			}
		}
	}
	return m
}

func buildMegaMatrix(m matrix) matrix {
	mm := make([][]*node, 5*len(m))
	// fmt.Println("mm", len(mm))
	for j, l := range m {
		for kj := 0; kj < 5; kj++ {
			new_j := kj*len(m) + j
			mm[new_j] = make([]*node, 5*len(l))
			for i, n := range l {
				risk := n.risk
				// fmt.Println("mmj", len(mm[new_j]))
				for ki := 0; ki < 5; ki++ {
					new_i := ki*len(l) + i
					new_risk := risk + uint8(ki) + uint8(kj)
					if new_risk > 9 {
						new_risk = 1
					}
					// fmt.Printf("(%d, %d) -> (%d, %d) risk=%d -> %d\n", i, j, new_i, new_j, risk, new_risk)
					// fmt.Printf("len(mm)=%d, len(mm[new_j])=%d\n", len(mm), len(mm[new_j]))
					mm[new_j][new_i] = &node{
						i:    new_i,
						j:    new_j,
						risk: new_risk,
					}
					// fmt.Println(mm[new_j][new_i])
				}
			}
		}
	}
	// for j := 0; j < len(mm); j++ {
	// 	for i := 0; i < len(mm[j]); i++ {
	// 		fmt.Printf("%d, %d: %v\n", i, j, mm[j][i])
	// 	}
	// }
	return mm
}

type node struct {
	i, j   int
	risk   uint8
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
	index  int
}

func (n *node) String() string {
	return fmt.Sprintf("(%d, %d) risk=%d", n.i, n.j, n.risk)
}

type pos struct{ i, j int }

func neighboors(m matrix, i, j int) []*node {
	res := []*node{}
	pos := []pos{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	for _, p := range pos {
		if p.i >= 0 && p.i < len(m) && p.j >= 0 && p.j < len(m[p.i]) {
			res = append(res, m[p.j][p.i])
		}
	}
	// fmt.Printf("neighboors(%d, %d) = %v\n", i, j, res)
	return res
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func manhattanDistance(from, to *node) float64 {
	absX := from.i - to.i
	if absX < 0 {
		absX = -absX
	}
	absY := from.j - to.j
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func path(from, to *node, m matrix) (path []node, distance float64, found bool) {
	nq := &PriorityQueue{}
	heap.Init(nq)
	from.open = true
	heap.Push(nq, from)

	for {
		if nq.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := heap.Pop(nq).(*node) // https://go.dev/ref/spec#Type_assertions
		current.open = false
		current.closed = true

		if current == to {
			// Found a path to the goal.
			p := []node{}
			curr := current
			for curr != nil {
				p = append(p, *curr)
				curr = curr.parent
			}
			return p, current.cost, true
		}

		for _, neighbor := range neighboors(m, current.i, current.j) {
			// fmt.Println("neighbor", neighbor.String())
			cost := current.cost + float64(neighbor.risk)
			if cost < neighbor.cost {
				if neighbor.open {
					heap.Remove(nq, neighbor.index)
				}
				neighbor.open = false
				neighbor.closed = false
			}
			if !neighbor.open && !neighbor.closed {
				neighbor.cost = cost
				neighbor.open = true
				neighbor.rank = cost + manhattanDistance(neighbor, to)
				neighbor.parent = current
				heap.Push(nq, neighbor)
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
	return pq[i].rank < pq[j].rank
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
