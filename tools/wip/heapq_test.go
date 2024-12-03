package wip

import (
	"testing"
)

func TestCapacity(t *testing.T) {
	hq := NewHeap[int](100_000, func(a, b *int) bool {
		return b == nil || a != nil && *a < *b
	})

	for idx := 0; idx < 100_000; idx++ {
		err := hq.Push(idx)
		if err != nil {
			t.Errorf("%v", err)
		}
	}

	for idx := 0; idx < 100_000; idx++ {
		el, err := hq.Pop()
		if err != nil {
			t.Errorf("%v", err)
		}
		if el != idx {
			t.Errorf("Result is incorrect, got: %d, want: %d.", el, idx)
		}
	}
}

func TestOrdering(t *testing.T) {
	hq := NewHeap[int](10, func(a, b *int) bool {
		return b == nil || a != nil && *a < *b
	})

	hq.Push(10)
	hq.Push(50)
	hq.Push(100)
	hq.Push(5)
	hq.Push(25)
	hq.Push(75)
	hq.Push(150)

	if e, _ := hq.Peek(); e != 5 {
		t.Errorf("Result is incorrect, got: %d, want: %d.", e, 5)
	}

	expected := []int{5, 10, 25, 50, 75}
	for idx := 0; idx < len(expected); idx++ {
		e, _ := hq.Pop()
		if e != expected[idx] {
			t.Errorf("Result is incorrect, got: %d, want: %d.", e, expected[idx])
		}
	}

	hq.Push(50)
	hq.Push(10)
	hq.Push(1000)

	expected = []int{10, 50, 100, 150, 1000}
	for idx := 0; idx < len(expected); idx++ {
		e, _ := hq.Pop()
		if e != expected[idx] {
			t.Errorf("Result is incorrect, got: %d, want: %d.", e, expected[idx])
		}
	}
}

func TestLeaf(t *testing.T) {
	hq := NewHeap[int](10, func(a, b *int) bool {
		return b == nil || a != nil && *a < *b
	})

	hq.Push(10)
	if !hq.isLeaf(0) {
		t.Errorf("Result is incorrect, got: %v, want: %v.", hq.isLeaf(0), true)
	}
	hq.Push(5)
	if hq.isLeaf(0) {
		t.Errorf("Result is incorrect, got: %v, want: %v.", hq.isLeaf(0), false)
	}
	if !hq.isLeaf(1) {
		t.Errorf("Result is incorrect, got: %v, want: %v.", hq.isLeaf(1), true)
	}

}
