package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type Handler interface {
	tick(x int)
	value() (int, string)
}

type Signal struct {
	strength int
	cycle    int
}

func NewSignal() *Signal {
	return &Signal{cycle: 1}
}

func (s *Signal) tick(x int) {
	if s.cycle%40 == 20 {
		s.strength += (s.cycle * x)
	}
	s.cycle++
}

func (s *Signal) value() (int, string) {
	return s.strength, ""
}

type Screen struct {
	lines       []string
	cursor      int
	currentLine string
}

func NewScreen() *Screen {
	return &Screen{
		lines: make([]string, 0),
	}
}

func (s *Screen) tick(x int) {
	if s.cursor >= x-1 && s.cursor <= x+1 {
		s.currentLine += "#"
	} else {
		s.currentLine += "."
	}
	s.cursor++
	if s.cursor >= 40 {
		s.lines = append(s.lines, s.currentLine)
		s.currentLine = ""
		s.cursor = 0
	}
}

func (s *Screen) String() string {
	return strings.Join(s.lines, "\n")
}

func (s *Screen) value() (int, string) {
	return 0, s.String()
}

func run(input string, h Handler) (int, string) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	x := 1
	for _, line := range lines {
		if strings.HasPrefix(line, "noop") {
			h.tick(x)
		} else {
			h.tick(x)
			h.tick(x)
			var v int
			fmt.Sscanf(line, "addx %d", &v)
			x += v
		}
	}
	return h.value()
}

func Part1(input string) int {
	i, _ := run(input, NewSignal())
	return i
}
func Part2(input string) string {
	_, s := run(input, NewScreen())
	return s
}

func main() {
	fmt.Println("--2022 day 10 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2:\n", Part2(inputDay))
	fmt.Println(time.Since(start))
}
