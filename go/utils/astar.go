package utils

import "container/heap"

type node struct {
	Pos
	priority int
	index    int
}

type heuristicFunction func(from, to Pos, m IntMatrix) int
type costFunction func(from, to Pos, m IntMatrix) int
type neighborsFunction func(m IntMatrix, i, j int) []Pos

func Path(start, to Pos, m IntMatrix, neighbors neighborsFunction, cost costFunction, heuristic heuristicFunction) (path []Pos, distance int) {
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

		for _, neighbor := range neighbors(m, current.X, current.Y) {
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
