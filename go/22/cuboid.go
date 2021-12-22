package main

import (
	"fmt"

	"github.com/pemoreau/advent-of-code-2021/go/utils"
)

type interval [2]int    // [min, max[
type Cuboid [3]interval // x, y, z

func CreateCuboid(xmin, xmax, ymin, ymax, zmin, zmax int) Cuboid {
	if IsEmptyCuboid(xmin, xmax, ymin, ymax, zmin, zmax) {
		panic(fmt.Sprint("Invalid interval: ", xmin, xmax, ymin, ymax, zmin, zmax))
	}
	c := Cuboid{}
	c[0] = interval{xmin, xmax}
	c[1] = interval{ymin, ymax}
	c[2] = interval{zmin, zmax}
	return c
}

func IsEmptyCuboid(xmin, xmax, ymin, ymax, zmin, zmax int) bool {
	return xmin >= xmax || ymin >= ymax || zmin >= zmax
}

func (c Cuboid) String() string {
	return fmt.Sprintf("x: %v, y: %v, z: %v size=%d", c[0], c[1], c[2], Size(c))
}

func Include(a, b Cuboid) bool {
	for d := 0; d < 3; d++ {
		if a[d][0] < b[d][0] {
			return false
		}
		if a[d][1] > b[d][1] {
			return false
		}
	}
	return true
}

func Intersection(a, b Cuboid) (Cuboid, bool) {
	xmin, xmax := utils.Max(a[0][0], b[0][0]), utils.Min(a[0][1], b[0][1])
	ymin, ymax := utils.Max(a[1][0], b[1][0]), utils.Min(a[1][1], b[1][1])
	zmin, zmax := utils.Max(a[2][0], b[2][0]), utils.Min(a[2][1], b[2][1])
	if !IsEmptyCuboid(xmin, xmax, ymin, ymax, zmin, zmax) {
		return CreateCuboid(xmin, xmax, ymin, ymax, zmin, zmax), true
	}
	return a, false
}

func DisjointList(list []Cuboid, b Cuboid) bool {
	for _, a := range list {
		if _, ok := Intersection(a, b); ok {
			return false
		}
	}
	return true
}

func AllDisjoint(list []Cuboid) bool {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if _, ok := Intersection(list[i], list[j]); ok {
				fmt.Printf("not disjoint:\n%d\t%v\n%d\t%v\n", i, list[i], j, list[j])
				return false
			}
		}
	}
	return true
}

// b over a
func Overlap(a, b Cuboid) []Cuboid {
	if _, ok := Intersection(a, b); !ok {
		return []Cuboid{a, b}
	}

	dimension := [3][]int{}
	for i := 0; i < 3; i++ {
		dimension[i] = []int{
			utils.Min(a[i][0], b[i][0]),
			utils.Max(a[i][0], b[i][0]),
			utils.Min(a[i][1], b[i][1]),
			utils.Max(a[i][1], b[i][1]),
		}
		// fmt.Printf("d=%d --> %v\n", i, dimension[i])
	}
	res := make([]Cuboid, 0, 27)
	for i := 0; i < 3; i++ {
		xmin, xmax := dimension[0][i], dimension[0][i+1]
		for j := 0; j < 3; j++ {
			ymin, ymax := dimension[1][j], dimension[1][j+1]
			for k := 0; k < 3; k++ {
				zmin, zmax := dimension[2][k], dimension[2][k+1]
				if !IsEmptyCuboid(xmin, xmax, ymin, ymax, zmin, zmax) {
					c := CreateCuboid(xmin, xmax, ymin, ymax, zmin, zmax)
					if !DisjointList(res, c) {
						panic(fmt.Sprint("Not disjoint: ", c, res))
					}
					if Include(c, b) || Include(c, a) {
						res = append(res, c)
					}
				}
			}
		}
	}
	if !AllDisjoint(res) {
		panic(fmt.Sprint("Not all disjoint: ", res))
	}
	return res
}

func makeUniq(list []Cuboid) []Cuboid {
	set := make(map[Cuboid]bool)
	for _, c := range list {
		set[c] = true
	}
	res := make([]Cuboid, 0, len(set))
	for c := range set {
		res = append(res, c)
	}
	return res
}

func Size(c Cuboid) int {
	return (c[0][1] - c[0][0]) * (c[1][1] - c[1][0]) * (c[2][1] - c[2][0])
}
