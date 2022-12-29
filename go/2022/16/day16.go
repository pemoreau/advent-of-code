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

const NBVALVES = 16

type Valve struct {
	name int
	rate int
	dest []int
	cost map[int]int
}

type Network map[int]*Valve

func neighbors(network Network, s State, maxTime int) []State {
	res := []State{}
	for name := range network {
		candidate := network[name]
		cost := network[s.name].cost[name]
		if candidate.rate > 0 && !s.path[candidate.name] && s.time+cost <= maxTime {
			newPath := s.path
			newPath[candidate.name] = true
			res = append(res, State{
				name: candidate.name,
				time: s.time + cost,
				path: newPath,
				prod: s.prod + (maxTime-(s.time+cost))*candidate.rate,
			})

		}
	}
	return res
}

func parse(input string) (Network, Network, int) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	valves := make(map[int]*Valve)

	toint := map[string]int{"AA": 0}
	cpt := 1
	for _, line := range lines {
		values := strings.Split(line, " ")
		name := values[1]
		var rate int
		fmt.Sscanf(values[4], "rate=%d;", &rate)
		if rate > 0 {
			toint[name] = cpt
			cpt++
		}
	}
	for _, line := range lines {
		values := strings.Split(line, " ")
		name := values[1]
		var rate int
		fmt.Sscanf(values[4], "rate=%d;", &rate)
		if _, ok := toint[name]; !ok && rate == 0 {
			toint[name] = cpt
			cpt++
		}
	}

	for _, line := range lines {
		values := strings.Split(line, " ")
		name := toint[values[1]]
		var rate int
		fmt.Sscanf(values[4], "rate=%d;", &rate)
		dest := []int{}
		for i := 9; i < len(values); i++ {
			dest = append(dest, toint[strings.TrimPrefix(strings.TrimSuffix(values[i], ","), " ")])
		}
		valves[name] = &Valve{
			name: name,
			rate: rate,
			dest: dest,
		}
	}

	activesValves := make(map[int]*Valve)
	for _, v := range valves {
		name := v.name
		if v.rate > 0 || v.name == 0 {
			activeValve := Valve{name: v.name, rate: v.rate, cost: make(map[int]int)}
			for d := range valves {
				if d != name && valves[d].rate > 0 {
					activeValve.dest = append(activeValve.dest, d)
					activeValve.cost[d] = len(distance(name, d, valves, []int{}))
				}
			}
			//fmt.Println(name, activeValve.cost)
			activesValves[name] = &activeValve
		}
	}
	return valves, activesValves, toint["AA"]
}

func listContains(list []int, name int) bool {
	for _, n := range list {
		if n == name {
			return true
		}
	}
	return false
}

func distance(start, end int, network Network, path []int) []int {
	path = append(path, start)
	if start == end {
		return path
	}
	if _, ok := network[start]; !ok {
		//fmt.Println("not found", start)
		return []int{}
	}
	shortest := []int{}
	for _, n := range network[start].dest {
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
	name int
	time int
	path [NBVALVES]bool
	prod int
}

func allPath(network Network, start State, maxTime int) map[[NBVALVES]bool]int {
	res := make(map[[NBVALVES]bool]int)
	queue := []State{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if max, ok := res[current.path]; !ok || max < current.prod {
			res[current.path] = current.prod
		}
		n := neighbors(network, current, maxTime)
		for _, s := range n {
			queue = append(queue, s)
		}
	}
	return res
}

func Part1(input string) int {
	_, actives, name := parse(input)
	start := State{
		name: name,
		time: 0,
		prod: 0,
		path: [NBVALVES]bool{},
	}

	paths := allPath(actives, start, 30)

	max := math.MinInt
	for _, p := range paths {
		max = utils.Max(max, p)
	}
	return max
}

func Part2(input string) int {
	_, actives, name := parse(input)

	res := math.MinInt

	start := State{
		name: name,
		time: 0,
		prod: 0,
		path: [NBVALVES]bool{},
	}

	paths := allPath(actives, start, 26)
	for p, max1 := range paths {
		for q, max2 := range paths {
			intersect := false
			for i := 0; i < len(actives); i++ {
				if p[i] && q[i] {
					intersect = true
					break
				}
			}
			if !intersect {
				res = utils.Max(res, max1+max2)
			}
		}
	}

	return res
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
