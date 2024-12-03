package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type matrix [26][26]byte

func buildRules(input string) matrix {
	res := matrix{}
	for _, l := range strings.Split(input, "\n") {
		rule := strings.Split(l, " -> ")
		res[rune(rule[0][0])-'A'][rune(rule[0][1])-'A'] = rule[1][0] - 'A'
	}
	return res
}

func step(o *[26][26]int, rules *matrix) [26][26]int {
	res := [26][26]int{}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if o[i][j] > 0 {
				r := rules[i][j]
				res[i][r] += o[i][j]
				res[r][j] += o[i][j]
			}
		}
	}
	return res
}

func computeOccurrences(o *[26][26]int, first, last byte) [26]int {
	res := [26]int{}
	res[first-'A'] = 1
	res[last-'A'] = 1
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			res[i] += o[i][j]
			res[j] += o[i][j]
		}
	}
	for i := 0; i < 26; i++ {
		res[i] /= 2
	}
	return res
}

func solve(subject string, part2 string, n int) int {
	rules := buildRules(part2)
	o := [26][26]int{}
	for i := 0; i < len(subject)-1; i++ {
		o[subject[i]-'A'][subject[i+1]-'A'] += 1
	}
	for i := 0; i < n; i++ {
		o = step(&o, &rules)
	}

	first := subject[0]
	last := subject[len(subject)-1]
	occurrence := computeOccurrences(&o, first, last)
	max := 0
	min := math.MaxInt64
	for _, n := range occurrence {
		if n > max {
			max = n
		}
		if n > 0 && n < min {
			min = n
		}
	}

	return max - min
}

func Part1(input string) int {
	parts := strings.SplitN(strings.TrimSuffix(input, "\n"), "\n\n", 2)
	return solve(parts[0], parts[1], 10)
}

func Part2(input string) int {
	parts := strings.SplitN(strings.TrimSuffix(input, "\n"), "\n\n", 2)
	return solve(parts[0], parts[1], 40)
}

func main() {
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
