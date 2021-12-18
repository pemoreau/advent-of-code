package main

import (
	"io/ioutil"
	"testing"
)

const input = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func TestPart1(t *testing.T) {
	result := Part1(input)
	expected := 1656
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(input)
	expected := 195
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	content, _ := ioutil.ReadFile("../../inputs/day11.txt")
	result := Part1(string(content))
	expected := 1681
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	content, _ := ioutil.ReadFile("../../inputs/day11.txt")
	result := Part2(string(content))
	expected := 276
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
