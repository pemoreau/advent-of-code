package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func fromSnafu(s string) int {
	digit := map[byte]int{'0': 0, '1': 1, '2': 2, '-': -1, '=': -2}
	res := 0
	for i := 0; i < len(s); i++ {
		res = 5*res + digit[s[i]]
	}
	return res
}

func toSnafu(i int) string {
	digit := map[int]byte{0: '0', 1: '1', 2: '2', -1: '-', -2: '='}
	res := ""
	for i > 0 {
		if i%5 > 2 {
			res = string(digit[i%5-5]) + res
			i = i/5 + 1
		} else {
			res = string(digit[i%5]) + res
			i = i / 5
		}
	}
	return res
}

func Part1(input string) string {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		res += fromSnafu(line)
	}

	return toSnafu(res)
}

func main() {
	fmt.Println("--2022 day 25 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))
}
