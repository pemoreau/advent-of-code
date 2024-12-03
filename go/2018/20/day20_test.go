package main

import (
	_ "embed"
	"github.com/pemoreau/advent-of-code/go/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	result := Part1("^ENWWW(NEEE|SSE(EE|N))$")
	expected := 10
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Part1("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$")
	expected = 18
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}

	result = Part1("^^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$")
	expected = 23
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part1(inputDay)
	expected := 3560
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part2(inputDay)
	expected := 8688
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	var inputDay = utils.Input()
	for range b.N {
		Part1(inputDay)
	}
}

func BenchmarkPart2(b *testing.B) {
	var inputDay = utils.Input()
	for range b.N {
		Part2(inputDay)
	}
}
