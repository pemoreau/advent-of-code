package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type matrix [][]byte

func (m matrix) String() string {
	var buf bytes.Buffer
	for _, l := range m {
		for _, c := range l {
			buf.WriteByte(c)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func BuildMatrix(lines []string) matrix {
	m := make([][]byte, len(lines))
	for j, l := range lines {
		l = strings.TrimSpace(l)
		m[j] = make([]byte, len(l))
		for i, c := range l {
			m[j][i] = byte(c)
		}
	}
	return m
}

func step(m matrix) (r matrix, moved bool) {
	M := len(m[0])
	N := len(m)
	var s matrix = make([][]byte, N)
	moved = false
	for j := 0; j < N; j++ {
		s[j] = make([]byte, M)
		for i := 0; i < M; i++ {
			if m[j][i] == '>' && m[j][(i+1)%M] == '.' {
				s[j][i] = '.'
				s[j][(i+1)%M] = '>'
				moved = true
			} else if s[j][i] == 0 {
				s[j][i] = m[j][i]
			}
		}
	}

	r = make([][]byte, N)
	for j := 0; j < N; j++ {
		r[j] = make([]byte, M)
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if s[j][i] == 'v' && s[(j+1)%N][i] == '.' {
				r[j][i] = '.'
				r[(j+1)%N][i] = 'v'
				moved = true
			} else if r[j][i] == 0 {
				r[j][i] = s[j][i]
			}
		}
	}
	return
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	m := BuildMatrix(lines)
	moved := true
	cpt := 0
	for moved {
		var r matrix
		r, moved = step(m)
		cpt++
		// fmt.Printf("after step %d r:\n%v", cpt, r)
		// fmt.Printf("moved: %v\n", moved)
		m = r
	}
	return cpt
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println("--2021 day 25 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
