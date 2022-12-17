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
	name int
	rate int
	dest []int
	cost map[int]int
}

type Network map[int]*Valve

func neighbors(network Network, s State) []State {
	res := []State{}
	for name := range network {
		candidate := network[name]
		cost := network[s.name].cost[name]
		if candidate.rate > 0 && !s.path[candidate.name] && s.time+cost <= 30 {
			newPath := s.path
			newPath[candidate.name] = true
			res = append(res, State{
				name: candidate.name,
				time: s.time + cost,
				path: newPath,
				prod: s.prod + (30-(s.time+cost))*candidate.rate,
			})

		}
	}
	return res
}

func parse(input string) (Network, int) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	valves := make(map[int]*Valve)

	cpt := 0
	toint := map[string]int{}
	for _, line := range lines {
		values := strings.Split(line, " ")
		toint[values[1]] = cpt
		cpt++
	}

	for _, line := range lines {
		values := strings.Split(line, " ")
		name := values[1]
		var rate int
		fmt.Sscanf(values[4], "rate=%d;", &rate)
		dest := []int{}
		for i := 9; i < len(values); i++ {
			dest = append(dest, toint[strings.TrimPrefix(strings.TrimSuffix(values[i], ","), " ")])
		}
		valves[toint[name]] = &Valve{
			name: toint[name],
			rate: rate,
			dest: dest,
		}
	}
	for k := range valves {
		valves[k].cost = make(map[int]int)
		for d := range valves {
			if valves[d].rate > 0 {
				valves[k].cost[d] = len(distance(k, d, valves, []int{}))
			}
		}
		fmt.Println(k, valves[k].cost)
	}
	return valves, toint["AA"]
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
	//fmt.Println("start", start, "end", end, "path", path)
	path = append(path, start)
	//fmt.Println("path", path)
	if start == end {
		return path
	}
	if _, ok := network[start]; !ok {
		//fmt.Println("not found", start)
		return []int{}
	}
	shortest := []int{}
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
	name int
	time int
	path [128]bool
	prod int
}

func Part1(input string) int {
	valves, name := parse(input)
	start := State{
		name: name,
		time: 0,
		prod: 0,
		path: [128]bool{},
	}

	res := findMaxProduction(valves, start)
	return res
}

var keys []int
var space map[State3]int

func Part2(input string) int {
	valves, name := parse(input)
	start := State2{
		name1: name,
		name2: name,
		time1: 0,
		time2: 0,
		prod:  0,
		path:  [128]bool{},
	}
	for k := range valves {
		keys = append(keys, k)
	}
	space = make(map[State3]int)
	sort.Ints(keys)
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
	name1 int
	name2 int
	time1 int
	time2 int
	path  [128]bool
	prod  int
}

func (s State2) String() string {
	bits := 0
	for i := 0; i < len(s.path); i++ {
		if s.path[i] {
			bits++
		}
	}
	return fmt.Sprintf("%d %d t1=%d t2=%d [%d] prod=%d", s.name1, s.name2, s.time1, s.time2, bits, s.prod)
}
func (s State3) String() string {
	return fmt.Sprintf("%d %d t1=%d t2=%d", s.name1, s.name2, s.time1, s.time2)
}

type State3 struct {
	name1 int
	name2 int
	time1 int
	time2 int
	path  [128]bool
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
			notVisited := !s.path[candidate1.name] && !s.path[candidate2.name]
			if activeCandidates && notVisited && s.time1+cost1 <= 26 && s.time2+cost2 <= 26 {
				prod1 := (26 - (s.time1 + cost1)) * candidate1.rate
				prod2 := (26 - (s.time2 + cost2)) * candidate2.rate
				newPath := s.path
				newPath[candidate1.name] = true
				newPath[candidate2.name] = true
				newProd := s.prod + prod1 + prod2
				s3 := State3{
					name1: candidate1.name,
					name2: candidate2.name,
					time1: s.time1 + cost1,
					time2: s.time2 + cost2,
					path:  newPath,
				}
				if oldProd, ok := space[s3]; ok && newProd <= oldProd {
					// do nothing
					//fmt.Println("skip", s3, newProd, oldProd)
				} else {
					space[s3] = newProd
					res = append(res, State2{
						name1: candidate1.name,
						name2: candidate2.name,
						time1: s.time1 + cost1,
						time2: s.time2 + cost2,
						path:  newPath,
						prod:  newProd,
					})
				}

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

		//sort.Strings(current.path)
		//s3 := State3{
		//	name1: current.name1,
		//	name2: current.name2,
		//	time1: current.time1,
		//	time2: current.time2,
		//	path:  strings.Join(current.path, ""),
		//}
		//if _, ok := space[s3]; ok {
		//	//fmt.Println("already visited", s3, n)
		//	continue
		//} else {
		//	space[s3]++
		//}

		//fmt.Println("current", current)
		maxProd = utils.Max(maxProd, current.prod)
		n := neighbors2(network, current)
		for _, s := range n {
			queue = append(queue, s)
		}
		cpt++
		if cpt%10000 == 0 {
			fmt.Println("cpt", cpt, "len", len(queue))
			fmt.Println("current", current)
		}
	}
	return maxProd
}
