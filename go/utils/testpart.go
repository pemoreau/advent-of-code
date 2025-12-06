package utils

import (
	_ "embed"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// DefaultTabWidth is the number of spaces a tab character represents.
const DefaultTabWidth = 4

type Test[I, E comparable] struct {
	Func     func(I) E
	Input    I
	Expected E
}

func TestPart[I, E comparable](t *testing.T, tests []Test[I, E]) {
	for i, tc := range tests {
		fnName := runtime.FuncForPC(reflect.ValueOf(tc.Func).Pointer()).Name()
		t.Run(fnName, func(t *testing.T) {
			got := tc.Func(tc.Input)
			if got != tc.Expected {
				t.Fatalf("%d-%s: input %v got %v, want %v", i, fnName, tc.Input, got, tc.Expected)
			}
		})
	}
}

// expandLeadingTabs converts leading tab characters in a line to spaces.
func expandLeadingTabs(line string, tabWidth int) string {
	var builder strings.Builder
	originalIndex := 0 // Track index in original string
	for originalIndex < len(line) {
		r := rune(line[originalIndex]) // Get rune at current index
		if r == '\t' {
			builder.WriteString(strings.Repeat(" ", tabWidth))
			originalIndex++ // Consume one character from original
		} else if r == ' ' {
			builder.WriteRune(r)
			originalIndex++ // Consume one character from original
		} else {
			break // Stop if non-whitespace character
		}
	}
	return builder.String() + line[originalIndex:]
}

// dedent supprime l'indentation commune des blocs multi-lignes pour plus de lisibilité.
func Dedent(s string) string {
	rawLines := strings.Split(s, "\n")
	var lines []string
	for _, line := range rawLines {
		lines = append(lines, expandLeadingTabs(line, DefaultTabWidth))
	}

	// Trim leading empty lines
	start := 0
	for start < len(lines) && strings.TrimSpace(lines[start]) == "" {
		start++
	}

	// Trim trailing empty lines
	end := len(lines) - 1
	for end >= start && strings.TrimSpace(lines[end]) == "" {
		end--
	}

	if start > end {
		return "" // All lines were empty
	}

	lines = lines[start : end+1]

	// calcule l'indent left (ignore les lignes vides)
	left := -1
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		i := 0
		for i < len(line) && line[i] == ' ' {
			i++
		}
		if left == -1 || i < left {
			left = i
		}
	}
	if left > 0 {
		for i, ln := range lines {
			actualLeadingSpaces := 0
			for actualLeadingSpaces < len(ln) && ln[actualLeadingSpaces] == ' ' {
				actualLeadingSpaces++
			}

			numToTrim := min(left, actualLeadingSpaces)
			lines[i] = ln[numToTrim:]
		}
	}
	return strings.Join(lines, "\n")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
