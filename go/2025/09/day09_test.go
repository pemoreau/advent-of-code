package main

import (
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func TestPart(t *testing.T) {
	var sample = utils.Dedent(`
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`)

	var tests = []utils.Test[string, int]{
		{Func: Part1, Input: sample, Expected: 50},
		{Func: Part2, Input: sample, Expected: 24},
		//
		{Func: Part1, Input: utils.Input(), Expected: 4741451444},
		{Func: Part2, Input: utils.Input(), Expected: 1562459680},
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
