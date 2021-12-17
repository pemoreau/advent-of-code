package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input_day string

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

func Part1(input string) int {
	target := Area{upperLeft: Pos{206, -57}, lowerRight: Pos{250, -105}}

	maxHeight := 0
	bestVelocity := Pos{0, 0}
	for vx := 0; vx <= 200; vx++ {
		for vy := 0; vy <= 500; vy++ {
			h, r := simulate(Pos{vx, vy}, target)
			if r && h > maxHeight {
				maxHeight = h
				bestVelocity = Pos{vx, vy}
			}
		}
	}
	fmt.Println(maxHeight, bestVelocity)

	return maxHeight
}

func Part2(input string) int {
	target := Area{upperLeft: Pos{206, -57}, lowerRight: Pos{250, -105}}
	// target = Area{upperLeft: Pos{20, -5}, lowerRight: Pos{30, -10}}

	reached := []Pos{}
	for vx := 0; vx <= 250; vx++ {
		for vy := -110; vy <= 500; vy++ {
			_, r := simulate(Pos{vx, vy}, target)
			if r {
				reached = append(reached, Pos{vx, vy})
			}
		}
	}
	// fmt.Println(reached)

	return len(reached)
}

func main() {
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}
