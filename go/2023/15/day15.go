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

func hash(input string) int {
	var res int
	for _, c := range input {
		res += int(c)
		res *= 17
		res &= 0xff // % 256
	}
	return res
}

type Box []Lens

type Lens struct {
	name   string
	number string
}

func index(b Box, name string) int {
	for i, l := range b {
		if l.name == name {
			return i
		}
	}
	return -1
}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	var parts = strings.Split(input, ",")
	var res int
	for _, part := range parts {
		res += hash(part)
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	var parts = strings.Split(input, ",")
	var boxes = make([]Box, 256)
	for _, part := range parts {
		if part[len(part)-1] == '-' {
			name := part[:len(part)-1]
			h := hash(name)
			box := boxes[h]
			if i := index(box, name); i != -1 {
				boxes[h] = append(box[:i], box[i+1:]...)
			}
		} else if name, number, found := strings.Cut(part, "="); found {
			h := hash(name)
			box := boxes[h]
			if i := index(box, name); i != -1 {
				boxes[h] = append(box[:i], Lens{name, number})
				boxes[h] = append(boxes[h], box[i+1:]...)
			} else {
				boxes[h] = append(box, Lens{name, number})
			}
		}
	}

	var res int
	for i, box := range boxes {
		for j, lens := range box {
			value, _ := strconv.Atoi(lens.number)
			res += (1 + i) * (j + 1) * value
		}
	}

	return res
}

func main() {
	fmt.Println("--2023 day 15 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
