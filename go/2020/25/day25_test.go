package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var input_test string

func TestPart1(t *testing.T) {
	result := Part1(string(input_test))
	expected := 14897079
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := Part1(string(input_day))
	expected := 9620012
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(input_day)
	}
}
