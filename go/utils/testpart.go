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
				t.Fatalf("%d-%s: input %s got %d, want %d", i, fnName, tc.Input, got, tc.Expected)
			}
		})
	}
}

// dedent supprime l'indentation commune des blocs multi-lignes pour plus de lisibilité.
func Dedent(s string) string {
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
