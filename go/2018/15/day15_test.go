package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/pemoreau/advent-of-code/go/utils"
)

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

func TestPart1Samples(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "Sample A -> 27730",
			input: dedent(`
				#######
				#.G...#
				#...EG#
				#.#.#G#
				#..G#E#
				#.....#
				#######
			`),
			expected: 27730,
		},
		{
			name: "Sample B -> 36334",
			input: dedent(`
				#######
				#G..#E#
				#E#E.E#
				#G.##.#
				#...#E#
				#...E.#
				#######
			`),
			expected: 36334,
		},
		{
			name: "Sample C -> 39514",
			input: dedent(`
				#######
				#E..EG#
				#.#G.E#
				#E.##E#
				#G..#.#
				#..E#.#
				#######
			`),
			expected: 39514,
		},
		{
			name: "Sample D -> 27755",
			input: dedent(`
				#######
				#E.G#.#
				#.#G..#
				#G.#.G#
				#G..#.#
				#...E.#
				#######
			`),
			expected: 27755,
		},
		{
			name: "Sample E -> 28944",
			input: dedent(`
				#######
				#.E...#
				#.#..G#
				#.###.#
				#E#G#G#
				#...#G#
				#######
			`),
			expected: 28944,
		},
		{
			name: "Sample F -> 18740",
			input: dedent(`
				#########
				#G......#
				#.E.#...#
				#..##..G#
				#...##..#
				#...#...#
				#.G...G.#
				#.....G.#
				#########
			`),
			expected: 18740,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Part1(tc.input)
			if got != tc.expected {
				t.Fatalf("Part1 %s: got %d, want %d", tc.name, got, tc.expected)
			}
		})
	}
}
func TestPart1Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part1(inputDay)
	expected := 198531
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	var inputDay = utils.Input()
	result := Part2(inputDay)
	expected := 90420
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
