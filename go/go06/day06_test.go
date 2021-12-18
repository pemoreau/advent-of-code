package main

import (
	"io/ioutil"
	"testing"
)

func TestPart1(t *testing.T) {
	input := "3,4,3,1,2"
	result := Part1(input)
	expected := 5934
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2(t *testing.T) {
	input := "3,4,3,1,2"
	result := Part2(input)
	expected := 26984457539
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	content, _ := ioutil.ReadFile("../../inputs/day06.txt")
	input := string(content)
	result := Part1(input)
	expected := 351092
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2Input(t *testing.T) {
	content, _ := ioutil.ReadFile("../../inputs/day06.txt")
	input := string(content)
	result := Part2(input)
	expected := 1595330616005
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
