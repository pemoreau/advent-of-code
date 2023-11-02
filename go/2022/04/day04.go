package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Filename = "example"
	s.Mode ^= scanner.SkipComments  // don't skip comments
	s.Whitespace |= 1<<'-' | 1<<',' // skip minus and comma

	res := 0
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		a, _ := strconv.Atoi(s.TokenText())
		//s.Scan() // skip -
		s.Scan()
		b, _ := strconv.Atoi(s.TokenText())
		//s.Scan() // skip ,
		s.Scan()
		c, _ := strconv.Atoi(s.TokenText())
		//s.Scan() // skip -
		s.Scan()
		d, _ := strconv.Atoi(s.TokenText())
		// check inclusion
		if (a <= c && d <= b) || (c <= a && b <= d) {
			res++
		}
	}
	return res
}

func Part2(input string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Whitespace |= 1<<'-' | 1<<',' // skip minus and comma
	res := 0
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		a, _ := strconv.Atoi(s.TokenText())
		s.Scan()
		b, _ := strconv.Atoi(s.TokenText())
		s.Scan()
		c, _ := strconv.Atoi(s.TokenText())
		s.Scan()
		d, _ := strconv.Atoi(s.TokenText())
		// check overlap
		if !(b < c || d < a) {
			res++
		}
	}
	return res
}

func Part1Slow(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		var a, b, c, d int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &c, &d)
		// check inclusion
		if (a <= c && d <= b) || (c <= a && b <= d) {
			res++
		}
	}
	return res
}

func Part2Slow(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		var a, b, c, d int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &c, &d)
		// check overlap
		if !(b < c || d < a) {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println("--2022 day 04 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
