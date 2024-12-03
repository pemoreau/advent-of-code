package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type hailstone struct {
	x, y, z    int
	vx, vy, vz int
}

func intersect2d(a, b hailstone) (float64, float64, int, bool) {
	var xa, ya float64 = float64(a.x), float64(a.y)
	var xb, yb float64 = float64(b.x), float64(b.y)

	var vxa, vya float64 = float64(a.vx), float64(a.vy)
	var vxb, vyb float64 = float64(b.vx), float64(b.vy)

	// xa + vxa * tx = xb + vxb * ty
	// ya + vya * tx = yb + vyb * ty
	// vxa * tx - vxb * ty = xb - xa
	// tx = (xb - xa + vxb * ty) / vxa
	// ya + vya * (xb - xa + vxb * ty) / vxa = yb + vyb * ty
	// ya + (vya/vxa)*(xb-xa) + vya*vxb*ty/vxa = yb + vyb * ty
	// ya + (vya/vxa)*(xb-xa) - yb =  ty * (vyb - vya*vxb/vxa)

	if xb-xa == 0 {
		//fmt.Println("xb-xa == 0")
		return 0, 0, 0, false
	}
	if vxa == 0 {
		//fmt.Println("vxa == 0")
		return 0, 0, 0, false
	}
	if vyb-vya*vxb/vxa == 0 {
		//fmt.Println("vyb - vya*vxb/vxa == 0")
		return 0, 0, 0, false
	}

	var ty float64 = (ya - yb + (vya/vxa)*(xb-xa)) / (vyb - vya*vxb/vxa)
	var tx float64 = (xb - xa + vxb*ty) / vxa

	if tx < 0 || ty < 0 {
		//fmt.Println("tx < 0 || ty < 0")
		return 0, 0, 0, false
	}
	//fmt.Printf("tx: %f, ty: %f, cross (at x=%f, y=%f)\n", tx, ty, xa+vxa*tx, ya+vya*tx)

	return xa + vxa*tx, ya + vya*tx, int(tx), true
}

func intersect3d(a, b hailstone) (float64, float64, float64, bool) {
	var xa, ya, za = float64(a.x), float64(a.y), float64(a.z)
	var xb, yb, _ = float64(b.x), float64(b.y), float64(b.z)

	var vxa, vya, vza = float64(a.vx), float64(a.vy), float64(a.vz)
	var vxb, vyb, _ = float64(b.vx), float64(b.vy), float64(b.vz)

	// xa + vxa * tx = xb + vxb * ty
	// ya + vya * tx = yb + vyb * ty
	// za + vza * tx = zb + vzb * ty

	// vxa * tx - vxb * ty = xb - xa
	// tx = (xb - xa + vxb * ty) / vxa
	// ya + vya * (xb - xa + vxb * ty) / vxa = yb + vyb * ty
	// ya + (vya/vxa)*(xb-xa) + vya*vxb*ty/vxa = yb + vyb * ty
	// ya + (vya/vxa)*(xb-xa) - yb =  ty * (vyb - vya*vxb/vxa)

	if xb-xa == 0 {
		fmt.Println("xb-xa == 0")
		return 0, 0, 0, false
	}
	if vxa == 0 {
		fmt.Println("vxa == 0")
		return 0, 0, 0, false
	}
	if vyb-vya*vxb/vxa == 0 {
		fmt.Println("vyb - vya*vxb/vxa == 0")
		return 0, 0, 0, false
	}

	var ty float64 = (ya - yb + (vya/vxa)*(xb-xa)) / (vyb - vya*vxb/vxa)
	var tx float64 = (xb - xa + vxb*ty) / vxa

	if tx < 0 || ty < 0 {
		return 0, 0, 0, false
	}
	//fmt.Printf("tx: %f, ty: %f, cross (at x=%f, y=%f)\n", tx, ty, xa+vxa*tx, ya+vya*tx)

	return xa + vxa*tx, ya + vya*tx, za + vza*tx, true
}

//func (a hailstone) constraints()  {
//	var a, b, c = float64(a.x), float64(a.y), float64(a.z)
//	var u, v, w = float64(a.vx), float64(a.vy), float64(a.vz)
//
//	// a + u * s = X + VX * T
//	// b + v * s = Y + VY * T
//	// c + w * s = Z + VZ * T
//
//	// u * s - VX * T = X - a
//	// s = (X - a + VX * T) / u
//	// b + v * (X - a + VX * T) / u = Y + VY * T
//	// b + (v/u)*(X-a) + v*VX*T/u = Y + VY * T
//	// b + (v/u)*(X-a) - Y =  T * (VY - v*VX/u)
//
//	if X-a == 0 {
//		fmt.Println("X-a == 0")
//		return 0, 0, 0, false
//	}
//	if u == 0 {
//		fmt.Println("u == 0")
//		return 0, 0, 0, false
//	}
//	if VY-v*VX/u == 0 {
//		fmt.Println("VY - v*VX/u == 0")
//		return 0, 0, 0, false
//	}
//
//	var T float64 = (b - Y + (v/u)*(X-a)) / (VY - v*VX/u)
//	var s float64 = (X - a + VX*T) / u
//
//	if s < 0 || T < 0 {
//		return 0, 0, 0, false
//	}
//
//}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var data []hailstone
	for _, line := range lines {
		var x, y, z, vx, vy, vz int
		fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &x, &y, &z, &vx, &vy, &vz)
		h := hailstone{x, y, z, vx, vy, vz}
		data = append(data, h)
	}

	aMin := 200000000000000
	aMax := 400000000000000
	var res int
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			a := data[i]
			b := data[j]
			//fmt.Printf("A: %v\nB: %v\n", a, b)
			x, y, t, ok := intersect2d(a, b)
			if ok && t >= 0 {
				if x >= float64(aMin) && x <= float64(aMax) && y >= float64(aMin) && y <= float64(aMax) {
					//fmt.Println("in range")
					res++
				} else {
					//fmt.Println("out of range")
				}
			}
		}
	}

	return res
}

func findXY(data []hailstone, n int) (int, int, int, int) {
	for vx := -n; vx < n; vx++ {
		for vy := -n; vy < n; vy++ {
			d0 := data[0]
			d0.vx -= vx
			d0.vy -= vy
			d1 := data[1]
			d1.vx -= vx
			d1.vy -= vy
			x, y, _, ok := intersect2d(d0, d1)
			ix, iy := math.Round(x), math.Round(y)
			if ok {
				var found = true
				//for i := 2; i < len(data); i++ {
				for i := 2; i < 100; i++ {
					d := data[i]
					d.vx -= vx
					d.vy -= vy
					x0, y0, _, ok0 := intersect2d(d, d0)
					ix0, iy0 := math.Round(x0), math.Round(y0)
					//x1, y1, _, ok1 := intersect2d(d, d1)
					if !ok0 || ix0 != ix || iy0 != iy {
						//if !ok1 || x1 != x || y1 != y {
						found = false
						break
						//}
					}
				}
				if !found {
					continue
				}

				return int(ix), int(iy), vx, vy
			}
		}
	}
	return 0, 0, 0, 0
}

func findZ(data []hailstone, x, y, vx, vy int, n int) (int, int) {
	for vz := -n; vz < n; vz++ {
		d0 := data[0]
		d0.vx -= vx
		d0.vy -= vy
		d0.vz -= vz

		d1 := data[1]
		d1.vx -= vx
		d1.vy -= vy
		d1.vz -= vz

		x2, y2, z2, ok := intersect3d(d0, d1)
		ix2, iy2, iz2 := math.Round(x2), math.Round(y2), math.Round(z2)
		if ok {
			var found = true
			//for i := 2; i < len(data); i++ {
			for i := 2; i < 100; i++ {
				d := data[i]
				d.vx -= vx
				d.vy -= vy
				d.vz -= vz
				xx, yy, zz, ok := intersect3d(d, d0)
				ixx, iyy, izz := math.Round(xx), math.Round(yy), math.Round(zz)
				if !ok || ixx != ix2 || iyy != iy2 || izz != iz2 {
					found = false
					break
				}
			}
			if !found {
				continue
			}
			return int(iz2), vz
		}
	}
	return 0, 0
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var data []hailstone
	for _, line := range lines {
		var x, y, z, vx, vy, vz int
		fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &x, &y, &z, &vx, &vy, &vz)
		h := hailstone{x, y, z, vx, vy, vz}
		data = append(data, h)
	}

	n := 1000
	x, y, vx, vy := findXY(data, n)
	fmt.Printf("x: %d, y: %d, vx: %d, vy: %d\n", x, y, vx, vy)

	z, vz := findZ(data, x, y, vx, vy, n)
	fmt.Printf("z: %d, vz: %d\n", z, vz)

	return x + y + z
}

func main() {
	fmt.Println("--2023 day 24 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
