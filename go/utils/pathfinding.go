package utils

import (
	"github.com/pemoreau/advent-of-code/go/utils/priorityqueue"
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
	type Item struct {
		State T
		Cost  int
	}

	itemCostCmp := func(a, b Item) int {
		// to implement a Min PriorityQueue
		return b.Cost - a.Cost
	}

	//frontier := lane.NewMinPriorityQueue[T, int]()
	frontier := priorityqueue.New(itemCostCmp)
	cameFrom := map[T]T{}
	costSoFar := map[T]int{}

	for _, start := range starts {
		//frontier.Push(start, 0)
		frontier.Insert(Item{State: start, Cost: 0})
		cameFrom[start] = start
		costSoFar[start] = 0
	}

	//for frontier.Size() > 0 {
	for frontier.Len() > 0 {
		//current, _, _ := frontier.Pop()
		item := frontier.PopMax()
		current := item.State
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
				//frontier.Push(neighbor, priority)
				frontier.Insert(Item{State: neighbor, Cost: priority})
				cameFrom[neighbor] = current
			}
		}
	}
	return
}
