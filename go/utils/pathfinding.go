package utils

import (
	"github.com/pemoreau/advent-of-code/go/utils/priorityqueue"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"math"
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

type PairPathCost[T comparable] struct {
	Path []T
	Cost int
}

func Dijkstra[T comparable](starts T, goal goalFunction[T], neighbors neighborsFunction[T], cost costFunction[T]) ([]T, int) {
	pairs := DijkstraMultipleStart([]T{starts}, goal, neighbors, cost, true)
	if len(pairs) > 0 {
		return pairs[0].Path, pairs[0].Cost
	}
	return nil, 0
}

func DijkstraAll[T comparable](starts T, goal goalFunction[T], neighbors neighborsFunction[T], cost costFunction[T]) []PairPathCost[T] {
	return DijkstraMultipleStart([]T{starts}, goal, neighbors, cost, false)
}

func DijkstraMultipleStart[T comparable](starts []T, goal goalFunction[T], neighbors neighborsFunction[T], cost costFunction[T], stopAfterFirst bool) []PairPathCost[T] {
	type Item struct {
		State T
		Cost  int
	}
	var res []PairPathCost[T]
	itemCostCmp := func(a, b Item) int {
		// to implement a Min PriorityQueue
		return b.Cost - a.Cost
	}
	frontier := priorityqueue.New(itemCostCmp)
	cameFrom := map[T][]T{}
	costSoFar := map[T]int{}
	var best = math.MaxInt
	var goalStates []T
	var goalFound = false
	var visited = set.NewSet[T]()

	for _, start := range starts {
		frontier.Insert(Item{State: start, Cost: 0})
		cameFrom[start] = append(cameFrom[start], start)
		costSoFar[start] = 0
	}
	for frontier.Len() > 0 {
		item := frontier.PopMax()
		current := item.State
		if visited.Contains(current) {
			continue
		}
		if item.Cost > best {
			continue
		}
		visited.Add(current)

		if goal(current) && costSoFar[current] < best {
			goalStates = append(goalStates, current)
			goalFound = true
			best = costSoFar[current]
			if stopAfterFirst {
				break
			}
		}

		for _, neighbor := range neighbors(current) {
			newCost := costSoFar[current] + cost(current, neighbor)
			if _, ok := costSoFar[neighbor]; !ok || newCost <= costSoFar[neighbor] {
				if newCost < costSoFar[neighbor] {
					cameFrom[neighbor] = []T{current}
				} else {
					cameFrom[neighbor] = append(cameFrom[neighbor], current)
				}
				priority := newCost
				frontier.Insert(Item{State: neighbor, Cost: priority})
				costSoFar[neighbor] = newCost
			}
		}
	}

	if goalFound {
		// Found a path to the goal.
		var paths = buildPathList[T](goalStates, cameFrom, starts)
		for _, path := range paths {
			res = append(res, PairPathCost[T]{Path: path, Cost: costSoFar[goalStates[0]]})
		}
	}
	return res
}

func buildPathList[T comparable](current []T, cameFrom map[T][]T, starts []T) [][]T {
	var res [][]T
	for _, c := range current {
		for _, path := range buildPath[T](c, cameFrom, starts) {
			res = append(res, path)
		}
	}
	return res
}

func buildPath[T comparable](current T, cameFrom map[T][]T, starts []T) [][]T {
	var res [][]T
	if slices.Contains(starts, current) {
		res = append(res, []T{current})
		return res
	}
	for _, prev := range cameFrom[current] {
		for _, path := range buildPath[T](prev, cameFrom, starts) {
			res = append(res, append(path, current))
		}
	}
	return res
}
