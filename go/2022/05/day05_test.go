package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func TestPart1(t *testing.T) {
	result := Part1(inputTest)
	expected := "CMZ"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := Part1(inputDay)
	expected := "MQTPGLLDN"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := "MCD"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := Part2(inputDay)
	expected := "LVZPSTTCZ"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
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
