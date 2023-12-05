package main

import (
	_ "embed"
	"github.com/pemoreau/advent-of-code/go/utils"
	"testing"
)

//go:embed input_test.txt
var inputTest string

func equals(a, b []utils.Interval) bool {
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
	_, other := transformer(utils.Interval{80, 80}, utils.Interval{98, 99}, 50)
	expectedOther := []utils.Interval{{80, 80}}
	if !equals(other, expectedOther) {
		t.Errorf("Other is incorrect, got: %d, want: %d.", other, expectedOther)
	}
	result, _ := transformer(utils.Interval{98, 98}, utils.Interval{98, 99}, 50)
	expected := utils.Interval{50, 50}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result, _ = transformer(utils.Interval{99, 99}, utils.Interval{98, 99}, 50)
	expected = utils.Interval{51, 51}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	_, other = transformer(utils.Interval{100, 100}, utils.Interval{98, 99}, 50)
	expectedOther = []utils.Interval{{100, 100}}
	if !equals(other, expectedOther) {
		t.Errorf("Other is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestTransformer2(t *testing.T) {
	result, other := transformer(utils.Interval{80, 98}, utils.Interval{98, 99}, 50)
	expected := utils.Interval{50, 50}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
	expectedOther := []utils.Interval{{80, 97}}
	if !equals(other, expectedOther) {
		t.Errorf("Other is incorrect, got: %d, want: %d.", result, expected)
	}

	result, _ = transformer(utils.Interval{98, 99}, utils.Interval{98, 99}, 50)
	expected = utils.Interval{50, 51}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result, other = transformer(utils.Interval{99, 101}, utils.Interval{98, 99}, 50)
	expected = utils.Interval{51, 51}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
	expectedOther = []utils.Interval{{100, 101}}
	if !equals(other, expectedOther) {
		t.Errorf("Other is incorrect, got: %d, want: %d.", result, expected)
	}

	result, other = transformer(utils.Interval{80, 110}, utils.Interval{98, 99}, 50)
	expected = utils.Interval{50, 51}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
	expectedOther = []utils.Interval{{80, 97}, {100, 110}}
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
