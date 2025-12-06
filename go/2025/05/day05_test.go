package main

import (
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func TestPart(t *testing.T) {
	var sample = utils.Dedent(`
	3-5
	10-14
	16-20
	12-18

	1
	5
	8
	11
	17
	32`)

	var tests = []utils.Test[string, int]{
		{Func: Part1, Input: sample, Expected: 3},
		{Func: Part2, Input: sample, Expected: 14},
		//
		{Func: Part1, Input: utils.Input(), Expected: 761},
		{Func: Part2, Input: utils.Input(), Expected: 345755049374932},
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
