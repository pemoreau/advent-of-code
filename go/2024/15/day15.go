package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

const (
	EMPTY = '.'
	WALL  = '#'
	ROBOT = '@'
	BOX   = 'O'
	LBOX  = '['
	RBOX  = ']'
)

func canMove(grid *game2d.MatrixChar, pos game2d.Pos, dir uint8) (game2d.Pos, bool) {
	var delta = map[uint8]game2d.Pos{'>': {X: 1, Y: 0}, '<': {X: -1, Y: 0}, '^': {X: 0, Y: -1}, 'v': {X: 0, Y: 1}}

	var nextPos = pos.Add(delta[dir])

	var v = grid.GetPos(nextPos)

	if v == EMPTY {
		return nextPos, true
	}
	if v == WALL {
		return pos, false
	}

	if v == BOX || dir == '<' || dir == '>' {
		return canMove(grid, nextPos, dir)
	}

	// dir is ^ or v
	var p1, cm1 = canMove(grid, nextPos, dir)
	if !cm1 {
		return pos, false
	}

	var p2 game2d.Pos
	var cm2 bool
	if v == LBOX {
		p2, cm2 = canMove(grid, nextPos.E(), dir)
	} else if v == RBOX {
		p2, cm2 = canMove(grid, nextPos.W(), dir)
	}
	if dir == '^' {
		return game2d.Pos{0, min(p1.Y, p2.Y)}, cm2
	} else if dir == 'v' {
		return game2d.Pos{0, max(p1.Y, p2.Y)}, cm2
	}

	return pos, false
}

func moverec(grid *game2d.MatrixChar, pos game2d.Pos, dir uint8, goal game2d.Pos) {
	if (dir == '<' || dir == '>') && pos.X == goal.X {
		return
	}
	if (dir == '^' || dir == 'v') && pos.Y == goal.Y {
		return
	}

	var delta = map[uint8]game2d.Pos{'>': {X: 1, Y: 0}, '<': {X: -1, Y: 0}, '^': {X: 0, Y: -1}, 'v': {X: 0, Y: 1}}
	var nextPos = pos.Add(delta[dir])

	if element := grid.GetPos(pos); element == BOX || element == ROBOT || dir == '<' || dir == '>' {
		moverec(grid, nextPos, dir, goal)
		grid.SetPos(nextPos, element)
		grid.SetPos(pos, EMPTY)
		return
	}

	// dir is ^ or v
	if grid.GetPos(pos) == LBOX && grid.GetPos(pos.E()) == RBOX {
		moverec(grid, nextPos, dir, goal)
		moverec(grid, nextPos.E(), dir, goal)
		grid.SetPos(nextPos, LBOX)
		grid.SetPos(nextPos.E(), RBOX)
		grid.SetPos(pos, EMPTY)
		grid.SetPos(pos.E(), EMPTY)
		return
	}
	if grid.GetPos(pos) == RBOX && grid.GetPos(pos.W()) == LBOX {
		moverec(grid, nextPos, dir, goal)
		moverec(grid, nextPos.W(), dir, goal)
		grid.SetPos(nextPos, RBOX)
		grid.SetPos(nextPos.W(), LBOX)
		grid.SetPos(pos, EMPTY)
		grid.SetPos(pos.W(), EMPTY)
		return
	}
}

func move(grid *game2d.MatrixChar, robot *game2d.Pos, dir uint8) bool {
	var delta = map[uint8]game2d.Pos{'>': {X: 1, Y: 0}, '<': {X: -1, Y: 0}, '^': {X: 0, Y: -1}, 'v': {X: 0, Y: 1}}
	if p, ok := canMove(grid, *robot, dir); ok {
		moverec(grid, *robot, dir, p)
		robot.X += delta[dir].X
		robot.Y += delta[dir].Y
		return true
	}
	return false
}

func sumGPS(grid *game2d.MatrixChar) int {
	var res int
	for p := range grid.AllPos() {
		if grid.Get(p.X, p.Y) == BOX || grid.Get(p.X, p.Y) == LBOX {
			res += 100*p.Y + p.X
		}
	}
	return res
}

func scaleGrid(grid *game2d.MatrixChar) *game2d.MatrixChar {
	var res = game2d.NewMatrixChar(2*grid.LenX(), grid.LenY())
	for y := 0; y < grid.LenY(); y++ {
		for x := 0; x < grid.LenX(); x++ {
			if grid.Get(x, y) == EMPTY {
				res.Set(2*x, y, EMPTY)
				res.Set(2*x+1, y, EMPTY)
				continue
			}
			if grid.Get(x, y) == WALL {
				res.Set(2*x, y, WALL)
				res.Set(2*x+1, y, WALL)
			}
			if grid.Get(x, y) == ROBOT {
				res.Set(2*x, y, ROBOT)
				res.Set(2*x+1, y, EMPTY)
			}
			if grid.Get(x, y) == BOX {
				res.Set(2*x, y, LBOX)
				res.Set(2*x+1, y, RBOX)
			}
		}
	}
	return res
}

func Part1(input string) int {
	var parts = strings.Split(input, "\n\n")
	var grid = game2d.BuildMatrixCharFromString(parts[0])

	var robot, _ = grid.Find(ROBOT)
	for _, lines := range strings.Split(parts[1], "\n") {
		for _, m := range lines {
			move(grid, &robot, uint8(m))
		}
	}

	return sumGPS(grid)
}

func Part2(input string) int {
	var parts = strings.Split(input, "\n\n")
	var grid = game2d.BuildMatrixCharFromString(parts[0])
	grid = scaleGrid(grid)

	var robot, _ = grid.Find(ROBOT)
	var cpt int

	for _, lines := range strings.Split(parts[1], "\n") {
		for _, m := range lines {
			cpt++
			//fmt.Printf("try %d: %c\n", cpt, m)
			move(grid, &robot, uint8(m))
			//fmt.Println(grid)
		}
	}

	return sumGPS(grid)
}

func main() {
	fmt.Println("--2024 day 14 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
