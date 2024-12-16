package utils

import (
	"reflect"
	"testing"
)

func TestFindNextPeriod(t *testing.T) {
	tests := []struct {
		cycle string
		start int
		want  int
	}{
		{"KKKKKKKKK", 0, 1},
		{"ABABABABAB", 0, 2},
		{"ABABABABA", 0, 2},
		{"ABCABCABCAB", 0, 3},
		{"ABABABCABABABCABAB", 0, 2},
		{"ABABABCABABABCABAB", 2, 7},
		{"ABABCABABCDABABCABABCDA", 0, 2},
		{"ABABCABABCDABABCABABCDA", 2, 5},
		{"ABABCABABCDABABCABABCDA", 5, 11},
	}

	for _, test := range tests {
		t.Run(test.cycle, func(t *testing.T) {
			got := FindNextPeriod([]byte(test.cycle), test.start)
			check(t, "period", got, test.want)
		})
	}
}

func TestCheckPeriod(t *testing.T) {
	tests := []struct {
		cycle  string
		period int
		dump   int
		want   bool
	}{
		{"KKKKKKKKK", 1, 0, true},
		{"KKKKKKKKK", 3, 0, true},
		{"KKKKKKKKK", 4, 0, true},
		{"KKKKKKKKK", 5, 0, false},
		{"ABABABABAB", 2, 0, true},
		{"ABABABABA", 2, 0, true},
		{"ABCABCABCAB", 3, 0, true},
		{"ABABABCABABABCABAB", 2, 0, false},
		{"ABABABCABABABCABAB", 7, 0, true},
		{"ABABCABABCDABABCABABCDA", 11, 0, true},
		{"ABABCABABCDABABCABABCDX", 11, 0, false},
		{"ABABCABABCDABABCABABCDX", 11, 1, true},
		{"ABABCABABCDABABCABABCDX", 11, 2, true},
		{"ABABCABABCDABABCABABCXX", 11, 2, false},
		{"ABABCABABCDABABCABABCDABABCABABCXX", 11, 2, true},
		{"ABABCABABCDABABCABABCDABABCABABXXX", 11, 2, false},
	}

	for _, test := range tests {
		t.Run(test.cycle, func(t *testing.T) {
			got := CheckPeriod([]byte(test.cycle), test.period, test.dump)
			check(t, "period", got, test.want)
		})
	}
}
func TestFindPeriod(t *testing.T) {
	tests := []struct {
		cycle string
		want  int
	}{
		{"KKKKKKKKK", 1},
		{"ABABABABAB", 2},
		{"ABABABABA", 2},
		{"ABCABCABCAB", 3},
		{"ABABABCABABABCABAB", 7},
		{"ABABCABABCDABABCABABCDA", 11},
		{"XXKKKKKKKKK"[2:], 1},
		{"XXABABABABAB"[2:], 2},
		{"XXABABABABA"[3:], 2},
		{"XXABCABCABCAB"[4:], 3},
		{"XXABABABCABABABCABAB"[3:], 7},
		{"XXABABCABABCDABABCABABCDA"[2:], 11},
	}

	for _, test := range tests {
		t.Run(test.cycle, func(t *testing.T) {
			got := FindPeriod([]byte(test.cycle))
			check(t, "period", got, test.want)
		})
	}
}

func check(t *testing.T, msg string, got, want any) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Result is incorrect, got: %v, want: %v.", got, want)
	}
}
