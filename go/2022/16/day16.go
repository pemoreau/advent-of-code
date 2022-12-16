package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

type Valve struct {
	name string
	rate int
	dest []string
	cost map[string]int
}

type State struct {
	name string
	time int
	path string
}

type Network map[string]*Valve

func contains(path string, name string) bool {
	for i := 0; i < len(path); i = i + 2 {
		if path[i:i+2] == name {
			return true
		}
	}
	return false
}

func neighbors(network Network, s State) []State {
	res := []State{}
	for name := range network {
		candidate := network[name]
		cost := network[s.name].cost[name]
		if candidate.rate > 0 && !contains(s.path, candidate.name) && s.time+cost <= 30 {
			res = append(res, State{
				name: candidate.name,
				time: s.time + cost,
				path: s.path + candidate.name,
			})

		}
	}
	return res
}

func parse(input string) Network {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	valves := make(map[string]*Valve)
	for _, line := range lines {
		values := strings.Split(line, " ")
		name := values[1]
		var rate int
		fmt.Sscanf(values[4], "rate=%d;", &rate)
		dest := []string{}
		for i := 9; i < len(values); i++ {
			dest = append(dest, strings.TrimPrefix(strings.TrimSuffix(values[i], ","), " "))
		}
		valves[name] = &Valve{
			name: name,
			rate: rate,
			dest: dest,
		}
	}
	for k := range valves {
		valves[k].cost = make(map[string]int)
		for d := range valves {
			if valves[d].rate > 0 {
				valves[k].cost[d] = len(distance(k, d, valves, []string{}))
			}
		}
		fmt.Println(k, valves[k].cost)
	}
	return valves
}

func listContains(list []string, name string) bool {
	for _, n := range list {
		if n == name {
			return true
		}
	}
	return false
}

func distance(start, end string, network Network, path []string) []string {
	//fmt.Println("start", start, "end", end, "path", path)
	path = append(path, start)
	//fmt.Println("path", path)
	if start == end {
		return path
	}
	if _, ok := network[start]; !ok {
		//fmt.Println("not found", start)
		return []string{}
	}
	shortest := []string{}
	for _, n := range network[start].dest {
		//fmt.Println("n", n)
		if !listContains(path, n) {
			newpath := distance(n, end, network, path)
			if len(newpath) > 0 {
				if len(shortest) == 0 || len(newpath) < len(shortest) {
					shortest = newpath
				}
			}
		}
	}
	return shortest
}

func Part1(input string) int {
	valves := parse(input)
	start := State{
		name: "AA",
		time: 0,
		path: "",
	}
	p, _ := path(start, valves)
	fmt.Println(p)
	res := 0
	for _, n := range p {
		res += (30 - n.time) * valves[n.name].rate
	}
	return res
	// 850 too low
}

func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	return 0

}

func main() {
	fmt.Println("--2022 day 16 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}

type node struct {
	State
	priority int
	index    int
}

func path(start State, network Network) (path []State, distance int) {
	frontier := &PriorityQueue{}
	heap.Init(frontier)
	heap.Push(frontier, &node{State: start, priority: 0})

	cameFrom := map[State]State{start: start}
	costSoFar := map[State]int{start: 0}

	for {
		if frontier.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := heap.Pop(frontier).(*node).State
		n := neighbors(network, current)
		if len(n) == 0 {
			// Found a path to the goal.
			path := []State{}
			curr := current
			for curr != start {
				path = append(path, curr)
				curr = cameFrom[curr]
			}
			return path, costSoFar[current]
		}

		for _, neighbor := range n {
			fmt.Println("neighbor", neighbor)
			prod := (30 - current.time) * network[current.name].rate
			newCost := costSoFar[current] + (100000 - prod)
			if _, ok := costSoFar[neighbor]; !ok || newCost < costSoFar[neighbor] {
				costSoFar[neighbor] = newCost
				priority := newCost //+ manhattanDistance(neighbor, to)
				heap.Push(frontier, &node{State: neighbor, priority: priority})
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
