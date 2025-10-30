package main

import (
	_ "embed"
	"reflect"
	"runtime"
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func TestPart(t *testing.T) {
	//var inputDay = utils.Input()
	tests := []struct {
		fn       func(string) int
		input    string
		expected int
	}{
		{fn: Part1, input: "1", expected: 0},
		{fn: Part1, input: "12", expected: 3},
		{fn: Part1, input: "23", expected: 2},
		{fn: Part1, input: "11", expected: 2},
		{fn: Part1, input: "11", expected: 2},
		{fn: Part1, input: "1024", expected: 31},
		{fn: Part1, input: utils.Input(), expected: 326},
		{fn: Part2, input: utils.Input(), expected: 363010},
	}

	for _, tc := range tests {
		fnName := runtime.FuncForPC(reflect.ValueOf(tc.fn).Pointer()).Name()

		t.Run(fnName, func(t *testing.T) {
			got := tc.fn(tc.input)
			if got != tc.expected {
				if len(tc.input) < 30 {
					t.Fatalf("%s: input %s got %d, want %d", fnName, tc.input, got, tc.expected)
				} else {
					t.Fatalf("%s: got %d, want %d", fnName, got, tc.expected)
				}
			}
		})
	}
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
