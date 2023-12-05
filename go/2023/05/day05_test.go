package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func equals(a, b []Interval) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestTransformer1(t *testing.T) {
	_, other := transformer(Interval{80, 80}, Interval{98, 99}, 50)
	expectedOther := []Interval{{80, 80}}
	if !equals(other, expectedOther) {
		t.Errorf("Other is incorrect, got: %d, want: %d.", other, expectedOther)
	}
	result, _ := transformer(Interval{98, 98}, Interval{98, 99}, 50)
	expected := Interval{50, 50}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result, _ = transformer(Interval{99, 99}, Interval{98, 99}, 50)
	expected = Interval{51, 51}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	_, other = transformer(Interval{100, 100}, Interval{98, 99}, 50)
	expectedOther = []Interval{{100, 100}}
	if !equals(other, expectedOther) {
		t.Errorf("Other is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestTransformer2(t *testing.T) {
	result, other := transformer(Interval{80, 98}, Interval{98, 99}, 50)
	expected := Interval{50, 50}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
	expectedOther := []Interval{{80, 97}}
	if !equals(other, expectedOther) {
		t.Errorf("Other is incorrect, got: %d, want: %d.", result, expected)
	}

	result, _ = transformer(Interval{98, 99}, Interval{98, 99}, 50)
	expected = Interval{50, 51}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result, other = transformer(Interval{99, 101}, Interval{98, 99}, 50)
	expected = Interval{51, 51}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
	expectedOther = []Interval{{100, 101}}
	if !equals(other, expectedOther) {
		t.Errorf("Other is incorrect, got: %d, want: %d.", result, expected)
	}

	result, other = transformer(Interval{80, 110}, Interval{98, 99}, 50)
	expected = Interval{50, 51}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
	expectedOther = []Interval{{80, 97}, {100, 110}}
	if !equals(other, expectedOther) {
		t.Errorf("Other is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := Part1(inputTest)
	expected := 35
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := 46
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := Part1(inputDay)
	expected := 157211394
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := Part2(inputDay)
	expected := 50855035
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
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
