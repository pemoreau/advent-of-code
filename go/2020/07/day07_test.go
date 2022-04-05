package main

import (
	_ "embed"
	"testing"
)

func TestPart1Input(t *testing.T) {
	result := Part1(input_day)
	expected := 128
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2Input(t *testing.T) {
	result := Part2(input_day)
	expected := 20189
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(input_day)
	}
}
func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(input_day)
	}
}
