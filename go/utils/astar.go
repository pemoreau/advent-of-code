package utils

import (
	"github.com/oleiade/lane/v2"
)

// import "container/heap"
//type node[T comparable] struct {
//	state    T
//	priority int
//	index    int
//}

type heuristicFunction[T comparable] func(from T) int
type goalFunction[T comparable] func(from T) bool
type costFunction[T comparable] func(from, to T) int
type neighborsFunction[T comparable] func(from T) []T

func Path[T comparable](start T, goal goalFunction[T], neighbors neighborsFunction[T], cost costFunction[T], heuristic heuristicFunction[T]) (path []T, distance int) {
	frontier := lane.NewMaxPriorityQueue[T, int]()
	frontier.Push(start, 0)

	cameFrom := map[T]T{start: start}
	costSoFar := map[T]int{start: 0}

	for {
		if frontier.Size() == 0 {
			// There's no path, return found false.
			return
		}
		current, _, _ := frontier.Pop()
		if goal(current) {
			// Found a path to the goal.
			path := []T{}
			curr := current
			for curr != start {
				path = append(path, curr)
				curr = cameFrom[curr]
			}
			return path, costSoFar[current]
		}

		for _, neighbor := range neighbors(current) {
			newCost := costSoFar[current] + cost(current, neighbor)
			if _, ok := costSoFar[neighbor]; !ok || newCost < costSoFar[neighbor] {
				costSoFar[neighbor] = newCost
				priority := newCost + heuristic(neighbor)
				frontier.Push(neighbor, priority)
				cameFrom[neighbor] = current
			}
		}
	}
}

//
//// A PriorityQueue implements heap.Interface and holds Items.
//// Code copied from https://pkg.go.dev/container/heap@go1.17.5
//type PriorityQueue []*node
//
//func (pq PriorityQueue) Len() int {
//	return len(pq)
//}
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	return pq[i].priority < pq[j].priority
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//	pq[i].index = i
//	pq[j].index = j
//}
//
//func (pq *PriorityQueue) Push(x interface{}) {
//	n := len(*pq)
//	item := x.(*node)
//	item.index = n
//	*pq = append(*pq, item)
//}
//
//func (pq *PriorityQueue) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	item := old[n-1]
//	old[n-1] = nil  // avoid memory leak
//	item.index = -1 // for safety
//	*pq = old[0 : n-1]
//	return item
//}
