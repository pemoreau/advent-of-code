package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Player struct {
	space uint16
	score uint16
}

type Dice struct {
	value  uint8
	nbRoll uint16
}

func (p Player) String() string {
	return fmt.Sprintf("space=%d score=%d", p.space, p.score)
}

func (p *Player) move(value uint16) {
	p.space = 1 + ((p.space + value - 1) % 10)
	p.score += p.space
}

func (d *Dice) deterministicRoll(n uint8) (sum uint16) {
	var i uint8
	for i = 0; i < n; i++ {
		sum += uint16(d.value)
		d.nbRoll += 1
		d.value += 1
		if d.value > 100 {
			d.value = 1
		}
	}
	return
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	space0, _ := strconv.Atoi(strings.Split(lines[0], " ")[4])
	space1, _ := strconv.Atoi(strings.Split(lines[1], " ")[4])

	p1 := Player{space: uint16(space0)}
	p2 := Player{space: uint16(space1)}
	d := Dice{value: 1}

	for {
		m := d.deterministicRoll(3)
		p1.move(m)
		if p1.score >= 1000 {
			return int(d.nbRoll) * int(p2.score)
		}
		m = d.deterministicRoll(3)
		p2.move(m)
		if p2.score >= 1000 {
			return int(d.nbRoll) * int(p1.score)
		}
	}
}

type state struct {
	player0        bool
	roll           uint
	dice           uint8
	score0, score1 uint
	space0, space1 uint8
}

type win struct{ win0, win1 uint }

func explore(s state) win {

	if res, found := cache[s]; found {
		return res
	}

	initS := s

	if s.player0 {
		s.space0 = 1 + ((s.space0 + s.dice - 1) % 10)
		if s.roll == 2 {
			s.score0 += uint(s.space0)
		}
		if s.score0 >= 21 {
			cache[initS] = win{1, 0}
			return win{1, 0}
		}
	} else {
		s.space1 = 1 + ((s.space1 + s.dice - 1) % 10)
		if s.roll == 2 {
			s.score1 += uint(s.space1)
		}
		if s.score1 >= 21 {
			cache[initS] = win{0, 1}
			return win{0, 1}
		}
	}

	if s.roll == 2 {
		s.player0 = !s.player0
	}
	s.roll = (s.roll + 1) % 3
	var win0, win1 uint
	for d := uint8(1); d <= 3; d++ {
		s.dice = d
		win := explore(s)
		win0 += win.win0
		win1 += win.win1
	}

	// fmt.Printf("store %v --> %v\n", initS, win{win0, win1})
	cache[initS] = win{win0, win1}
	return win{win0, win1}
}

var cache map[state]win

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	cache = make(map[state]win) // new cache for each benchmark

	space0, _ := strconv.Atoi(strings.Split(lines[0], " ")[4])
	space1, _ := strconv.Atoi(strings.Split(lines[1], " ")[4])
	player0 := true
	roll := uint(0)
	var win0, win1 uint
	for d := uint8(1); d <= 3; d++ {
		s := state{player0, roll, d, 0, 0, uint8(space0), uint8(space1)}
		win := explore(s)
		win0 += win.win0
		win1 += win.win1
	}

	if win0 > win1 {
		return int(win0)
	} else {
		return int(win1)
	}
}

func main() {
	fmt.Println("--2021 day 21 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
