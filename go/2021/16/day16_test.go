package main

import (
	_ "embed"
	"github.com/pemoreau/advent-of-code/go/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	inputs := []string{"8A004A801A8002F478", "620080001611562C8802118E34", "C0015000016115A2E0802F182340", "A0016C880162017C3686B18A3D4780"}
	expected := []int{16, 12, 23, 31}
	for i, input := range inputs {
		result := Part1(input)
		if result != expected[i] {
			t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected[i])
		}
	}
}

func TestPart2(t *testing.T) {
	inputs := []string{"C200B40A82", "04005AC33890", "880086C3E88112", "CE00C43D881120", "D8005AC2A8F0", "F600BC2D8F", "9C005AC2F8F0", "9C0141080250320F1802104A08"}
	expected := []int{3, 54, 7, 9, 1, 0, 0, 1}
	for i, input := range inputs {
		result := Part2(input)
		if result != expected[i] {
			t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected[i])
		}
	}
}

func TestPart1Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part1(inputDay)
	expected := 974
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part2(inputDay)
	expected := 180616437720
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
