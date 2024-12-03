package main

import (
	_ "embed"
	"testing"
)

func TestPart1Input(t *testing.T) {
	result := Part1(inputDay)
	expected := "1041411104"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := Part2(inputDay)
	expected := 20174745
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
