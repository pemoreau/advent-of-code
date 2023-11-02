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

type Color struct {
	name    string
	contain []ColorWithMultiplicity
}

type ColorWithMultiplicity struct {
	name         string
	multiplicity int
}

func parseLine(line string) Color {
	parts := strings.Split(line, "bags")
	color := Color{
		name:    strings.TrimSpace(parts[0]),
		contain: []ColorWithMultiplicity{},
	}
	parts = strings.Split(line, "contain")
	parts = strings.Split(strings.TrimSpace(parts[1]), " ")
	for i := 0; i < len(parts); i = i + 4 {
		multiplicity, err := strconv.Atoi(parts[i])
		if err == nil {
			colorm := ColorWithMultiplicity{
				multiplicity: multiplicity,
				name:         parts[i+1] + " " + parts[i+2],
			}
			color.contain = append(color.contain, colorm)
		}
	}
	return color
}

type Store = map[string][]ColorWithMultiplicity

func readFile() Store {
	lines := strings.Split(strings.TrimSuffix(inputDay, "\n"), "\n")
	res := make(Store)
	for _, line := range lines {
		color := parseLine(line)
		res[color.name] = color.contain
	}

	return res
}

// returns true if name or a descendent contains colorName
func containsColor(s Store, name string, colorName string) bool {
	// fmt.Printf("<%s> contains? <%s>\n", name, colorName)
	colors := s[name]
	for _, c := range colors {
		if c.name == colorName || containsColor(s, c.name, colorName) {
			return true
		}
	}
	return false
}
func countBags(s Store, name string) int {
	// fmt.Printf("<%s> contains? <%s>\n", name, colorName)
	colors := s[name]
	res := 1
	for _, c := range colors {
		res += c.multiplicity * countBags(s, c.name)
	}
	return res
}

func Part1(input string) (res int) {
	colors := readFile()
	for key := range colors {
		b := containsColor(colors, key, "shiny gold")
		if b {
			res++
		}
	}
	return
}

func Part2(input string) int {
	colors := readFile()
	return countBags(colors, "shiny gold") - 1
}

func main() {
	fmt.Println("--2020 day 07 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
