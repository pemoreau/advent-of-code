package main

import (
	_ "embed"
	"fmt"
	"github.com/oleiade/lane/v2"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"slices"
	"time"
)

//go:embed input.txt
var inputDay string

func neighbors1(grid game2d.Grid, s state) []state {
	p := s.pos
	c, _ := grid[p]
	if c == '#' {
		return nil
	}

	var tmp []game2d.Pos
	switch c {
	case '.':
		tmp = append(tmp, p.Neighbors4()...)
	case '>':
		tmp = append(tmp, game2d.Pos{X: p.X + 1, Y: p.Y})
	case '<':
		tmp = append(tmp, game2d.Pos{X: p.X - 1, Y: p.Y})
	case '^':
		tmp = append(tmp, game2d.Pos{X: p.X, Y: p.Y - 1})
	case 'v':
		tmp = append(tmp, game2d.Pos{X: p.X, Y: p.Y + 1})
	}

	var res []state
	for _, n := range tmp {
		if c, ok := grid[n]; ok && c != '#' && !s.path.Contains(n) {
			copyPath := set.NewSet[game2d.Pos]()
			for p := range s.path {
				copyPath.Add(p)
			}
			copyPath.Add(n)
			newState := state{
				pos:  n,
				path: copyPath,
			}
			res = append(res, newState)
		}
	}
	return res
}

type state struct {
	pos  game2d.Pos
	path set.Set[game2d.Pos]
}

func Part11(input string) int {
	grid := game2d.BuildGridFromString(input)

	minX, maxX, minY, maxY := game2d.GridBounds(grid)
	start := game2d.Pos{X: minX + 1, Y: minY}
	end := game2d.Pos{X: maxX - 1, Y: maxY}
	//fmt.Printf("start: %v - %c\n", start, grid[start])
	//fmt.Printf("end:   %v - %c\n", end, grid[end])

	var costSoFar = make(map[game2d.Pos]int)
	costSoFar[start] = 0

	var todo = lane.NewMaxPriorityQueue[state, int]()

	todo.Push(state{pos: start, path: set.NewSet[game2d.Pos]()}, 0)

	for todo.Size() > 0 {
		s, _, _ := todo.Pop()
		p := s.pos

		//fmt.Printf("visit %d, %d", p.Y+1, p.X+1)
		for _, n := range neighbors1(grid, s) {
			newCost := costSoFar[p] + 1
			if _, ok := costSoFar[n.pos]; !ok || newCost > costSoFar[n.pos] {
				costSoFar[n.pos] = newCost
				todo.Push(n, newCost)
				//fmt.Printf(" - Push %d, %d cost:%d\n", n.Y+1, n.X+1, newCost)
			}
		}
	}

	return costSoFar[end]
}

type PosCost struct {
	pos  game2d.Pos
	cost int
}

func exploreSinglePath(grid game2d.Grid, previous game2d.Pos, current game2d.Pos, cost int, part2 bool) (PosCost, bool) {
	if c, ok := grid[current]; ok && c != '#' {
		var cpt int
		for _, ne := range current.Neighbors4() {
			if c, ok := grid[ne]; ok && c != '#' {
				cpt++
			}
		}
		if cpt > 2 {
			return PosCost{pos: current, cost: cost}, true
		}
	}

	if !part2 {
		// cut branches in part1
		if c, ok := grid[current]; ok && c != '.' {
			if current.X > previous.X && c != '>' ||
				current.X < previous.X && c != '<' ||
				current.Y > previous.Y && c != 'v' ||
				current.Y < previous.Y && c != '^' {
				return PosCost{}, false
			}
		}
	}

	for _, n := range current.Neighbors4() {
		if c, ok := grid[n]; ok && c != '#' && n != previous {
			return exploreSinglePath(grid, current, n, cost+1, part2)
		}
	}

	return PosCost{pos: current, cost: cost}, true
}

func explore(neighbors Graph, p, end game2d.Pos, visited map[game2d.Pos]bool, path int, bestPath int) int {
	if p == end {
		if path > bestPath {
			bestPath = path
		}
		return bestPath
	}

	visited[p] = true
	for _, pc := range neighbors[p] {
		if !visited[pc.pos] {
			bestPath = explore(neighbors, pc.pos, end, visited, path+pc.cost, bestPath)
		}
	}
	visited[p] = false
	return bestPath
}

type Graph map[game2d.Pos][]PosCost

func buildGraph(grid game2d.Grid, start game2d.Pos, part2 bool) Graph {
	var res = make(map[game2d.Pos][]PosCost)

	var todo = []game2d.Pos{}
	todo = append(todo, start)

	for len(todo) > 0 {
		p := todo[0]
		todo = todo[1:]
		if c, ok := grid[p]; !ok || c == '#' {
			continue
		}
		for _, n := range p.Neighbors4() {
			if c, ok := grid[n]; !ok || c == '#' {
				continue
			}
			pc, ok := exploreSinglePath(grid, p, n, 1, part2)
			if ok && !slices.Contains(res[p], pc) {
				res[p] = append(res[p], pc)
				todo = append(todo, pc.pos)
			}
		}
	}

	return res
}

func solve(input string, part2 bool) int {
	grid := game2d.BuildGridFromString(input)

	minX, maxX, minY, maxY := game2d.GridBounds(grid)
	start := game2d.Pos{X: minX + 1, Y: minY}
	end := game2d.Pos{X: maxX - 1, Y: maxY}

	neighbors := buildGraph(grid, start, part2)
	//for k, v := range neighbors {
	//	fmt.Printf("%d, %d\n", k.Y+1, k.X+1)
	//	for _, pc := range v {
	//		fmt.Printf("  %d, %d - %d\n", pc.pos.Y+1, pc.pos.X+1, pc.cost)
	//	}
	//}

	visited := make(map[game2d.Pos]bool)
	return explore(neighbors, start, end, visited, 0, 0)
}

func Part1(input string) int {
	return solve(input, false)
}

func Part2(input string) int {
	return solve(input, true)
}

func main() {
	fmt.Println("--2023 day 24 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
