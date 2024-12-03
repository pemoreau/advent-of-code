package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func TestPart1(t *testing.T) {
	result := Part1(inputTest)
	expected := 1147
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

//func TestPart2(t *testing.T) {
//	result := Part2(inputTest)
//	expected := 29
//	if result != expected {
//		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
//	}
//}

func TestPart1Input(t *testing.T) {
	result := Part1(inputDay)
	expected := 506385
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := Part2(inputDay)
	expected := 215404
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
