package main

import (
	_ "embed"
	"testing"
)

func TestPart2Input(t *testing.T) {
	result := Part2(inputDay)
	expected := 10101
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for range b.N {
		Part1(inputDay)
	}
}

func BenchmarkPart2(b *testing.B) {
	for range b.N {
		Part2(inputDay)
	}
}
