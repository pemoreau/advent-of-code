package main

import (
	"fmt"

	"github.com/pemoreau/advent-of-code-2021/go/utils"
)

type interval [2]int // [min, max]
type Cuboid struct {
	intervals   [3]interval // x, y, z
	information bool
}

func CreateCuboid(xmin, xmax, ymin, ymax, zmin, zmax int, info bool) Cuboid {
	c := Cuboid{}
	c.intervals[0] = interval{xmin, xmax + 1}
	c.intervals[1] = interval{ymin, ymax + 1}
	c.intervals[2] = interval{zmin, zmax + 1}
	c.information = info
	return c
}

func (c Cuboid) String() string {
	return fmt.Sprintf("x: %v, y: %v, z: %v i: %v", c.intervals[0], c.intervals[1], c.intervals[2], c.information)
}

func Include(a, b Cuboid) bool {
	for d := 0; d < 3; d++ {
		if a.intervals[d][0] < b.intervals[d][0] {
			return false
		}
		if a.intervals[d][1] > b.intervals[d][1] {
			return false
		}
	}
	return true
}

func Intersection(a, b Cuboid) []Cuboid {
	dimension := [3][]int{}
	for i := 0; i < 3; i++ {
		dimension[i] = []int{
			utils.Min(a.intervals[i][0], b.intervals[i][0]),
			utils.Max(a.intervals[i][0], b.intervals[i][0]),
			utils.Min(a.intervals[i][1], b.intervals[i][1]),
			utils.Max(a.intervals[i][1], b.intervals[i][1]),
		}
	}
	res := make([]Cuboid, 0, 27)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				c := CreateCuboid(dimension[0][i], dimension[0][i+1]-1, dimension[1][j], dimension[1][j+1]-1, dimension[2][k], dimension[2][k+1]-1, false)
				if Include(c, b) {
					c.information = b.information
					res = append(res, c)
				} else if Include(c, a) {
					c.information = a.information
					res = append(res, c)
				}
			}
		}
	}

	// ix := []int{
	// 	utils.Min(a.intervals[0][0], b.intervals[0][0]),
	// 	utils.Max(a.intervals[0][0], b.intervals[0][0]),
	// 	utils.Min(a.intervals[0][1], b.intervals[0][1]),
	// 	utils.Max(a.intervals[0][1], b.intervals[0][1]),
	// }
	return res
}

func IntersectionList(list []Cuboid, b Cuboid) []Cuboid {
	res := make([]Cuboid, 0)
	for _, a := range list {
		res = append(res, Intersection(a, b)...)
	}
	return res
}

func Size(c Cuboid) int {
	return (c.intervals[0][1] - c.intervals[0][0]) * (c.intervals[1][1] - c.intervals[1][0]) * (c.intervals[2][1] - c.intervals[2][0])
}
