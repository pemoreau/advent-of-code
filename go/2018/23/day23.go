package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"math"
	"slices"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type pos struct {
	x, y, z int
}

type bot struct {
	pos
	r int
}

func (b *bot) String() string {
	return fmt.Sprintf("<%d,%d,%d>, r=%d", b.x, b.y, b.z, b.r)
}

func parseInput(input string) []bot {
	input = strings.Trim(input, "\n")
	var lines = strings.Split(input, "\n")

	var bots []bot
	var b bot
	for _, line := range lines {
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &b.x, &b.y, &b.z, &b.r)
		bots = append(bots, b)
	}
	return bots
}

func distance(a, b pos) int {
	return utils.Abs(a.x-b.x) + utils.Abs(a.y-b.y) + utils.Abs(a.z-b.z)
}

func Part1(input string) int {
	var bots = parseInput(input)
	var selected = slices.MaxFunc(bots, func(a, b bot) int { return a.r - b.r })
	var res int
	for _, b := range bots {
		if distance(selected.pos, b.pos) <= selected.r {
			res++
		}
	}

	return res
}

func tangent(a, b bot) (pos, bool) {
	var d = distance(a.pos, b.pos)
	if d == a.r+b.r {
		//fmt.Println("outer tangency")
		var x = a.x + (b.x-a.x)*a.r/d
		var y = a.y + (b.y-a.y)*a.r/d
		var z = a.z + (b.z-a.z)*a.r/d
		return pos{x, y, z}, true
	}
	if d == utils.Abs(a.r-b.r) {
		//fmt.Println("inner tangency")
		var x = a.x + (b.x-a.x)*a.r/d
		var y = a.y + (b.y-a.y)*a.r/d
		var z = a.z + (b.z-a.z)*a.r/d
		return pos{x, y, z}, true
	}

	return pos{}, false
}

func neighborsn(p pos, n int) []pos {
	var res []pos
	for i := -n; i <= n; i++ {
		for j := -n; j <= n; j++ {
			for k := -n; k <= n; k++ {
				res = append(res, pos{p.x + i, p.y + j, p.z + k})
			}
		}
	}
	return res
}

func Part2(input string) int {
	var bots = parseInput(input)

	var maxi = 0
	var minDist = math.MaxInt
	var best pos
	for i, b1 := range bots {
		for j, b2 := range bots {
			if i == j {
				continue
			}
			if t, ok := tangent(b1, b2); ok {
				for _, p := range neighborsn(t, 1) {
					if nb := nbInRange(p, bots); nb >= maxi {
						d := utils.Abs(p.x) + utils.Abs(p.y) + utils.Abs(p.z)
						if nb > maxi {
							fmt.Println("point", p, "nbInRange", nbInRange(p, bots))
							maxi = nbInRange(p, bots)
							minDist = d
							best = p
						} else if d < minDist {
							minDist = d
							best = p
						}
					}
				}

			}
		}
	}

	fmt.Println("best", best, "dist", minDist)

	var visited = set.NewSet[pos]()
	var todo = []pos{best}
	for _, b := range bots {
		//var d = distance(b.pos, pos{})
		// p is on line 0--b.x at distance b.r from b
		//var p = pos{b.x * (d - b.r), b.y * (d - b.r), b.z * (d - b.r)}
		todo = append(todo, b.pos)
	}

	for len(todo) > 0 {
		//fmt.Println("todo", len(todo))
		var p = todo[0]
		todo = todo[1:]
		if visited.Contains(p) {
			continue
		}
		visited.Add(p)
		for _, n := range neighborsn(p, 1) {
			if d := distance(n, pos{}); nbInRange(n, bots) >= maxi && d < minDist {
				todo = append(todo, n)
				if nbInRange(n, bots) > maxi {
					fmt.Println("point", n, "nbInRange", nbInRange(n, bots))
					maxi = nbInRange(n, bots)
					minDist = d
					best = n
				} else if d < minDist {
					minDist = d
					best = n
				}
			}
		}
	}

	fmt.Println("best2", best, "dist", minDist)

	return minDist
}

func nbInRange(p pos, bots []bot) int {
	var res int
	for _, b := range bots {
		if utils.Abs(b.x-p.x)+utils.Abs(b.y-p.y)+utils.Abs(b.z-p.z) <= b.r {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println("--2018 day 23 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
