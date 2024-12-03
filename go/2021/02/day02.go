package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func parse(line string) (command string, value int) {
	s := strings.SplitN(line, " ", 2)
	command = s[0]
	value, _ = strconv.Atoi(s[1])
	return
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	horizontal := 0
	depth := 0
	for l := range lines {
		c, v := parse(lines[l])
		switch c {
		case "forward":
			horizontal += v
		case "up":
			depth -= v
		case "down":
			depth += v
		}
	}
	return horizontal * depth
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	horizontal := 0
	depth := 0
	aim := 0
	for l := range lines {
		c, v := parse(lines[l])
		switch c {
		case "forward":
			horizontal += v
			depth += (v * aim)
		case "up":
			aim -= v
		case "down":
			aim += v
		}
	}
	return horizontal * depth
}

func main() {
	fmt.Println("--2021 day 02 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
