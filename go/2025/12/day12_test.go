package main

import (
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func TestPart(t *testing.T) {
	var tests = []utils.Test[string, int]{
		//{Func: Part1, Input: sample, Expected: 0},
		//{ Func: Part2, Input: sample, Expected: 0 },
		//
		{Func: Part1, Input: utils.Input(), Expected: 425},
		//{Func: Part2, Input: utils.Input(), Expected: 0},
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
