package main

import (
	_ "embed"
	"github.com/pemoreau/advent-of-code/go/utils"
	"testing"
)

func TestPart1Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part1(inputDay)
	expected := 1801
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part2(inputDay)
	expected := 2060
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
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
