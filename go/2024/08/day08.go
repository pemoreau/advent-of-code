package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func solve(input string, addAntinodes func(ax, ay int, dx, dy int, maxX, maxY int, antinodes set.Set[game2d.Pos])) int {
	var lines = strings.Split(input, "\n")
	var maxX = len(lines[0]) - 1
	var maxY = len(lines) - 1
	//var antenna = make(map[uint8][]game2d.Pos)
	var antenna [256][]game2d.Pos // use an array to see if it is faster than a map
	var antinodes = set.NewSet[game2d.Pos]()
	for j, l := range lines {
		for i, c := range l {
			if c == '.' {
				continue
			}
			antenna[uint8(c)] = append(antenna[uint8(c)], game2d.Pos{X: i, Y: j})
		}
	}

	for _, s := range antenna {
		if len(s) > 0 {
			for i, a1 := range s[:len(s)-1] {
				for _, a2 := range s[i+1:] {
					dx := a2.X - a1.X
					dy := a2.Y - a1.Y
					addAntinodes(a2.X, a2.Y, dx, dy, maxX, maxY, antinodes)
					addAntinodes(a1.X, a1.Y, -dx, -dy, maxX, maxY, antinodes)
				}
			}
		}
	}

	return antinodes.Len()
}

func Part1(input string) int {
	var addAntinodes = func(ax, ay int, dx, dy int, maxX, maxY int, antinodes set.Set[game2d.Pos]) {
		ax += dx
		ay += dy
		if ax >= 0 && ax <= maxX && ay >= 0 && ay <= maxY {
			antinodes.Add(game2d.Pos{X: ax, Y: ay})
		}
	}
	return solve(input, addAntinodes)
}

func Part2(input string) int {
	var addAntinodes = func(ax, ay int, dx, dy int, maxX, maxY int, antinodes set.Set[game2d.Pos]) {
		for ax >= 0 && ax <= maxX && ay >= 0 && ay <= maxY {
			antinodes.Add(game2d.Pos{X: ax, Y: ay})
			ax += dx
			ay += dy
		}
	}
	return solve(input, addAntinodes)
}

func main() {
	fmt.Println("--2024 day 08 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
