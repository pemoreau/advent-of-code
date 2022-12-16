package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
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
				prod: s.prod + (30-(s.time+cost))*candidate.rate,
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

type State struct {
	name string
	time int
	path string
	prod int
}

func Part1(input string) int {
	valves := parse(input)
	start := State{
		name: "AA",
		time: 0,
		prod: 0,
		path: "",
	}

	res := findMaxProduction(valves, start)
	//p, _ := path(start, valves)
	//fmt.Println(p)
	//res := 0
	//for _, n := range p {
	//	res += (30 - n.time) * valves[n.name].rate
	//}
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

func findMaxProduction(network Network, start State) int {
	queue := []State{start}
	maxProd := math.MinInt
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println("current", current)

		maxProd = utils.Max(maxProd, current.prod)
		n := neighbors(network, current)
		for _, s := range n {
			queue = append(queue, s)
		}
	}
	return maxProd
}

type node struct {
	State
	priority int
	index    int
}

type State2 struct {
	name string
	time int
	path utils.Set[string]
	prod int
}

func neighbors2(network Network, s State) []State {
	res := []State{}
	for name := range network {
		candidate := network[name]
		cost := network[s.name].cost[name]
		if candidate.rate > 0 && !contains(s.path, candidate.name) && s.time+cost <= 30 {
			res = append(res, State{
				name: candidate.name,
				time: s.time + cost,
				path: s.path + candidate.name,
				prod: s.prod + (30-(s.time+cost))*candidate.rate,
			})

		}
	}
	return res
}
