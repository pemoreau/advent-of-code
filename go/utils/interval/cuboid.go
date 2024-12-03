package interval

import (
	"fmt"
)

type interval [2]int    // [min, max[
type Cuboid [3]interval // x, y, z

func CreateCuboid(xmin, xmax, ymin, ymax, zmin, zmax int) Cuboid {
	return Cuboid{interval{xmin, xmax}, interval{ymin, ymax}, interval{zmin, zmax}}
}

func (c Cuboid) IsEmpty() bool {
	return c[0][0] >= c[0][1] || c[1][0] >= c[1][1] || c[2][0] >= c[2][1]
}

func (c Cuboid) String() string {
	return fmt.Sprintf("x: %v, y: %v, z: %v size=%d", c[0], c[1], c[2], c.Size())
}

func (b Cuboid) Contains(a Cuboid) bool {
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
	xmin, xmax := max(a[0][0], b[0][0]), min(a[0][1], b[0][1])
	ymin, ymax := max(a[1][0], b[1][0]), min(a[1][1], b[1][1])
	zmin, zmax := max(a[2][0], b[2][0]), min(a[2][1], b[2][1])
	c := CreateCuboid(xmin, xmax, ymin, ymax, zmin, zmax)
	return c, !c.IsEmpty()
}

func (a Cuboid) Disjoint(b Cuboid) bool {
	xmin, xmax := max(a[0][0], b[0][0]), min(a[0][1], b[0][1])
	ymin, ymax := max(a[1][0], b[1][0]), min(a[1][1], b[1][1])
	zmin, zmax := max(a[2][0], b[2][0]), min(a[2][1], b[2][1])
	return xmin >= xmax || ymin >= ymax || zmin >= zmax
}

// b over a
func (b Cuboid) Overlap(a Cuboid) []Cuboid {
	if a.Disjoint(b) {
		return []Cuboid{a, b}
	}

	dimension := [3][]int{}
	for i := 0; i < 3; i++ {
		dimension[i] = []int{
			min(a[i][0], b[i][0]),
			max(a[i][0], b[i][0]),
			min(a[i][1], b[i][1]),
			max(a[i][1], b[i][1]),
		}
	}
	var res []Cuboid
	for i := 0; i < 3; i++ {
		xmin, xmax := dimension[0][i], dimension[0][i+1]
		for j := 0; j < 3; j++ {
			ymin, ymax := dimension[1][j], dimension[1][j+1]
			for k := 0; k < 3; k++ {
				zmin, zmax := dimension[2][k], dimension[2][k+1]
				c := CreateCuboid(xmin, xmax, ymin, ymax, zmin, zmax)
				if !c.IsEmpty() {
					if b.Contains(c) || a.Contains(c) {
						res = append(res, c)
					}
				}
			}
		}
	}
	return res
}

func (c Cuboid) Size() int {
	return (c[0][1] - c[0][0]) * (c[1][1] - c[1][0]) * (c[2][1] - c[2][0])
}
