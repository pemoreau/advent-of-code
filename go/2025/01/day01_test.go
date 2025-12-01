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
              L68
              L30
              R48
              L5
              R60
              L55
              L1
              L99
              R14
              L82
			`),
			Expected: 3,
		},
		{
			Func: Part2,
			Input: utils.Dedent(`
              L68
              L30
              R48
              L5
              R60
              L55
              L1
              L99
              R14
              L82
			`),
			Expected: 6,
		},

		{Func: Part1, Input: utils.Input(), Expected: 1043},
		{Func: Part2, Input: utils.Input(), Expected: 5963},
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
