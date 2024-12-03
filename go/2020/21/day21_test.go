package main

import (
	_ "embed"
	"github.com/pemoreau/advent-of-code/go/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	result := Part1(inputTest)
	expected := 5
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(inputTest)
	expected := "mxmxvkd,sqjhc,fvjkl"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part1(inputDay)
	expected := 2614
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part2(inputDay)
	expected := "qhvz,kbcpn,fzsl,mjzrj,bmj,mksmf,gptv,kgkrhg"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
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
