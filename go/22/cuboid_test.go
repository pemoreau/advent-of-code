package main

import (
	"fmt"
	"testing"
)

func TestSize(t *testing.T) {
	list := []Cuboid{
		CreateCuboid(10, 12, 10, 12, 10, 12, false),
		CreateCuboid(11, 11, 11, 11, 11, 11, false),
		CreateCuboid(10, 9, 10, 12, 10, 12, false),
	}
	expected := []int{27, 1, 0}
	for i, c := range list {
		result := Size(c)
		if result != expected[i] {
			t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
		}
	}
}

func TestInformation(t *testing.T) {
	c := CreateCuboid(10, 12, 10, 12, 10, 12, true)
	result := c.information
	expected := true
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}
func TestInclude(t *testing.T) {
	c1 := CreateCuboid(10, 12, 10, 12, 10, 12, true)
	c2 := CreateCuboid(11, 11, 11, 11, 11, 11, true)
	result := Include(c2, c1)
	expected := true
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}
func TestIntersection1(t *testing.T) {
	c1 := CreateCuboid(10, 12, 10, 12, 10, 12, true)
	c2 := CreateCuboid(11, 11, 11, 11, 11, 11, true)
	result := len(Intersection(c1, c2))
	expected := 27
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestIntersection2(t *testing.T) {
	c1 := CreateCuboid(10, 12, 10, 12, 10, 12, true)
	c2 := CreateCuboid(11, 13, 11, 13, 11, 13, true)
	list := Intersection(c1, c2)
	result := 0
	for _, c := range list {
		result += Size(c)
	}
	expected := 27 + 19
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestIntersection3(t *testing.T) {
	c1 := CreateCuboid(10, 12, 10, 12, 10, 12, true)
	c2 := CreateCuboid(11, 13, 11, 13, 11, 13, true)
	c3 := CreateCuboid(9, 11, 9, 11, 9, 11, false)
	c4 := CreateCuboid(10, 10, 10, 10, 10, 10, true)

	list := IntersectionList(IntersectionList(Intersection(c1, c2), c3), c4)
	size := 0
	result := 0
	for _, c := range list {
		fmt.Println(c)
		size += Size(c)
		if c.information {
			result += Size(c)
		}
	}
	expected := 39
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
