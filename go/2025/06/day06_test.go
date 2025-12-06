package main

import (
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func TestPart(t *testing.T) {
	var sample = utils.Dedent(`
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`)

	var tests = []utils.Test[string, int]{
		{Func: Part1, Input: sample, Expected: 4277556},
		{Func: Part2, Input: sample, Expected: 3263827},
		//
		{Func: Part1, Input: utils.Input(), Expected: 6299564383938},
		{Func: Part2, Input: utils.Input(), Expected: 11950004808442},
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
