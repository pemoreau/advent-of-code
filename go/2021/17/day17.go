package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

//go:embed input.txt
var inputDay string

type Pos struct{ x, y int }
type Area struct{ upperLeft, lowerRight Pos }
type State struct {
	pos      Pos
	velocity Pos
}

func step(s *State) {
	s.pos.x += s.velocity.x
	s.pos.y += s.velocity.y

	if s.velocity.x > 0 {
		s.velocity.x -= 1
	} else if s.velocity.x < 0 {
		s.velocity.x += 1
	}
	s.velocity.y -= 1
}

func simulate(initialVelocity Pos, target Area) (maxHeight int, reached bool) {
	state := State{Pos{0, 0}, initialVelocity}
	maxHeight = 0
	for state.pos.x <= 1+target.lowerRight.x && state.pos.y >= target.lowerRight.y-1 {
		step(&state)
		if state.pos.y > maxHeight {
			maxHeight = state.pos.y
		}
		if state.pos.x >= target.upperLeft.x && state.pos.x <= target.lowerRight.x &&
			state.pos.y <= target.upperLeft.y && state.pos.y >= target.lowerRight.y {
			return maxHeight, true
		}
	}
	reached = false
	return
}

func parseInput(in string) Area {
	reg := regexp.MustCompile(`target area: x=([0-9\-]+)\.\.([0-9\-]+), y=([0-9\-]+)\.\.([0-9\-]+)`)
	g := reg.FindStringSubmatch(in)
	x1, _ := strconv.Atoi(g[1])
	x2, _ := strconv.Atoi(g[2])
	y1, _ := strconv.Atoi(g[3])
	y2, _ := strconv.Atoi(g[4])
	return Area{
		upperLeft:  Pos{x1, y2},
		lowerRight: Pos{x2, y1},
	}

}

func stepX(x, vx, min, max int) int {
	res := vx
	for x <= max && vx > 0 {
		x += vx
		if vx > 0 {
			vx -= 1
		} else if vx < 0 {
			vx += 1
		}
		if x >= min && x <= max {
			return res
		}
	}
	return 0
}

func Part1(input string) int {
	target := parseInput(input)
	y := target.lowerRight.y
	return y * (y + 1) / 2
}

func Part2(input string) int {
	target := parseInput(input)

	vxs := []int{}
	for vx := 0; vx <= target.lowerRight.x; vx++ {
		if v := stepX(0, vx, target.upperLeft.x, target.lowerRight.x); v != 0 {
			vxs = append(vxs, v)
		}
	}

	reached := 0
	for _, vx := range vxs {
		for vy := target.lowerRight.y; vy <= -target.lowerRight.y; vy++ {
			_, r := simulate(Pos{vx, vy}, target)
			if r {
				reached += 1
			}
		}
	}

	return reached

}

func main() {
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
