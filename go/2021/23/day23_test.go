package main

import (
	_ "embed"
	"github.com/pemoreau/advent-of-code/go/utils"
	"testing"
)

//go:embed input2.txt
var input2 string

//go:embed input3.txt
var input3 string

func TestPart1Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part1(inputDay)
	expected := 15516
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part2(inputDay)
	expected := 45272
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

func TestPart1Input2(t *testing.T) {
	result := Part1(string(input2))
	expected := 16157
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input2(t *testing.T) {
	result := Part2(string(input2))
	expected := 43481
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input3(t *testing.T) {
	result := Part1(string(input3))
	expected := 11854
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input3(t *testing.T) {
	result := Part2(string(input3))
	expected := 64434
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
