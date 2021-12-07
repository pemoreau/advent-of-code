package main

import (
	"io/ioutil"
	"testing"
)

func TestPart1(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	result := Part1(input)
	expected := 37
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	result := Part2(input)
	expected := 168
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	content, _ := ioutil.ReadFile("../../inputs/day07.txt")
	input := string(content)
	result := Part1(input)
	expected := 355989
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2Input(t *testing.T) {
	content, _ := ioutil.ReadFile("../../inputs/day07.txt")
	input := string(content)
	result := Part2(input)
	expected := 102245489
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
