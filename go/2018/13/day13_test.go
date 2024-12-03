package main

import (
	_ "embed"
	"github.com/pemoreau/advent-of-code/go/utils"
	"testing"
)

//go:embed sample2.txt
var inputTest2 string

func TestPart1(t *testing.T) {
	result := Part1(inputTest)
	expected := "7,3"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest2)
	expected := "6,4"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part1(inputDay)
	expected := "83,49"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part2(inputDay)
	expected := "73,36"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
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
