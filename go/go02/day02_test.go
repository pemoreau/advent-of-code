package main

import (
	"testing"
)

const input_test = "input_test.txt"
const input = "../../inputs/day02.txt"

func TestPart1(t *testing.T) {
	result := Part1(input_test)
	expected := 150
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2(t *testing.T) {
	result := Part2(input_test)
	expected := 900
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := Part1(input)
	expected := 1648020
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2Input(t *testing.T) {
	result := Part2(input)
	expected := 1759818555
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
