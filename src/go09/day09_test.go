package main

import (
	"strings"
	"testing"
)

const input = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestPart1(t *testing.T) {
	m := BuildMatrix(strings.Split(input, "\n"))
	result := Part1(m)
	expected := 15
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	m := BuildMatrix(strings.Split(input, "\n"))
	result := Part2(m)
	expected := 1134
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	lines := ReadFile("../../inputs/day09.txt")
	m := BuildMatrix(lines)
	result := Part1(m)
	expected := 562
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	lines := ReadFile("../../inputs/day09.txt")
	m := BuildMatrix(lines)
	result := Part2(m)
	expected := 1076922
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
