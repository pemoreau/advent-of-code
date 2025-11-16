package utils

import (
	_ "embed"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

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

// dedent supprime l'indentation commune des blocs multi-lignes pour plus de lisibilité.
func Dedent(s string) string {
	lines := strings.Split(strings.Trim(s, "\n"), "\n")
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
			if len(ln) >= left {
				lines[i] = ln[left:]
			}
		}
	}
	return strings.Join(lines, "\n")
}
