package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Pos struct {
	x int
	y int
}

type Instr struct {
	axis  rune
	value int
}

func BuildPos(input string) set.Set[Pos] {
	res := set.NewSet[Pos]()
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

func step(screen set.Set[Pos], inst Instr) {
	d := inst.value
	for p := range screen {
		switch inst.axis {
		case 'x':
			if p.x > d {
				screen.Add(Pos{2*d - p.x, p.y})
				screen.Remove(p)
			}
		case 'y':
			if p.y > d {
				screen.Add(Pos{p.x, 2*d - p.y})
				screen.Remove(p)
			}
		}
	}
}

func display(pos set.Set[Pos]) {
	screen := [6][40]rune{}
	for p := range pos {
		screen[p.y][p.x] = '#'
	}

	var buf bytes.Buffer
	for _, line := range screen {
		for _, c := range line {
			if c == 0 {
				buf.WriteRune(' ')
			} else {
				buf.WriteRune(c)
			}
		}
		buf.WriteRune('\n')
	}
	fmt.Println(buf.String())
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
	for _, inst := range inst {
		step(pos, inst)
	}
	display(pos)
	return 0
}

func main() {

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
