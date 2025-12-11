package main

import (
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func TestPart(t *testing.T) {
	var sample = utils.Dedent(`
aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
`)
	var sample2 = utils.Dedent(`
svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
`)

	var tests = []utils.Test[string, int]{
		{Func: Part1, Input: sample, Expected: 5},
		{Func: Part2, Input: sample2, Expected: 2},
		//
		{Func: Part1, Input: utils.Input(), Expected: 690},
		{Func: Part2, Input: utils.Input(), Expected: 557332758684000},
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
