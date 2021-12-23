package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

type matrix [][]Octopus
type Octopus struct {
	energy  uint8
	flashed bool
}

func BuildMatrix(lines []string) matrix {
	m := make([][]Octopus, len(lines))
	for j, l := range lines {
		l = strings.TrimSpace(l)
		m[j] = make([]Octopus, len(l))
		for i, c := range l {
			m[j][i] = Octopus{energy: uint8(c - '0'), flashed: false}
		}
	}
	return m
}

func increase_energy(m matrix) {
	for j := range m {
		for i := range m[j] {
			m[j][i].energy += 1
		}
	}
}

func increase_neighbours_energy(m matrix, x, y int) {
	for j := y - 1; j <= y+1; j++ {
		for i := x - 1; i <= x+1; i++ {
			if j >= 0 && j < len(m) && i >= 0 && i < len(m[j]) {
				m[j][i].energy += 1
			}
		}
	}
	m[y][x].energy -= 1
}

func clear_flashed(m matrix) {
	for j := range m {
		for i := range m[j] {
			if m[j][i].flashed {
				m[j][i].flashed = false
				m[j][i].energy = 0
			}
		}
	}
}

func flash(m matrix) int {
	continue_flashing := true
	flashed := 0
	for continue_flashing {
		continue_flashing = false
		for j := range m {
			for i := range m[j] {
				if m[j][i].energy > 9 && !m[j][i].flashed {
					m[j][i].flashed = true
					flashed++
					continue_flashing = true
					increase_neighbours_energy(m, i, j)
				}
			}
		}
	}
	clear_flashed(m)
	return flashed
}

func step(m matrix) int {
	increase_energy(m)
	return flash(m)
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := BuildMatrix(lines)

	res := 0
	for i := 1; i <= 100; i++ {
		res += step(m)
	}
	return res
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	m := BuildMatrix(lines)
	n := len(m) * len(m[0])

	for i := 1; ; i++ {
		flashed := step(m)
		if flashed == n {
			return i
		}
	}
}

func main() {
	content, _ := ioutil.ReadFile("../../inputs/day11.txt")

	start := time.Now()
	fmt.Println("part1: ", Part1(string(content)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(content)))
	fmt.Println(time.Since(start))
}
