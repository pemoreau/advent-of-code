package main

import (
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func TestPart(t *testing.T) {
	var sample = utils.Dedent(`
			987654321111111
			811111111111119
			234234234234278
			818181911112111
            `)

	var tests = []utils.Test[string, int]{
		{Func: Part1, Input: sample, Expected: 357},
		{Func: Part2, Input: sample, Expected: 3121910778619},
		//
		{Func: Part1, Input: utils.Input(), Expected: 17311},
		{Func: Part2, Input: utils.Input(), Expected: 171419245422055},
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
