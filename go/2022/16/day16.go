package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

const NBVALVES = 16

type Valve struct {
	name int
	rate int
	dest []int
	cost map[int]int
}

func (v *Valve) String() string {
	return fmt.Sprintf("%d(%d) cost:%v", v.name, v.rate, v.cost)
}

type Network map[int]*Valve

func neighbors(network Network, s State, maxTime int) []State {
	var res []State
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
		var dest []int
		for i := 9; i < len(values); i++ {
			dest = append(dest, toint[strings.TrimPrefix(strings.TrimSuffix(values[i], ","), " ")])
		}
		valves[name] = &Valve{
			name: name,
			rate: rate,
			dest: dest,
		}
	}

	// Floyd–Warshall algorithm
	dist := [128][128]int{}

	//fmt.Println("len", len(valves))
	for v := range valves {
		for w := range valves {
			dist[v][w] = math.MaxInt16
		}
	}
	for _, v := range valves {
		for _, d := range v.dest {
			dist[v.name][d] = 1
		}
	}
	for v := range valves {
		dist[v][v] = 0
	}
	for k := range valves {
		for i := range valves {
			for j := range valves {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	activesValves := make(map[int]*Valve)
	for _, v := range valves {
		if v.rate > 0 || v.name == 0 {
			activeValve := Valve{name: v.name, rate: v.rate, cost: make(map[int]int)}
			for d := range valves {
				if d != v.name && valves[d].rate > 0 {
					activeValve.dest = append(activeValve.dest, d)
					activeValve.cost[d] = dist[v.name][d] + 1
				}
			}
			activesValves[v.name] = &activeValve
		}
	}

	return valves, activesValves, toint["AA"]
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

	res := math.MinInt
	for _, p := range paths {
		res = max(res, p)
	}
	return res
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
				res = max(res, max1+max2)
			}
		}
	}

	return res
}

func main() {
	fmt.Println("--2022 day 16 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
