package main

import (
	_ "embed"
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func TestPart(t *testing.T) {
	var tests = []utils.Test[string, int]{
		{
			Func: Part1,
			Input: utils.Dedent(`
				5 1 9 5
                7 5 3
                2 4 6 8
			`),
			Expected: 18,
		},
		{Func: Part1, Input: utils.Input(), Expected: 34925},
		{
			Func: Part2,
			Input: utils.Dedent(`
                5 9 2 8
    			9 4 7 3
    			3 8 6 5
            `),
			Expected: 9,
		},
		{Func: Part2, Input: utils.Input(), Expected: 221},
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
