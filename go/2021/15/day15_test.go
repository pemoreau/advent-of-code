package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input_test.txt
var input_test string

//go:embed matrix_test.txt
var matrix_test string

func TestPart1(t *testing.T) {
	result := Part1(input_test)
	expected := 40
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(input_test)
	expected := 315
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestMegaMatrix(t *testing.T) {
	lines := strings.Split(strings.TrimSuffix(matrix_test, "\n"), "\n")

	m := BuildMatrix(strings.Split(strings.TrimSuffix(input_test, "\n"), "\n"))
	mm := buildMegaMatrix(m)

	expected := len(lines)
	result := len(mm)
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	expected = len(lines[0])
	result = len(mm[0])
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	for j := 0; j < len(mm); j++ {
		for i := 0; i < len(mm[j]); i++ {
			expected = int(lines[j][i] - '0')
			result = int(mm[j][i])
			if result != expected {
				t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
			}
		}
	}
}

func TestPart1Input(t *testing.T) {
	result := Part1(string(input_day))
	expected := 824
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2Input(t *testing.T) {
	result := Part2(string(input_day))
	expected := 3063
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(input_day)
	}
}
func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(input_day)
	}
}
