package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func TestPart1(t *testing.T) {
	result := Part1(inputTest)
	expected := 19114
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := 167409079868000
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart1Input(t *testing.T) {
	result := Part1(inputDay)
	expected := 420739
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := Part2(inputDay)
	expected := 130251901420382
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
