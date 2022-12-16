package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"sort"
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
	return res
}

var keys []string
var space map[State3]int

func Part2(input string) int {
	valves := parse(input)
	start := State2{
		name1: "AA",
		name2: "AA",
		time1: 0,
		time2: 0,
		prod:  0,
		path:  []string{},
	}
	for k := range valves {
		keys = append(keys, k)
	}
	space = make(map[State3]int)
	sort.Strings(keys)
	fmt.Println("keys", keys)
	res := findMaxProduction2(valves, start)
	return res
	//2184 too low
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
		//fmt.Println("current", current)
		maxProd = utils.Max(maxProd, current.prod)
		n := neighbors(network, current)
		for _, s := range n {
			queue = append(queue, s)
		}
	}
	return maxProd
}

type State2 struct {
	name1 string
	name2 string
	time1 int
	time2 int
	path  []string
	prod  int
}

type State3 struct {
	name1 string
	name2 string
	time1 int
	time2 int
	path  string
}

func neighbors2(network Network, s State2) []State2 {
	res := []State2{}
	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(keys); j++ {
			if i == j {
				continue
			}
			name1 := keys[i]
			name2 := keys[j]
			candidate1 := network[name1]
			candidate2 := network[name2]
			cost1 := network[s.name1].cost[name1]
			cost2 := network[s.name2].cost[name2]
			activeCandidates := candidate1.rate > 0 && candidate2.rate > 0
			notVisited := !listContains(s.path, candidate1.name) && !listContains(s.path, candidate2.name)
			if activeCandidates && notVisited && s.time1+cost1 <= 26 && s.time2+cost2 <= 26 {
				prod1 := (26 - (s.time1 + cost1)) * candidate1.rate
				prod2 := (26 - (s.time2 + cost2)) * candidate2.rate
				newPath := make([]string, len(s.path)+2)
				copy(newPath, s.path)
				newPath[len(s.path)] = candidate1.name
				newPath[len(s.path)+1] = candidate2.name
				//newPath := s.path + candidate1.name + candidate2.name
				res = append(res, State2{
					name1: candidate1.name,
					name2: candidate2.name,
					time1: s.time1 + cost1,
					time2: s.time2 + cost2,
					path:  newPath,
					prod:  s.prod + prod1 + prod2,
				})

			}
		}
	}
	return res
}
func findMaxProduction2(network Network, start State2) int {
	queue := []State2{start}
	maxProd := math.MinInt
	cpt := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		sort.Strings(current.path)
		s3 := State3{
			name1: current.name1,
			name2: current.name2,
			time1: current.time1,
			time2: current.time2,
			path:  strings.Join(current.path, ""),
		}
		if _, ok := space[s3]; ok {
			//fmt.Println("already visited", s3, n)
			continue
		} else {
			space[s3]++
		}

		//fmt.Println("current", current)
		maxProd = utils.Max(maxProd, current.prod)
		n := neighbors2(network, current)
		for _, s := range n {
			queue = append(queue, s)
		}
		cpt++
		if cpt%1000 == 0 {
			fmt.Println("cpt", cpt, "len", len(queue))
			fmt.Println("current", current)
		}
	}
	return maxProd
}
