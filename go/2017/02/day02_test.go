package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func TestPart1(t *testing.T) {
	result := Part1(inputTest)
	expected := 18
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

// dedent supprime l'indentation commune des blocs multi-lignes pour plus de lisibilité.
func dedent(s string) string {
	lines := strings.Split(strings.Trim(s, "\n"), "\n")
	// calcule l'indent min (ignore les lignes vides)
	min := -1
	for _, ln := range lines {
		if strings.TrimSpace(ln) == "" {
			continue
		}
		i := 0
		for i < len(ln) && ln[i] == ' ' {
			i++
		}
		if min == -1 || i < min {
			min = i
		}
	}
	if min > 0 {
		for i, ln := range lines {
			if len(ln) >= min {
				lines[i] = ln[min:]
			}
		}
	}
	return strings.Join(lines, "\n")
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "part2",
			input: dedent(`
				5 9 2 8
  				9 4 7 3
				3 8 6 5
			`),
			expected: 9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Part2(tc.input)
			if got != tc.expected {
				t.Fatalf("Part2 %s: got %d, want %d", tc.name, got, tc.expected)
			}
		})
	}
}

func TestPart1Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part1(inputDay)
	expected := 34925
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part2(inputDay)
	expected := 221
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
