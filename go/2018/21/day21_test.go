package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

//func TestPart1(t *testing.T) {
//	result := Part2(inputTest)
//	expected := 1
//	if result != expected {
//		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
//	}
//}

//func TestPart2(t *testing.T) {
//	result := Part2(inputTest)
//	expected := 66
//	if result != expected {
//		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
//	}
//}

func TestPart1Input(t *testing.T) {
	result := Part1(inputDay)
	expected := 11513432
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := Part2(inputDay)
	expected := 7434231
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
