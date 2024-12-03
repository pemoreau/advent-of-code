package interval

import (
	"testing"
)

func TestSize(t *testing.T) {
	list := []Cuboid{
		CreateCuboid(10, 13, 10, 13, 10, 13),
		CreateCuboid(11, 12, 11, 12, 11, 12),
	}
	expected := []int{27, 1}
	for i, c := range list {
		result := c.Size()
		if result != expected[i] {
			t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected[i])
		}
	}
}

func TestContains(t *testing.T) {
	c1 := CreateCuboid(10, 13, 10, 13, 10, 13)
	c2 := CreateCuboid(11, 12, 11, 12, 11, 12)
	result := c1.Contains(c2)
	expected := true
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}
func TestIntersection(t *testing.T) {
	c1 := CreateCuboid(0, 1, 0, 1, 0, 1)
	c2 := CreateCuboid(1, 2, 1, 2, 1, 2)
	_, result := Intersection(c1, c2)
	expected := false
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestOverlap1(t *testing.T) {
	c1 := CreateCuboid(10, 13, 10, 13, 10, 13)
	c2 := CreateCuboid(11, 12, 11, 12, 11, 12)
	list := c2.Overlap(c1)
	result := len(list)
	expected := 27
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestOverlap2(t *testing.T) {
	c1 := CreateCuboid(10, 13, 10, 13, 10, 13)
	c2 := CreateCuboid(11, 14, 11, 14, 11, 14)
	list := c2.Overlap(c1)
	result := 0
	for _, c := range list {
		result += c.Size()
	}
	expected := 27 + 19
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
