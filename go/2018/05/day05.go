package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var inputDay string

func findRedex(s []byte, fromIndex int) int {
	if fromIndex < 0 {
		fromIndex = 0
	}
	for i := fromIndex; i < len(s)-1; i++ {
		if s[i] == s[i+1]+32 || s[i] == s[i+1]-32 {
			return i
		}
	}
	return -1
}

func rewrite(s []byte, index int) ([]byte, int) {
	if index >= 0 && index+1 < len(s) {
		if s[index] == s[index+1]+32 || s[index] == s[index+1]-32 {
			s = append(s[:index], s[index+2:]...)
			return rewrite(s, index-1)
		}
	}
	return s, index
}

func normalize(s []byte) []byte {
	var index = findRedex(s, 0)
	for index >= 0 {
		s, index = rewrite(s, index)
		index = findRedex(s, index)
	}
	return s
}

func stackReduce(input string, skip int32) int {
	var stack = make([]int32, len(input))
	var sp = 0
	for _, c := range input {
		if c == skip || c == skip-32 {
			continue
		}
		if sp > 0 {
			var top = stack[sp-1]
			if top == c+32 || top == c-32 {
				sp = sp - 1
			} else {
				stack[sp] = c
				sp = sp + 1
			}
		} else {
			stack[sp] = c
			sp = sp + 1
		}
	}
	return sp
}

func Part1(input string) int {
	//var s = []byte(input)
	//return len(normalize(s))
	return stackReduce(input, 0)
}

func Part2(input string) int {
	var minLen = len(input)
	//for letter := 'a'; letter <= 'z'; letter++ {
	//	var s []byte
	//	for _, c := range input {
	//		if c != letter && c != letter-32 {
	//			s = append(s, byte(c))
	//		}
	//	}
	//	s = normalize(s)
	//	minLen = min(minLen, len(s))
	//}
	for letter := 'a'; letter <= 'z'; letter++ {
		minLen = min(minLen, stackReduce(input, letter))
	}

	return minLen
}

func main() {
	fmt.Println("--2018 day 05 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
