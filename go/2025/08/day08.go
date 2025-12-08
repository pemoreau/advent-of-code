package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
)

type cube struct {
	x, y, z int
}

type dist struct {
	c1 cube
	c2 cube
	d  int
}

type clique = set.Set[cube]

func parse(input string) []dist {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var cubes []cube
	for _, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		cubes = append(cubes, cube{x, y, z})
	}
	var dists []dist
	var stored set.Set[cube]
	for i := 0; i < len(cubes)-1; i++ {
		for j := i + 1; j < len(cubes); j++ {
			if stored.Contains(cubes[i]) && !stored.Contains(cubes[j]) {
				continue
			}
			var a = cubes[j].x - cubes[i].x
			var b = cubes[j].y - cubes[i].y
			var c = cubes[j].z - cubes[i].z
			var d = a*a + b*b + c*c
			dists = append(dists, dist{cubes[i], cubes[j], d})
		}
	}
	return dists
}

func seachClique(c cube, cliques []clique) (int, bool) {
	for i, clique := range cliques {
		if clique.Contains(c) {
			return i, true
		}
	}
	return 0, false
}

func solve(dists []dist, n int, part1 bool) int {
	var cliques []set.Set[cube]
	var cpt int
	for _, d := range dists {
		if part1 && cpt >= n {
			slices.SortFunc(cliques, func(a, b clique) int { return cmp.Compare(len(b), len(a)) })
			return len(cliques[0]) * len(cliques[1]) * len(cliques[2])
		}
		cpt++

		i1, ok1 := seachClique(d.c1, cliques)
		i2, ok2 := seachClique(d.c2, cliques)
		if !ok1 && !ok2 {
			var s = set.NewSet[cube]()
			s.Add(d.c1)
			s.Add(d.c2)
			cliques = append(cliques, s)
		} else if ok1 && ok2 && i1 == i2 {
			// noting happens
		} else if ok1 && ok2 && i1 != i2 {
			// merge 2 cliques
			cliques[i1].Add(d.c2)
			for cube := range cliques[i2] {
				cliques[i1].Add(cube)
			}
			cliques = append(cliques[:i2], cliques[i2+1:]...)
		} else if ok1 {
			cliques[i1].Add(d.c2)
		} else if ok2 {
			cliques[i2].Add(d.c1)
		}
		if !part1 && len(cliques) == 1 && len(cliques[0]) == n {
			return d.c1.x * d.c2.x
		}
	}
	return 0
}

func Part1(input string) int {
	var dists = parse(input)
	slices.SortFunc(dists, func(a, b dist) int { return cmp.Compare(a.d, b.d) })
	return solve(dists, 1000, true)
}

func Part2(input string) int {
	var dists = parse(input)
	slices.SortFunc(dists, func(a, b dist) int { return cmp.Compare(a.d, b.d) })
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	return solve(dists, len(lines), false)
}

func main() {
	fmt.Println("--2025 day 08 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
