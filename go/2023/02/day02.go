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

type rgb struct {
	r, g, b int
}

func parse(input string) [][]rgb {
	var res [][]rgb
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		_, err := strconv.Atoi(parts[0][5:])
		if err != nil {
			panic(err)
		}

		parts = strings.Split(parts[1], ";")
		var rbgLine []rgb
		for _, part := range parts {
			var c = rgb{0, 0, 0}
			colors := strings.Split(part, ",")
			for _, color := range colors {
				nc := strings.Split(strings.TrimSpace(color), " ")
				n, _ := strconv.Atoi(nc[0])
				switch strings.TrimSpace(nc[1]) {
				case "red":
					c.r = n
				case "green":
					c.g = n
				case "blue":
					c.b = n
				}
			}
			rbgLine = append(rbgLine, c)
		}
		res = append(res, rbgLine)
	}
	return res
}

func maxRgb(l []rgb) rgb {
	var res rgb
	for _, c := range l {
		res.r = max(res.r, c.r)
		res.g = max(res.g, c.g)
		res.b = max(res.b, c.b)
	}
	return res
}

func Part1(input string) int {
	var res int
	data := parse(input)

	for id, line := range data {
		mrgb := maxRgb(line)
		if !(mrgb.r > 12 || mrgb.g > 13 || mrgb.b > 14) {
			res += id + 1
		}
	}

	return res
}

func Part2(input string) int {
	var res int
	data := parse(input)

	for _, line := range data {
		mrgb := maxRgb(line)
		power := mrgb.r * mrgb.g * mrgb.b
		res += power
	}

	return res
}

func main() {
	fmt.Println("--2023 day 02 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
