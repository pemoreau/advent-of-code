package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

//go:embed input_test2_1.txt
var inputTest21 string

//go:embed input_test2_2.txt
var inputTest22 string

//go:embed input_test2_3.txt
var inputTest23 string

func TestPart1(t *testing.T) {
	result := Part1(inputTest)
	expected := 8
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart21(t *testing.T) {
	result := Part2(inputTest21)
	expected := 4
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart22(t *testing.T) {
	result := Part2(inputTest22)
	expected := 8
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart23(t *testing.T) {
	result := Part2(inputTest23)
	expected := 10
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := Part1(inputDay)
	expected := 6968
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := Part2(inputDay)
	expected := 413
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
