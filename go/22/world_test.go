package main

import (
	"testing"
)

func TestCount(t *testing.T) {
	world := World{}
	world.Add(CreateCuboid(10, 13, 10, 13, 10, 13), 1)
	expected := 27
	result := world.Count(1)
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestAdd1(t *testing.T) {
	world := World{}
	world.Add(CreateCuboid(10, 13, 10, 13, 10, 13), 1)
	world.Add(CreateCuboid(11, 12, 11, 12, 11, 12), 1)
	expected := 27
	result := world.Count(1)
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestAdd2(t *testing.T) {
	world := World{}
	world.Add(CreateCuboid(10, 13, 10, 13, 0, 1), 1)
	world.Add(CreateCuboid(11, 12, 11, 12, 0, 1), 2)
	expected := 9 - 1
	result := world.Count(1)
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestAdd3(t *testing.T) {
	world := World{}
	world.Add(CreateCuboid(10, 13, 10, 13, 10, 13), 1)
	world.Add(CreateCuboid(11, 12, 11, 12, 11, 12), 1)
	world.Add(CreateCuboid(11, 14, 11, 14, 11, 14), 1)
	expected := 27 + 19
	result := world.Count(1)
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestAdd4(t *testing.T) {
	world := World{}
	world.Add(CreateCuboid(10, 13, 10, 13, 10, 13), 1)
	world.Add(CreateCuboid(11, 14, 11, 14, 11, 14), 1)
	world.Add(CreateCuboid(9, 12, 9, 12, 9, 12), 2)
	world.Add(CreateCuboid(10, 11, 10, 11, 10, 11), 1)
	expected := 39
	result := world.Count(1)
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
