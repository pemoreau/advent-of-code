package main

import (
	_ "embed"
	"fmt"
	"github.com/bits-and-blooms/bitset"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func winning(line string) int {
	split := func(c rune) bool { return c == ':' || c == '|' }
	fields := strings.FieldsFunc(line, split)

	var winningNumbers = set.NewSet[string]()
	for _, n := range strings.Fields(fields[1]) {
		winningNumbers.Add(n)
	}
	var numbers = set.NewSet[string]()
	for _, n := range strings.Fields(fields[2]) {
		numbers.Add(n)
	}
	return winningNumbers.Intersect(numbers).Len()
}

func winning2(line string) int {
	split := func(c rune) bool { return c == ':' || c == '|' }
	fields := strings.FieldsFunc(line, split)

	var winningNumbers = bitset.New(100)
	for _, n := range strings.Fields(fields[1]) {
		v, _ := strconv.Atoi(n)
		winningNumbers.Set(uint(v))
	}
	var numbers = bitset.New(100)
	for _, n := range strings.Fields(fields[2]) {
		v, _ := strconv.Atoi(n)
		numbers.Set(uint(v))
	}
	return int(winningNumbers.IntersectionCardinality(numbers))
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	var res int
	for _, line := range lines {
		v := winning2(line)
		if v > 0 {
			res += 1 << (v - 1)
		}
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	var cards []int
	var matchs []int
	for _, line := range lines {
		v := winning2(line)
		cards = append(cards, 1)
		matchs = append(matchs, v)
	}
	for i := 0; i < len(cards); i++ {
		if matchs[i] > 0 {
			for j := i + 1; j < i+1+matchs[i]; j++ {
				cards[j] += cards[i]
			}
		}
	}
	var res int
	for _, c := range cards {
		res += c
	}
	return res
}

func main() {
	fmt.Println("--2023 day 04 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
