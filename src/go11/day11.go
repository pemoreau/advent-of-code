package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type matrix [][]Octopus
type Octopus struct {
	energy  uint8
	flashed bool
}

func BuildMatrix(lines []string) matrix {
	m := make([][]Octopus, len(lines))
	for i, l := range lines {
		l = strings.TrimSpace(l)
		m[i] = make([]Octopus, len(l))
		for j, c := range l {
			m[i][j] = Octopus{energy: uint8(c - '0'), flashed: false}
		}
	}
	return m
}

func increase_energy(m matrix) {
	for i := range m {
		for j := range m[i] {
			m[i][j].energy += 1
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
	for y := range m {
		for x := range m[y] {
			if m[y][x].flashed {
				m[y][x].flashed = false
				m[y][x].energy = 0
			}
		}
	}
}

func flash(m matrix) int {
	continue_flashing := true
	flashed := 0
	for continue_flashing {
		continue_flashing = false
		for y := range m {
			for x := range m[y] {
				if m[y][x].energy > 9 && !m[y][x].flashed {
					m[y][x].flashed = true
					flashed++
					continue_flashing = true
					increase_neighbours_energy(m, x, y)
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
