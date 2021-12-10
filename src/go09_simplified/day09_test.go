package main

import (
	"io/ioutil"
	"testing"
)

const input = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestPart1(t *testing.T) {
	result := Part1(input)
	expected := 15
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(input)
	expected := 1134
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	content, _ := ioutil.ReadFile("../../inputs/day09.txt")
	result := Part1(string(content))
	expected := 562
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	content, _ := ioutil.ReadFile("../../inputs/day09.txt")
	result := Part2(string(content))
	expected := 1076922
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
