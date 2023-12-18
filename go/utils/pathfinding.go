package utils

import (
	"github.com/oleiade/lane/v2"
	"slices"
)

type heuristicFunction[T comparable] func(from T) int
type goalFunction[T comparable] func(from T) bool
type costFunction[T comparable] func(from, to T) int
type neighborsFunction[T comparable] func(from T) []T

func Astar[T comparable](start T, goal goalFunction[T], neighbors neighborsFunction[T], cost costFunction[T], heuristic heuristicFunction[T]) (path []T, distance int) {
	return AstarMultipleStart([]T{start}, goal, neighbors, cost, heuristic)
}

func AstarMultipleStart[T comparable](starts []T, goal goalFunction[T], neighbors neighborsFunction[T], cost costFunction[T], heuristic heuristicFunction[T]) (path []T, distance int) {
	frontier := lane.NewMinPriorityQueue[T, int]()
	cameFrom := map[T]T{}
	costSoFar := map[T]int{}

	for _, start := range starts {
		frontier.Push(start, 0)
		cameFrom[start] = start
		costSoFar[start] = 0
	}

	for frontier.Size() > 0 {
		current, _, _ := frontier.Pop()
		//fmt.Println("current", current, "priority", priority)
		if goal(current) {
			// Found a path to the goal.
			var path []T
			curr := current
			for !slices.Contains(starts, curr) {
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
	return
}
