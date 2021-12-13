package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Pos struct {
	x int
	y int
}
type Instr struct {
	axis  rune
	value int
}

type set map[Pos]struct{}

func BuildSet() set {
	return make(map[Pos]struct{})
}

func (s set) Add(value Pos) {
	s[value] = struct{}{}
}

func (s set) Remove(value Pos) {
	delete(s, value)
}

func (s set) Contains(value Pos) bool {
	_, ok := s[value]
	return ok
}
func (s set) Len() int {
	return len(s)
}

func BuildPos(input string) set {
	res := BuildSet()
	for _, l := range strings.Split(input, "\n") {
		coords := strings.SplitN(l, ",", 2)
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		res.Add(Pos{x, y})
	}

	return res
}

func BuildInstr(input string) []Instr {
	res := []Instr{}
	for _, l := range strings.Split(input, "\n") {
		instr := strings.Split(l, " ")
		param := strings.SplitN(instr[2], "=", 2)
		value, _ := strconv.Atoi(param[1])
		res = append(res, Instr{axis: rune(param[0][0]), value: value})
	}

	return res
}

func step(screen set, inst Instr) {
	d := inst.value
	for p := range screen {
		switch inst.axis {
		case 'x':
			if p.x > d {
				screen.Add(Pos{d - (p.x - d), p.y})
				screen.Remove(p)
			}
		case 'y':
			if p.y > d {
				screen.Add(Pos{p.x, d - (p.y - d)})
				screen.Remove(p)
			}
		}
	}
}

func display(pos set, inst []Instr) {
	screen := [6][40]rune{}
	for _, inst := range inst {
		step(pos, inst)
	}
	for p := range pos {
		screen[p.y][p.x] = '#'
	}
	for _, line := range screen {
		for _, c := range line {
			if c == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
}

func Part1(input string) int {
	parts := strings.SplitN(strings.TrimSuffix(input, "\n"), "\n\n", 2)
	pos := BuildPos(parts[0])
	inst := BuildInstr(parts[1])
	step(pos, inst[0])
	return pos.Len()
}

func Part2(input string) int {
	parts := strings.SplitN(strings.TrimSuffix(input, "\n"), "\n\n", 2)
	pos := BuildPos(parts[0])
	inst := BuildInstr(parts[1])
	display(pos, inst)
	return 0
}

func main() {
	content, _ := ioutil.ReadFile("../../inputs/day13.txt")

	start := time.Now()
	fmt.Println("part1: ", Part1(string(content)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(content)))
	fmt.Println(time.Since(start))
}
