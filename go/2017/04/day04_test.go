package main

import (
	_ "embed"
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func TestPart(t *testing.T) {
	var tests = []utils.Test[string, int]{
		{Func: Part1, Input: "aa bb cc dd ee", Expected: 1},
		{Func: Part1, Input: "aa bb cc dd aa", Expected: 0},
		{Func: Part1, Input: "aa bb cc dd aaa", Expected: 1},
		{Func: Part2, Input: "abcde fghij", Expected: 1},
		{Func: Part2, Input: "abcde xyz ecdab", Expected: 0},
		{Func: Part2, Input: "a ab abc abd abf abj", Expected: 1},
		{Func: Part2, Input: "iiii oiii ooii oooi oooo", Expected: 1},
		{Func: Part2, Input: "oiii ioii iioi iiio", Expected: 0},
		{Func: Part1, Input: utils.Input(), Expected: 325},
		{Func: Part2, Input: utils.Input(), Expected: 119},
	}
	utils.TestPart(t, tests)
}

func BenchmarkPart1(b *testing.B) {
	var inputDay = utils.Input()
	for range b.N {
		Part1(inputDay)
	}
}
func BenchmarkPart2(b *testing.B) {
	var inputDay = utils.Input()
	for range b.N {
		Part2(inputDay)
	}
}
