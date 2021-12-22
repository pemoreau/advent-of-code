package main

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
		result := Size(c)
		if result != expected[i] {
			t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected[i])
		}
	}
}

// func TestInformation(t *testing.T) {
// 	c := CreateCuboid(10, 12, 10, 12, 10, 12, true)
// 	result := c.information
// 	expected := true
// 	if result != expected {
// 		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
// 	}
// }
func TestInclude(t *testing.T) {
	c1 := CreateCuboid(10, 13, 10, 13, 10, 13)
	c2 := CreateCuboid(11, 12, 11, 12, 11, 12)
	result := Include(c2, c1)
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
	list := Overlap(c1, c2)
	result := len(list)
	expected := 27
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestOverlap2(t *testing.T) {
	c1 := CreateCuboid(10, 13, 10, 13, 10, 13)
	c2 := CreateCuboid(11, 14, 11, 14, 11, 14)
	list := Overlap(c1, c2)
	result := 0
	for _, c := range list {
		result += Size(c)
	}
	expected := 27 + 19
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

// func TestOverlap3(t *testing.T) {
// 	c1 := CreateCuboid(10, 13, 10, 13, 0, 1)
// 	c2 := CreateCuboid(11, 14, 11, 14, 11, 14)
// 	list := Overlap(c1, c2)
// 	result := 0
// 	for _, c := range list {
// 		result += Size(c)
// 	}
// 	expected := 27 + 19
// 	if result != expected {
// 		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
// 	}
// }
// func TestOverlap3(t *testing.T) {
// 	c1 := CreateCuboid(10, 13, 10, 13, 10, 13, true)
// 	c2 := CreateCuboid(11, 12, 11, 12, 11, 12, false)
// 	list := Overlap(c1, c2)
// 	result := 0
// 	for _, c := range list {
// 		if c.information {
// 			result += Size(c)
// 		}
// 	}
// 	expected := 27 - 1
// 	if result != expected {
// 		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
// 	}
// }

// func TestOverlap4(t *testing.T) {
// 	c1 := CreateCuboid(10, 13, 10, 13, 10, 13, true)
// 	c2 := CreateCuboid(11, 14, 11, 14, 11, 14, false)
// 	list := Overlap(c1, c2)
// 	on := 0
// 	off := 0
// 	for _, c := range list {
// 		// fmt.Printf("c=%v size=%v\n", c, Size(c))
// 		if c.information {
// 			on += Size(c)
// 		} else {
// 			off += Size(c)
// 		}
// 	}
// 	expected := 27 - 8
// 	if on != expected {
// 		t.Errorf("Result is incorrect, got: %d, want: %d.", on, expected)
// 	}
// 	expected = 27
// 	if off != expected {
// 		t.Errorf("Result is incorrect, got: %d, want: %d.", off, expected)
// 	}
// }

// func TestOverlap5(t *testing.T) {
// 	c1 := CreateCuboid(10, 13, 10, 13, 10, 13, true)
// 	c2 := CreateCuboid(11, 14, 11, 14, 11, 14, true)
// 	c3 := CreateCuboid(11, 12, 11, 12, 11, 12, false)
// 	list := Overlap(c1, c2)
// 	for _, c := range list {
// 		fmt.Printf("c=%v size=%v\n", c, Size(c))
// 	}

// 	list = OverlapList(list, c3)
// 	size := 0
// 	result := 0
// 	for _, c := range list {
// 		// fmt.Printf("c=%v size=%v\n", c, Size(c))
// 		size += Size(c)
// 		if c.information {
// 			// fmt.Printf("c=%v size=%v\n", c, Size(c))
// 			result += Size(c)
// 		}
// 	}
// 	expected := 27 + 19 - 1
// 	if result != expected {
// 		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
// 	}
// }

// func TestOverlap6(t *testing.T) {
// 	c1 := CreateCuboid(10, 13, 10, 13, 10, 13, true)
// 	c2 := CreateCuboid(11, 14, 11, 14, 11, 14, true)
// 	c3 := CreateCuboid(9, 12, 9, 12, 9, 12, false)
// 	c4 := CreateCuboid(10, 11, 10, 11, 10, 11, true)

// 	list := OverlapList(Overlap(c1, c2), c3)
// 	// list := OverlapList(OverlapList(Overlap(c1, c2), c3), c4)
// 	size := 0
// 	result := 0
// 	for _, c := range list {
// 		if _, ok := Intersection(c, c4); ok {
// 			fmt.Printf("c4=%v c=%v\n", c4, c)
// 		}
// 		size += Size(c)
// 		if c.information {
// 			// fmt.Printf("c=%v size=%v\n", c, Size(c))
// 			result += Size(c)
// 		}
// 	}
// 	expected := 27 + 19 - 8
// 	if result != expected {
// 		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
// 	}
// }
