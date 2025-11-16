package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func validPassphrase(passphrase string) bool {
	var words = strings.Fields(passphrase)
	slices.Sort(words)
	var previous string
	for _, word := range words {
		if word == previous {
			return false
		}
		previous = word
	}
	return true
}

func solve(lines []string) int {
	var res int
	for _, line := range lines {
		if validPassphrase(line) {
			res++
		}
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	var lines = strings.Split(input, "\n")
	return solve(lines)
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	var lines = strings.Split(input, "\n")
	var orderedLines []string
	for _, line := range lines {
		words := strings.Fields(line)
		var orderedWords []string
		for _, word := range words {
			var letters = []byte(word)
			slices.Sort(letters)
			orderedWords = append(orderedWords, string(letters))
		}
		slices.Sort(orderedWords)
		orderedLines = append(orderedLines, strings.Join(orderedWords, " "))
	}

	return solve(orderedLines)
}

func main() {
	fmt.Println("--2017 day 04 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
