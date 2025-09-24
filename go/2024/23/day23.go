package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"slices"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

// Bron–Kerbosch
func BronKerbosch(R, P, X set.Set[string], graph map[string]set.Set[string], C *[][]string) {
	if P.Len() == 0 && X.Len() == 0 {
		if R.Len() > 2 {
			var Rlist []string
			for v := range R.All() {
				Rlist = append(Rlist, v)
			}
			slices.Sort(Rlist)
			*C = append(*C, Rlist)
		}
		return
	}
	for v := range P.All() {
		var vset = set.NewSet[string]()
		vset.Add(v)
		BronKerbosch(R.Union(vset), P.Intersect(graph[v]), X.Intersect(graph[v]), graph, C)
		P.Remove(v)
		X.Add(v)
	}
}

func Part1(input string) int {
	var lines = strings.Split(input, "\n")

	var graph = make(map[string][]string)
	var pairs = set.NewSet[[2]string]()

	for _, line := range lines {
		var parts = strings.Split(line, "-")
		graph[parts[0]] = append(graph[parts[0]], parts[1])
		graph[parts[1]] = append(graph[parts[1]], parts[0])
		pairs.Add([2]string{parts[0], parts[1]})
		pairs.Add([2]string{parts[1], parts[0]})
	}

	var cliques = set.NewSet[string]()

	// node iterator
	for node, successors := range graph {
		for i, a := range successors {
			for _, b := range successors[i+1:] {
				// check if a, b are connected
				if pairs.Contains([2]string{a, b}) {
					if node[0] == 't' || a[0] == 't' || b[0] == 't' {
						var t = []string{node, a, b}
						slices.Sort(t)
						cliques.Add(strings.Join(t, ","))
					}
				}
			}
		}
	}
	return cliques.Len()
}

func Part2(input string) string {
	var lines = strings.Split(input, "\n")

	var graph = make(map[string]set.Set[string])
	var V = set.NewSet[string]()
	for _, line := range lines {
		var parts = strings.Split(line, "-")
		if _, ok := graph[parts[0]]; !ok {
			graph[parts[0]] = set.NewSet[string]()
		}
		if _, ok := graph[parts[1]]; !ok {
			graph[parts[1]] = set.NewSet[string]()
		}
		graph[parts[0]].Add(parts[1])
		graph[parts[1]].Add(parts[0])
		V.Add(parts[0])
		V.Add(parts[1])
	}

	var C = &[][]string{}
	BronKerbosch(set.NewSet[string](), V, set.NewSet[string](), graph, C)

	var res string
	var n = 0
	for _, c := range *C {
		if len(c) > n {
			n = len(c)
			res = strings.Join(c, ",")
		}
	}
	return res
}

func main() {
	fmt.Println("--2024 day 23 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
