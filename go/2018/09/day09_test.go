package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func TestPart1(t *testing.T) {
	result := Solve(10, 1618)
	expected := 8317
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Solve(13, 7999)
	expected = 146373
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Solve(17, 1104)
	expected = 2764
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Solve(21, 6111)
	expected = 54718
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Solve(30, 5807)
	expected = 37305
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

//func TestPart2(t *testing.T) {
//	result := Part2(inputTest)
//	expected := 66
//	if result != expected {
//		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
//	}
//}

func TestPart1Input(t *testing.T) {
	result := Part1(inputDay)
	expected := 439635
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := Part2(inputDay)
	expected := 3562722971
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(inputDay)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(inputDay)
	}
}
