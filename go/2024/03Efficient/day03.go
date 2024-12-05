package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"time"
	"unicode"
)

func parseDo(input string, i int) (int, bool) {
	var bytes = []byte(input)
	var end = len(bytes)
	var j = i
	if !(j < end && bytes[j] == 'd') {
		return j - i + 1, false
	}
	j++
	if !(j < end && bytes[j] == 'o') {
		return j - i + 1, false
	}
	j++
	if !(j < end && bytes[j] == '(') {
		return j - i + 1, false
	}
	j++
	if !(j < end && bytes[j] == ')') {
		return j - i + 1, false
	}
	j++
	return j - i, true
}

func parseDont(input string, i int) (int, bool) {
	var bytes = []byte(input)
	var end = len(bytes)
	var j = i
	if !(j < end && bytes[j] == 'd') {
		return j - i + 1, false
	}
	j++
	if !(j < end && bytes[j] == 'o') {
		return j - i + 1, false
	}
	j++
	if !(j < end && bytes[j] == 'n') {
		return j - i + 1, false
	}
	j++
	if !(j < end && bytes[j] == '\'') {
		return j - i + 1, false
	}
	j++
	if !(j < end && bytes[j] == 't') {
		return j - i + 1, false
	}
	j++
	if !(j < end && bytes[j] == '(') {
		return j - i + 1, false
	}
	j++
	if !(j < end && bytes[j] == ')') {
		return j - i + 1, false
	}
	j++
	return j - i, true
}

func parseMul(input string, i int) (int, bool, int) {
	var bytes = []byte(input)
	var end = len(bytes)
	var j = i
	//fmt.Printf("parseMul: i=%d '%c'\n", i, bytes[i])
	if !(j < end && bytes[j] == 'm') {
		return j - i + 1, false, 0
	}
	j++
	if !(j < end && bytes[j] == 'u') {
		return j - i + 1, false, 0
	}
	j++
	if !(j < end && bytes[j] == 'l') {
		return j - i + 1, false, 0
	}
	j++
	if !(j < end && bytes[j] == '(') {
		return j - i + 1, false, 0
	}
	j++
	var x, y int
	for j < end && unicode.IsDigit(rune(input[j])) {
		x = x*10 + int(input[j]-'0')
		j++
	}
	if !(j < end && bytes[j] == ',') {
		return j - i, false, 0
	}
	j++
	for j < end && unicode.IsDigit(rune(input[j])) {
		y = y*10 + int(input[j]-'0')
		j++
	}
	if !(j < end && bytes[j] == ')') {
		return j - i + 1, false, 0
	}
	j++
	return j - i, true, x * y
}

func parse(input string, part2 bool) int {
	var res int
	var enable = true
	var i int
	for i < len(input) {

		if part2 {
			if j, ok := parseDo(input, i); ok {
				enable = true
				i += j
				continue
			} else {
				//i += j
			}
			if j, ok := parseDont(input, i); ok {
				enable = false
				i += j
				continue
			} else {
				//i += j
			}
		}

		if !enable {
			i++
			continue
		}
		if j, ok, v := parseMul(input, i); ok {
			res += v
			i += j
		} else {
			i += j
		}
	}
	return res
}

func Part1(input string) int {
	return parse(input, false)
}

func Part2(input string) int {
	return parse(input, true)
}

func main() {
	fmt.Println("--2024 day 03 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
